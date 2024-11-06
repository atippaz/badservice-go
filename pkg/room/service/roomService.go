package service

import (
	_AdminModel "github.com/atippaz/isekai-shop/pkg/admin/model"
)

type AdminService interface {
	Insert(itemFilter *_AdminModel.AdminInsertRequest) (*_AdminModel.AdminResult, error)
	FindById(AdminId string) (*_AdminModel.AdminResult, error)
}
