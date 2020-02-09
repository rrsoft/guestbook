package core

import (
	"errors"
	"time"

	"github.com/rrsoft/guestbook/data"
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

func GetList(page, size int) ([]*Greeting, error) {
	rows, err := data.Query(data.AppStting.DSNUser, SELECT_GUESTBOOK_LIST, page*size, size)
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

func GetDetails(id int) (*Greeting, error) {
	row := data.QueryRow(data.AppStting.DSNUser, SELECT_DETAILS, id)
	if row == nil {
		return nil, errors.New("id is not found")
	}
	var g = new(Greeting)
	if err := row.Scan(&g.Id, &g.Author, &g.Content, &g.PostDate); err != nil {
		return nil, err
	}
	return g, nil
}

func Count() int {
	row := data.QueryRow(data.AppStting.DSNUser, SELECT_COUNT_GUESTBOOK)
	if row == nil {
		return 0
	}
	var count int
	if err := row.Scan(&count); err != nil {
		return 0
	}
	return count
}

func Comment(g *Greeting) error {
	tx, err := data.Begin(data.AppStting.DSNUser)
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
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	g.Id = int(id)
	return nil
}

func Create(g *Greeting) error {
	id, err := data.ExecInsertId(data.AppStting.DSNUser,
		INSERT_GUESTBOOK, g.Author, g.Content, g.PostDate)
	if err != nil {
		return err
	}
	g.Id = int(id)
	return nil
}

func Delete(id int) error {
	_, err := data.Exec(data.AppStting.DSNUser, DELETE_GUESTBOOK, id)
	if err != nil {
		return err
	}
	//n := res.RowsAffected()
	return nil
}
