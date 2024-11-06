package service

import (
	"github.com/atippaz/isekai-shop/entities"
	_adminModel "github.com/atippaz/isekai-shop/pkg/admin/model"
	_adminRepository "github.com/atippaz/isekai-shop/pkg/admin/repository"
)

type adminServiceImpl struct {
	adminRepository _adminRepository.AdminRepository
}

func NewAdminServiceImpl(adminRepository _adminRepository.AdminRepository) AdminService {
	return &adminServiceImpl{
		adminRepository: adminRepository,
	}
}

func (r *adminServiceImpl) Insert(adminRequest *_adminModel.AdminInsertRequest) (*_adminModel.AdminResult, error) {
	newAdmin := entities.Admin{
		Email:  adminRequest.Email,
		Avatar: adminRequest.Avatar,
		Name:   adminRequest.Name,
	}
	result, err := r.adminRepository.Insert(newAdmin)
	if err != nil {
		return nil, err
	}
	return &_adminModel.AdminResult{
		ID:     result.ID,
		Email:  result.Email,
		Name:   result.Name,
		Avatar: result.Avatar,
	}, nil
}
func (r *adminServiceImpl) FindById(adminId string) (*_adminModel.AdminResult, error) {
	result, error := r.adminRepository.FindById(adminId)
	if error != nil {
		return nil, error
	}
	return &_adminModel.AdminResult{
		ID:     result.ID,
		Email:  result.Email,
		Name:   result.Name,
		Avatar: result.Avatar,
	}, nil
}
