package repositories

import (
	"log"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/infra"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) GetLogins() (logins []*domain.LoginData, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	var entityLogins []entities.Login
	res := db.Find(&entityLogins)
	if res.Error != nil {
		return nil, res.Error
	}
	logins = make([]*domain.LoginData, 0)
	toDeleteLogins := make([]int, 0)
	for _, login := range entityLogins {
		logins = append(logins, &domain.LoginData{
			UserName:      login.UserName,
			UserId:        login.MappaUserId,
			PasswordHash:  login.PasswordHash,
			LastLogin:     login.LastLogin,
			Authorization: login.MappaAuth,
			ValidUntil:    login.MappaValidUntil,
		})
	}
	if len(toDeleteLogins) > 0 {
		res := db.Delete(&entities.Login{}, "mappa_user_id IN (?)", toDeleteLogins)

		if res.Error != nil {
			log.Printf("Error deleting invalid logins: %s", res.Error)
		} else {
			log.Printf("Deleted %d invalid logins", res.RowsAffected)
		}
	}
	return logins, err
}

func (r *DBRepository) SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error {
	r.DBLock()
	defer r.DBUnlock()

	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities.Login{
		UserName:        username,
		PasswordHash:    infra.GetPasswordHash(password),
		LastLogin:       last_login,
		MappaUserId:     loginResponse.Userid,
		MappaAuth:       loginResponse.ID,
		MappaValidUntil: loginResponse.ValidUntil(),
	})

	return res.Error
}

func (r *DBRepository) DeleteLogin(username string) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	deletedLogin := entities.Login{
		UserName: username,
	}
	res := db.Delete(&deletedLogin)
	return res.Error
}
