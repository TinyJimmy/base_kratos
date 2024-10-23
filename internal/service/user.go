package service

import (
	"context"
	"fmt"

	pb "test_kratos/api/helloworld/v1"
	"test_kratos/internal/biz"
	"test_kratos/internal/zlog"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	name := req.Name
	age := req.Age
	email := req.Email
	phone := req.Phone
	fmt.Println("name:", name, "age:", age, "email:", email, "phone:", phone)
	code := "0000"
	msg := "success"
	s.uc.Create(ctx, &biz.User{
		Name:  name,
		Age:   int(age),
		Email: email,
		Phone: phone,
	})
	zlog.ServiceLogger.Info("创建用户", zap.String("name", name), zap.Int32("age", int32(age)), zap.String("email", email), zap.String("phone", phone))
	return &pb.CreateUserReply{Code: code, Msg: msg}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	a, _ := s.uc.Get(ctx, req.UserId)
	fmt.Println("a:", a)
	zlog.ServiceLogger.Info("创建用户", zap.String("name", a.Name), zap.Int32("age", int32(a.Age)), zap.String("email", a.Email), zap.String("phone",
		a.Phone))
	zlog.RouterLogger.Info("创建用户", zap.String("name", a.Name), zap.Int32("age", int32(a.Age)), zap.String("email", a.Email), zap.String("phone",
		a.Phone))
	log.Info("创建用户", "name", a.Name, "age", a.Age, "email", a.Email, "phone", a.Phone)
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
