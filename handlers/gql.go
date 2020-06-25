package handlers

import (
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gin-gonic/gin"
    "gqlgen-todos/graph/generated"
    "gqlgen-todos/graph/resolvers"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
    // NewExecutableSchema and Config are in the generated.go file
    // Resolver is in the resolver.go file
    h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

// Defining the Playground handler
func PlaygroundHandler(path string) gin.HandlerFunc {
    
    h := playground.Handler("GraphQL", path)

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}