package asclepius

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent"
	"github.com/minskylab/asclepius/ent/history"
	"github.com/minskylab/asclepius/ent/patient"
	"github.com/minskylab/asclepius/ent/test"
	"github.com/pkg/errors"
)

type ClinicSymptoms struct {
	GeneralDiscomfort bool
	Fever             bool
	ThirdAge          bool
	Dyspnea           bool
}

type EpidemicConditions struct {
	VisitedPlaces               []string
	InfectedFamily              bool
	FromInfectedPlaceExposition int
	ToInfectedPlaceExposition   int
}

func (core *Core) registerNewPatientFromFacebook(facebookID string, name string, phone string) (*ent.Patient, error) {
	name = strings.Trim(name, " \n,.'\"\\/")
	phone = strings.Trim(name, " \n,.'\"\\/")

	tx, err := core.client.Tx(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error at init transaction")
	}

	h, err := tx.History.
		Create().
		SetID(uuid.New()).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at try to create history for new patient")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	p, err := tx.Patient.
		Create().
		SetID(uuid.New()).
		SetName(name).
		SetPhone(phone).
		SetFacebookID(facebookID).
		SetFirstContact(time.Now()).
		SetHistory(h).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at try to create a new patient")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "error at commit transaction")
	}

	return p, nil
}

func (core *Core) registerNewTestToPatientByAlias(alias string, clinic ClinicSymptoms, epidemic EpidemicConditions, notes ...string) (*ent.Test, error) {
	h, err := core.client.History.
		Query().
		Where(
			history.HasPatientWith(
				patient.Or(
					patient.FacebookIDEQ(alias),
					patient.WatsonIDEQ(alias),
				),
			),
		).Only(context.Background())

	if err != nil {
		return nil, errors.Wrap(err, "error at get one patient by alias")
	}

	tx, err := core.client.Tx(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error at init transaction")
	}

	clinical, err := tx.ClinicalResults.Create().
		SetID(uuid.New()).
		SetFever(clinic.Fever).
		SetGeneralDiscomfort(clinic.GeneralDiscomfort).
		SetThirdAge(clinic.ThirdAge).
		SetDyspnea(clinic.Dyspnea).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at create new clinical result")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	epidemiologic, err := tx.EpidemiologicResults.Create().
		SetID(uuid.New()).
		SetVisitedPlaces(epidemic.VisitedPlaces).
		SetInfectedFamily(epidemic.InfectedFamily).
		SetFromInfectedPlace(epidemic.FromInfectedPlaceExposition).
		SetToInfectedPlace(epidemic.ToInfectedPlaceExposition).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at create new epidemiologic result")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	t, err := tx.Test.Create().
		SetID(uuid.New()).
		SetHistory(h).
		AddClinical(clinical).
		AddEpidemiologic(epidemiologic).
		SetNotes(notes).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at create new test")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "error at perform transaction commit")
	}

	return t, nil
}

func (core *Core) addEpidemiologicToTest(testID string, epidemic EpidemicConditions) (*ent.Test, error) {
	id, err := uuid.Parse(testID)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id, cannot parse it")
	}

	tx, err := core.client.Tx(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error at create transaction")
	}

	epidemiologic, err := tx.EpidemiologicResults.Create().
		SetID(uuid.New()).
		SetVisitedPlaces(epidemic.VisitedPlaces).
		SetInfectedFamily(epidemic.InfectedFamily).
		SetFromInfectedPlace(epidemic.FromInfectedPlaceExposition).
		SetToInfectedPlace(epidemic.ToInfectedPlaceExposition).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at create new epidemiologic result")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	t, err := tx.Test.
		Update().
		Where(test.IDEQ(id)).
		AddEpidemiologic(epidemiologic).
		Save(context.Background())
	if err != nil {
		err = errors.Wrap(err, "error at create new test with epidemiological data")
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errors.Wrap(err, "error at perform tx rollback")
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "error at commit last transaction")
	}

	if t < 1 {
		return nil, errors.New("update test with epidemiological was not performed correctly")
	}

	return core.client.Test.Get(context.Background(), id)
}


func (core *Core) AddEpidemiologicFactorsToTest(testID string, epidemic EpidemicConditions) (*ent.Test, error) {
	return core.addEpidemiologicToTest(testID, epidemic)
}
func (core *Core) RegisterPatientFromFacebook(fbID, name, phone string) (*ent.Patient, error) {
	return core.registerNewPatientFromFacebook(fbID, name, phone)
}

func (core *Core) RegisterTestFromBot(fbID string, clinic ClinicSymptoms, epidemic EpidemicConditions, notes ...string) (*ent.Test, error) {
	return core.registerNewTestToPatientByAlias(fbID, clinic, epidemic)
}
