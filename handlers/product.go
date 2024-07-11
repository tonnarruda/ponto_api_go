package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// GetProdutos retorna todos os produtos do banco de dados
func GetProdutos(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, nome FROM produtos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}
		defer rows.Close()

		var produtos []Produto
		for rows.Next() {
			var produto Produto
			if err := rows.Scan(&produto.ID, &produto.Nome); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product"})
				return
			}
			produtos = append(produtos, produto)
		}
		if err = rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during iteration"})
			return
		}

		c.JSON(http.StatusOK, produtos)
	}
}

func GetProdutoByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var produto Produto
		err = db.QueryRow("SELECT id, nome FROM produtos WHERE id = $1", id).Scan(&produto.ID, &produto.Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			}
			return
		}

		c.JSON(http.StatusOK, produto)
	}
}

// CreateProduto cria um novo produto no banco de dados
func CreateProduto(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produto Produto
		if err := c.ShouldBindJSON(&produto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}

		stmt, err := tx.Prepare("INSERT INTO produtos (nome) VALUES ($1) RETURNING id")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
			return
		}
		defer stmt.Close()

		err = stmt.QueryRow(produto.Nome).Scan(&produto.ID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute statement"})
			return
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Produto criado com sucesso", "produto": produto})
	}
}

func DeleteProdutoByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var produto Produto
		err = db.QueryRow("DELETE FROM produtos WHERE id = $1 RETURNING id, nome", id).Scan(&produto.ID, &produto.Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "produto": produto})
	}
}
