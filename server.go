package main

import (
	"crypto/rand"
	"github.com/brandon-julio-t/graph-gongular-backend/factories"
	"github.com/brandon-julio-t/graph-gongular-backend/graph"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"github.com/brandon-julio-t/graph-gongular-backend/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
)

const defaultPort = "8080"
const graphqlEndpoint = "/graphql"

var secret []byte

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file")
	}

	secret = []byte(os.Getenv("APP_KEY"))
	if len(secret) == 0 {
		secret = makeSecret()
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := setupDatabase()
	router := setupRouter(db)
	runServer(port, router)
}

func makeSecret() []byte {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Error while generating APP_KEY")
	}

	return key
}

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.UserRole{}, &model.User{}); err != nil {
		panic("Error while auto migrating User")
	}

	//seedDatabase(db)

	return db
}

func seedDatabase(db *gorm.DB) {
	adminRole := &model.UserRole{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Name: "Admin",
	}

	userRole := &model.UserRole{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Name: "User",
	}

	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	adminDOB, _ := time.Parse("2006-01-02", "1970-01-01")

	userHash, _ := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	userDOB, _ := time.Parse("2006-01-02", "1970-01-01")

	adminUser := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Admin",
		Email:       "admin@admin.com",
		Password:    string(adminHash),
		DateOfBirth: adminDOB,
		Gender:      "Male",
		Address:     "Admin Address",
		UserRole:    adminRole,
	}

	regularUser := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "User",
		Email:       "user@user.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "User Address",
		UserRole:    userRole,
	}

	db.Create(adminUser)
	db.Create(regularUser)
}

func setupRouter(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()

	resolvers := &graph.Resolver{
		UserService: &services.UserService{
			UserRepository:     &repository.UserRepository{DB: db},
			UserRoleRepository: &repository.UserRoleRepository{DB: db},
		},
		JwtService: &services.JwtService{
			Secret:           secret,
			JwtCookieFactory: new(factories.JwtCookieFactory),
		},
	}

	setupMiddlewares(router, resolvers)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolvers,
			},
		),
	)

	router.Handle("/", playground.Handler("GraphQL playground", graphqlEndpoint))
	router.Handle(graphqlEndpoint, srv)

	return router
}

func setupMiddlewares(router *chi.Mux, resolvers *graph.Resolver) {
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", "http://localhost:4200", "https://graph-gongular-frontend.netlify.app"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(middlewares.CookieWriterProviderMiddleware())

	// Inject services one by one to prevent circular import
	router.Use(middlewares.AuthMiddleware(resolvers.JwtService, resolvers.UserService))
}

func runServer(port string, router *chi.Mux) {
	log.Printf("Starting app with JWT secret: %v\n", secret)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
