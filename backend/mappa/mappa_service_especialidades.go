package mappa

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/dtos"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/tools"
)

func (svc *MappaService) GetEspecialidades() (especialidades []*responses.MappaEspecialidadeResponse, err error) {
	keyMarcacoes := fmt.Sprintf("especialidades")
	lastFetch := svc.Cache.GetLastEventTime(keyMarcacoes)
	if lastFetch.Before(tools.DaysAgo(PROGRESSOES_UPDATE_INTERVAL)) {
		especialidades, err = svc.API.GetEspecialidades()
		if err == nil {
			if especialidades == nil || len(especialidades) == 0 {
				err = fmt.Errorf("No especialidades found")
			}
		}
		if err != nil {
			return
		}
		if err = svc.updateEspecialidades(especialidades); err == nil {
			svc.Cache.SetLastEventTime(keyMarcacoes, time.Now())
		}
	} else {
		var objEsp []*entities.MappaEspecialidade
		if objEsp, err = svc.Repository.GetEspecialidades(); err != nil {
			return
		}
		especialidades = make([]*responses.MappaEspecialidadeResponse, len(objEsp))
		for i, esp := range objEsp {
			especialidades[i] = dtos.MappaEspecialidadeToResponse(esp)
		}
	}
	return
}

// 		objMarcacoes, err := svc.API.GetMarcacoes(codigoSecao, lastFetch, authorization)
// 		if err == nil && objMarcacoes != nil && objMarcacoes.Values != nil && len(objMarcacoes.Values) > 0 {
// 			especialidades = objMarcacoes.Values
// 			if err = svc.updateMarcacoes(codigoSecao, especialidades); err == nil {
// 				svc.Cache.SetLastEventTime(keyMarcacoes, time.Now())
// 			}
// 		}

// 		err = svc.updateMappaProgressoes(svc.Repository)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		objMarcacoes, err := svc.Repository.GetMarcacoes(codigoSecao)
// 		if err == nil && objMarcacoes != nil && len(objMarcacoes) > 0 {
// 			especialidades = make([]*responses.MappaMarcacaoResponse, len(objMarcacoes))
// 			for i, marcacao := range objMarcacoes {
// 				especialidades[i] = dtos.MappaMarcacaoToResponse(marcacao)
// 			}
// 		}

// 	}
// 	return
// }

func (svc *MappaService) updateEspecialidades(especialidades []*responses.MappaEspecialidadeResponse) error {
	eEspecialidades := make([]*entities.MappaEspecialidade, len(especialidades))
	for i, especialidade := range especialidades {
		eEspecialidades[i] = dtos.MappaEspecialidadeToEntity(especialidade)
	}
	return svc.Repository.UpdateMappaEspecialidades(eEspecialidades)
}
