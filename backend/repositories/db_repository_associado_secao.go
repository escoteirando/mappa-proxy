package repositories

import (
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
)

func (r *DBRepository) setAssociadoSecao(codigoAssociado int, subSecao bool, codsSecao ...int) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()

	// Remove
	res := db.Delete(&entities.AssociadoSecao{}, "cod_associado = ? and subsecao = ?", codigoAssociado, subSecao)
	log.Printf("Delete sessao para associado %d: %d (%v)", codigoAssociado, res.RowsAffected, res.Error)

	// Add
	for _, codSecao := range codsSecao {
		res := db.Create(&entities.AssociadoSecao{
			CodigoAssociado: codigoAssociado,
			CodigoSecao:     codSecao,
			SubSecao:        subSecao,
		})
		if res.Error != nil {
			return res.Error
		}
		log.Printf("Create sessao para associado %d (%d): %d (%v)", codigoAssociado, codSecao, res.RowsAffected, res.Error)
	}
	return nil
}

func (r *DBRepository) getAssociadoSecoes(codigoAssociado int, subSecao bool) (secoesIds []int, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()

	res := db.Model(&entities.AssociadoSecao{}).Select("cod_secao").Where("cod_associado = ? and subsecao = ?", codigoAssociado, subSecao).Find(&secoesIds)
	err = res.Error
	return
}

func (r *DBRepository) SetAssociadoSecao(codigoAssociado int, codigosSubSecao ...int) error {
	return r.setAssociadoSecao(codigoAssociado, false, codigosSubSecao...)
}

func (r *DBRepository) GetAssociadoSecoes(codigoAssociado int) (secoes []*entities.Secao, err error) {
	ids, err := r.getAssociadoSecoes(codigoAssociado, false)
	if err != nil || len(ids) == 0 {
		return
	}

	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()

	res := db.Where("codigo in ?", ids).Find(&secoes)
	err = res.Error
	return
}

func (r *DBRepository) SetAssociadoSubSecao(codigoAssociado int, codigoSubSecao int) error {
	return r.setAssociadoSecao(codigoAssociado, true, codigoSubSecao)
}

func (r *DBRepository) GetAssociadoSubSecao(codigoAssociado int) (subsecao *entities.SubSecao, err error) {
	ids, err := r.getAssociadoSecoes(codigoAssociado, true)
	if err != nil || len(ids) == 0 {
		return
	}

	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()

	res := db.Where("codigo = ?", ids[0]).First(&subsecao)
	err = res.Error
	return
}
