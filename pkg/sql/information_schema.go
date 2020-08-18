// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package sql

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/security"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/catalogkv"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/catconstants"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/resolver"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgcode"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/cockroachdb/cockroach/pkg/sql/privilege"
	"github.com/cockroachdb/cockroach/pkg/sql/schemaexpr"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sessiondata"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/sql/vtable"
	"github.com/cockroachdb/errors"
)

const (
	pgCatalogName = sessiondata.PgCatalogName
)

var pgCatalogNameDString = tree.NewDString(pgCatalogName)

// informationSchema lists all the table definitions for
// information_schema.
var informationSchema = virtualSchema{
	name: sessiondata.InformationSchemaName,
	allTableNames: buildStringSet(
		// Generated with:
		// select distinct '"'||table_name||'",' from information_schema.tables
		//    where table_schema='information_schema' order by table_name;
		"_pg_foreign_data_wrappers",
		"_pg_foreign_servers",
		"_pg_foreign_table_columns",
		"_pg_foreign_tables",
		"_pg_user_mappings",
		"administrable_role_authorizations",
		"applicable_roles",
		"attributes",
		"character_sets",
		"check_constraint_routine_usage",
		"check_constraints",
		"collation_character_set_applicability",
		"collations",
		"column_domain_usage",
		"column_options",
		"column_privileges",
		"column_udt_usage",
		"columns",
		"constraint_column_usage",
		"constraint_table_usage",
		"data_type_privileges",
		"domain_constraints",
		"domain_udt_usage",
		"domains",
		"element_types",
		"enabled_roles",
		"foreign_data_wrapper_options",
		"foreign_data_wrappers",
		"foreign_server_options",
		"foreign_servers",
		"foreign_table_options",
		"foreign_tables",
		"information_schema_catalog_name",
		"key_column_usage",
		"parameters",
		"referential_constraints",
		"role_column_grants",
		"role_routine_grants",
		"role_table_grants",
		"role_udt_grants",
		"role_usage_grants",
		"routine_privileges",
		"routines",
		"schemata",
		"sequences",
		"sql_features",
		"sql_implementation_info",
		"sql_languages",
		"sql_packages",
		"sql_parts",
		"sql_sizing",
		"sql_sizing_profiles",
		"table_constraints",
		"table_privileges",
		"tables",
		"transforms",
		"triggered_update_columns",
		"triggers",
		"udt_privileges",
		"usage_privileges",
		"user_defined_types",
		"user_mapping_options",
		"user_mappings",
		"view_column_usage",
		"view_routine_usage",
		"view_table_usage",
		"views",
	),
	tableDefs: map[descpb.ID]virtualSchemaDef{
		catconstants.InformationSchemaAdministrableRoleAuthorizationsID: informationSchemaAdministrableRoleAuthorizations,
		catconstants.InformationSchemaApplicableRolesID:                 informationSchemaApplicableRoles,
		catconstants.InformationSchemaCheckConstraints:                  informationSchemaCheckConstraints,
		catconstants.InformationSchemaColumnPrivilegesID:                informationSchemaColumnPrivileges,
		catconstants.InformationSchemaColumnsTableID:                    informationSchemaColumnsTable,
		catconstants.InformationSchemaConstraintColumnUsageTableID:      informationSchemaConstraintColumnUsageTable,
		catconstants.InformationSchemaEnabledRolesID:                    informationSchemaEnabledRoles,
		catconstants.InformationSchemaKeyColumnUsageTableID:             informationSchemaKeyColumnUsageTable,
		catconstants.InformationSchemaParametersTableID:                 informationSchemaParametersTable,
		catconstants.InformationSchemaReferentialConstraintsTableID:     informationSchemaReferentialConstraintsTable,
		catconstants.InformationSchemaRoleTableGrantsID:                 informationSchemaRoleTableGrants,
		catconstants.InformationSchemaRoutineTableID:                    informationSchemaRoutineTable,
		catconstants.InformationSchemaSchemataTableID:                   informationSchemaSchemataTable,
		catconstants.InformationSchemaSchemataTablePrivilegesID:         informationSchemaSchemataTablePrivileges,
		catconstants.InformationSchemaSequencesID:                       informationSchemaSequences,
		catconstants.InformationSchemaStatisticsTableID:                 informationSchemaStatisticsTable,
		catconstants.InformationSchemaTableConstraintTableID:            informationSchemaTableConstraintTable,
		catconstants.InformationSchemaTablePrivilegesID:                 informationSchemaTablePrivileges,
		catconstants.InformationSchemaTablesTableID:                     informationSchemaTablesTable,
		catconstants.InformationSchemaViewsTableID:                      informationSchemaViewsTable,
		catconstants.InformationSchemaUserPrivilegesID:                  informationSchemaUserPrivileges,
	},
	tableValidator:             validateInformationSchemaTable,
	validWithNoDatabaseContext: true,
}

func buildStringSet(ss ...string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, s := range ss {
		m[s] = struct{}{}
	}
	return m
}

var (
	emptyString = tree.NewDString("")
	// information_schema was defined before the BOOLEAN data type was added to
	// the SQL specification. Because of this, boolean values are represented as
	// STRINGs. The BOOLEAN data type should NEVER be used in information_schema
	// tables. Instead, define columns as STRINGs and map bools to STRINGs using
	// yesOrNoDatum.
	yesString = tree.NewDString("YES")
	noString  = tree.NewDString("NO")
)

func yesOrNoDatum(b bool) tree.Datum {
	if b {
		return yesString
	}
	return noString
}

func dNameOrNull(s string) tree.Datum {
	if s == "" {
		return tree.DNull
	}
	return tree.NewDName(s)
}

func dIntFnOrNull(fn func() (int32, bool)) tree.Datum {
	if n, ok := fn(); ok {
		return tree.NewDInt(tree.DInt(n))
	}
	return tree.DNull
}

func validateInformationSchemaTable(table *descpb.TableDescriptor) error {
	// Make sure no tables have boolean columns.
	for i := range table.Columns {
		if table.Columns[i].Type.Family() == types.BoolFamily {
			return errors.Errorf("information_schema tables should never use BOOL columns. "+
				"See the comment about yesOrNoDatum. Found BOOL column in %s.", table.Name)
		}
	}
	return nil
}

var informationSchemaAdministrableRoleAuthorizations = virtualSchemaTable{
	comment: `roles for which the current user has admin option
` + base.DocsURL("information-schema.html#administrable_role_authorizations") + `
https://www.postgresql.org/docs/9.5/infoschema-administrable-role-authorizations.html`,
	schema: vtable.InformationSchemaAdministrableRoleAuthorizations,
	populate: func(ctx context.Context, p *planner, _ *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		currentUser := p.SessionData().User
		memberMap, err := p.MemberOfWithAdminOption(ctx, currentUser)
		if err != nil {
			return err
		}

		grantee := tree.NewDString(currentUser)
		for roleName, isAdmin := range memberMap {
			if !isAdmin {
				// We only show memberships with the admin option.
				continue
			}

			if err := addRow(
				grantee,                   // grantee: always the current user
				tree.NewDString(roleName), // role_name
				yesString,                 // is_grantable: always YES
			); err != nil {
				return err
			}
		}

		return nil
	},
}

var informationSchemaApplicableRoles = virtualSchemaTable{
	comment: `roles available to the current user
` + base.DocsURL("information-schema.html#applicable_roles") + `
https://www.postgresql.org/docs/9.5/infoschema-applicable-roles.html`,
	schema: vtable.InformationSchemaApplicableRoles,
	populate: func(ctx context.Context, p *planner, _ *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		currentUser := p.SessionData().User
		memberMap, err := p.MemberOfWithAdminOption(ctx, currentUser)
		if err != nil {
			return err
		}

		grantee := tree.NewDString(currentUser)

		for roleName, isAdmin := range memberMap {
			if err := addRow(
				grantee,                   // grantee: always the current user
				tree.NewDString(roleName), // role_name
				yesOrNoDatum(isAdmin),     // is_grantable
			); err != nil {
				return err
			}
		}

		return nil
	},
}

