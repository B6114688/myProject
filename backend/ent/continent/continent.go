// Code generated by entc, DO NOT EDIT.

package continent

const (
	// Label holds the string label denoting the continent type in the database.
	Label = "continent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContinent holds the string denoting the continent field in the database.
	FieldContinent = "continent"

	// EdgeContProv holds the string denoting the cont_prov edge name in mutations.
	EdgeContProv = "cont_prov"

	// Table holds the table name of the continent in the database.
	Table = "continents"
	// ContProvTable is the table the holds the cont_prov relation/edge.
	ContProvTable = "provinces"
	// ContProvInverseTable is the table name for the Province entity.
	// It exists in this package in order to avoid circular dependency with the "province" package.
	ContProvInverseTable = "provinces"
	// ContProvColumn is the table column denoting the cont_prov relation/edge.
	ContProvColumn = "continent_cont_prov"
)

// Columns holds all SQL columns for continent fields.
var Columns = []string{
	FieldID,
	FieldContinent,
}

var (
	// ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	ContinentValidator func(string) error
)
