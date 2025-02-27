package tests

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/rawsql"
)

func TestSqlGen(t *testing.T) {
	//rawsql := []string{
	//	"CREATE TABLE `users` (\n                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,\n                         `created_at` datetime(3) DEFAULT NULL,\n                         `name` varchar(255) DEFAULT NULL COMMENT 'oneline',\n                         `address` varchar(255) DEFAULT '',\n                         `register_time` datetime(3) DEFAULT NULL,\n                         `alive` tinyint(1) DEFAULT NULL COMMENT 'multiline\\nline1\\nline2',\n                         `company_id` bigint(20) unsigned DEFAULT '666',\n                         `private_url` varchar(255) DEFAULT 'https://a.b.c ',\n                         PRIMARY KEY (`id`),\n                         KEY `idx_name` (`name`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;",
	//}
	db, err := gorm.Open(rawsql.New(rawsql.Config{
		//SQL:      rawsql,                      //create table sql
		FilePath: []string{
			//"./sql/user.sql", // create table sql file
			"./sql", // create table sql file directory
		},
	}))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(db.Migrator().GetTables())
	fmt.Println(db.Migrator().ColumnTypes("users"))
}
