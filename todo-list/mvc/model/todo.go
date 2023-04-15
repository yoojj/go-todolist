package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Tid  string `json:"-"    gorm:"AUTO_INCREMENT"`
	Uuid string `json:"uuid" gorm:"UNIQUE"`
	Task string `json:"task" binding:"required"`

	Important string `json:"important" gorm:"size:1"`
	Star      string `json:"star"      gorm:"size:1"`

	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`

	CreateName string `json:"name"`

	TaskYN   string `json:"taskYN"     gorm:"size:1"`
	DeleteYN string `json:"deleteYN"   gorm:"size:1"`

	CreateDate time.Time `json:"createDate"`
	ModifyDate time.Time `json:"-"`
	DeleteDate time.Time `json:"-"`

	Status int `json:"status"`
}

func (t *Todo) CreateUUID() {
	t.Uuid = strings.Replace(uuid.New().String(), "-", "", -1)
}

func (t Todo) GetUUID() string {
	return t.Uuid
}
