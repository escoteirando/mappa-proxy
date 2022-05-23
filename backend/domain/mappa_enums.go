package domain

import "strings"

type MappaRamoEnum uint8

const (
	RamoLobinho   MappaRamoEnum = 1
	RamoEscoteiro MappaRamoEnum = 2
	RamoSenior    MappaRamoEnum = 3
	RamoPioneiro  MappaRamoEnum = 4
	RamoInvalido  MappaRamoEnum = 0
)

func (enum MappaRamoEnum) String() string {
	switch enum {
	case RamoLobinho:
		return "Lobinho"
	case RamoEscoteiro:
		return "Escoteiro"
	case RamoSenior:
		return "Senior"
	case RamoPioneiro:
		return "Pioneiro"
	default:
		return "Inválido"
	}
}

func ParseRamo(ramo string) MappaRamoEnum {
	switch strings.ToLower(ramo)[0:1] {
	case "l", "1":
		return RamoLobinho
	case "e", "2":
		return RamoEscoteiro
	case "s", "3":
		return RamoSenior
	case "p", "4":
		return RamoPioneiro
	default:
		return RamoInvalido
	}
}

func CaminhoToRamo(caminho int) MappaRamoEnum {
	switch caminho {
	case 1, 2, 3:
		return RamoLobinho
	case 4, 5, 6:
		return RamoEscoteiro
	case 11, 12:
		return RamoSenior
	case 15, 16:
		return RamoPioneiro
	default:
		return RamoInvalido
	}
}

func (enum MappaRamoEnum) Caminhos() []int {
	switch enum {
	case RamoLobinho:
		return []int{1, 2, 3}
	case RamoEscoteiro:
		return []int{4, 5, 6}
	case RamoSenior:
		return []int{11, 12}
	case RamoPioneiro:
		return []int{15, 16}
	default:
		return []int{}
	}
}

func (enum MappaRamoEnum) ToCode() string {
	switch enum {
	case RamoLobinho:
		return "L"
	case RamoEscoteiro:
		return "E"
	case RamoSenior:
		return "S"
	case RamoPioneiro:
		return "P"
	default:
		return ""
	}
}
