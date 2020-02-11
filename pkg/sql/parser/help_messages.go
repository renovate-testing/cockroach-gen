// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1139
	`ALTER`: {
		//line sql.y: 1140
		Category: hGroup,
		//line sql.y: 1141
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1156
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1157
		Category: hDDL,
		//line sql.y: 1158
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... ALTER PRIMARY KEY USING INDEX <name>
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1196
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1210
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1211
		Category: hDDL,
		//line sql.y: 1212
		Text: `
ALTER PARTITION <name> <command>

Commands:
  -- Alter a single partition which exists on any of a table's indexes.
  ALTER PARTITION <partition> OF TABLE <tablename> CONFIGURE ZONE <zoneconfig>

  -- Alter a partition of a specific index.
  ALTER PARTITION <partition> OF INDEX <tablename>@<indexname> CONFIGURE ZONE <zoneconfig>

  -- Alter all partitions with the same name across a table's indexes.
  ALTER PARTITION <partition> OF INDEX <tablename>@* CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1231
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1236
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1237
		Category: hDDL,
		//line sql.y: 1238
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1240
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1247
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1248
		Category: hDDL,
		//line sql.y: 1249
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1272
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1273
		Category: hPriv,
		//line sql.y: 1274
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1276
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1281
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1282
		Category: hDDL,
		//line sql.y: 1283
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1285
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1293
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1294
		Category: hDDL,
		//line sql.y: 1295
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1307
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1312
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1313
		Category: hDDL,
		//line sql.y: 1314
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1330
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1831
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1832
		Category: hCCL,
		//line sql.y: 1833
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1850
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1862
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1863
		Category: hCCL,
		//line sql.y: 1864
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1880
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1918
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1919
		Category: hCCL,
		//line sql.y: 1920
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   DELIMITED
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1948
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1992
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1993
		Category: hCCL,
		//line sql.y: 1994
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 2003
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2097
	`CANCEL`: {
		//line sql.y: 2098
		Category: hGroup,
		//line sql.y: 2099
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2106
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2107
		Category: hMisc,
		//line sql.y: 2108
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2111
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2129
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2130
		Category: hMisc,
		//line sql.y: 2131
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2134
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2165
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2166
		Category: hMisc,
		//line sql.y: 2167
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2170
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2240
	`CREATE`: {
		//line sql.y: 2241
		Category: hGroup,
		//line sql.y: 2242
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2323
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2324
		Category: hMisc,
		//line sql.y: 2325
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2468
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2469
		Category: hDML,
		//line sql.y: 2470
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2474
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2494
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2495
		Category: hCfg,
		//line sql.y: 2496
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2508
	`DROP`: {
		//line sql.y: 2509
		Category: hGroup,
		//line sql.y: 2510
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2527
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2528
		Category: hDDL,
		//line sql.y: 2529
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2530
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2542
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2543
		Category: hDDL,
		//line sql.y: 2544
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2545
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2557
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2558
		Category: hDDL,
		//line sql.y: 2559
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2560
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2572
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2573
		Category: hDDL,
		//line sql.y: 2574
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2575
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2595
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2596
		Category: hDDL,
		//line sql.y: 2597
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2598
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2618
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2619
		Category: hPriv,
		//line sql.y: 2620
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2621
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2633
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2634
		Category: hPriv,
		//line sql.y: 2635
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2636
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2660
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2661
		Category: hMisc,
		//line sql.y: 2662
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2675
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2758
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2759
		Category: hMisc,
		//line sql.y: 2760
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2761
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2792
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2793
		Category: hMisc,
		//line sql.y: 2794
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2795
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2825
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2826
		Category: hMisc,
		//line sql.y: 2827
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2828
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2848
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2849
		Category: hPriv,
		//line sql.y: 2850
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2863
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2879
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2880
		Category: hPriv,
		//line sql.y: 2881
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2894
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2948
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2949
		Category: hCfg,
		//line sql.y: 2950
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2951
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2963
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2964
		Category: hCfg,
		//line sql.y: 2965
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2966
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2975
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2976
		Category: hCfg,
		//line sql.y: 2977
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2980
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3001
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3002
		Category: hExperimental,
		//line sql.y: 3003
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3011
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3017
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3018
		Category: hExperimental,
		//line sql.y: 3019
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3027
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3035
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3036
		Category: hExperimental,
		//line sql.y: 3037
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 3048
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3103
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3104
		Category: hCfg,
		//line sql.y: 3105
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3106
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3127
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3128
		Category: hCfg,
		//line sql.y: 3129
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3135
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3152
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3153
		Category: hTxn,
		//line sql.y: 3154
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3161
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3353
	`SHOW`: {
		//line sql.y: 3354
		Category: hGroup,
		//line sql.y: 3355
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3391
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3392
		Category: hCfg,
		//line sql.y: 3393
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3394
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3415
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3416
		Category: hExperimental,
		//line sql.y: 3417
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3424
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3437
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3438
		Category: hExperimental,
		//line sql.y: 3439
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3443
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3456
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3457
		Category: hCCL,
		//line sql.y: 3458
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3459
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3498
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3499
		Category: hCfg,
		//line sql.y: 3500
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3503
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3529
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3530
		Category: hDDL,
		//line sql.y: 3531
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3532
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3540
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3541
		Category: hDDL,
		//line sql.y: 3542
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3543
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3563
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3564
		Category: hDDL,
		//line sql.y: 3565
		Text: `SHOW DATABASES
`,
		//line sql.y: 3566
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3574
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3575
		Category: hPriv,
		//line sql.y: 3576
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3582
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3595
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3596
		Category: hDDL,
		//line sql.y: 3597
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3598
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3628
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3629
		Category: hDDL,
		//line sql.y: 3630
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3631
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3644
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3645
		Category: hMisc,
		//line sql.y: 3646
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3647
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3668
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3669
		Category: hMisc,
		//line sql.y: 3670
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3673
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3713
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3714
		Category: hMisc,
		//line sql.y: 3715
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3717
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3740
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3741
		Category: hMisc,
		//line sql.y: 3742
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3743
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3756
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3757
		Category: hDDL,
		//line sql.y: 3758
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3759
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3791
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3792
		Category: hDDL,
		//line sql.y: 3793
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3805
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3806
		Category: hDDL,
		//line sql.y: 3807
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3819
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3820
		Category: hMisc,
		//line sql.y: 3821
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3830
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3831
		Category: hCfg,
		//line sql.y: 3832
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3833
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3852
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3853
		Category: hDDL,
		//line sql.y: 3854
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3855
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3873
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3874
		Category: hPriv,
		//line sql.y: 3875
		Text: `SHOW USERS
`,
		//line sql.y: 3876
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3884
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3885
		Category: hPriv,
		//line sql.y: 3886
		Text: `SHOW ROLES
`,
		//line sql.y: 3887
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3943
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3944
		Category: hMisc,
		//line sql.y: 3945
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3966
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3967
		Category: hMisc,
		//line sql.y: 3968
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4205
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4206
		Category: hMisc,
		//line sql.y: 4207
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4210
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4227
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4228
		Category: hDDL,
		//line sql.y: 4229
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> ) [USING HASH WITH BUCKET_COUNT = <shard_buckets>]
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4256
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5016
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5017
		Category: hDDL,
		//line sql.y: 5018
		Text: `
CREATE [TEMPORARY | TEMP] SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 5028
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5093
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5094
		Category: hDML,
		//line sql.y: 5095
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5096
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5104
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5105
		Category: hPriv,
		//line sql.y: 5106
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5107
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5136
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5137
		Category: hPriv,
		//line sql.y: 5138
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5139
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5157
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5158
		Category: hDDL,
		//line sql.y: 5159
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5160
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5207
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5208
		Category: hDDL,
		//line sql.y: 5209
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5217
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5448
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5449
		Category: hTxn,
		//line sql.y: 5450
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5451
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5459
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5460
		Category: hMisc,
		//line sql.y: 5461
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5464
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5481
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5482
		Category: hTxn,
		//line sql.y: 5483
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5484
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5499
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5500
		Category: hTxn,
		//line sql.y: 5501
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5509
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5522
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5523
		Category: hTxn,
		//line sql.y: 5524
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5527
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5551
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5552
		Category: hTxn,
		//line sql.y: 5553
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5554
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5672
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5673
		Category: hDDL,
		//line sql.y: 5674
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5675
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5744
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5745
		Category: hDML,
		//line sql.y: 5746
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5751
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5770
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5771
		Category: hDML,
		//line sql.y: 5772
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5776
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5887
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5888
		Category: hDML,
		//line sql.y: 5889
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5896
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6121
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6122
		Category: hDML,
		//line sql.y: 6123
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6134
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6135
		Category: hDML,
		//line sql.y: 6136
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 6148
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6223
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6224
		Category: hDML,
		//line sql.y: 6225
		Text: `TABLE <tablename>
`,
		//line sql.y: 6226
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6538
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6539
		Category: hDML,
		//line sql.y: 6540
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6541
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6650
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6651
		Category: hDML,
		//line sql.y: 6652
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 6674
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
