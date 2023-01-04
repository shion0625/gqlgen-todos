package loader

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/shion0625/gqlgen-todos/graph/model"
	"gorm.io/gorm"
	"log"
)

// UserLoader はデータベースからユーザーを読み取ります
type UserLoader struct {
	DB *gorm.DB
}

// BatchGetUsers は、ID によって多くのユーザーを取得できるバッチ関数を実装します。
func (u *UserLoader) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// 単一のクエリで要求されたすべてのユーザーを読み取ります
	userIDs := make([]string, len(keys))
	users := []*model.User{}
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	if err := u.DB.Debug().Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		err := fmt.Errorf("fail get users, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index := range users {
		user := users[index]
		output[index] = &dataloader.Result{Data: user, Error: nil}
	}
	return output
}

// dataloader.Loadをwrapして型づけした実装
func LoadUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
}
