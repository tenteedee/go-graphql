package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/tenteedee/go-graphql/env"
	"github.com/tenteedee/go-graphql/graph"
	"github.com/tenteedee/go-graphql/postgres"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	env.Init()

	DB := postgres.New(&pg.Options{
		User:     env.DB_USER,
		Password: env.DB_PASSWORD,
		Addr:     env.DB_ADDR,
		Database: env.DB_NAME,
	})

	defer func() {
		if err := DB.Close(); err != nil {
			log.Fatalf("failed to close database connection: %v", err)
		}
	}()

	DB.AddQueryHook(&postgres.DBLogger{})

	port := env.PORT
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				UserRepo:   &postgres.UserRepo{DB: DB},
				MeetupRepo: &postgres.MeetupRepo{DB: DB},
			}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
