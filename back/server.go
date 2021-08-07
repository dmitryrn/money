package main

import (
	context "context"
	"github.com/dmitryrn/money/proto"
)

type Server struct {
	proto.UnimplementedBudgetingAppServer
}

func (s Server) GetAccounts(ctx context.Context, request *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	panic("implement me")
}

func (s Server) GetCategories(ctx context.Context, request *proto.GetCategoriesRequest) (*proto.GetCategoriesResponse, error) {
	panic("implement me")
}

func (s Server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	panic("implement me")
}

func (s Server) CreateCategory(ctx context.Context, request *proto.CreateCategoryRequest) (*proto.CreateCategoryResponse, error) {
	panic("implement me")
}

func (s Server) CreateTransaction(ctx context.Context, request *proto.CreateTransactionRequest) (*proto.CreateTransactionResponse, error) {
	panic("implement me")
}
