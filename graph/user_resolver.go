package graph

import (
	"context"

	"github.com/tenteedee/go-graphql/models"
)

type userResolver struct{ *Resolver }

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.UserRepo.GetMeetupsByUserId(obj.ID)
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }
