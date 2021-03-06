// Code generated by entc, DO NOT EDIT.

package region

const (
	// Label holds the string label denoting the region type in the database.
	Label = "region"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeRegiProv holds the string denoting the regi_prov edge name in mutations.
	EdgeRegiProv = "regi_prov"

	// Table holds the table name of the region in the database.
	Table = "regions"
	// RegiProvTable is the table the holds the regi_prov relation/edge.
	RegiProvTable = "provinces"
	// RegiProvInverseTable is the table name for the Province entity.
	// It exists in this package in order to avoid circular dependency with the "province" package.
	RegiProvInverseTable = "provinces"
	// RegiProvColumn is the table column denoting the regi_prov relation/edge.
	RegiProvColumn = "region_regi_prov"
)

// Columns holds all SQL columns for region fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
