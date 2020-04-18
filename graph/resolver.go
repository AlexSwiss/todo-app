//go:generate go run github.com/99designs/gqlgen

package graph


// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/AlexSwiss/todo/graph/model"


type Resolver struct{
	todos []*model.Todo
}
