package data

import (
	"context"
	"fmt"
	"test_kratos/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}

func (u *User) ToBizUser() *biz.User {
	return &biz.User{
		ID:    u.ID,
		Name:  u.Name,
		Age:   u.Age,
		Email: u.Email,
		Phone: u.Phone,
	}
}

func (u *User) FromBizUser(user *biz.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
		Phone: user.Phone,
	}
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *userRepo) GetUser(ctx context.Context, id string) (*biz.User, error) {
	userQuery := ar.data.mongo.Database("test_go").Collection("user").FindOne(ctx, bson.M{"_id": id})
	var result User
	err := userQuery.Decode(&result)
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("查找结果为: %v\n", result)
	return result.ToBizUser(), nil
}

func (ar *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	data_user := &User{}
	data_user = data_user.FromBizUser(user)
	data_user.ID = uuid.New().String()
	_, err := ar.data.mongo.Database("test_go").Collection("user").
		InsertOne(ctx, data_user)
	return err
}

func (ar *userRepo) UpdateUser(ctx context.Context, id string, user *biz.User) error {
	_, err := ar.data.mongo.Database("test_go").Collection("user").
		UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return err
}

func (ar *userRepo) DeleteUser(ctx context.Context, id string) error {
	_, err := ar.data.mongo.Database("test_go").Collection("user").
		DeleteOne(ctx, bson.M{"id": id})
	return err
}
