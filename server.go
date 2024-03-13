package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/tabed23/social-media-api/config"
	"github.com/tabed23/social-media-api/graph"
	"github.com/tabed23/social-media-api/internal/store"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.StartMigration(); err != nil {
		log.Fatal(err)
	}
	userLayer := store.NewUserStore(db.GetDB())
	commnetLayer := store.NewCommentStore(db.GetDB())
	likeLayer := store.NewLikeStore(db.GetDB())
	postLayer := store.NewPostStore(db.GetDB())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserRepository:    userLayer,
		CommentRepository: commnetLayer,
		LikeRepository:    likeLayer,
		PostRepository:    postLayer,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
