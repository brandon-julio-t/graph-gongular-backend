package graphql_server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/brandon-julio-t/graph-gongular-backend/graph"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
	"github.com/gorilla/websocket"
	"net/http"
)

type Factory struct{}

func (*Factory) Create(resolver *graph.Resolver) *handler.Server {
	server := handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
	)

	server.AddTransport(transport.POST{})
	server.AddTransport(transport.Options{})
	server.Use(extension.Introspection{})

	const mb int64 = 1 << 20
	server.AddTransport(transport.MultipartForm{
		MaxUploadSize: 100 * mb,
		MaxMemory:     50 * mb,
	})

	server.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 0,
	})

	return server
}