var informationSchemaCheckConstraints = virtualSchemaTable{
	comment: `check constraints
` + base.DocsURL("information-schema.html#check_constraints") + `
https://www.postgresql.org/docs/9.5/infoschema-check-constraints.html`,
	schema: vtable.InformationSchemaCheckConstraints,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		h := makeOidHasher()
		return forEachTableDescWithTableLookup(ctx, p, dbContext, hideVirtual /* no constraints in virtual tables */, func(
			db *sqlbase.ImmutableDatabaseDescriptor,
			scName string,
			table *sqlbase.ImmutableTableDescriptor,
			tableLookup tableLookupFn,
		) error {
			conInfo, err := table.GetConstraintInfoWithLookup(tableLookup.getTableByID)
			if err != nil {
				return err
			}
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			for conName, con := range conInfo {
				// Only Check constraints are included.
				if con.Kind != descpb.ConstraintTypeCheck {
					continue
				}
				conNameStr := tree.NewDString(conName)
				// Like with pg_catalog.pg_constraint, Postgres wraps the check
				// constraint expression in two pairs of parentheses.
				chkExprStr := tree.NewDString(fmt.Sprintf("((%s))", con.Details))
				if err := addRow(
					dbNameStr,  // constraint_catalog
					scNameStr,  // constraint_schema
					conNameStr, // constraint_name
					chkExprStr, // check_clause
				); err != nil {
					return err
				}
			}

			// Unlike with pg_catalog.pg_constraint, Postgres also includes NOT
			// NULL column constraints in information_schema.check_constraints.
			// Cockroach doesn't track these constraints as check constraints,
			// but we can pull them off of the table's column descriptors.
			colNum := 0
			return forEachColumnInTable(table, func(column *descpb.ColumnDescriptor) error {
				colNum++
				// Only visible, non-nullable columns are included.
				if column.Hidden || column.Nullable {
					return nil
				}
				// Generate a unique name for each NOT NULL constraint. Postgres
				// uses the format <namespace_oid>_<table_oid>_<col_idx>_not_null.
				// We might as well do the same.
				conNameStr := tree.NewDString(fmt.Sprintf(
					"%s_%s_%d_not_null", h.NamespaceOid(db, scName), tableOid(table.ID), colNum,
				))
				chkExprStr := tree.NewDString(fmt.Sprintf(
					"%s IS NOT NULL", column.Name,
				))
				return addRow(
					dbNameStr,  // constraint_catalog
					scNameStr,  // constraint_schema
					conNameStr, // constraint_name
					chkExprStr, // check_clause
				)
			})
		})
	},
}

var informationSchemaColumnPrivileges = virtualSchemaTable{
	comment: `column privilege grants (incomplete)
` + base.DocsURL("information-schema.html#column_privileges") + `
https://www.postgresql.org/docs/9.5/infoschema-column-privileges.html`,
	schema: vtable.InformationSchemaColumnPrivileges,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, virtualMany, func(
			db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor,
		) error {
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			columndata := privilege.List{privilege.SELECT, privilege.INSERT, privilege.UPDATE} // privileges for column level granularity
			for _, u := range table.Privileges.Users {
				for _, priv := range columndata {
					if priv.Mask()&u.Privileges != 0 {
						for i := range table.Columns {
							cd := &table.Columns[i]
							if err := addRow(
								tree.DNull,                     // grantor
								tree.NewDString(u.User),        // grantee
								dbNameStr,                      // table_catalog
								scNameStr,                      // table_schema
								tree.NewDString(table.Name),    // table_name
								tree.NewDString(cd.Name),       // column_name
								tree.NewDString(priv.String()), // privilege_type
								tree.DNull,                     // is_grantable
							); err != nil {
								return err
							}
						}
					}
				}
			}
			return nil
		})
	},
}

var informationSchemaColumnsTable = virtualSchemaTable{
	comment: `table and view columns (incomplete)
` + base.DocsURL("information-schema.html#columns") + `
https://www.postgresql.org/docs/9.5/infoschema-columns.html`,
	schema: vtable.InformationSchemaColumns,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, virtualMany, func(
			db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor,
		) error {
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			return forEachColumnInTable(table, func(column *descpb.ColumnDescriptor) error {
				collationCatalog := tree.DNull
				collationSchema := tree.DNull
				collationName := tree.DNull
				if locale := column.Type.Locale(); locale != "" {
					collationCatalog = dbNameStr
					collationSchema = pgCatalogNameDString
					collationName = tree.NewDString(locale)
				}
				colDefault := tree.DNull
				if column.DefaultExpr != nil {
					colExpr, err := schemaexpr.FormatExprForDisplay(ctx, table, *column.DefaultExpr, &p.semaCtx)
					if err != nil {
						return err
					}
					colDefault = tree.NewDString(colExpr)
				}
				colComputed := emptyString
				if column.ComputeExpr != nil {
					colExpr, err := schemaexpr.FormatExprForDisplayWithoutTypeAnnotations(ctx, table, *column.ComputeExpr, &p.semaCtx)
					if err != nil {
						return err
					}
					colComputed = tree.NewDString(colExpr)
				}
				return addRow(
					dbNameStr,                    // table_catalog
					scNameStr,                    // table_schema
					tree.NewDString(table.Name),  // table_name
					tree.NewDString(column.Name), // column_name
					tree.NewDInt(tree.DInt(column.GetPGAttributeNum())), // ordinal_position
					colDefault,                    // column_default
					yesOrNoDatum(column.Nullable), // is_nullable
					tree.NewDString(column.Type.InformationSchemaName()), // data_type
					characterMaximumLength(column.Type),                  // character_maximum_length
					characterOctetLength(column.Type),                    // character_octet_length
					numericPrecision(column.Type),                        // numeric_precision
					numericPrecisionRadix(column.Type),                   // numeric_precision_radix
					numericScale(column.Type),                            // numeric_scale
					datetimePrecision(column.Type),                       // datetime_precision
					tree.DNull,                                           // interval_type
					tree.DNull,                                           // interval_precision
					tree.DNull,                                           // character_set_catalog
					tree.DNull,                                           // character_set_schema
					tree.DNull,                                           // character_set_name
					collationCatalog,                                     // collation_catalog
					collationSchema,                                      // collation_schema
					collationName,                                        // collation_name
					tree.DNull,                                           // domain_catalog
					tree.DNull,                                           // domain_schema
					tree.DNull,                                           // domain_name
					dbNameStr,                                            // udt_catalog
					pgCatalogNameDString,                                 // udt_schema
					tree.NewDString(column.Type.PGName()),                // udt_name
					tree.DNull,                                           // scope_catalog
					tree.DNull,                                           // scope_schema
					tree.DNull,                                           // scope_name
					tree.DNull,                                           // maximum_cardinality
					tree.DNull,                                           // dtd_identifier
					tree.DNull,                                           // is_self_referencing
					tree.DNull,                                           // is_identity
					tree.DNull,                                           // identity_generation
					tree.DNull,                                           // identity_start
					tree.DNull,                                           // identity_increment
					tree.DNull,                                           // identity_maximum
					tree.DNull,                                           // identity_minimum
					tree.DNull,                                           // identity_cycle
					yesOrNoDatum(column.IsComputed()),                    // is_generated
					colComputed,                                          // generation_expression
					yesOrNoDatum(table.IsTable() &&
						!table.IsVirtualTable() &&
						!column.IsComputed(),
					), // is_updatable
					yesOrNoDatum(column.Hidden),              // is_hidden
					tree.NewDString(column.Type.SQLString()), // crdb_sql_type
				)
			})
		})
	},
}

