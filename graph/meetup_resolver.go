package graph

import (
	"context"

	"github.com/tenteedee/go-graphql/models"
)

type meetupResolver struct{ *Resolver }

// Organizer is the resolver for the organizer field.
// func (r *meetupResolver) Organizer(ctx context.Context, obj *models.Meetup) (*models.User, error) {
// 	fmt.Println("OrganizerID:", obj.UserID)
// 	// return GetUserLoader(ctx).Load(obj.OrganizerID)
// 	return r.UserRepo.GetUserById(obj.UserID)
// }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	// return r.UserRepo.GetUserById(obj.UserID)
	return GetUserLoader(ctx).Load(obj.UserID)
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}
