package handler

import (
	"eff_mob_test/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GraphqlHandler() gin.HandlerFunc {

	han := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		han.ServeHTTP(c.Writer, c.Request)
	}

}

func (h *Handler) PlaygroundHandler() gin.HandlerFunc {
	han := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		han.ServeHTTP(c.Writer, c.Request)
	}
}