var informationSchemaEnabledRoles = virtualSchemaTable{
	comment: `roles for the current user
` + base.DocsURL("information-schema.html#enabled_roles") + `
https://www.postgresql.org/docs/9.5/infoschema-enabled-roles.html`,
	schema: `
CREATE TABLE information_schema.enabled_roles (
	ROLE_NAME STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, _ *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		currentUser := p.SessionData().User
		memberMap, err := p.MemberOfWithAdminOption(ctx, currentUser)
		if err != nil {
			return err
		}

		// The current user is always listed.
		if err := addRow(
			tree.NewDString(currentUser), // role_name: the current user
		); err != nil {
			return err
		}

		for roleName := range memberMap {
			if err := addRow(
				tree.NewDString(roleName), // role_name
			); err != nil {
				return err
			}
		}

		return nil
	},
}

// characterMaximumLength returns the declared maximum length of
// characters if the type is a character or bit string data
// type. Returns false if the data type is not a character or bit
// string, or if the string's length is not bounded.
func characterMaximumLength(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.StringFamily, types.CollatedStringFamily, types.BitFamily:
			if colType.Width() > 0 {
				return colType.Width(), true
			}
		}
		return 0, false
	})
}

// characterOctetLength returns the maximum possible length in
// octets of a datum if the T is a character string. Returns
// false if the data type is not a character string, or if the
// string's length is not bounded.
func characterOctetLength(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.StringFamily, types.CollatedStringFamily:
			if colType.Width() > 0 {
				return colType.Width() * utf8.UTFMax, true
			}
		}
		return 0, false
	})
}

// numericPrecision returns the declared or implicit precision of numeric
// data types. Returns false if the data type is not numeric, or if the precision
// of the numeric type is not bounded.
func numericPrecision(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.IntFamily:
			return colType.Width(), true
		case types.FloatFamily:
			if colType.Width() == 32 {
				return 24, true
			}
			return 53, true
		case types.DecimalFamily:
			if colType.Precision() > 0 {
				return colType.Precision(), true
			}
		}
		return 0, false
	})
}

// numericPrecisionRadix returns the implicit precision radix of
// numeric data types. Returns false if the data type is not numeric.
func numericPrecisionRadix(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.IntFamily:
			return 2, true
		case types.FloatFamily:
			return 2, true
		case types.DecimalFamily:
			return 10, true
		}
		return 0, false
	})
}

// NumericScale returns the declared or implicit precision of exact numeric
// data types. Returns false if the data type is not an exact numeric, or if the
// scale of the exact numeric type is not bounded.
func numericScale(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.IntFamily:
			return 0, true
		case types.DecimalFamily:
			if colType.Precision() > 0 {
				return colType.Width(), true
			}
		}
		return 0, false
	})
}

// datetimePrecision returns the declared or implicit precision of Time,
// Timestamp or Interval data types. Returns false if the data type is not
// a Time, Timestamp or Interval.
func datetimePrecision(colType *types.T) tree.Datum {
	return dIntFnOrNull(func() (int32, bool) {
		switch colType.Family() {
		case types.TimeFamily, types.TimeTZFamily, types.TimestampFamily, types.TimestampTZFamily, types.IntervalFamily:
			return colType.Precision(), true
		}
		return 0, false
	})
}

var informationSchemaConstraintColumnUsageTable = virtualSchemaTable{
	comment: `columns usage by constraints
https://www.postgresql.org/docs/9.5/infoschema-constraint-column-usage.html`,
	schema: `
CREATE TABLE information_schema.constraint_column_usage (
	TABLE_CATALOG      STRING NOT NULL,
	TABLE_SCHEMA       STRING NOT NULL,
	TABLE_NAME         STRING NOT NULL,
	COLUMN_NAME        STRING NOT NULL,
	CONSTRAINT_CATALOG STRING NOT NULL,
	CONSTRAINT_SCHEMA  STRING NOT NULL,
	CONSTRAINT_NAME    STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDescWithTableLookup(ctx, p, dbContext, hideVirtual /* no constraints in virtual tables */, func(
			db *sqlbase.ImmutableDatabaseDescriptor,
			scName string,
			table *sqlbase.ImmutableTableDescriptor,
			tableLookup tableLookupFn,
		) error {
			conInfo, err := table.GetConstraintInfoWithLookup(tableLookup.getTableByID)
			if err != nil {
				return err
			}
			scNameStr := tree.NewDString(scName)
			dbNameStr := tree.NewDString(db.GetName())

			for conName, con := range conInfo {
				conTable := table
				conCols := con.Columns
				conNameStr := tree.NewDString(conName)
				if con.Kind == descpb.ConstraintTypeFK {
					// For foreign key constraint, constraint_column_usage
					// identifies the table/columns that the foreign key
					// references.
					conTable = sqlbase.NewImmutableTableDescriptor(*con.ReferencedTable)
					conCols, err = conTable.NamesForColumnIDs(con.FK.ReferencedColumnIDs)
					if err != nil {
						return err
					}
				}
				tableNameStr := tree.NewDString(conTable.Name)
				for _, col := range conCols {
					if err := addRow(
						dbNameStr,            // table_catalog
						scNameStr,            // table_schema
						tableNameStr,         // table_name
						tree.NewDString(col), // column_name
						dbNameStr,            // constraint_catalog
						scNameStr,            // constraint_schema
						conNameStr,           // constraint_name
					); err != nil {
						return err
					}
				}
			}
			return nil
		})
	},
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/key-column-usage-table.html
var informationSchemaKeyColumnUsageTable = virtualSchemaTable{
	comment: `column usage by indexes and key constraints
` + base.DocsURL("information-schema.html#key_column_usage") + `
https://www.postgresql.org/docs/9.5/infoschema-key-column-usage.html`,
	schema: `
CREATE TABLE information_schema.key_column_usage (
	CONSTRAINT_CATALOG STRING NOT NULL,
	CONSTRAINT_SCHEMA  STRING NOT NULL,
	CONSTRAINT_NAME    STRING NOT NULL,
	TABLE_CATALOG      STRING NOT NULL,
	TABLE_SCHEMA       STRING NOT NULL,
	TABLE_NAME         STRING NOT NULL,
	COLUMN_NAME        STRING NOT NULL,
	ORDINAL_POSITION   INT NOT NULL,
	POSITION_IN_UNIQUE_CONSTRAINT INT
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDescWithTableLookup(ctx, p, dbContext, hideVirtual /* no constraints in virtual tables */, func(
			db *sqlbase.ImmutableDatabaseDescriptor,
			scName string,
			table *sqlbase.ImmutableTableDescriptor,
			tableLookup tableLookupFn,
		) error {
			conInfo, err := table.GetConstraintInfoWithLookup(tableLookup.getTableByID)
			if err != nil {
				return err
			}
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			tbNameStr := tree.NewDString(table.Name)
			for conName, con := range conInfo {
				// Only Primary Key, Foreign Key, and Unique constraints are included.
				switch con.Kind {
				case descpb.ConstraintTypePK:
				case descpb.ConstraintTypeFK:
				case descpb.ConstraintTypeUnique:
				default:
					continue
				}

				cstNameStr := tree.NewDString(conName)

				for pos, col := range con.Columns {
					ordinalPos := tree.NewDInt(tree.DInt(pos + 1))
					uniquePos := tree.DNull
					if con.Kind == descpb.ConstraintTypeFK {
						uniquePos = ordinalPos
					}
					if err := addRow(
						dbNameStr,            // constraint_catalog
						scNameStr,            // constraint_schema
						cstNameStr,           // constraint_name
						dbNameStr,            // table_catalog
						scNameStr,            // table_schema
						tbNameStr,            // table_name
						tree.NewDString(col), // column_name
						ordinalPos,           // ordinal_position, 1-indexed
						uniquePos,            // position_in_unique_constraint
					); err != nil {
						return err
					}
				}
			}
			return nil
		})
	},
}

// Postgres: https://www.postgresql.org/docs/9.6/static/infoschema-parameters.html
// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/parameters-table.html
var informationSchemaParametersTable = virtualSchemaTable{
	comment: `built-in function parameters (empty - introspection not yet supported)
https://www.postgresql.org/docs/9.5/infoschema-parameters.html`,
	schema: `
CREATE TABLE information_schema.parameters (
	SPECIFIC_CATALOG STRING,
	SPECIFIC_SCHEMA STRING,
	SPECIFIC_NAME STRING,
	ORDINAL_POSITION INT,
	PARAMETER_MODE STRING,
	IS_RESULT STRING,
	AS_LOCATOR STRING,
	PARAMETER_NAME STRING,
	DATA_TYPE STRING,
	CHARACTER_MAXIMUM_LENGTH INT,
	CHARACTER_OCTET_LENGTH INT,
	CHARACTER_SET_CATALOG STRING,
	CHARACTER_SET_SCHEMA STRING,
	CHARACTER_SET_NAME STRING,
	COLLATION_CATALOG STRING,
	COLLATION_SCHEMA STRING,
	COLLATION_NAME STRING,
	NUMERIC_PRECISION INT,
	NUMERIC_PRECISION_RADIX INT,
	NUMERIC_SCALE INT,
	DATETIME_PRECISION INT,
	INTERVAL_TYPE STRING,
	INTERVAL_PRECISION INT,
	UDT_CATALOG STRING,
	UDT_SCHEMA STRING,
	UDT_NAME STRING,
	SCOPE_CATALOG STRING,
	SCOPE_SCHEMA STRING,
	SCOPE_NAME STRING,
	MAXIMUM_CARDINALITY INT,
	DTD_IDENTIFIER STRING,
	PARAMETER_DEFAULT STRING
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return nil
	},
}

var (
	matchOptionFull    = tree.NewDString("FULL")
	matchOptionPartial = tree.NewDString("PARTIAL")
	matchOptionNone    = tree.NewDString("NONE")

	matchOptionMap = map[descpb.ForeignKeyReference_Match]tree.Datum{
		descpb.ForeignKeyReference_SIMPLE:  matchOptionNone,
		descpb.ForeignKeyReference_FULL:    matchOptionFull,
		descpb.ForeignKeyReference_PARTIAL: matchOptionPartial,
	}

	refConstraintRuleNoAction   = tree.NewDString("NO ACTION")
	refConstraintRuleRestrict   = tree.NewDString("RESTRICT")
	refConstraintRuleSetNull    = tree.NewDString("SET NULL")
	refConstraintRuleSetDefault = tree.NewDString("SET DEFAULT")
	refConstraintRuleCascade    = tree.NewDString("CASCADE")
)

func dStringForFKAction(action descpb.ForeignKeyReference_Action) tree.Datum {
	switch action {
	case descpb.ForeignKeyReference_NO_ACTION:
		return refConstraintRuleNoAction
	case descpb.ForeignKeyReference_RESTRICT:
		return refConstraintRuleRestrict
	case descpb.ForeignKeyReference_SET_NULL:
		return refConstraintRuleSetNull
	case descpb.ForeignKeyReference_SET_DEFAULT:
		return refConstraintRuleSetDefault
	case descpb.ForeignKeyReference_CASCADE:
		return refConstraintRuleCascade
	}
	panic(errors.Errorf("unexpected ForeignKeyReference_Action: %v", action))
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/referential-constraints-table.html
var informationSchemaReferentialConstraintsTable = virtualSchemaTable{
	comment: `foreign key constraints
` + base.DocsURL("information-schema.html#referential_constraints") + `
https://www.postgresql.org/docs/9.5/infoschema-referential-constraints.html`,
	schema: `
CREATE TABLE information_schema.referential_constraints (
	CONSTRAINT_CATALOG        STRING NOT NULL,
	CONSTRAINT_SCHEMA         STRING NOT NULL,
	CONSTRAINT_NAME           STRING NOT NULL,
	UNIQUE_CONSTRAINT_CATALOG STRING NOT NULL,
	UNIQUE_CONSTRAINT_SCHEMA  STRING NOT NULL,
	UNIQUE_CONSTRAINT_NAME    STRING,
	MATCH_OPTION              STRING NOT NULL,
	UPDATE_RULE               STRING NOT NULL,
	DELETE_RULE               STRING NOT NULL,
	TABLE_NAME                STRING NOT NULL,
	REFERENCED_TABLE_NAME     STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDescWithTableLookup(ctx, p, dbContext, hideVirtual /* no constraints in virtual tables */, func(
			db *sqlbase.ImmutableDatabaseDescriptor,
			scName string,
			table *sqlbase.ImmutableTableDescriptor,
			tableLookup tableLookupFn,
		) error {
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			tbNameStr := tree.NewDString(table.Name)
			for i := range table.OutboundFKs {
				fk := &table.OutboundFKs[i]
				refTable, err := tableLookup.getTableByID(fk.ReferencedTableID)
				if err != nil {
					return err
				}
				var matchType = tree.DNull
				if r, ok := matchOptionMap[fk.Match]; ok {
					matchType = r
				}
				referencedIdx, err := sqlbase.FindFKReferencedIndex(refTable, fk.ReferencedColumnIDs)
				if err != nil {
					return err
				}
				if err := addRow(
					dbNameStr,                           // constraint_catalog
					scNameStr,                           // constraint_schema
					tree.NewDString(fk.Name),            // constraint_name
					dbNameStr,                           // unique_constraint_catalog
					scNameStr,                           // unique_constraint_schema
					tree.NewDString(referencedIdx.Name), // unique_constraint_name
					matchType,                           // match_option
					dStringForFKAction(fk.OnUpdate),     // update_rule
					dStringForFKAction(fk.OnDelete),     // delete_rule
					tbNameStr,                           // table_name
					tree.NewDString(refTable.GetName()), // referenced_table_name
				); err != nil {
					return err
				}
			}
			return nil
		})
	},
}

// Postgres: https://www.postgresql.org/docs/9.6/static/infoschema-role-table-grants.html
// MySQL:    missing
var informationSchemaRoleTableGrants = virtualSchemaTable{
	comment: `privileges granted on table or views (incomplete; see also information_schema.table_privileges; may contain excess users or roles)
` + base.DocsURL("information-schema.html#role_table_grants") + `
https://www.postgresql.org/docs/9.5/infoschema-role-table-grants.html`,
	schema: `
CREATE TABLE information_schema.role_table_grants (
	GRANTOR        STRING,
	GRANTEE        STRING NOT NULL,
	TABLE_CATALOG  STRING NOT NULL,
	TABLE_SCHEMA   STRING NOT NULL,
	TABLE_NAME     STRING NOT NULL,
	PRIVILEGE_TYPE STRING NOT NULL,
	IS_GRANTABLE   STRING,
	WITH_HIERARCHY STRING
)`,
	// This is the same as information_schema.table_privileges. In postgres, this virtual table does
	// not show tables with grants provided through PUBLIC, but table_privileges does.
	// Since we don't have the PUBLIC concept, the two virtual tables are identical.
	populate: populateTablePrivileges,
}

