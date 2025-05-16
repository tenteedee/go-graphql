package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/tenteedee/go-graphql/models"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	var user models.User

	err := u.DB.Model(&user).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetMeetupsByUserId(id string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := u.DB.Model(&meetups).Where("user_id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}
