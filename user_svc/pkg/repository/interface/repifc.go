package RepoIfc

import (
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/domain"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/models"
)

type Repo interface {
	CheckUserExistsEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	FindUserByPhone(phone string) (*domain.User, error)
	AddAddress(Address models.Address, UserId uint) (models.AddressRes, error)
	ViewAddress(id uint) ([]models.AddressRes, error)
}