// MySQL:    https://dev.mysql.com/doc/mysql-infoschema-excerpt/5.7/en/routines-table.html
var informationSchemaRoutineTable = virtualSchemaTable{
	comment: `built-in functions (empty - introspection not yet supported)
https://www.postgresql.org/docs/9.5/infoschema-routines.html`,
	schema: `
CREATE TABLE information_schema.routines (
	SPECIFIC_CATALOG STRING,
	SPECIFIC_SCHEMA STRING,
	SPECIFIC_NAME STRING,
	ROUTINE_CATALOG STRING,
	ROUTINE_SCHEMA STRING,
	ROUTINE_NAME STRING,
	ROUTINE_TYPE STRING,
	MODULE_CATALOG STRING,
	MODULE_SCHEMA STRING,
	MODULE_NAME STRING,
	UDT_CATALOG STRING,
	UDT_SCHEMA STRING,
	UDT_NAME STRING,
	DATA_TYPE STRING,
	CHARACTER_MAXIMUM_LENGTH INT,
	CHARACTER_OCTET_LENGTH INT,
	CHARACTER_SET_CATALOG STRING,
	CHARACTER_SET_SCHEMA STRING,
	CHARACTER_SET_NAME STRING,
	COLLATION_CATALOG STRING,
	COLLATION_SCHEMA STRING,
	COLLATION_NAME STRING,
	NUMERIC_PRECISION INT,
	NUMERIC_PRECISION_RADIX INT,
	NUMERIC_SCALE INT,
	DATETIME_PRECISION INT,
	INTERVAL_TYPE STRING,
	INTERVAL_PRECISION STRING,
	TYPE_UDT_CATALOG STRING,
	TYPE_UDT_SCHEMA STRING,
	TYPE_UDT_NAME STRING,
	SCOPE_CATALOG STRING,
	SCOPE_NAME STRING,
	MAXIMUM_CARDINALITY INT,
	DTD_IDENTIFIER STRING,
	ROUTINE_BODY STRING,
	ROUTINE_DEFINITION STRING,
	EXTERNAL_NAME STRING,
	EXTERNAL_LANGUAGE STRING,
	PARAMETER_STYLE STRING,
	IS_DETERMINISTIC STRING,
	SQL_DATA_ACCESS STRING,
	IS_NULL_CALL STRING,
	SQL_PATH STRING,
	SCHEMA_LEVEL_ROUTINE STRING,
	MAX_DYNAMIC_RESULT_SETS INT,
	IS_USER_DEFINED_CAST STRING,
	IS_IMPLICITLY_INVOCABLE STRING,
	SECURITY_TYPE STRING,
	TO_SQL_SPECIFIC_CATALOG STRING,
	TO_SQL_SPECIFIC_SCHEMA STRING,
	TO_SQL_SPECIFIC_NAME STRING,
	AS_LOCATOR STRING,
	CREATED  TIMESTAMPTZ,
	LAST_ALTERED TIMESTAMPTZ,
	NEW_SAVEPOINT_LEVEL  STRING,
	IS_UDT_DEPENDENT STRING,
	RESULT_CAST_FROM_DATA_TYPE STRING,
	RESULT_CAST_AS_LOCATOR STRING,
	RESULT_CAST_CHAR_MAX_LENGTH  INT,
	RESULT_CAST_CHAR_OCTET_LENGTH STRING,
	RESULT_CAST_CHAR_SET_CATALOG STRING,
	RESULT_CAST_CHAR_SET_SCHEMA  STRING,
	RESULT_CAST_CHAR_SET_NAME STRING,
	RESULT_CAST_COLLATION_CATALOG STRING,
	RESULT_CAST_COLLATION_SCHEMA STRING,
	RESULT_CAST_COLLATION_NAME STRING,
	RESULT_CAST_NUMERIC_PRECISION INT,
	RESULT_CAST_NUMERIC_PRECISION_RADIX INT,
	RESULT_CAST_NUMERIC_SCALE INT,
	RESULT_CAST_DATETIME_PRECISION STRING,
	RESULT_CAST_INTERVAL_TYPE STRING,
	RESULT_CAST_INTERVAL_PRECISION INT,
	RESULT_CAST_TYPE_UDT_CATALOG STRING,
	RESULT_CAST_TYPE_UDT_SCHEMA  STRING,
	RESULT_CAST_TYPE_UDT_NAME STRING,
	RESULT_CAST_SCOPE_CATALOG STRING,
	RESULT_CAST_SCOPE_SCHEMA STRING,
	RESULT_CAST_SCOPE_NAME STRING,
	RESULT_CAST_MAXIMUM_CARDINALITY INT,
	RESULT_CAST_DTD_IDENTIFIER STRING
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return nil
	},
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/schemata-table.html
var informationSchemaSchemataTable = virtualSchemaTable{
	comment: `database schemas (may contain schemata without permission)
` + base.DocsURL("information-schema.html#schemata") + `
https://www.postgresql.org/docs/9.5/infoschema-schemata.html`,
	schema: vtable.InformationSchemaSchemata,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachDatabaseDesc(ctx, p, dbContext, true, /* requiresPrivileges */
			func(db *sqlbase.ImmutableDatabaseDescriptor) error {
				return forEachSchemaName(ctx, p, db, func(sc string, userDefined bool) error {
					return addRow(
						tree.NewDString(db.GetName()), // catalog_name
						tree.NewDString(sc),           // schema_name
						tree.DNull,                    // default_character_set_name
						tree.DNull,                    // sql_path
						yesOrNoDatum(userDefined),     // crdb_is_user_defined
					)
				})
			})
	},
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/schema-privileges-table.html
var informationSchemaSchemataTablePrivileges = virtualSchemaTable{
	comment: `schema privileges (incomplete; may contain excess users or roles)
` + base.DocsURL("information-schema.html#schema_privileges"),
	schema: `
CREATE TABLE information_schema.schema_privileges (
	GRANTEE         STRING NOT NULL,
	TABLE_CATALOG   STRING NOT NULL,
	TABLE_SCHEMA    STRING NOT NULL,
	PRIVILEGE_TYPE  STRING NOT NULL,
	IS_GRANTABLE    STRING
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachDatabaseDesc(ctx, p, dbContext, true, /* requiresPrivileges */
			func(db *sqlbase.ImmutableDatabaseDescriptor) error {
				return forEachSchemaName(ctx, p, db, func(scName string, _ bool) error {
					privs := db.Privileges.Show(privilege.Schema)
					dbNameStr := tree.NewDString(db.GetName())
					scNameStr := tree.NewDString(scName)
					// TODO(knz): This should filter for the current user, see
					// https://github.com/cockroachdb/cockroach/issues/35572
					for _, u := range privs {
						userNameStr := tree.NewDString(u.User)
						for _, priv := range u.Privileges {
							if err := addRow(
								userNameStr,           // grantee
								dbNameStr,             // table_catalog
								scNameStr,             // table_schema
								tree.NewDString(priv), // privilege_type
								tree.DNull,            // is_grantable
							); err != nil {
								return err
							}
						}
					}
					return nil
				})
			})
	},
}

var (
	indexDirectionNA   = tree.NewDString("N/A")
	indexDirectionAsc  = tree.NewDString(descpb.IndexDescriptor_ASC.String())
	indexDirectionDesc = tree.NewDString(descpb.IndexDescriptor_DESC.String())
)

func dStringForIndexDirection(dir descpb.IndexDescriptor_Direction) tree.Datum {
	switch dir {
	case descpb.IndexDescriptor_ASC:
		return indexDirectionAsc
	case descpb.IndexDescriptor_DESC:
		return indexDirectionDesc
	}
	panic("unreachable")
}

var informationSchemaSequences = virtualSchemaTable{
	comment: `sequences
` + base.DocsURL("information-schema.html#sequences") + `
https://www.postgresql.org/docs/9.5/infoschema-sequences.html`,
	schema: `
CREATE TABLE information_schema.sequences (
    SEQUENCE_CATALOG         STRING NOT NULL,
    SEQUENCE_SCHEMA          STRING NOT NULL,
    SEQUENCE_NAME            STRING NOT NULL,
    DATA_TYPE                STRING NOT NULL,
    NUMERIC_PRECISION        INT NOT NULL,
    NUMERIC_PRECISION_RADIX  INT NOT NULL,
    NUMERIC_SCALE            INT NOT NULL,
    START_VALUE              STRING NOT NULL,
    MINIMUM_VALUE            STRING NOT NULL,
    MAXIMUM_VALUE            STRING NOT NULL,
    INCREMENT                STRING NOT NULL,
    CYCLE_OPTION             STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, hideVirtual, /* no sequences in virtual schemas */
			func(db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor) error {
				if !table.IsSequence() {
					return nil
				}
				return addRow(
					tree.NewDString(db.GetName()),    // catalog
					tree.NewDString(scName),          // schema
					tree.NewDString(table.GetName()), // name
					tree.NewDString("bigint"),        // type
					tree.NewDInt(64),                 // numeric precision
					tree.NewDInt(2),                  // numeric precision radix
					tree.NewDInt(0),                  // numeric scale
					tree.NewDString(strconv.FormatInt(table.SequenceOpts.Start, 10)),     // start value
					tree.NewDString(strconv.FormatInt(table.SequenceOpts.MinValue, 10)),  // min value
					tree.NewDString(strconv.FormatInt(table.SequenceOpts.MaxValue, 10)),  // max value
					tree.NewDString(strconv.FormatInt(table.SequenceOpts.Increment, 10)), // increment
					noString, // cycle
				)
			})
	},
}

// Postgres: missing
// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/statistics-table.html
var informationSchemaStatisticsTable = virtualSchemaTable{
	comment: `index metadata and statistics (incomplete)
` + base.DocsURL("information-schema.html#statistics"),
	schema: `
CREATE TABLE information_schema.statistics (
	TABLE_CATALOG STRING NOT NULL,
	TABLE_SCHEMA  STRING NOT NULL,
	TABLE_NAME    STRING NOT NULL,
	NON_UNIQUE    STRING NOT NULL,
	INDEX_SCHEMA  STRING NOT NULL,
	INDEX_NAME    STRING NOT NULL,
	SEQ_IN_INDEX  INT NOT NULL,
	COLUMN_NAME   STRING NOT NULL,
	"COLLATION"   STRING,
	CARDINALITY   INT,
	DIRECTION     STRING NOT NULL,
	STORING       STRING NOT NULL,
	IMPLICIT      STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, hideVirtual, /* virtual tables have no indexes */
			func(db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor) error {
				dbNameStr := tree.NewDString(db.GetName())
				scNameStr := tree.NewDString(scName)
				tbNameStr := tree.NewDString(table.GetName())

				appendRow := func(index *descpb.IndexDescriptor, colName string, sequence int,
					direction tree.Datum, isStored, isImplicit bool,
				) error {
					return addRow(
						dbNameStr,                         // table_catalog
						scNameStr,                         // table_schema
						tbNameStr,                         // table_name
						yesOrNoDatum(!index.Unique),       // non_unique
						scNameStr,                         // index_schema
						tree.NewDString(index.Name),       // index_name
						tree.NewDInt(tree.DInt(sequence)), // seq_in_index
						tree.NewDString(colName),          // column_name
						tree.DNull,                        // collation
						tree.DNull,                        // cardinality
						direction,                         // direction
						yesOrNoDatum(isStored),            // storing
						yesOrNoDatum(isImplicit),          // implicit
					)
				}

				return forEachIndexInTable(table, func(index *descpb.IndexDescriptor) error {
					// Columns in the primary key that aren't in index.ColumnNames or
					// index.StoreColumnNames are implicit columns in the index.
					var implicitCols map[string]struct{}
					var hasImplicitCols bool
					if index.HasOldStoredColumns() {
						// Old STORING format: implicit columns are extra columns minus stored
						// columns.
						hasImplicitCols = len(index.ExtraColumnIDs) > len(index.StoreColumnNames)
					} else {
						// New STORING format: implicit columns are extra columns.
						hasImplicitCols = len(index.ExtraColumnIDs) > 0
					}
					if hasImplicitCols {
						implicitCols = make(map[string]struct{})
						for _, col := range table.PrimaryIndex.ColumnNames {
							implicitCols[col] = struct{}{}
						}
					}

					sequence := 1
					for i, col := range index.ColumnNames {
						// We add a row for each column of index.
						dir := dStringForIndexDirection(index.ColumnDirections[i])
						if err := appendRow(index, col, sequence, dir, false, false); err != nil {
							return err
						}
						sequence++
						delete(implicitCols, col)
					}
					for _, col := range index.StoreColumnNames {
						// We add a row for each stored column of index.
						if err := appendRow(index, col, sequence,
							indexDirectionNA, true, false); err != nil {
							return err
						}
						sequence++
						delete(implicitCols, col)
					}
					for col := range implicitCols {
						// We add a row for each implicit column of index.
						if err := appendRow(index, col, sequence,
							indexDirectionAsc, false, true); err != nil {
							return err
						}
						sequence++
					}
					return nil
				})
			})
	},
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/table-constraints-table.html
var informationSchemaTableConstraintTable = virtualSchemaTable{
	comment: `table constraints
` + base.DocsURL("information-schema.html#table_constraints") + `
https://www.postgresql.org/docs/9.5/infoschema-table-constraints.html`,
	schema: `
CREATE TABLE information_schema.table_constraints (
	CONSTRAINT_CATALOG STRING NOT NULL,
	CONSTRAINT_SCHEMA  STRING NOT NULL,
	CONSTRAINT_NAME    STRING NOT NULL,
	TABLE_CATALOG      STRING NOT NULL,
	TABLE_SCHEMA       STRING NOT NULL,
	TABLE_NAME         STRING NOT NULL,
	CONSTRAINT_TYPE    STRING NOT NULL,
	IS_DEFERRABLE      STRING NOT NULL,
	INITIALLY_DEFERRED STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		h := makeOidHasher()
		return forEachTableDescWithTableLookup(ctx, p, dbContext, hideVirtual, /* virtual tables have no constraints */
			func(
				db *sqlbase.ImmutableDatabaseDescriptor,
				scName string,
				table *sqlbase.ImmutableTableDescriptor,
				tableLookup tableLookupFn,
			) error {
				conInfo, err := table.GetConstraintInfoWithLookup(tableLookup.getTableByID)
				if err != nil {
					return err
				}

				dbNameStr := tree.NewDString(db.GetName())
				scNameStr := tree.NewDString(scName)
				tbNameStr := tree.NewDString(table.Name)

				for conName, c := range conInfo {
					if err := addRow(
						dbNameStr,                       // constraint_catalog
						scNameStr,                       // constraint_schema
						tree.NewDString(conName),        // constraint_name
						dbNameStr,                       // table_catalog
						scNameStr,                       // table_schema
						tbNameStr,                       // table_name
						tree.NewDString(string(c.Kind)), // constraint_type
						yesOrNoDatum(false),             // is_deferrable
						yesOrNoDatum(false),             // initially_deferred
					); err != nil {
						return err
					}
				}

				// Unlike with pg_catalog.pg_constraint, Postgres also includes NOT
				// NULL column constraints in information_schema.check_constraints.
				// Cockroach doesn't track these constraints as check constraints,
				// but we can pull them off of the table's column descriptors.
				colNum := 0
				return forEachColumnInTable(table, func(col *descpb.ColumnDescriptor) error {
					colNum++
					// NOT NULL column constraints are implemented as a CHECK in postgres.
					conNameStr := tree.NewDString(fmt.Sprintf(
						"%s_%s_%d_not_null", h.NamespaceOid(db, scName), tableOid(table.ID), colNum,
					))
					if !col.Nullable {
						if err := addRow(
							dbNameStr,                // constraint_catalog
							scNameStr,                // constraint_schema
							conNameStr,               // constraint_name
							dbNameStr,                // table_catalog
							scNameStr,                // table_schema
							tbNameStr,                // table_name
							tree.NewDString("CHECK"), // constraint_type
							yesOrNoDatum(false),      // is_deferrable
							yesOrNoDatum(false),      // initially_deferred
						); err != nil {
							return err
						}
					}
					return nil
				})
			})
	},
}

// Postgres: not provided
// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/user-privileges-table.html
// TODO(knz): this introspection facility is of dubious utility.
var informationSchemaUserPrivileges = virtualSchemaTable{
	comment: `grantable privileges (incomplete)`,
	schema: `
CREATE TABLE information_schema.user_privileges (
	GRANTEE        STRING NOT NULL,
	TABLE_CATALOG  STRING NOT NULL,
	PRIVILEGE_TYPE STRING NOT NULL,
	IS_GRANTABLE   STRING
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachDatabaseDesc(ctx, p, dbContext, true, /* requiresPrivileges */
			func(dbDesc *sqlbase.ImmutableDatabaseDescriptor) error {
				dbNameStr := tree.NewDString(dbDesc.GetName())
				for _, u := range []string{security.RootUser, security.AdminRole} {
					grantee := tree.NewDString(u)
					for _, p := range privilege.DBSchemaTablePrivileges.SortedNames() {
						if err := addRow(
							grantee,            // grantee
							dbNameStr,          // table_catalog
							tree.NewDString(p), // privilege_type
							tree.DNull,         // is_grantable
						); err != nil {
							return err
						}
					}
				}
				return nil
			})
	},
}

// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/table-privileges-table.html
var informationSchemaTablePrivileges = virtualSchemaTable{
	comment: `privileges granted on table or views (incomplete; may contain excess users or roles)
` + base.DocsURL("information-schema.html#table_privileges") + `
https://www.postgresql.org/docs/9.5/infoschema-table-privileges.html`,
	schema: `
CREATE TABLE information_schema.table_privileges (
	GRANTOR        STRING,
	GRANTEE        STRING NOT NULL,
	TABLE_CATALOG  STRING NOT NULL,
	TABLE_SCHEMA   STRING NOT NULL,
	TABLE_NAME     STRING NOT NULL,
	PRIVILEGE_TYPE STRING NOT NULL,
	IS_GRANTABLE   STRING,
	WITH_HIERARCHY STRING NOT NULL
)`,
	populate: populateTablePrivileges,
}

// populateTablePrivileges is used to populate both table_privileges and role_table_grants.
func populateTablePrivileges(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	addRow func(...tree.Datum) error,
) error {
	return forEachTableDesc(ctx, p, dbContext, virtualMany,
		func(db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor) error {
			dbNameStr := tree.NewDString(db.GetName())
			scNameStr := tree.NewDString(scName)
			tbNameStr := tree.NewDString(table.Name)
			// TODO(knz): This should filter for the current user, see
			// https://github.com/cockroachdb/cockroach/issues/35572
			for _, u := range table.Privileges.Show(privilege.Table) {
				for _, priv := range u.Privileges {
					if err := addRow(
						tree.DNull,                     // grantor
						tree.NewDString(u.User),        // grantee
						dbNameStr,                      // table_catalog
						scNameStr,                      // table_schema
						tbNameStr,                      // table_name
						tree.NewDString(priv),          // privilege_type
						tree.DNull,                     // is_grantable
						yesOrNoDatum(priv == "SELECT"), // with_hierarchy
					); err != nil {
						return err
					}
				}
			}
			return nil
		})
}

