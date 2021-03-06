// Code generated by entc, DO NOT EDIT.

package clinicalresults

const (
	// Label holds the string label denoting the clinicalresults type in the database.
	Label = "clinical_results"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGeneralDiscomfort holds the string denoting the generaldiscomfort vertex property in the database.
	FieldGeneralDiscomfort = "general_discomfort"
	// FieldFever holds the string denoting the fever vertex property in the database.
	FieldFever = "fever"
	// FieldThirdAge holds the string denoting the thirdage vertex property in the database.
	FieldThirdAge = "third_age"
	// FieldDyspnea holds the string denoting the dyspnea vertex property in the database.
	FieldDyspnea = "dyspnea"

	// Table holds the table name of the clinicalresults in the database.
	Table = "clinical_results"
	// TestTable is the table the holds the test relation/edge.
	TestTable = "clinical_results"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "tests"
	// TestColumn is the table column denoting the test relation/edge.
	TestColumn = "test_clinical"
)

// Columns holds all SQL columns for clinicalresults fields.
var Columns = []string{
	FieldID,
	FieldGeneralDiscomfort,
	FieldFever,
	FieldThirdAge,
	FieldDyspnea,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ClinicalResults type.
var ForeignKeys = []string{
	"test_clinical",
}
