package main

const (
	MySQLInfo           = "SELECT VERSION(), USER(), DATABASE()"
	MySQLDatabases      = "SELECT SCHEMA_NAME, DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME FROM information_schema.SCHEMATA ORDER BY schema_name;"
	MySQLDatabaseTables = "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = '%s' ORDER BY TABLE_NAME;"
	MySQLTables         = "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test' ORDER BY TABLE_NAME;"
	MySQLTable          = "SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH, CHARACTER_SET_NAME, COLUMN_DEFAULT FROM information_schema.columns WHERE TABLE_NAME = '%s';"
	MySQLTableInfo      = "SELECT DATA_LENGTH AS data_length, INDEX_LENGTH AS index_length, (DATA_LENGTH + INDEX_LENGTH) AS total_size, TABLE_ROWS AS row_count FROM information_schema.TABLES WHERE TABLE_NAME = '%s';"
	MySQLTableIndexs    = "SELECT INDEX_NAME, INDEX_TYPE FROM information_schema.statistics WHERE TABLE_NAME = '%s';"
)
