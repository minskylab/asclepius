package bot

import (
	"strings"
	"time"

	"github.com/minskylab/asclepius"
	neo "github.com/minskylab/neocortex"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (bot *Bot) actionsForFb(fb neo.CommunicationChannel) {
	title := neo.IfDialogNodeTitleIs("Episodio 3")

	bot.engine.Resolve(fb, title, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		c.SetContextVariable("name", c.Person.Name)
		dni := c.GetStringContextVariable("dni", "00000000")
		phone := c.GetStringContextVariable("phone", "+51957821858")

		phone = strings.TrimSpace(phone)
		if !strings.HasPrefix(phone, "+51") {
			phone = "+51" + phone
		}

		log.WithFields(log.Fields{
			"dni":   dni,
			"phone": phone,
		}).Info("generating sms alert")

		_ = response(c, out)

		go func(fbID string) {
			if _, err := bot.core.RegisterTaskFromBotPatient(
				fbID,
				"Seguimiento",
				"Este paciente necesita seguimiento personalizado, la convocatoria es" +
					"pública para todo doctor voluntario.",
				24 * time.Hour,
			); err != nil {
				log.Error(errors.Wrap(err, "error at register task from bot patient"))
			}
		}(c.Person.ID)

		timer, err := bot.emitter.LogMeasureBySMS(phone, c.Person.Name, 15*time.Minute)
		if err != nil {
			return errors.Wrap(err, "error at emmit message on episodio 3 resolve")
		}

		go func(to, name, dni string, timer *time.Timer) {
			<-timer.C
			if err := bot.emitter.SendRemember(to, name, dni); err != nil {
				panic(err)
			}
		}(phone, c.Person.Name, dni, timer)

		return nil
	})

	clinicResults := neo.IfDialogNodeTitleIs("Resultados Clínicos")
	bot.engine.Resolve(fb, clinicResults, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		test, err := bot.core.RegisterTestFromBot(c.Person.ID, asclepius.ClinicSymptoms{
			GeneralDiscomfort: c.Variables["b1"].(bool),
			Fever:             c.Variables["b2"].(bool),
			ThirdAge:          c.Variables["b3"].(bool),
			Dyspnea:           c.Variables["b4"].(bool),
		}, asclepius.EpidemicConditions{})
		if err != nil {
			log.WithField("session", c.SessionID).Error(errors.Cause(err))
		} else {
			c.SetContextVariable("last_test_id", test.ID.String())
		}

		return response(c, out)
	})

	epidemiologicResults := neo.IfDialogNodeTitleIs("Resultados Finales")
	bot.engine.Resolve(fb, epidemiologicResults, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		testID := c.GetStringContextVariable("last_test_id", "")
		if testID == "" {
			return response(c, out)
		}

		exposition := 0

		intimacyExposition := c.Variables["b5"].(bool)
		regionalExposition := c.Variables["b6"].(bool)
		globalExposition := c.Variables["b7"].(bool)

		if intimacyExposition {
			exposition = 4
		}

		if regionalExposition {
			exposition = 3
		}

		if globalExposition {
			exposition = 2
		}

		_, err := bot.core.AddEpidemiologicFactorsToTest(testID, asclepius.EpidemicConditions{
			VisitedPlaces:               []string{},
			InfectedFamily:              false,
			FromInfectedPlaceExposition: exposition,
			ToInfectedPlaceExposition:   exposition,
		})
		if err != nil {
			log.WithField("session", c.SessionID).Error(errors.Cause(err))
		}
		return response(c, out)
	})

	bot.engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		c.SetContextVariable("name", c.Person.Name)
		go func(c *neo.Context) {
			phone := c.GetStringContextVariable("phone", "")
			phoneSaved, _ := c.Variables["phone_saved"].(bool)
			if phone != "" && !phoneSaved {
				_, err := bot.core.UpdatePatientPhoneByFbAlias(c.Person.ID, phone)
				if err != nil {
					log.WithField("session", c.SessionID).Error(err)
					return
				}
				c.SetContextVariable("phone_saved", true)
			}

			if c.Variables["identified"] != nil {
				return
			}

			phone = c.GetStringContextVariable("phone", "+51000000000")

			log.WithField("session", c.SessionID).Info("verifying if user exist")

			isPatient, err := bot.core.UserOfFacebookIsPatient(c.Person.ID)
			if err != nil {
				log.WithField("session", c.SessionID).Error(err)
				return
			}

			if isPatient {
				c.SetContextVariable("identified", true)
				return
			}

			if _, err = bot.core.RegisterPatientFromFacebook(c.Person.ID, c.Person.Name, phone); err != nil {
				log.WithField("session", c.SessionID).Error(err)
			}
			c.SetContextVariable("identified", true)
		}(c)
		return response(c, out)
	})

}
