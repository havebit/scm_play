package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"myapi/todo"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	_ "modernc.org/sqlite"
)

var hmacSampleSecret = []byte("drowssap")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	os.Remove("./todo.sqlite")

	db, err := sql.Open("sqlite", "./todo.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sqlStmt := `
	// create table todos (id integer primary key autoincrement, task text, done integer);
	// delete from todos;
	// `
	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	r := gin.Default()
	r.POST("/ping", pingpongHandler)
	r.GET("/ping/:id/:foo", getPingPongHandler)
	r.GET("/login", func(c *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	todo.SetDB(db)
	r.POST("/todos", todo.AddTaskHandler)

	// todoHandler := todo.NewTodoHandler(db)
	// r.POST("/todos", todoHandler.AddTask)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func getPingPongHandler(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"id":      c.Param("id"),
		"foo":     c.Param("foo"),
	})
}

type PingpongRequest struct {
	Name string `json:"name" binding:"required"`
}

func pingpongHandler(c *gin.Context) {
	var m PingpongRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(m)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"request": m,
	})
}