var (
	tableTypeSystemView = tree.NewDString("SYSTEM VIEW")
	tableTypeBaseTable  = tree.NewDString("BASE TABLE")
	tableTypeView       = tree.NewDString("VIEW")
	tableTypeTemporary  = tree.NewDString("LOCAL TEMPORARY")
)

var informationSchemaTablesTable = virtualSchemaTable{
	comment: `tables and views
` + base.DocsURL("information-schema.html#tables") + `
https://www.postgresql.org/docs/9.5/infoschema-tables.html`,
	schema: vtable.InformationSchemaTables,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, virtualMany, addTablesTableRow(addRow))
	},
	indexes: []virtualIndex{
		{
			populate: func(ctx context.Context, constraint tree.Datum, p *planner, db *sqlbase.ImmutableDatabaseDescriptor,
				addRow func(...tree.Datum) error) (bool, error) {
				// This index is on the TABLE_NAME column.
				name := tree.MustBeDString(constraint)
				flags := tree.ObjectLookupFlags{}
				flags.DesiredTableDescKind = tree.ResolveAnyTableKind
				desc, err := resolver.ResolveExistingTableObject(ctx, p, tree.NewUnqualifiedTableName(tree.Name(name)), flags)
				if err != nil || desc == nil {
					return false, err
				}
				schemaName, err := resolver.ResolveSchemaNameByID(ctx, p.txn, p.ExecCfg().Codec, db.GetID(), desc.GetParentSchemaID())
				if err != nil {
					return false, err
				}
				return true, addTablesTableRow(addRow)(db, schemaName, desc)
			},
		},
	},
}

func addTablesTableRow(
	addRow func(...tree.Datum) error,
) func(db *sqlbase.ImmutableDatabaseDescriptor, scName string,
	table *sqlbase.ImmutableTableDescriptor) error {
	return func(db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor) error {
		if table.IsSequence() {
			return nil
		}
		tableType := tableTypeBaseTable
		insertable := yesString
		if table.IsVirtualTable() {
			tableType = tableTypeSystemView
			insertable = noString
		} else if table.IsView() {
			tableType = tableTypeView
			insertable = noString
		} else if table.Temporary {
			tableType = tableTypeTemporary
		}
		dbNameStr := tree.NewDString(db.GetName())
		scNameStr := tree.NewDString(scName)
		tbNameStr := tree.NewDString(table.Name)
		return addRow(
			dbNameStr,                              // table_catalog
			scNameStr,                              // table_schema
			tbNameStr,                              // table_name
			tableType,                              // table_type
			insertable,                             // is_insertable_into
			tree.NewDInt(tree.DInt(table.Version)), // version
		)
	}
}

