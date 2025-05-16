package graph

import (
	"context"
	"fmt"

	"github.com/tenteedee/go-graphql/graph/model"
	"github.com/tenteedee/go-graphql/models"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, fmt.Errorf("name too short")
	}

	if len(input.Description) < 10 {
		return nil, fmt.Errorf("description too short")
	}

	return r.MeetupRepo.CreateMeetup(&models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	})
}

// DeleteMeetup is the resolver for the deleteMeetup field.
func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteMeetup - deleteMeetup"))
}

// UpdateMeetup is the resolver for the updateMeetup field.
func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, fmt.Errorf("name too short")
	}

	if len(input.Description) < 10 {
		return nil, fmt.Errorf("description too short")
	}

	return r.MeetupRepo.UpdateMeetup(
		id,
		&model.UpdateMeetup{
			Name:        input.Name,
			Description: input.Description,
		})
}
