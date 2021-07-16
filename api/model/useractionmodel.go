package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	useractionFieldNames          = builderx.RawFieldNames(&Useraction{})
	useractionRows                = strings.Join(useractionFieldNames, ",")
	useractionRowsExpectAutoSet   = strings.Join(stringx.Remove(useractionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	useractionRowsWithPlaceHolder = strings.Join(stringx.Remove(useractionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UseractionModel interface {
		Insert(data Useraction) (sql.Result, error)
		FindOne(id int64) (*Useraction, error)
		Update(data Useraction) error
		Delete(id int64) error
	}

	defaultUseractionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Useraction struct {
		Id   int64          `db:"id"`
		Name sql.NullString `db:"name"`
	}
)

func NewUseractionModel(conn sqlx.SqlConn) UseractionModel {
	return &defaultUseractionModel{
		conn:  conn,
		table: "`useraction`",
	}
}

func (m *defaultUseractionModel) Insert(data Useraction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, useractionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name)
	return ret, err
}

func (m *defaultUseractionModel) FindOne(id int64) (*Useraction, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", useractionRows, m.table)
	var resp Useraction
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

func (m *defaultUseractionModel) Update(data Useraction) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, useractionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Id)
	return err
}

func (m *defaultUseractionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
