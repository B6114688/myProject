// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "activityname", Type: field.TypeString},
		{Name: "added", Type: field.TypeTime},
		{Name: "hours", Type: field.TypeString},
		{Name: "agency_agen_acti", Type: field.TypeInt, Nullable: true},
		{Name: "place_place_acti", Type: field.TypeInt, Nullable: true},
		{Name: "student_stud_acti", Type: field.TypeInt, Nullable: true},
		{Name: "term_term_acti", Type: field.TypeInt, Nullable: true},
		{Name: "year_year_acti", Type: field.TypeInt, Nullable: true},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "activities_agencies_agen_acti",
				Columns: []*schema.Column{ActivitiesColumns[4]},

				RefColumns: []*schema.Column{AgenciesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_places_place_acti",
				Columns: []*schema.Column{ActivitiesColumns[5]},

				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_students_stud_acti",
				Columns: []*schema.Column{ActivitiesColumns[6]},

				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_terms_term_acti",
				Columns: []*schema.Column{ActivitiesColumns[7]},

				RefColumns: []*schema.Column{TermsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_years_year_acti",
				Columns: []*schema.Column{ActivitiesColumns[8]},

				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AgenciesColumns holds the columns for the "agencies" table.
	AgenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "agency", Type: field.TypeString},
	}
	// AgenciesTable holds the schema information for the "agencies" table.
	AgenciesTable = &schema.Table{
		Name:        "agencies",
		Columns:     AgenciesColumns,
		PrimaryKey:  []*schema.Column{AgenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ContinentsColumns holds the columns for the "continents" table.
	ContinentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "continent", Type: field.TypeString, Unique: true},
	}
	// ContinentsTable holds the schema information for the "continents" table.
	ContinentsTable = &schema.Table{
		Name:        "continents",
		Columns:     ContinentsColumns,
		PrimaryKey:  []*schema.Column{ContinentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "country", Type: field.TypeString, Unique: true},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:        "countries",
		Columns:     CountriesColumns,
		PrimaryKey:  []*schema.Column{CountriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "course", Type: field.TypeString},
		{Name: "degree_degr_cour", Type: field.TypeInt, Nullable: true},
		{Name: "faculty_facu_cour", Type: field.TypeInt, Nullable: true},
		{Name: "institution_inst_cour", Type: field.TypeInt, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "courses_degrees_degr_cour",
				Columns: []*schema.Column{CoursesColumns[2]},

				RefColumns: []*schema.Column{DegreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courses_faculties_facu_cour",
				Columns: []*schema.Column{CoursesColumns[3]},

				RefColumns: []*schema.Column{FacultiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courses_institutions_inst_cour",
				Columns: []*schema.Column{CoursesColumns[4]},

				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DegreesColumns holds the columns for the "degrees" table.
	DegreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "degree", Type: field.TypeString},
	}
	// DegreesTable holds the schema information for the "degrees" table.
	DegreesTable = &schema.Table{
		Name:        "degrees",
		Columns:     DegreesColumns,
		PrimaryKey:  []*schema.Column{DegreesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmpsColumns holds the columns for the "emps" table.
	EmpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString},
		{Name: "pass", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
	}
	// EmpsTable holds the schema information for the "emps" table.
	EmpsTable = &schema.Table{
		Name:        "emps",
		Columns:     EmpsColumns,
		PrimaryKey:  []*schema.Column{EmpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FacultiesColumns holds the columns for the "faculties" table.
	FacultiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "faculty", Type: field.TypeString},
	}
	// FacultiesTable holds the schema information for the "faculties" table.
	FacultiesTable = &schema.Table{
		Name:        "faculties",
		Columns:     FacultiesColumns,
		PrimaryKey:  []*schema.Column{FacultiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString, Unique: true},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "institution", Type: field.TypeString},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:        "institutions",
		Columns:     InstitutionsColumns,
		PrimaryKey:  []*schema.Column{InstitutionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PlacesColumns holds the columns for the "places" table.
	PlacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "place", Type: field.TypeString},
	}
	// PlacesTable holds the schema information for the "places" table.
	PlacesTable = &schema.Table{
		Name:        "places",
		Columns:     PlacesColumns,
		PrimaryKey:  []*schema.Column{PlacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PrefixesColumns holds the columns for the "prefixes" table.
	PrefixesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "prefix", Type: field.TypeString},
	}
	// PrefixesTable holds the schema information for the "prefixes" table.
	PrefixesTable = &schema.Table{
		Name:        "prefixes",
		Columns:     PrefixesColumns,
		PrimaryKey:  []*schema.Column{PrefixesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProfessorsColumns holds the columns for the "professors" table.
	ProfessorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "tel", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "faculty_facu_prof", Type: field.TypeInt, Nullable: true},
		{Name: "prefix_ID", Type: field.TypeInt, Nullable: true},
		{Name: "professorship_ID", Type: field.TypeInt, Nullable: true},
	}
	// ProfessorsTable holds the schema information for the "professors" table.
	ProfessorsTable = &schema.Table{
		Name:       "professors",
		Columns:    ProfessorsColumns,
		PrimaryKey: []*schema.Column{ProfessorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "professors_faculties_facu_prof",
				Columns: []*schema.Column{ProfessorsColumns[4]},

				RefColumns: []*schema.Column{FacultiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "professors_prefixes_pre_prof",
				Columns: []*schema.Column{ProfessorsColumns[5]},

				RefColumns: []*schema.Column{PrefixesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "professors_professorships_pros_prof",
				Columns: []*schema.Column{ProfessorsColumns[6]},

				RefColumns: []*schema.Column{ProfessorshipsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProfessorshipsColumns holds the columns for the "professorships" table.
	ProfessorshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "professorship", Type: field.TypeString},
	}
	// ProfessorshipsTable holds the schema information for the "professorships" table.
	ProfessorshipsTable = &schema.Table{
		Name:        "professorships",
		Columns:     ProfessorshipsColumns,
		PrimaryKey:  []*schema.Column{ProfessorshipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProvincesColumns holds the columns for the "provinces" table.
	ProvincesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "province", Type: field.TypeString, Unique: true},
		{Name: "district", Type: field.TypeString},
		{Name: "subdistrict", Type: field.TypeString},
		{Name: "postal", Type: field.TypeInt},
		{Name: "continent_cont_prov", Type: field.TypeInt, Nullable: true},
		{Name: "country_coun_prov", Type: field.TypeInt, Nullable: true},
		{Name: "region_regi_prov", Type: field.TypeInt, Nullable: true},
	}
	// ProvincesTable holds the schema information for the "provinces" table.
	ProvincesTable = &schema.Table{
		Name:       "provinces",
		Columns:    ProvincesColumns,
		PrimaryKey: []*schema.Column{ProvincesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "provinces_continents_cont_prov",
				Columns: []*schema.Column{ProvincesColumns[5]},

				RefColumns: []*schema.Column{ContinentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "provinces_countries_coun_prov",
				Columns: []*schema.Column{ProvincesColumns[6]},

				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "provinces_regions_regi_prov",
				Columns: []*schema.Column{ProvincesColumns[7]},

				RefColumns: []*schema.Column{RegionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RegionsColumns holds the columns for the "regions" table.
	RegionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// RegionsTable holds the schema information for the "regions" table.
	RegionsTable = &schema.Table{
		Name:        "regions",
		Columns:     RegionsColumns,
		PrimaryKey:  []*schema.Column{RegionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ResultsColumns holds the columns for the "results" table.
	ResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "grade", Type: field.TypeFloat64},
		{Name: "student_stud_resu", Type: field.TypeInt, Nullable: true},
		{Name: "subject_subj_resu", Type: field.TypeInt, Nullable: true},
		{Name: "term_term_resu", Type: field.TypeInt, Nullable: true},
		{Name: "year_year_resu", Type: field.TypeInt, Nullable: true},
	}
	// ResultsTable holds the schema information for the "results" table.
	ResultsTable = &schema.Table{
		Name:       "results",
		Columns:    ResultsColumns,
		PrimaryKey: []*schema.Column{ResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "results_students_stud_resu",
				Columns: []*schema.Column{ResultsColumns[2]},

				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "results_subjects_subj_resu",
				Columns: []*schema.Column{ResultsColumns[3]},

				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "results_terms_term_resu",
				Columns: []*schema.Column{ResultsColumns[4]},

				RefColumns: []*schema.Column{TermsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "results_years_year_resu",
				Columns: []*schema.Column{ResultsColumns[5]},

				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fname", Type: field.TypeString},
		{Name: "lname", Type: field.TypeString},
		{Name: "schoolname", Type: field.TypeString},
		{Name: "recent_address", Type: field.TypeString},
		{Name: "telephone", Type: field.TypeInt},
		{Name: "email", Type: field.TypeString},
		{Name: "degree_degr_stud", Type: field.TypeInt, Nullable: true},
		{Name: "gender_gend_stud", Type: field.TypeInt, Nullable: true},
		{Name: "prefix_pref_stud", Type: field.TypeInt, Nullable: true},
		{Name: "province_prov_stud", Type: field.TypeInt, Nullable: true},
		{Name: "province_dist_stud", Type: field.TypeInt, Nullable: true},
		{Name: "province_subd_stud", Type: field.TypeInt, Nullable: true},
		{Name: "province_post_stud", Type: field.TypeInt, Nullable: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "students_degrees_degr_stud",
				Columns: []*schema.Column{StudentsColumns[7]},

				RefColumns: []*schema.Column{DegreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_genders_gend_stud",
				Columns: []*schema.Column{StudentsColumns[8]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_prefixes_pref_stud",
				Columns: []*schema.Column{StudentsColumns[9]},

				RefColumns: []*schema.Column{PrefixesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_provinces_prov_stud",
				Columns: []*schema.Column{StudentsColumns[10]},

				RefColumns: []*schema.Column{ProvincesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_provinces_dist_stud",
				Columns: []*schema.Column{StudentsColumns[11]},

				RefColumns: []*schema.Column{ProvincesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_provinces_subd_stud",
				Columns: []*schema.Column{StudentsColumns[12]},

				RefColumns: []*schema.Column{ProvincesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "students_provinces_post_stud",
				Columns: []*schema.Column{StudentsColumns[13]},

				RefColumns: []*schema.Column{ProvincesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeInt},
		{Name: "subjects", Type: field.TypeString},
		{Name: "creditpiont", Type: field.TypeInt},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:        "subjects",
		Columns:     SubjectsColumns,
		PrimaryKey:  []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TermsColumns holds the columns for the "terms" table.
	TermsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "semester", Type: field.TypeInt},
		{Name: "year_year_term", Type: field.TypeInt, Nullable: true},
	}
	// TermsTable holds the schema information for the "terms" table.
	TermsTable = &schema.Table{
		Name:       "terms",
		Columns:    TermsColumns,
		PrimaryKey: []*schema.Column{TermsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "terms_years_year_term",
				Columns: []*schema.Column{TermsColumns[2]},

				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// YearsColumns holds the columns for the "years" table.
	YearsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "years", Type: field.TypeInt},
	}
	// YearsTable holds the schema information for the "years" table.
	YearsTable = &schema.Table{
		Name:        "years",
		Columns:     YearsColumns,
		PrimaryKey:  []*schema.Column{YearsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		AgenciesTable,
		ContinentsTable,
		CountriesTable,
		CoursesTable,
		DegreesTable,
		EmpsTable,
		FacultiesTable,
		GendersTable,
		InstitutionsTable,
		PlacesTable,
		PrefixesTable,
		ProfessorsTable,
		ProfessorshipsTable,
		ProvincesTable,
		RegionsTable,
		ResultsTable,
		StudentsTable,
		SubjectsTable,
		TermsTable,
		YearsTable,
	}
)

func init() {
	ActivitiesTable.ForeignKeys[0].RefTable = AgenciesTable
	ActivitiesTable.ForeignKeys[1].RefTable = PlacesTable
	ActivitiesTable.ForeignKeys[2].RefTable = StudentsTable
	ActivitiesTable.ForeignKeys[3].RefTable = TermsTable
	ActivitiesTable.ForeignKeys[4].RefTable = YearsTable
	CoursesTable.ForeignKeys[0].RefTable = DegreesTable
	CoursesTable.ForeignKeys[1].RefTable = FacultiesTable
	CoursesTable.ForeignKeys[2].RefTable = InstitutionsTable
	ProfessorsTable.ForeignKeys[0].RefTable = FacultiesTable
	ProfessorsTable.ForeignKeys[1].RefTable = PrefixesTable
	ProfessorsTable.ForeignKeys[2].RefTable = ProfessorshipsTable
	ProvincesTable.ForeignKeys[0].RefTable = ContinentsTable
	ProvincesTable.ForeignKeys[1].RefTable = CountriesTable
	ProvincesTable.ForeignKeys[2].RefTable = RegionsTable
	ResultsTable.ForeignKeys[0].RefTable = StudentsTable
	ResultsTable.ForeignKeys[1].RefTable = SubjectsTable
	ResultsTable.ForeignKeys[2].RefTable = TermsTable
	ResultsTable.ForeignKeys[3].RefTable = YearsTable
	StudentsTable.ForeignKeys[0].RefTable = DegreesTable
	StudentsTable.ForeignKeys[1].RefTable = GendersTable
	StudentsTable.ForeignKeys[2].RefTable = PrefixesTable
	StudentsTable.ForeignKeys[3].RefTable = ProvincesTable
	StudentsTable.ForeignKeys[4].RefTable = ProvincesTable
	StudentsTable.ForeignKeys[5].RefTable = ProvincesTable
	StudentsTable.ForeignKeys[6].RefTable = ProvincesTable
	TermsTable.ForeignKeys[0].RefTable = YearsTable
}
