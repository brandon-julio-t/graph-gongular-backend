package chi_router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	resolverFactory "github.com/brandon-julio-t/graph-gongular-backend/factories/resolver"
	"github.com/brandon-julio-t/graph-gongular-backend/graph"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/brandon-julio-t/graph-gongular-backend/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const graphqlEndpoint = "/graphql"

type Factory struct{}

func (*Factory) Create(secret []byte, db *gorm.DB) *chi.Mux {
	resolver := new(resolverFactory.Factory).Create(secret, db)
	router := setupRouterWithMiddlewares(resolver)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
	)

	router.Handle("/", playground.Handler("GraphQL playground", graphqlEndpoint))
	router.Handle(graphqlEndpoint, srv)

	return router
}

func setupRouterWithMiddlewares(resolver *graph.Resolver) *chi.Mux {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(setupCorsHandler())
	router.Use(middlewares.JwtRenewMiddleware(resolver.JwtService))
	router.Use(middlewares.CookieWriterProviderMiddleware())

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// Inject services one by one to prevent circular import
	router.Use(middlewares.AuthMiddleware(resolver.JwtService, resolver.UserService))

	return router
}

func setupCorsHandler() func(next http.Handler) http.Handler {
	return cors.New(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:8080",
				"http://localhost:4200",
				"https://graph-gongular-frontend.netlify.app",
			},
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
		},
	).Handler
}
