package todo

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type AddTaskRequest struct {
	Task string `json:"task"`
}

var db *sql.DB

func SetDB(_db *sql.DB) {
	db = _db
}

var hmacSampleSecret = []byte("drowssap")

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODg3MTg3NDF9.xx1ervpDARQpVXf8w9EwheeQgcUYpzAAz8fofto8FKM
func AddTaskHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")[7:]
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	var task AddTaskRequest
	if err = c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = db.Exec("insert into todos(task) values(?)", task.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

type TodoHandler struct {
	db *sql.DB
}

func NewTodoHandler(db *sql.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) AddTask(c *gin.Context) {
	var task AddTaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err := t.db.Exec("insert into todos(task) values(?)", task.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
