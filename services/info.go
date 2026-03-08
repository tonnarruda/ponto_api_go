package services

// PodeTirarCNH verifica se a pessoa pode tirar a CNH. Retorna true se tiver 18 anos ou mais.
func PodeTirarCNH(idade int) bool {
	return idade >= 18
}
