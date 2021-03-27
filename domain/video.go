package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init()  {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Video struct {
	ID         string `valid:"uuid"`
	ResourceID string `valid:"notnull"`
	FilePath   string `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}

func NewVideo() *Video {
	return &Video{}
}