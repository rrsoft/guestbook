package data

import (
	"database/sql"
)

// 执行命令返回影响行数和自增Id
func Exec(dsn, query string, args ...interface{}) (sql.Result, error) {
	db, err := Open(dsn)
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

// 执行命令返回自增Id
func ExecInsertId(dsn, query string, args ...interface{}) (int64, error) {
	res, err := Exec(dsn, query, args...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 执行命令返回影响行数
func ExecRowsAffected(dsn, query string, args ...interface{}) (int64, error) {
	res, err := Exec(dsn, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// 执行命令返回行
func Query(dsn, query string, args ...interface{}) (*sql.Rows, error) {
	db, err := Open(dsn)
	if err != nil {
		return nil, err
	}
	return db.Query(query, args...)
}

// 执行命令返回第一行第一列数据
func QueryRow(dsn, query string, args ...interface{}) *sql.Row {
	db, err := Open(dsn)
	if err != nil {
		return nil
	}
	return db.QueryRow(query, args...)
}

// 开始事务
func Begin(dsn string) (*sql.Tx, error) {
	db, err := Open(dsn)
	if err != nil {
		return nil, err
	}
	return db.Begin()
}
