package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/tenteedee/go-graphql/graph/model"
	"github.com/tenteedee/go-graphql/models"
)

type MeetupRepo struct {
	DB *pg.DB
}

func (m *MeetupRepo) GetMeetups(page int, pageSize int) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Limit(pageSize).Offset(pageSize * (page - 1)).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return meetup, nil
}

func (m *MeetupRepo) DeleteMeetup(id string) (bool, error) {
	_, err := m.DB.Model(&models.Meetup{}).Where("id = ?", id).Delete()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MeetupRepo) UpdateMeetup(id string, meetup *model.UpdateMeetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", id).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	// return meetup, nil
	return nil, nil
}
