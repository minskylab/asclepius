package asclepius

import (
	"context"

	"github.com/minskylab/asclepius/ent/patient"
	"github.com/pkg/errors"
)

type UserRealm string

const Unknown UserRealm = "unknown"
const Patient UserRealm = "patient"
const Doctor UserRealm = "doctor"

func (core *Core) userReportByAlias(alias string) (UserRealm, error) {
	isPatient, err := core.client.Patient.
		Query().
		Where(patient.Or(patient.FacebookIDEQ(alias), patient.WatsonIDEQ(alias))).
		Exist(context.Background())
	if err != nil {
		return Unknown, errors.Wrap(err, "error at verify if patient with determined alias exist")
	}

	if isPatient {
		return Patient, nil
	}

	// errors.New("report implementation not complete")
	return Unknown, nil
}

func (core *Core) UserOfFacebookIsPatient(fbID string) (bool, error) {
	realm, err := core.userReportByAlias(fbID)
	if err != nil {
		return false, errors.Wrap(err, "error at generate report of user by alias")
	}

	if realm == Patient {
		return true, nil
	}

	return false, nil
}