// Postgres: https://www.postgresql.org/docs/9.6/static/infoschema-views.html
// MySQL:    https://dev.mysql.com/doc/refman/5.7/en/views-table.html
var informationSchemaViewsTable = virtualSchemaTable{
	comment: `views (incomplete)
` + base.DocsURL("information-schema.html#views") + `
https://www.postgresql.org/docs/9.5/infoschema-views.html`,
	schema: `
CREATE TABLE information_schema.views (
    TABLE_CATALOG              STRING NOT NULL,
    TABLE_SCHEMA               STRING NOT NULL,
    TABLE_NAME                 STRING NOT NULL,
    VIEW_DEFINITION            STRING NOT NULL,
    CHECK_OPTION               STRING,
    IS_UPDATABLE               STRING NOT NULL,
    IS_INSERTABLE_INTO         STRING NOT NULL,
    IS_TRIGGER_UPDATABLE       STRING NOT NULL,
    IS_TRIGGER_DELETABLE       STRING NOT NULL,
    IS_TRIGGER_INSERTABLE_INTO STRING NOT NULL
)`,
	populate: func(ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor, addRow func(...tree.Datum) error) error {
		return forEachTableDesc(ctx, p, dbContext, hideVirtual, /* virtual schemas have no views */
			func(db *sqlbase.ImmutableDatabaseDescriptor, scName string, table *sqlbase.ImmutableTableDescriptor) error {
				if !table.IsView() {
					return nil
				}
				// Note that the view query printed will not include any column aliases
				// specified outside the initial view query into the definition returned,
				// unlike Postgres. For example, for the view created via
				//  `CREATE VIEW (a) AS SELECT b FROM foo`
				// we'll only print `SELECT b FROM foo` as the view definition here,
				// while Postgres would more accurately print `SELECT b AS a FROM foo`.
				// TODO(a-robinson): Insert column aliases into view query once we
				// have a semantic query representation to work with (#10083).
				return addRow(
					tree.NewDString(db.GetName()),    // table_catalog
					tree.NewDString(scName),          // table_schema
					tree.NewDString(table.Name),      // table_name
					tree.NewDString(table.ViewQuery), // view_definition
					tree.DNull,                       // check_option
					noString,                         // is_updatable
					noString,                         // is_insertable_into
					noString,                         // is_trigger_updatable
					noString,                         // is_trigger_deletable
					noString,                         // is_trigger_insertable_into
				)
			})
	},
}

// forEachSchemaName iterates over the physical and virtual schemas.
func forEachSchemaName(
	ctx context.Context,
	p *planner,
	db *sqlbase.ImmutableDatabaseDescriptor,
	fn func(scName string, userDefined bool) error,
) error {
	userDefinedSchemas := make(map[string]struct{})
	schemaNames, err := getSchemaNames(ctx, p, db)
	if err != nil {
		return err
	}
	for _, name := range schemaNames {
		if !strings.HasPrefix(name, sessiondata.PgTempSchemaName) && name != tree.PublicSchema {
			userDefinedSchemas[name] = struct{}{}
		}
	}
	vtableEntries := p.getVirtualTabler().getEntries()
	scNames := make([]string, 0, len(schemaNames)+len(vtableEntries))
	for _, name := range schemaNames {
		scNames = append(scNames, name)
	}
	for _, schema := range vtableEntries {
		scNames = append(scNames, schema.desc.GetName())
	}
	sort.Strings(scNames)
	for _, sc := range scNames {
		_, userDefined := userDefinedSchemas[sc]
		if err := fn(sc, userDefined); err != nil {
			return err
		}
	}
	return nil
}

// forEachDatabaseDesc calls a function for the given DatabaseDescriptor, or if
// it is nil, retrieves all database descriptors and iterates through them in
// lexicographical order with respect to their name. If privileges are required,
// the function is only called if the user has privileges on the database.
func forEachDatabaseDesc(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	requiresPrivileges bool,
	fn func(*sqlbase.ImmutableDatabaseDescriptor) error,
) error {
	var dbDescs []*sqlbase.ImmutableDatabaseDescriptor
	if dbContext == nil {
		allDbDescs, err := p.Descriptors().GetAllDatabaseDescriptors(ctx, p.txn)
		if err != nil {
			return err
		}
		dbDescs = allDbDescs
	} else {
		// We can't just use dbContext here because we need to fetch the descriptor
		// with privileges from kv.
		fetchedDbDesc, err := catalogkv.GetDatabaseDescriptorsFromIDs(ctx, p.txn, p.ExecCfg().Codec, []descpb.ID{dbContext.GetID()})
		if err != nil {
			if errors.Is(err, sqlbase.ErrDescriptorNotFound) {
				return pgerror.Newf(pgcode.UndefinedDatabase, "database %s does not exist", dbContext.GetName())
			}
			return err
		}
		dbDescs = fetchedDbDesc
	}

	// Ignore databases that the user cannot see.
	for _, dbDesc := range dbDescs {
		if !requiresPrivileges || userCanSeeDatabase(ctx, p, dbDesc) {
			if err := fn(dbDesc); err != nil {
				return err
			}
		}
	}

	return nil
}

// forEachTypeDesc calls a function for each TypeDescriptor. If dbContext is
// not nil, then the function is called for only TypeDescriptors within the
// given database.
func forEachTypeDesc(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	fn func(db *sqlbase.ImmutableDatabaseDescriptor, sc string, typ *sqlbase.ImmutableTypeDescriptor) error,
) error {
	descs, err := p.Descriptors().GetAllDescriptors(ctx, p.txn)
	if err != nil {
		return err
	}
	schemaNames, err := getSchemaNames(ctx, p, dbContext)
	if err != nil {
		return err
	}
	lCtx := newInternalLookupCtx(descs, dbContext)
	for _, id := range lCtx.typIDs {
		typ := lCtx.typDescs[id]
		dbDesc, parentExists := lCtx.dbDescs[typ.ParentID]
		if !parentExists {
			continue
		}
		scName, ok := schemaNames[typ.GetParentSchemaID()]
		if !ok {
			return errors.AssertionFailedf("schema id %d not found", typ.GetParentSchemaID())
		}
		if err := fn(dbDesc, scName, typ); err != nil {
			return err
		}
	}
	return nil
}

// forEachTableDesc retrieves all table descriptors from the current
// database and all system databases and iterates through them. For
// each table, the function will call fn with its respective database
// and table descriptor.
//
// The dbContext argument specifies in which database context we are
// requesting the descriptors. In context nil all descriptors are
// visible, in non-empty contexts only the descriptors of that
// database are visible.
//
// The virtualOpts argument specifies how virtual tables are made
// visible.
func forEachTableDesc(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	virtualOpts virtualOpts,
	// TODO(ajwerner): Introduce TableDescriptor.
	fn func(*sqlbase.ImmutableDatabaseDescriptor, string, *sqlbase.ImmutableTableDescriptor) error,
) error {
	return forEachTableDescWithTableLookup(ctx, p, dbContext, virtualOpts, func(
		db *sqlbase.ImmutableDatabaseDescriptor,
		scName string,
		table *sqlbase.ImmutableTableDescriptor,
		_ tableLookupFn,
	) error {
		return fn(db, scName, table)
	})
}

type virtualOpts int

const (
	// virtualMany iterates over virtual schemas in every catalog/database.
	virtualMany virtualOpts = iota
	// virtualOnce iterates over virtual schemas once, in the nil database.
	virtualOnce
	// hideVirtual completely hides virtual schemas during iteration.
	hideVirtual
)

// forEachTableDescAll does the same as forEachTableDesc but also
// includes newly added non-public descriptors.
func forEachTableDescAll(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	virtualOpts virtualOpts,
	fn func(*sqlbase.ImmutableDatabaseDescriptor, string, *sqlbase.ImmutableTableDescriptor) error,
) error {
	return forEachTableDescAllWithTableLookup(ctx,
		p, dbContext, virtualOpts,
		func(
			db *sqlbase.ImmutableDatabaseDescriptor,
			scName string,
			table *sqlbase.ImmutableTableDescriptor,
			_ tableLookupFn,
		) error {
			return fn(db, scName, table)
		})
}

// forEachTableDescAllWithTableLookup is like forEachTableDescAll, but it also
// provides a tableLookupFn like forEachTableDescWithTableLookup.
func forEachTableDescAllWithTableLookup(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	virtualOpts virtualOpts,
	fn func(*sqlbase.ImmutableDatabaseDescriptor, string, *sqlbase.ImmutableTableDescriptor, tableLookupFn) error,
) error {
	return forEachTableDescWithTableLookupInternal(ctx,
		p, dbContext, virtualOpts, true /* allowAdding */, fn)
}

