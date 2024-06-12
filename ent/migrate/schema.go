// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OtherRequirementsColumns holds the columns for the "other_requirements" table.
	OtherRequirementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_member", Type: field.TypeInt},
		{Name: "is_quest", Type: field.TypeInt},
		{Name: "is_skill", Type: field.TypeInt},
		{Name: "id_of_requirement", Type: field.TypeInt},
		{Name: "id_unlock", Type: field.TypeInt, Nullable: true},
	}
	// OtherRequirementsTable holds the schema information for the "other_requirements" table.
	OtherRequirementsTable = &schema.Table{
		Name:       "other_requirements",
		Columns:    OtherRequirementsColumns,
		PrimaryKey: []*schema.Column{OtherRequirementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "other_requirements_unlocks_other_requirements",
				Columns:    []*schema.Column{OtherRequirementsColumns[6]},
				RefColumns: []*schema.Column{UnlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_member", Type: field.TypeInt},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// UnlocksColumns holds the columns for the "unlocks" table.
	UnlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_member", Type: field.TypeInt},
		{Name: "level", Type: field.TypeInt},
		{Name: "id_skill", Type: field.TypeInt, Nullable: true},
	}
	// UnlocksTable holds the schema information for the "unlocks" table.
	UnlocksTable = &schema.Table{
		Name:       "unlocks",
		Columns:    UnlocksColumns,
		PrimaryKey: []*schema.Column{UnlocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "unlocks_skills_unlocks",
				Columns:    []*schema.Column{UnlocksColumns[5]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OtherRequirementsTable,
		SkillsTable,
		UnlocksTable,
	}
)

func init() {
	OtherRequirementsTable.ForeignKeys[0].RefTable = UnlocksTable
	UnlocksTable.ForeignKeys[0].RefTable = SkillsTable
}
