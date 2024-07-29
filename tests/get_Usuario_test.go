package tests

import (
	"encoding/json"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsuarios(t *testing.T) {

	testCases := []struct {
		description string
		expected    int
	}{
		{
			description: "Buscar Usuarios com Sucesso",
			expected:    http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := SetupApi()

			resp, err := api.Client.R().
				Get("/usuario")

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)

			// Validar o contrato da resposta
			var usuarios []map[string]interface{}
			err = json.Unmarshal(resp.Body(), &usuarios)
			assert.NoError(t, err, "Erro ao deserializar resposta para %s", tc.description)

			for _, usuario := range usuarios {
				// Validar que o campo "id" está presente, é uma string, segue o formato de UUID e não está vazio
				assert.Contains(t, usuario, "id", "Campo 'id' ausente na resposta para %s", tc.description)
				assert.IsType(t, "", usuario["id"], "Campo 'id' deve ser uma string para %s", tc.description)
				assert.Regexp(t, regexp.MustCompile("^[a-f0-9-]{36}$"), usuario["id"], "ID inválido para %s", tc.description)
				assert.NotEmpty(t, usuario["id"], "Campo 'id' não deve ser vazio para %s", tc.description)

				// Validar que o campo "codigo" está presente, é uma string e não está vazio
				assert.Contains(t, usuario, "codigo", "Campo 'codigo' ausente na resposta para %s", tc.description)
				assert.IsType(t, "", usuario["codigo"], "Campo 'codigo' deve ser uma string para %s", tc.description)
				assert.NotEmpty(t, usuario["codigo"], "Campo 'codigo' não deve ser vazio para %s", tc.description)
			}
		})
	}
}
