package repositories

import (
	"encoding/json"
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error {
	detalheEscotistaJson, err := json.Marshal(detalheEscotista)
	if err != nil {
		return err
	}
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities.DetalhesEscotista{
		UserId:   userId,
		Detalhes: string(detalheEscotistaJson),
	})

	return res.Error
}
func (r *DBRepository) GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error) {
	r.DBLock()
	defer r.DBUnlock()
	detalheEscotista := entities.DetalhesEscotista{
		UserId: userId,
	}
	db := r.getDBFunc()
	res := db.First(&detalheEscotista, userId)
	if res.Error != nil {
		return nil, res.Error
	}
	var detalhe responses.MappaDetalhesResponse
	if err := json.Unmarshal([]byte(detalheEscotista.Detalhes), &detalhe); err != nil {
		log.Printf("Erro ao deserializar detalhe escotista: %s %s", err, detalheEscotista.Detalhes)
		return nil, err
	}
	return &detalhe, nil
}

func (r *DBRepository) SetEscotista(escotista *entities.Escotista) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(escotista)
	return res.Error
}

func (r *DBRepository) GetEscotista(userId int) (escotista *entities.Escotista, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.First(&escotista, userId)
	err = res.Error
	return
}
