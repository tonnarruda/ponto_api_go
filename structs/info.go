package structs

type Info struct {
	VersaoBD       int    `json:"versao_bd"`
	StatusBD       string `json:"status_bd"`
	Sistema        string `json:"sistema"`
	VersaoBDBeta   int    `json:"versao_bd_beta"`
	Atualizando    int    `json:"atualizando"`
	Fortes         int    `json:"fortes"`
	ConvertePonto3 int    `json:"converte_ponto3"`
}
