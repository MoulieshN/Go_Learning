package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type User struct {
	ID   string
	Name string
	Age  int
}

var users = []User{
	{
		ID:   "1",
		Name: "Rithin",
		Age:  18,
	},
	{
		ID:   "2",
		Name: "Bob",
		Age:  30,
	},
	{
		ID:   "3",
		Name: "Alice",
		Age:  25,
	},
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return users, nil
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(string)
				for _, user := range users {
					if user.ID == id {
						return user, nil
					}
				}
				return nil, fmt.Errorf("no user found with id %s", id)
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: addUser,
		},
	},
})

func addUser(p graphql.ResolveParams) (any, error) {
	name, _ := p.Args["name"].(string)
	age, _ := p.Args["age"].(int)

	totalUsers := len(users) + 1

	newUser := User{
		ID:   strconv.Itoa(totalUsers),
		Name: name,
		Age:  age,
	}
	users = append(users, newUser)
	return users, nil
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Graphql server has been started......")
	http.ListenAndServe(":3035", nil)
}
