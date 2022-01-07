package models

import "testing"

func NewReview(stars int, comment string, beerid int64) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   	stars,
		Comment: 	comment,
		BeerID:		beerid,
	}
}
func Test_withCorrectParams(t *testing.T) {
	r := NewReview(5, "Amazing", 1)

	err := r.validate()

	if err != nil {
		t.Error("the validation did not past")
		t.Fail()
	}
}

func Test_shouldFailWithWrongNumberOfStars(t *testing.T) {
	r := NewReview(8, "Nice Beer", 1)
	err := r.validate()
	if err != nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}

func Test_shouldFailWithExtensionComment(t *testing.T) {
	r := NewReview(4, "I am so glad this was asked. I bought the Seedlip Garden from Web about a month ago, and took it to a party over the weekend; it was amazing! Sadly, I left my bottle there, thinking I could just get another to replace it. And even more sadly, found out that shipping to Pennsylvania is also being blocked. C'mon, Amazon! Resolve this! I want to start experimenting with it with aquafaba and other ingredients!! SO.MUCH.FUN.", 1)
	err := r.validate()
	if err != nil {
		t.Error("comment lenth is more than 400 chars")
		t.Fail()
	}
}
