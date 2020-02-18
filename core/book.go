package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/rrsoft/guestbook/data"
	"github.com/rrsoft/guestbook/utils"
)

type Greeting struct {
	Id       int
	Author   string
	Content  string
	PostDate time.Time
}

var (
	SELECT_GUESTBOOK_LIST  = "SELECT * FROM `guestbook` ORDER BY `id` DESC LIMIT ?,?"
	SELECT_DETAILS         = "SELECT * FROM `guestbook` WHERE `id`=?"
	SELECT_COUNT_GUESTBOOK = "SELECT COUNT(*) FROM `guestbook`"
	INSERT_GUESTBOOK       = "INSERT INTO `guestbook` VALUES (null, ?, ?, ?)"
	DELETE_GUESTBOOK       = "DELETE FROM `guestbook` WHERE `id`=?"
)

// GetList 留言分页  page从0开始
func GetList(page, size int) ([]*Greeting, error) {
	settings := utils.GetSetting()
	rows, err := data.Query(settings.DSNUser, SELECT_GUESTBOOK_LIST, page*size, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Greeting
	for rows.Next() {
		var g = new(Greeting)
		if e := rows.Scan(&g.Id, &g.Author, &g.Content, &g.PostDate); e != nil {
			return nil, e
		}
		res = append(res, g)
	}
	return res, nil
}

// GetDetails 留言详情
func GetDetails(id int) (*Greeting, error) {
	settings := utils.GetSetting()
	row := data.QueryRow(settings.DSNUser, SELECT_DETAILS, id)
	if row == nil {
		return nil, errors.New("id is not found")
	}
	var g = new(Greeting)
	if err := row.Scan(&g.Id, &g.Author, &g.Content, &g.PostDate); err != nil {
		return nil, err
	}
	return g, nil
}

// Count 留言数量
func Count() int {
	settings := utils.GetSetting()
	row := data.QueryRow(settings.DSNUser, SELECT_COUNT_GUESTBOOK)
	if row == nil {
		return 0
	}
	var count int
	if err := row.Scan(&count); err != nil {
		return 0
	}
	return count
}

// Comment 发表留言-事务
func Comment(g *Greeting) error {
	settings := utils.GetSetting()
	tx, err := data.Begin(settings.DSNUser)
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(INSERT_GUESTBOOK)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(g.Author, g.Content, g.PostDate)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		tx.Commit()
		g.Id = int(id)
	} else {
		tx.Rollback()
	}
	return err
}

// Create 发表留言
func Create(g *Greeting) error {
	settings := utils.GetSetting()
	id, err := data.ExecInsertId(settings.DSNUser, INSERT_GUESTBOOK, g.Author, g.Content, g.PostDate)
	if err == nil {
		g.Id = int(id)
	}
	return err
}

// Delete 删除留言
func Delete(id int) error {
	settings := utils.GetSetting()
	res, err := data.Exec(settings.DSNUser, DELETE_GUESTBOOK, id)
	if err == nil {
		n, _ := res.RowsAffected()
		fmt.Println(n)
	}
	return err
}
