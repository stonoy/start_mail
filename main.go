package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stonoy/start_mail/internal/database"
)

type apiConfig struct {
	jwtSecret string
	db        *database.Queries
}

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can not load env variables %v", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("no port env set %v", err)
		return
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Fatalf("no jwt_secret env set %v", err)
		return
	}

	apiCfg := &apiConfig{
		jwtSecret: jwt_secret,
	}

	if conn := os.Getenv("CONN"); conn != "" {
		dbConn, err := sql.Open("postgres", conn)
		if err != nil {
			log.Fatalf("can establish a connection with database -> %v", err)
			return
		}

		dbQ := database.New(dbConn)

		apiCfg.db = dbQ
	} else {
		log.Println("server stated without db connection")
	}

	// create a new router
	r := chi.NewRouter()

	// basic cors
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// created a sub router for api
	apiRouter := chi.NewRouter()

	// user
	apiRouter.Post("/register", apiCfg.register)
	apiRouter.Post("/login", apiCfg.login)

	// emails
	apiRouter.Post("/createemails", apiCfg.checkUserMiddleware(apiCfg.createEmail))
	apiRouter.Get("/inboxemails", apiCfg.checkUserMiddleware(apiCfg.Inbox))
	apiRouter.Get("/sentboxemails", apiCfg.checkUserMiddleware(apiCfg.SentBox))
	apiRouter.Get("/getemail/{emailID}", apiCfg.checkUserMiddleware(apiCfg.getSingleMail))
	apiRouter.Get("/getmailboxnums", apiCfg.checkUserMiddleware(apiCfg.getMAilBoxNums))

	// favourite
	apiRouter.Get("/createFav/{emailID}", apiCfg.checkUserMiddleware(apiCfg.createFav))
	apiRouter.Get("/getallfavuser", apiCfg.checkUserMiddleware(apiCfg.getAllFav))
	apiRouter.Delete("/deletefav/{favID}", apiCfg.checkUserMiddleware(apiCfg.deleteFav))

	// mount
	r.Mount("/api/v1", apiRouter)

	the_server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("server is listening on port -> %v", port)

	log.Fatal(the_server.ListenAndServe())
}
