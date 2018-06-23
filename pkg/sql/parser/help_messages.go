// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1079
	`ALTER`: {
		//line sql.y: 1080
		Category: hGroup,
		//line sql.y: 1081
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1095
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1096
		Category: hDDL,
		//line sql.y: 1097
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
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1121
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1133
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1134
		Category: hDDL,
		//line sql.y: 1135
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1137
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1144
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1145
		Category: hDDL,
		//line sql.y: 1146
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
	//line sql.y: 1169
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1170
		Category: hPriv,
		//line sql.y: 1171
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1173
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1178
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1179
		Category: hDDL,
		//line sql.y: 1180
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1182
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1193
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1194
		Category: hDDL,
		//line sql.y: 1195
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1203
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1545
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1546
		Category: hCCL,
		//line sql.y: 1547
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
		//line sql.y: 1564
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1572
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1573
		Category: hCCL,
		//line sql.y: 1574
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
		//line sql.y: 1590
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1608
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1609
		Category: hCCL,
		//line sql.y: 1610
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
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
		//line sql.y: 1632
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1653
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1654
		Category: hCCL,
		//line sql.y: 1655
		Text: `
EXPORT <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1664
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1751
	`CANCEL`: {
		//line sql.y: 1752
		Category: hGroup,
		//line sql.y: 1753
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1760
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1761
		Category: hMisc,
		//line sql.y: 1762
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1765
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1783
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1784
		Category: hMisc,
		//line sql.y: 1785
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1788
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1819
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1820
		Category: hMisc,
		//line sql.y: 1821
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1824
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1871
	`CREATE`: {
		//line sql.y: 1872
		Category: hGroup,
		//line sql.y: 1873
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1895
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1896
		Category: hMisc,
		//line sql.y: 1897
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1932
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1933
		Category: hDML,
		//line sql.y: 1934
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1938
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1953
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1954
		Category: hCfg,
		//line sql.y: 1955
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1967
	`DROP`: {
		//line sql.y: 1968
		Category: hGroup,
		//line sql.y: 1969
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 1985
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1986
		Category: hDDL,
		//line sql.y: 1987
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1988
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2000
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2001
		Category: hDDL,
		//line sql.y: 2002
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2003
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2015
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2016
		Category: hDDL,
		//line sql.y: 2017
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2018
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2030
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2031
		Category: hDDL,
		//line sql.y: 2032
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2033
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2053
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2054
		Category: hDDL,
		//line sql.y: 2055
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2056
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2076
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2077
		Category: hPriv,
		//line sql.y: 2078
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2079
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2091
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2092
		Category: hPriv,
		//line sql.y: 2093
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2094
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2116
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2117
		Category: hMisc,
		//line sql.y: 2118
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, VERBOSE

`,
		//line sql.y: 2130
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2195
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2196
		Category: hMisc,
		//line sql.y: 2197
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2198
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2220
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2221
		Category: hMisc,
		//line sql.y: 2222
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2223
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2246
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2247
		Category: hMisc,
		//line sql.y: 2248
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2249
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2269
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2270
		Category: hPriv,
		//line sql.y: 2271
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
		//line sql.y: 2284
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2300
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2301
		Category: hPriv,
		//line sql.y: 2302
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
		//line sql.y: 2315
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2370
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2371
		Category: hCfg,
		//line sql.y: 2372
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2373
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2385
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2386
		Category: hCfg,
		//line sql.y: 2387
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2388
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2397
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2398
		Category: hCfg,
		//line sql.y: 2399
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2402
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2419
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2420
		Category: hExperimental,
		//line sql.y: 2421
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2429
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2435
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2436
		Category: hExperimental,
		//line sql.y: 2437
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2445
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2453
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2454
		Category: hExperimental,
		//line sql.y: 2455
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
		//line sql.y: 2466
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2521
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2522
		Category: hCfg,
		//line sql.y: 2523
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2524
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2545
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2546
		Category: hCfg,
		//line sql.y: 2547
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { off | cluster | on | kv | local } [,...]

`,
		//line sql.y: 2553
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2570
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2571
		Category: hTxn,
		//line sql.y: 2572
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2579
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2754
	`SHOW`: {
		//line sql.y: 2755
		Category: hGroup,
		//line sql.y: 2756
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW CREATE SEQUENCE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2790
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2791
		Category: hCfg,
		//line sql.y: 2792
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2793
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2814
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2815
		Category: hMisc,
		//line sql.y: 2816
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2823
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2835
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2836
		Category: hMisc,
		//line sql.y: 2837
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2841
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2854
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2855
		Category: hCCL,
		//line sql.y: 2856
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 2857
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2882
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2883
		Category: hCfg,
		//line sql.y: 2884
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2887
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2904
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2905
		Category: hDDL,
		//line sql.y: 2906
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2907
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2915
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2916
		Category: hDDL,
		//line sql.y: 2917
		Text: `SHOW DATABASES
`,
		//line sql.y: 2918
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2926
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2927
		Category: hPriv,
		//line sql.y: 2928
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2934
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2947
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2948
		Category: hDDL,
		//line sql.y: 2949
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2950
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2968
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2969
		Category: hDDL,
		//line sql.y: 2970
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2971
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2984
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2985
		Category: hMisc,
		//line sql.y: 2986
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2987
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3003
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3004
		Category: hMisc,
		//line sql.y: 3005
		Text: `SHOW JOBS
`,
		//line sql.y: 3006
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3014
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3015
		Category: hMisc,
		//line sql.y: 3016
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
SHOW [COMPACT] [KV] TRACE FOR <statement>
`,
		//line sql.y: 3019
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3049
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3050
		Category: hMisc,
		//line sql.y: 3051
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3052
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3068
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3069
		Category: hDDL,
		//line sql.y: 3070
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3071
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3097
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3098
		Category: hDDL,
		//line sql.y: 3099
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3111
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3112
		Category: hMisc,
		//line sql.y: 3113
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3122
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3123
		Category: hCfg,
		//line sql.y: 3124
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3125
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3144
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 3145
		Category: hDDL,
		//line sql.y: 3146
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 3147
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3155
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 3156
		Category: hDDL,
		//line sql.y: 3157
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 3158
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 3166
	`SHOW CREATE SEQUENCE`: {
		ShortDescription: `display the CREATE SEQUENCE statement for a sequence`,
		//line sql.y: 3167
		Category: hDDL,
		//line sql.y: 3168
		Text: `SHOW CREATE SEQUENCE <seqname>
`,
	},
	//line sql.y: 3176
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3177
		Category: hPriv,
		//line sql.y: 3178
		Text: `SHOW USERS
`,
		//line sql.y: 3179
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3187
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3188
		Category: hPriv,
		//line sql.y: 3189
		Text: `SHOW ROLES
`,
		//line sql.y: 3190
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3242
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3243
		Category: hMisc,
		//line sql.y: 3244
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3480
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3481
		Category: hMisc,
		//line sql.y: 3482
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3485
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3502
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3503
		Category: hDDL,
		//line sql.y: 3504
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
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
		//line sql.y: 3531
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4042
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4043
		Category: hDDL,
		//line sql.y: 4044
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]

