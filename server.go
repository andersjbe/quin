package main

import (
	"database/sql"
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

	_ "github.com/go-sql-driver/mysql"
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

	conn, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
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
