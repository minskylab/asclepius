package asclepius

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent"
	"github.com/minskylab/asclepius/ent/patient"
	"github.com/pkg/errors"
)

func (core *Core) updatePhoneOfPatient(patientID uuid.UUID, phone string) (*ent.Patient, error){
	phone = strings.Trim(phone, " \n,.'\"\\/")
	return core.client.Patient.UpdateOneID(patientID).SetPhone(phone).Save(context.Background())
}

func (core *Core) UpdatePatientPhoneByFbAlias(fbID string, phone string) (*ent.Patient, error) {
	pID, err := core.client.Patient.Query().Where(patient.FacebookIDEQ(fbID)).OnlyID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error at find patient by fb alias")
	}

	p, err := core.updatePhoneOfPatient(pID, phone)
	if err != nil {
		return nil, errors.Wrap(err, "error at update patient phone")
	}

	return p, nil
}