`,
		//line sql.y: 4053
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4106
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4107
		Category: hDML,
		//line sql.y: 4108
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4109
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4117
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4118
		Category: hPriv,
		//line sql.y: 4119
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4120
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4142
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4143
		Category: hPriv,
		//line sql.y: 4144
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4145
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4157
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4158
		Category: hDDL,
		//line sql.y: 4159
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4160
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4174
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4175
		Category: hDDL,
		//line sql.y: 4176
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4184
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4386
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4387
		Category: hTxn,
		//line sql.y: 4388
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4389
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4397
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4398
		Category: hMisc,
		//line sql.y: 4399
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4402
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4419
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4420
		Category: hTxn,
		//line sql.y: 4421
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4422
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4437
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4438
		Category: hTxn,
		//line sql.y: 4439
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4447
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4460
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4461
		Category: hTxn,
		//line sql.y: 4462
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4465
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4489
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4490
		Category: hTxn,
		//line sql.y: 4491
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4492
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4605
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4606
		Category: hDDL,
		//line sql.y: 4607
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4608
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4677
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4678
		Category: hDML,
		//line sql.y: 4679
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4684
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4703
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4704
		Category: hDML,
		//line sql.y: 4705
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4709
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4814
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4815
		Category: hDML,
		//line sql.y: 4816
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4823
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 4993
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 4994
		Category: hDML,
		//line sql.y: 4995
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5006
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5007
		Category: hDML,
		//line sql.y: 5008
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
		//line sql.y: 5020
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5095
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5096
		Category: hDML,
		//line sql.y: 5097
		Text: `TABLE <tablename>
`,
		//line sql.y: 5098
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5364
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5365
		Category: hDML,
		//line sql.y: 5366
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5367
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5468
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5469
		Category: hDML,
		//line sql.y: 5470
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5488
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
