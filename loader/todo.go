package loader

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/shion0625/gqlgen-todos/graph/model"
	"gorm.io/gorm"
	"log"
)

// TodoLoader はデータベースからtodoを読み取ります
type TodoLoader struct {
	DB *gorm.DB
}

// BatchGetTodos は、ID によって多くのtodoを取得できるバッチ関数を実装します。
func (u *TodoLoader) BatchGetTodos(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// 単一のクエリで要求されたすべてのtodoを読み取ります
	userIDs := make([]string, len(keys))

	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	todosTemp := []*model.Todo{}
	if err := u.DB.Debug().Where("user_id IN ?", userIDs).Find(&todosTemp).Error; err != nil {
		err := fmt.Errorf("fail get todos, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	todoByUserId := map[string][]*model.Todo{}
	for _, todo := range todosTemp {
		todoByUserId[todo.UserId] = append(todoByUserId[todo.UserId], todo)
	}

	todos := make([][]*model.Todo, len(userIDs))

	for i, id := range userIDs {
		todos[i] = todoByUserId[id]
	}

	output := make([]*dataloader.Result, len(todos))
	for index := range todos {
		todo := todos[index]
		output[index] = &dataloader.Result{Data: todo, Error: nil}
	}
	return output
}

// dataloader.Loadをwrapして型づけした実装
func LoadTodo(ctx context.Context, todoID string) ([]*model.Todo, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.TodoLoader.Load(ctx, dataloader.StringKey(todoID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.([]*model.Todo), nil
}
