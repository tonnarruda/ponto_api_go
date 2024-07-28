package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostAprovaAbono(t *testing.T) {

	testCases := []struct {
		description  string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Buscar Usuarios com Sucesso",
			expected:     http.StatusOK,
			expectedDesc: "Codigo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := SetupApi()

			resp, err := api.Client.R().
				Get("/usuario")

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
