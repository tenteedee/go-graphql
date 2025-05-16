//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/tenteedee/go-graphql/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo   *postgres.UserRepo
	MeetupRepo *postgres.MeetupRepo
}
