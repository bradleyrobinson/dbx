package dbxdialect

import (
	"text/template"

	"bitbucket.org/pkg/inflect"
	"github.com/spacemonkeygo/dbx/internal/dbx"
)

func PostgresTableName(table *dbx.Table) string {
	return inflect.Underscore(inflect.Pluralize(table.Name))
}

func PostgresColumnName(column *dbx.Column) string {
	return inflect.Underscore(column.Name)
}

func PostgresColumnNames(columns []*dbx.Column) (out []string) {
	for _, column := range columns {
		out = append(out, PostgresColumnName(column))
	}
	return out
}

func PostgresColumnType(column *dbx.Column) string {
	switch column.Type {
	case "text":
		return "text"
	case "int":
		return "integer"
	case "serial":
		return "serial"
	case "int64":
		return "bigint"
	case "serial64":
		return "serial64"
	case "blob":
		return "bytea"
	case "timestamp":
		return "timestamp with time zone"
	case "bool":
		return "boolean"
	}
	panic("unhandled column type " + "%s")
}

type PostgresSchema struct {
	Tables []PostgresTable
}

func PostgresSchemaFromSchema(schema *dbx.Schema) PostgresSchema {
	return PostgresSchema{
		Tables: PostgresTablesFromTables(schema.Tables),
	}
}

type PostgresTable struct {
	Name       string
	Columns    []PostgresColumn
	PrimaryKey []string
	Unique     [][]string
}

func PostgresTableFromTable(table *dbx.Table) PostgresTable {
	t := PostgresTable{
		Name:       PostgresTableName(table),
		Columns:    PostgresColumnsFromColumns(table.Columns),
		PrimaryKey: PostgresColumnNames(table.PrimaryKey),
	}
	for _, unique := range table.Unique {
		t.Unique = append(t.Unique, PostgresColumnNames(unique))
	}
	return t
}

func PostgresTablesFromTables(tables []*dbx.Table) (out []PostgresTable) {
	for _, table := range tables {
		out = append(out, PostgresTableFromTable(table))
	}
	return out
}

type PostgresReference struct {
	Table    string
	Column   string
	OnDelete string
}

func PostgresReferenceFromRelation(relation *dbx.Relation) *PostgresReference {
	if relation == nil {
		return nil
	}
	r := &PostgresReference{
		Table:  PostgresTableName(relation.Column.Table),
		Column: PostgresColumnName(relation.Column),
	}

	switch relation.Kind {
	case dbx.HasA:
		r.OnDelete = "RESTRICT"
	case dbx.OwnedBy:
		r.OnDelete = "CASCADE"
	}
	return r
}

type PostgresColumn struct {
	Name      string
	NotNull   bool
	Reference *PostgresReference
}

func PostgresColumnFromColumn(column *dbx.Column) PostgresColumn {
	return PostgresColumn{
		Name:      PostgresColumnName(column),
		NotNull:   column.NotNull,
		Reference: PostgresReferenceFromRelation(column.Relation),
	}
}

func PostgresColumnsFromColumns(columns []*dbx.Column) (out []PostgresColumn) {
	for _, column := range columns {
		out = append(out, PostgresColumnFromColumn(column))
	}
	return out
}

type Postgres struct {
	tmpl *template.Template
}

func NewPostgres(loader dbx.Loader) (*Postgres, error) {
	tmpl, err := loader.Load("postgres.tmpl")
	if err != nil {
		return nil, err
	}

	return &Postgres{
		tmpl: tmpl,
	}, nil
}

func (s *Postgres) Name() string {
	return "postgres"
}

func (s *Postgres) RenderSchema(schema *dbx.Schema) (
	string, error) {

	return dbx.RenderTemplateString(s.tmpl, "schema",
		PostgresSchemaFromSchema(schema))
}

type PostgresColumnRef struct {
	Table  string
	Column string
}

func PostgresColumnRefFromColumn(column *dbx.Column) PostgresColumnRef {
	return PostgresColumnRef{
		Table:  column.Table.Name,
		Column: column.Name,
	}
}

type PostgresColumnCmp struct {
	Left     PostgresColumnRef
	Operator string
}

func PostgresColumnCmpFromColumnCmp(
	params *dbx.ColumnCmpParams) *PostgresColumnCmp {

	if params == nil {
		return nil
	}
	return &PostgresColumnCmp{
		Left:     PostgresColumnRefFromColumn(params.Left),
		Operator: params.Operator,
	}
}

type PostgresColumnCmpColumn struct {
	Left     PostgresColumnRef
	Right    PostgresColumnRef
	Operator string
}

func PostgresColumnCmpColumnFromColumnCmpColumn(
	params *dbx.ColumnCmpColumnParams) *PostgresColumnCmpColumn {

	if params == nil {
		return nil
	}
	return &PostgresColumnCmpColumn{
		Left:     PostgresColumnRefFromColumn(params.Left),
		Operator: params.Operator,
	}
}

type PostgresColumnIn struct {
	Left PostgresColumnRef
	In   PostgresSelect
}

func PostgresColumnInFromColumnIn(
	params *dbx.ColumnInParams) *PostgresColumnIn {

	if params == nil {
		return nil
	}
	return &PostgresColumnIn{
		Left: PostgresColumnRefFromColumn(params.Left),
		In:   PostgresSelectFromSelect(params.In),
	}
}

type PostgresCondition struct {
	ColumnCmp       *PostgresColumnCmp
	ColumnCmpColumn *PostgresColumnCmpColumn
	ColumnIn        *PostgresColumnIn
}

func PostgresConditionFromCondition(
	params *dbx.ConditionParams) PostgresCondition {
	return PostgresCondition{
		ColumnCmp:       PostgresColumnCmpFromColumnCmp(params.ColumnCmp),
		ColumnCmpColumn: PostgresColumnCmpColumnFromColumnCmpColumn(params.ColumnCmpColumn),
		ColumnIn:        PostgresColumnInFromColumnIn(params.ColumnIn),
	}
}

func PostgresConditionsFromConditions(conditions []*dbx.ConditionParams) (
	out []PostgresCondition) {

	for _, condition := range conditions {
		out = append(out, PostgresConditionFromCondition(condition))
	}
	return out
}

type PostgresOrderBy struct {
	Column PostgresColumnRef
}

type PostgresJoin struct {
	Left  PostgresColumnRef
	Right PostgresColumnRef
}

func PostgresJoinFromJoin(params *dbx.JoinParams) PostgresJoin {
	return PostgresJoin{
		Left:  PostgresColumnRefFromColumn(params.Left),
		Right: PostgresColumnRefFromColumn(params.Right),
	}
}

func PostgresJoinsFromJoins(joins []*dbx.JoinParams) (out []PostgresJoin) {
	for _, join := range joins {
		out = append(out, PostgresJoinFromJoin(join))
	}
	return out
}

type PostgresSelect struct {
	Table      string
	What       []string
	LeftJoins  []PostgresJoin
	Conditions []PostgresCondition
	OrderBy    *PostgresOrderBy
	Limit      bool
}

func PostgresSelectFromSelect(params *dbx.SelectParams) PostgresSelect {
	out := PostgresSelect{
		Table:      PostgresTableName(params.Table),
		What:       PostgresColumnNames(params.What),
		LeftJoins:  PostgresJoinsFromJoins(params.LeftJoins),
		Conditions: PostgresConditionsFromConditions(params.Conditions),
	}
	if params.PagedOn != nil {
		out.Conditions = append(out.Conditions, PostgresConditionFromCondition(
			&dbx.ConditionParams{
				ColumnCmp: &dbx.ColumnCmpParams{
					Left:     params.PagedOn,
					Operator: ">",
				},
			}))
		out.OrderBy = &PostgresOrderBy{
			Column: PostgresColumnRefFromColumn(params.PagedOn),
		}
		out.Limit = true
	}
	return out
}

func (s *Postgres) RenderSelect(params *dbx.SelectParams) (
	string, error) {

	return dbx.RenderTemplateString(s.tmpl, "select",
		PostgresSelectFromSelect(params))
}

func (s *Postgres) RenderCount(params *dbx.SelectParams) (
	string, error) {

	return dbx.RenderTemplateString(s.tmpl, "count",
		PostgresSelectFromSelect(params))
}

type PostgresDelete struct {
	Table      string
	Conditions []PostgresCondition
}

func PostgresDeleteFromDelete(params *dbx.DeleteParams) PostgresDelete {
	return PostgresDelete{
		Table:      params.Table.Name,
		Conditions: PostgresConditionsFromConditions(params.Conditions),
	}
}

func (s *Postgres) RenderDelete(params *dbx.DeleteParams) (
	string, error) {

	return dbx.RenderTemplateString(s.tmpl, "delete",
		PostgresDeleteFromDelete(params))
}

type PostgresInsert struct {
	Table   string
	Columns []string
}

func PostgresInsertFromInsert(params *dbx.InsertParams) PostgresInsert {
	return PostgresInsert{
		Table:   params.Table.Name,
		Columns: PostgresColumnNames(params.Columns),
	}
}

func (s *Postgres) RenderInsert(params *dbx.InsertParams) (
	string, error) {

	return dbx.RenderTemplateString(s.tmpl, "insert",
		PostgresInsertFromInsert(params))
}

func (s *Postgres) InsertReturns() bool {
	return true
}