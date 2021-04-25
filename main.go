package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	// Service 层错误
	err1 := Service("Service_1", 102, 3)
	fmt.Printf("---------------\n%v\n----------------", err1)

	err2 := Service("Service_1", 102, 0)
	fmt.Printf("---------------\n%v\n----------------", err2)

	err3 := Service("Service_2", 103, 3)
	fmt.Printf("---------------\n%v\n----------------", err3)

	err4 := Service("Service_3", 104, 3)
	fmt.Printf("---------------\n%v\n----------------", err4)

	// Dao层错误
	err5 := Service("Dao_1", 99, 0)
	fmt.Printf("---------------\n%v\n----------------", err5)

	err6 := Service("Dao_2", 200, 0)
	fmt.Printf("---------------\n%v\n----------------", err6)

}

// type OkToDo func(rows *DaoRows, mod int) error

// ------------------------------------------------

type ServiceError error

var (
	ErrServiceUndefinedBehavior ServiceError = errors.New("Service_UndefinedBehavior")
	ErrServiceUnexpectedData    ServiceError = errors.New("Service_UnexpectedData")
)

func OkHandler(rows *DaoRows, mod int) (errNew error) {
	// 根据期望处理数据
	if mod == 0 {
		return ErrServiceUndefinedBehavior
	}
	switch rows.DaoData % mod {
	case 0:
		return
	case 1:
		return ErrServiceUnexpectedData
	default:
		return ErrServiceUndefinedBehavior
	}
}

// Service Layer
func Service(callerIdent string, sql int, expectMod int) (err error) {

	rows, err := DaoQuery(sql)
	if err != nil {
		return errors.Wrapf(err, "Caller:%s", callerIdent)
	}
	return OkHandler(rows, expectMod)

}

// ------------------------------------------------
// DAO Layer
type DaoRows struct {
	DaoData int
}

type DaoError error

var (
	// Dao层错误
	ErrDaoNoRow = errors.New("ErrDaoNoRow")
	ErrDaoOther = errors.New("ErrDaoOther")

	// 映射DB层错误至Dao层错误，不直接返回DB错误给用户，用户侧只处理Dao层的错误
	ErrMap = map[DBError]DaoError{
		ErrDBNoRows: ErrDaoNoRow,
		ErrDBOther:  ErrDaoOther,
	}
)

func DaoQuery(sql int) (rows *DaoRows, err error) {
	dbRows, err := DBQuery(sql)
	if err != nil {
		// 返回Dao层映射的错误给用户侧
		return nil, ErrMap[err]
	}

	return &DaoRows{DaoData: dbRows.Data}, nil
}

// ------------------------------------------------
// DB Layer
type DBError error

var (
	// DB 层错误
	ErrDBNoRows DBError = errors.New("DBErrNoRows")
	ErrDBOther  DBError = errors.New("DBErrOther")
)

type DBRows struct {
	Data int
}

func DBQuery(sql int) (result *DBRows, err error) {
	if sql >= 100 && sql < 200 {
		return &DBRows{sql}, nil
	} else if sql > 0 && sql < 100 {
		return nil, ErrDBNoRows
	} else {
		return nil, ErrDBOther
	}
}
