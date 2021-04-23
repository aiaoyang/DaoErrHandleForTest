package db

import (
	"errors"
	"math/rand"
	"time"
)

type DBError error

var (
	// DB 层错误
	ErrDBNoRows DBError = errors.New("DBErrNoRows")
	ErrDBOther  DBError = errors.New("DBErrOther")
)

type Rows struct {
	Data int
}

func Query(_sql string) (result *Rows, err error) {

	// 生成随机数，使多次运行结果不一致以观察效果
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(11)

	switch rnd % 3 {
	case 0:
		return &Rows{999}, nil
	case 1:
		return nil, ErrDBNoRows
	default:
		return nil, ErrDBOther
	}
}
