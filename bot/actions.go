package bot

import (
	"strings"
	"time"

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

		response(c, out)

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

	bot.engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		c.SetContextVariable("name", c.Person.Name)
		go func(sID, fbID, name, phone string) {
			log.WithField("session", sID).Info("verifying if user exist")
			isPatient, err := bot.core.UserOfFacebookIsPatient(fbID)
			if err != nil {
				log.WithField("session", sID).Error(err)
				return
			}

			if isPatient {
				return
			}

			if _, err = bot.core.RegisterPatientFromFacebook(fbID, name, phone); err != nil {
				log.WithField("session", sID).Error(err)
			}
		}(c.SessionID, c.Person.ID, c.Person.Name, c.GetStringContextVariable("phone", "+51000000000"))
		return response(c, out)
	})

}

