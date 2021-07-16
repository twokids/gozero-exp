package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userinfoFieldNames          = builderx.RawFieldNames(&Userinfo{})
	userinfoRows                = strings.Join(userinfoFieldNames, ",")
	userinfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userinfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userinfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userinfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserinfoModel interface {
		Insert(data Userinfo) (sql.Result, error)
		FindOne(id int64) (*Userinfo, error)
		Update(data Userinfo) error
		Delete(id int64) error
	}

	defaultUserinfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Userinfo struct {
		Id        int64        `db:"id"`
		Username  string       `db:"username"`
		Nickname  string       `db:"nickname"`
		Age       int64        `db:"age"`
		Pwd       string       `db:"pwd"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)

func NewUserinfoModel(conn sqlx.SqlConn) UserinfoModel {
	return &defaultUserinfoModel{
		conn:  conn,
		table: "`userinfo`",
	}
}

func (m *defaultUserinfoModel) Insert(data Userinfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, userinfoRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Username, data.Nickname, data.Age, data.Pwd, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultUserinfoModel) FindOne(id int64) (*Userinfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userinfoRows, m.table)
	var resp Userinfo
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserinfoModel) Update(data Userinfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userinfoRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Username, data.Nickname, data.Age, data.Pwd, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultUserinfoModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
