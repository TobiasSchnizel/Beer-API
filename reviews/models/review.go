package models

import (
	"errors"
	"time"
)

const maxLengthInComments = 400

type Review struct {
	Id      int64
	Stars   int    //1-5
	Comment string //400 chars at max
	Date    time.Time
	BeerID  int64
}

type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
	BeerID  int64  `json:"beer_id"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 - 5")
	}

	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("comment must be less than 400 chars")
	}
	return nil

}
