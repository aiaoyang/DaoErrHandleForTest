package service

import (
	"log"

	"github.com/aiaoyang/errhandle/pkg/dao"
)

func Query(_sql string) {

	rows, err := dao.DaoQuery("_somethingUseless")
	if err != nil {
		log.Printf("failed, err: %+v\n", err)
		return
	}
	log.Printf("success, rows: %+v\n", rows)

}
