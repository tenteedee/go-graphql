package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/tenteedee/go-graphql/models"
)

type contextKey string

const userloaderKey = contextKey("userloader")

func DataloaderMiddleware(DB *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				err := DB.Model(&users).Where("id IN (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				userMap := make(map[string]*models.User)
				for _, user := range users {
					userMap[user.ID] = user
				}

				result := make([]*models.User, len(ids))
				for i, id := range ids {
					if user, ok := userMap[id]; ok {
						result[i] = user
					}
				}

				return result, nil
			},
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userloaderKey, &userloader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	loader, ok := ctx.Value(userloaderKey).(*UserLoader)
	if !ok {
		return nil
	}
	return loader
}
