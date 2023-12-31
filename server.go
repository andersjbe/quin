package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andersjbe/quin/graph"
	"github.com/andersjbe/quin/internal/auth"
	"github.com/andersjbe/quin/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	conn, err := sql.Open(
		"postgres", 
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=verify-full", 
			os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE")
		))
	if err != nil {
		log.Fatal("Could not establish database connection", err)
	}
	db := database.New(conn)

	router.Use(auth.Middleware(db))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db, Conn: conn}}))

	router.Get("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	http.ListenAndServe(":"+port, router)
}