// forEachTableDescWithTableLookup acts like forEachTableDesc, except it also provides a
// tableLookupFn when calling fn to allow callers to lookup fetched table descriptors
// on demand. This is important for callers dealing with objects like foreign keys, where
// the metadata for each object must be augmented by looking at the referenced table.
//
// The dbContext argument specifies in which database context we are
// requesting the descriptors.  In context "" all descriptors are
// visible, in non-empty contexts only the descriptors of that
// database are visible.
func forEachTableDescWithTableLookup(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	virtualOpts virtualOpts,
	fn func(*sqlbase.ImmutableDatabaseDescriptor, string, *sqlbase.ImmutableTableDescriptor, tableLookupFn) error,
) error {
	return forEachTableDescWithTableLookupInternal(ctx, p, dbContext, virtualOpts, false /* allowAdding */, fn)
}

func getSchemaNames(
	ctx context.Context, p *planner, dbContext *sqlbase.ImmutableDatabaseDescriptor,
) (map[descpb.ID]string, error) {
	if dbContext != nil {
		return p.Descriptors().GetSchemasForDatabase(ctx, p.txn, dbContext.GetID())
	}
	ret := make(map[descpb.ID]string)
	dbs, err := p.Descriptors().GetAllDatabaseDescriptors(ctx, p.txn)
	if err != nil {
		return nil, err
	}
	for _, db := range dbs {
		schemas, err := p.Descriptors().GetSchemasForDatabase(ctx, p.txn, db.GetID())
		if err != nil {
			return nil, err
		}
		for id, name := range schemas {
			ret[id] = name
		}
	}
	return ret, nil
}

// forEachTableDescWithTableLookupInternal is the logic that supports
// forEachTableDescWithTableLookup.
//
// The allowAdding argument if true includes newly added tables that
// are not yet public.
func forEachTableDescWithTableLookupInternal(
	ctx context.Context,
	p *planner,
	dbContext *sqlbase.ImmutableDatabaseDescriptor,
	virtualOpts virtualOpts,
	allowAdding bool,
	fn func(*sqlbase.ImmutableDatabaseDescriptor, string, *ImmutableTableDescriptor, tableLookupFn) error,
) error {
	descs, err := p.Descriptors().GetAllDescriptors(ctx, p.txn)
	if err != nil {
		return err
	}
	lCtx := newInternalLookupCtx(descs, dbContext)

	if virtualOpts == virtualMany || virtualOpts == virtualOnce {
		// Virtual descriptors first.
		vt := p.getVirtualTabler()
		vEntries := vt.getEntries()
		vSchemaNames := vt.getSchemaNames()
		iterate := func(dbDesc *sqlbase.ImmutableDatabaseDescriptor) error {
			for _, virtSchemaName := range vSchemaNames {
				e := vEntries[virtSchemaName]
				for _, tName := range e.orderedDefNames {
					te := e.defs[tName]
					if err := fn(dbDesc, virtSchemaName, sqlbase.NewImmutableTableDescriptor(*te.desc), lCtx); err != nil {
						return err
					}
				}
			}
			return nil
		}

		switch virtualOpts {
		case virtualOnce:
			if err := iterate(nil); err != nil {
				return err
			}
		case virtualMany:
			for _, dbID := range lCtx.dbIDs {
				dbDesc := lCtx.dbDescs[dbID]
				if err := iterate(dbDesc); err != nil {
					return err
				}
			}
		}
	}

	// Generate all schema names, and keep a mapping.
	schemaNames, err := getSchemaNames(ctx, p, dbContext)
	if err != nil {
		return err
	}

	// Physical descriptors next.
	for _, tbID := range lCtx.tbIDs {
		table := lCtx.tbDescs[tbID]
		dbDesc, parentExists := lCtx.dbDescs[table.GetParentID()]
		if table.Dropped() || !userCanSeeTable(ctx, p, table, allowAdding) || !parentExists {
			continue
		}
		scName, ok := schemaNames[table.GetParentSchemaID()]
		if !ok {
			return errors.AssertionFailedf("schema id %d not found", table.GetParentSchemaID())
		}
		if err := fn(dbDesc, scName, table, lCtx); err != nil {
			return err
		}
	}
	return nil
}

func forEachIndexInTable(
	table *sqlbase.ImmutableTableDescriptor, fn func(*descpb.IndexDescriptor) error,
) error {
	if table.IsPhysicalTable() {
		if err := fn(&table.PrimaryIndex); err != nil {
			return err
		}
	}
	for i := range table.Indexes {
		if err := fn(&table.Indexes[i]); err != nil {
			return err
		}
	}
	return nil
}

func forEachColumnInTable(
	table *sqlbase.ImmutableTableDescriptor, fn func(*descpb.ColumnDescriptor) error,
) error {
	// Table descriptors already hold columns in-order.
	for i := range table.Columns {
		if err := fn(&table.Columns[i]); err != nil {
			return err
		}
	}
	return nil
}

func forEachColumnInIndex(
	table *sqlbase.ImmutableTableDescriptor,
	index *descpb.IndexDescriptor,
	fn func(*descpb.ColumnDescriptor) error,
) error {
	colMap := make(map[descpb.ColumnID]*descpb.ColumnDescriptor, len(table.Columns))
	for i := range table.Columns {
		id := table.Columns[i].ID
		colMap[id] = &table.Columns[i]
	}
	for _, columnID := range index.ColumnIDs {
		column := colMap[columnID]
		if err := fn(column); err != nil {
			return err
		}
	}
	return nil
}

func forEachRole(
	ctx context.Context,
	p *planner,
	fn func(username string, isRole bool, noLogin bool, rolValidUntil *time.Time) error,
) error {
	query := `
SELECT
	u.username,
	"isRole",
	EXISTS(
		SELECT
			option
		FROM
			system.role_options AS r
		WHERE
			r.username = u.username AND option = 'NOLOGIN'
	)
		AS nologin,
	ro.value::TIMESTAMPTZ AS rolvaliduntil
FROM
	system.users AS u
	LEFT JOIN system.role_options AS ro ON
			ro.username = u.username
			AND option = 'VALID UNTIL';
`
	rows, err := p.ExtendedEvalContext().ExecCfg.InternalExecutor.Query(
		ctx, "read-roles", p.txn, query,
	)

	if err != nil {
		return err
	}

	for _, row := range rows {
		username := tree.MustBeDString(row[0])
		isRole, ok := row[1].(*tree.DBool)
		if !ok {
			return errors.Errorf("isRole should be a boolean value, found %s instead", row[1].ResolvedType())
		}
		noLogin, ok := row[2].(*tree.DBool)
		if !ok {
			return errors.Errorf("noLogin should be a boolean value, found %s instead", row[2].ResolvedType())
		}
		var rolValidUntil *time.Time
		if rolValidUntilDatum, ok := row[3].(*tree.DTimestampTZ); ok {
			rolValidUntil = &rolValidUntilDatum.Time
		} else if row[3] != tree.DNull {
			return errors.Errorf("rolValidUntil should be a timestamp or null value, found %s instead", row[3].ResolvedType())
		}

		if err := fn(string(username), bool(*isRole), bool(*noLogin), rolValidUntil); err != nil {
			return err
		}
	}

	return nil
}

func forEachRoleMembership(
	ctx context.Context, p *planner, fn func(role, member string, isAdmin bool) error,
) error {
	query := `SELECT "role", "member", "isAdmin" FROM system.role_members`
	rows, err := p.ExtendedEvalContext().ExecCfg.InternalExecutor.Query(
		ctx, "read-members", p.txn, query,
	)
	if err != nil {
		return err
	}

	for _, row := range rows {
		roleName := tree.MustBeDString(row[0])
		memberName := tree.MustBeDString(row[1])
		isAdmin := row[2].(*tree.DBool)

		if err := fn(string(roleName), string(memberName), bool(*isAdmin)); err != nil {
			return err
		}
	}
	return nil
}

func userCanSeeDatabase(
	ctx context.Context, p *planner, db *sqlbase.ImmutableDatabaseDescriptor,
) bool {
	return p.CheckAnyPrivilege(ctx, db) == nil
}

func userCanSeeTable(
	ctx context.Context, p *planner, table sqlbase.TableDescriptor, allowAdding bool,
) bool {
	return tableIsVisible(table, allowAdding) && p.CheckAnyPrivilege(ctx, table) == nil
}

func tableIsVisible(table sqlbase.TableDescriptor, allowAdding bool) bool {
	return table.GetState() == descpb.TableDescriptor_PUBLIC ||
		(allowAdding && table.GetState() == descpb.TableDescriptor_ADD)
}
