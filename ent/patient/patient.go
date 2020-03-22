// Code generated by entc, DO NOT EDIT.

package patient

import (
	"time"

	"github.com/minskylab/asclepius/ent/schema"
)

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone vertex property in the database.
	FieldPhone = "phone"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldEmail holds the string denoting the email vertex property in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword = "password"
	// FieldFacebookID holds the string denoting the facebookid vertex property in the database.
	FieldFacebookID = "facebook_id"
	// FieldWatsonID holds the string denoting the watsonid vertex property in the database.
	FieldWatsonID = "watson_id"
	// FieldFirstContact holds the string denoting the first_contact vertex property in the database.
	FieldFirstContact = "first_contact"
	// FieldConditions holds the string denoting the conditions vertex property in the database.
	FieldConditions = "conditions"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// HistoryTable is the table the holds the history relation/edge.
	HistoryTable = "histories"
	// HistoryInverseTable is the table name for the History entity.
	// It exists in this package in order to avoid circular dependency with the "history" package.
	HistoryInverseTable = "histories"
	// HistoryColumn is the table column denoting the history relation/edge.
	HistoryColumn = "patient_history"
	// ScheduleTable is the table the holds the schedule relation/edge.
	ScheduleTable = "schedules"
	// ScheduleInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	ScheduleInverseTable = "schedules"
	// ScheduleColumn is the table column denoting the schedule relation/edge.
	ScheduleColumn = "patient_schedule"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhone,
	FieldAge,
	FieldEmail,
	FieldPassword,
	FieldFacebookID,
	FieldWatsonID,
	FieldFirstContact,
	FieldConditions,
}

var (
	fields = schema.Patient{}.Fields()

	// descAge is the schema descriptor for age field.
	descAge = fields[3].Descriptor()
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator = descAge.Validators[0].(func(int) error)

	// descPassword is the schema descriptor for password field.
	descPassword = fields[5].Descriptor()
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator = descPassword.Validators[0].(func(string) error)

	// descFirstContact is the schema descriptor for first_contact field.
	descFirstContact = fields[8].Descriptor()
	// DefaultFirstContact holds the default value on creation for the first_contact field.
	DefaultFirstContact = descFirstContact.Default.(func() time.Time)
)
