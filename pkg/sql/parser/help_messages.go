// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1178
	`ALTER`: {
		//line sql.y: 1179
		Category: hGroup,
		//line sql.y: 1180
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER, ALTER ROLE
`,
	},
	//line sql.y: 1195
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1196
		Category: hDDL,
		//line sql.y: 1197
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
		//line sql.y: 1235
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1249
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1250
		Category: hDDL,
		//line sql.y: 1251
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
		//line sql.y: 1270
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1275
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1276
		Category: hDDL,
		//line sql.y: 1277
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1279
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1286
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1287
		Category: hDDL,
		//line sql.y: 1288
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
	//line sql.y: 1311
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1312
		Category: hDDL,
		//line sql.y: 1313
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1315
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1323
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1324
		Category: hDDL,
		//line sql.y: 1325
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
		//line sql.y: 1337
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1342
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1343
		Category: hDDL,
		//line sql.y: 1344
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
		//line sql.y: 1360
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1862
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1863
		Category: hCCL,
		//line sql.y: 1864
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
		//line sql.y: 1881
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1893
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1894
		Category: hCCL,
		//line sql.y: 1895
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
		//line sql.y: 1911
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1949
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1950
		Category: hCCL,
		//line sql.y: 1951
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
		//line sql.y: 1979
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 2023
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 2024
		Category: hCCL,
		//line sql.y: 2025
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 2034
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2128
	`CANCEL`: {
		//line sql.y: 2129
		Category: hGroup,
		//line sql.y: 2130
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2137
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2138
		Category: hMisc,
		//line sql.y: 2139
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2142
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2160
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2161
		Category: hMisc,
		//line sql.y: 2162
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2165
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2196
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2197
		Category: hMisc,
		//line sql.y: 2198
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2201
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2271
	`CREATE`: {
		//line sql.y: 2272
		Category: hGroup,
		//line sql.y: 2273
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE, CREATE TYPE
`,
	},
	//line sql.y: 2352
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2353
		Category: hMisc,
		//line sql.y: 2354
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2497
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2498
		Category: hDML,
		//line sql.y: 2499
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2503
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2523
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2524
		Category: hCfg,
		//line sql.y: 2525
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2537
	`DROP`: {
		//line sql.y: 2538
		Category: hGroup,
		//line sql.y: 2539
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE, DROP TYPE
`,
	},
	//line sql.y: 2556
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2557
		Category: hDDL,
		//line sql.y: 2558
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2559
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2571
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2572
		Category: hDDL,
		//line sql.y: 2573
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2574
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2586
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2587
		Category: hDDL,
		//line sql.y: 2588
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2589
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2601
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2602
		Category: hDDL,
		//line sql.y: 2603
		Text: `DROP INDEX [CONCURRENTLY] [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2604
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2626
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2627
		Category: hDDL,
		//line sql.y: 2628
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2629
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2649
	`DROP TYPE`: {
		ShortDescription: `remove a type`,
		//line sql.y: 2650
		Category: hDDL,
		//line sql.y: 2651
		Text: `DROP TYPE [IF EXISTS] <type_name> [, ...] [CASCASE | RESTRICT]
`,
		//line sql.y: 2652
		SeeAlso: `WEBDOCS/drop-type.html
`,
	},
	//line sql.y: 2684
	`DROP ROLE`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2685
		Category: hPriv,
		//line sql.y: 2686
		Text: `DROP ROLE [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2687
		SeeAlso: `CREATE ROLE, SHOW ROLE
`,
	},
	//line sql.y: 2711
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2712
		Category: hMisc,
		//line sql.y: 2713
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
		//line sql.y: 2726
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2833
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2834
		Category: hMisc,
		//line sql.y: 2835
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2836
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2867
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2868
		Category: hMisc,
		//line sql.y: 2869
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2870
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2900
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2901
		Category: hMisc,
		//line sql.y: 2902
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2903
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2923
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2924
		Category: hPriv,
		//line sql.y: 2925
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
		//line sql.y: 2938
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2954
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2955
		Category: hPriv,
		//line sql.y: 2956
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
		//line sql.y: 2969
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 3023
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 3024
		Category: hCfg,
		//line sql.y: 3025
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 3026
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3038
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 3039
		Category: hCfg,
		//line sql.y: 3040
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 3041
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 3050
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 3051
		Category: hCfg,
		//line sql.y: 3052
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 3055
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3076
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3077
		Category: hExperimental,
		//line sql.y: 3078
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3086
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3092
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3093
		Category: hExperimental,
		//line sql.y: 3094
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3102
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3110
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3111
		Category: hExperimental,
		//line sql.y: 3112
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
		//line sql.y: 3123
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3178
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3179
		Category: hCfg,
		//line sql.y: 3180
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3181
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3202
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3203
		Category: hCfg,
		//line sql.y: 3204
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3210
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3227
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3228
		Category: hTxn,
		//line sql.y: 3229
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3236
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3428
	`SHOW`: {
		//line sql.y: 3429
		Category: hGroup,
		//line sql.y: 3430
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3498
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3499
		Category: hCfg,
		//line sql.y: 3500
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3501
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3522
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3523
		Category: hExperimental,
		//line sql.y: 3524
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3531
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3544
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3545
		Category: hExperimental,
		//line sql.y: 3546
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3550
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3563
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3564
		Category: hCCL,
		//line sql.y: 3565
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3566
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3605
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3606
		Category: hCfg,
		//line sql.y: 3607
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3610
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3636
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3637
		Category: hDDL,
		//line sql.y: 3638
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3639
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3647
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3648
		Category: hDDL,
		//line sql.y: 3649
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3650
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3670
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3671
		Category: hDDL,
		//line sql.y: 3672
		Text: `SHOW DATABASES
`,
		//line sql.y: 3673
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3681
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3682
		Category: hPriv,
		//line sql.y: 3683
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3689
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3702
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3703
		Category: hDDL,
		//line sql.y: 3704
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3705
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3735
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3736
		Category: hDDL,
		//line sql.y: 3737
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3738
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3751
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3752
		Category: hMisc,
		//line sql.y: 3753
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3754
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3775
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3776
		Category: hMisc,
		//line sql.y: 3777
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3780
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3820
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3821
		Category: hMisc,
		//line sql.y: 3822
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3824
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3847
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3848
		Category: hMisc,
		//line sql.y: 3849
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3850
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3863
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3864
		Category: hDDL,
		//line sql.y: 3865
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3866
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3898
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3899
		Category: hDDL,
		//line sql.y: 3900
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3912
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3913
		Category: hDDL,
		//line sql.y: 3914
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3926
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3927
		Category: hMisc,
		//line sql.y: 3928
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3937
	`SHOW SAVEPOINT`: {
		ShortDescription: `display current savepoint properties`,
		//line sql.y: 3938
		Category: hCfg,
		//line sql.y: 3939
		Text: `SHOW SAVEPOINT STATUS
`,
	},
	//line sql.y: 3947
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3948
		Category: hCfg,
		//line sql.y: 3949
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3950
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3969
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3970
		Category: hDDL,
		//line sql.y: 3971
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3972
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3990
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3991
		Category: hPriv,
		//line sql.y: 3992
		Text: `SHOW USERS
`,
		//line sql.y: 3993
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 4001
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 4002
		Category: hPriv,
		//line sql.y: 4003
		Text: `SHOW ROLES
`,
		//line sql.y: 4004
		SeeAlso: `CREATE ROLE, ALTER ROLE, DROP ROLE
`,
	},
	//line sql.y: 4060
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 4061
		Category: hMisc,
		//line sql.y: 4062
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 4083
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 4084
		Category: hMisc,
		//line sql.y: 4085
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4322
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4323
		Category: hMisc,
		//line sql.y: 4324
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4327
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4344
	`CREATE SCHEMA`: {
		ShortDescription: `create a new schema (not yet supported)`,
		//line sql.y: 4345
		Category: hDDL,
		//line sql.y: 4346
		Text: `
CREATE SCHEMA [IF NOT EXISTS] <schemaname>
`,
	},
	//line sql.y: 4364
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4365
		Category: hDDL,
		//line sql.y: 4366
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>] [<on_commit>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source> [<interleave>] [<on commit>]

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [{STORING | INCLUDE | COVERING} ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> ) [USING HASH WITH BUCKET_COUNT = <shard_buckets>]
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [{STORING | INCLUDE | COVERING} ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

On commit clause:
   ON COMMIT {PRESERVE ROWS | DROP | DELETE ROWS}

`,
		//line sql.y: 4396
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5211
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5212
		Category: hDDL,
		//line sql.y: 5213
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
		//line sql.y: 5223
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5288
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5289
		Category: hDML,
		//line sql.y: 5290
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5291
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5309
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5310
		Category: hPriv,
		//line sql.y: 5311
		Text: `CREATE ROLE [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5312
		SeeAlso: `ALTER ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5324
	`ALTER ROLE`: {
		ShortDescription: `alter a role`,
		//line sql.y: 5325
		Category: hPriv,
		//line sql.y: 5326
		Text: `ALTER ROLE <name> [WITH] <options...>
`,
		//line sql.y: 5327
		SeeAlso: `CREATE ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5356
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5357
		Category: hDDL,
		//line sql.y: 5358
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5359
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5456
	`CREATE TYPE`: {
		ShortDescription: `- create a type`,
		//line sql.y: 5457
		Category: hDDL,
		//line sql.y: 5458
		Text: `CREATE TYPE <type_name> AS ENUM (...)
`,
		//line sql.y: 5459
		SeeAlso: `WEBDOCS/create-type.html
`,
	},
	//line sql.y: 5502
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5503
		Category: hDDL,
		//line sql.y: 5504
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [CONCURRENTLY] [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5512
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5746
	`RELEASE`: {
		ShortDescription: `complete a sub-transaction`,
		//line sql.y: 5747
		Category: hTxn,
		//line sql.y: 5748
		Text: `RELEASE [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5749
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5757
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5758
		Category: hMisc,
		//line sql.y: 5759
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5762
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5779
	`SAVEPOINT`: {
		ShortDescription: `start a sub-transaction`,
		//line sql.y: 5780
		Category: hTxn,
		//line sql.y: 5781
		Text: `SAVEPOINT <savepoint name>
`,
		//line sql.y: 5782
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5797
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5798
		Category: hTxn,
		//line sql.y: 5799
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5807
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5820
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5821
		Category: hTxn,
		//line sql.y: 5822
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5825
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5849
	`ROLLBACK`: {
		ShortDescription: `abort the current (sub-)transaction`,
		//line sql.y: 5850
		Category: hTxn,
		//line sql.y: 5851
		Text: `
ROLLBACK [TRANSACTION]
ROLLBACK [TRANSACTION] TO [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5854
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5954
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5955
		Category: hDDL,
		//line sql.y: 5956
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5957
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 6026
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 6027
		Category: hDML,
		//line sql.y: 6028
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 6033
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 6052
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 6053
		Category: hDML,
		//line sql.y: 6054
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 6058
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 6169
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 6170
		Category: hDML,
		//line sql.y: 6171
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 6178
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6403
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6404
		Category: hDML,
		//line sql.y: 6405
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6416
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6417
		Category: hDML,
		//line sql.y: 6418
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
		//line sql.y: 6430
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6505
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6506
		Category: hDML,
		//line sql.y: 6507
		Text: `TABLE <tablename>
`,
		//line sql.y: 6508
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6830
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6831
		Category: hDML,
		//line sql.y: 6832
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6833
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6942
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6943
		Category: hDML,
		//line sql.y: 6944
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
		//line sql.y: 6966
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
