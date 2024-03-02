package repository

import (
	"errors"

	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/domain"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/models"
	RepoIfc "github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) RepoIfc.Repo {
	return &UserRepo{DB: db}
}

func (clean *UserRepo) CheckUserExistsEmail(email string) (*domain.User, error) {
	var user domain.User
	result := clean.DB.Where(&domain.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil

}

func (clean *UserRepo) CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := clean.DB.Where(&domain.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (clean *UserRepo) FindUserByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := clean.DB.Raw("SELECT * FROM users WHERE phone = ?", phone).Scan(&user)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}
	return &user, nil
}

func (clean *UserRepo) AddAddress(Address models.Address, UserId uint) (models.AddressRes, error) {
	var AddressRes models.AddressRes
	query := clean.DB.Raw(`INSERT INTO addresses(user_id,name,house_name,phone,street,city,state,pin) VALUES (?,?,?,?,?,?,?,?) RETURNING id,name,house_name,phone,street,city,state,pin`, UserId, Address.Name, Address.Housename, Address.Phone, Address.Street, Address.City, Address.State, Address.Pin).Scan(&AddressRes)
	if query.Error != nil {
		return models.AddressRes{}, query.Error
	}
	return AddressRes, nil
}

func (clean *UserRepo) ViewAddress(id uint) ([]models.AddressRes, error) {
	var Address []models.AddressRes
	query := clean.DB.Raw(`SELECT * FROM addresses WHERE user_id = ?`, id).Scan(&Address)
	if query.Error != nil {
		return []models.AddressRes{}, query.Error
	}

	if query.RowsAffected < 1 {
		return []models.AddressRes{}, errors.New("no address found. add new address")
	}

	return Address, nil
}
