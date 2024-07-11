package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/produtos", GetProdutos(db))
	router.POST("/produtos", CreateProduto(db))
	return router
}

func TestGetProdutos(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nome"}).
		AddRow(1, "Produto 1").
		AddRow(2, "Produto 2")

	mock.ExpectQuery("SELECT id, nome FROM produtos").WillReturnRows(rows)

	router := setupRouter(db)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/produtos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var produtos []Produto
	err = json.Unmarshal(w.Body.Bytes(), &produtos)
	assert.NoError(t, err)
	assert.Len(t, produtos, 2)
	assert.Equal(t, "Produto 1", produtos[0].Nome)
	assert.Equal(t, "Produto 2", produtos[1].Nome)
}

func TestCreateProduto(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO produtos").
		ExpectQuery().
		WithArgs("Produto Teste").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	router := setupRouter(db)
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"nome":"Produto Teste"}`)
	req, _ := http.NewRequest("POST", "/produtos", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Produto criado com sucesso", response["message"])
	assert.Equal(t, float64(1), response["produto"].(map[string]interface{})["id"])
	assert.Equal(t, "Produto Teste", response["produto"].(map[string]interface{})["nome"])
}
