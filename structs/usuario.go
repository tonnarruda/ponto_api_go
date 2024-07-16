package structs

import "time"

type Usuario struct {
	ID                   string    `json:"id"`
	Codigo               string    `json:"codigo"`
	Senha                int       `json:"senha"`
	UltimoAcesso         time.Time `json:"ultimo_acesso"`
	Bloqueado            int       `json:"bloqueado"`
	UserRegistrationDate time.Time `json:"user_registration_date"`
	LimiteEpgData        time.Time `json:"limite_epg_data"`
}
