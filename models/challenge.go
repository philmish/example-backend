package models

import (
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	Name  string    `json:"name" gorm:"type:varchar(50);column:name;unique"`
	Start time.Time `json:"start" gorm:"type:date;column:start"`
	End   time.Time `json:"end" gorm:"type:date;column:end"`
}

func newChallenge(name string, start, end time.Time) (Challenge, error) {
	if end.Before(start) {
		return Challenge{}, errors.New("Challenge ends before start")
	}
	return Challenge{Name: name, Start: start, End: end}, nil
}

type ChallengeCreateReq struct {
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (c ChallengeCreateReq) ChallengeFromReq() (Challenge, error) {
	start, err := time.Parse("02.01.2006", c.Start)
	if err != nil {
		return Challenge{}, err
	}
	end, err := time.Parse("02.01.2006", c.End)
	if err != nil {
		return Challenge{}, err
	}
	challenge, err := newChallenge(c.Name, start, end)
	return challenge, err
}

func CreateChallenge(challenge Challenge, dbName string) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.Create(&challenge).Error
	return err
}
