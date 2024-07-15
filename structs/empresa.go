package structs

import "time"

type Empresa struct {
	ID                  string     `json:"id"`
	Codigo              string     `json:"codigo"`
	Nome                string     `json:"nome"`
	RazaoSocial         string     `json:"razao_social"`
	CNPJBase            string     `json:"cnpj_base"`
	USUCodigo           *string    `json:"usu_codigo,omitempty"`
	ConvertTipoHe       int        `json:"convert_tipo_he"`
	CPF                 string     `json:"cpf"`
	DataEncerramento    *time.Time `json:"dt_encerramento,omitempty"`
	UltimaAtualizacaoAC *time.Time `json:"ultima_atualizacao_ac,omitempty"`
	FaltaAjustarNoAC    int        `json:"falta_ajustar_no_ac"`
	AderiuESocial       int        `json:"aderiu_esocial"`
	DataAdesaoESocial   *time.Time `json:"data_adesao_esocial,omitempty"`
	DataAdesaoESocialF2 *time.Time `json:"data_adesao_esocial_f2,omitempty"`
	TpAmbESocial        int        `json:"tp_amb_esocial"`
	StatusEnvioApp      int        `json:"status_envio_app"`
	NomeFantasia        string     `json:"nmfantasia"`
	CNPJLicenciado      string     `json:"cnpj_licenciado"`
	FreemiumLastUpdate  string     `json:"freemium_last_update"`
}
