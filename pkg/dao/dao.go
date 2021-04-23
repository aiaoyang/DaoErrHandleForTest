package dao

import (
	"errors"
	"log"

	"github.com/aiaoyang/errhandle/pkg/db"
)

type Rows struct {
	daoData int
}

type DaoError error

var (
	// Dao层错误
	ErrDaoNoRow = errors.New("ErrDaoNoRow")
	ErrDaoOther = errors.New("ErrDaoOther")

	// 映射DB层错误至Dao层错误，不直接返回DB错误给用户，用户侧只处理Dao层的错误
	ErrMap = map[db.DBError]DaoError{
		db.ErrDBNoRows: ErrDaoNoRow,
		db.ErrDBOther:  ErrDaoOther,
	}
)

func DaoQuery(_sql string) (rows *Rows, err error) {

	dbRows, err := db.Query("_somethingUseless")
	log.Printf("%+v", err)
	if err != nil {
		// 返回Dao层映射的错误给用户侧
		return nil, ErrMap[err]
	}

	return &Rows{daoData: dbRows.Data}, nil
}
