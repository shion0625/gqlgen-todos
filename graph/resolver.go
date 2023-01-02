package graph

import "github.com/shion0625/gqlgen-todos/graph/model"

//go:generate go run github.com/99designs/gqlgen generate //追加
//今後schema.graphqlsを変更した際に go generate ./...で更新することができるようになる。

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo//追加
}
