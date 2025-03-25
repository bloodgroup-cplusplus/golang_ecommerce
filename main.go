package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// define promethus metrics

var (
	addGoalCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "add_goal_requests_total",
		Help: "Total number of add goal requests",
	})
	removeGoalCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "remove_goal_requests_total",
		Help: "Total number of remove goal requests",
	})
	httpRequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	// Register prometheus metrics
	prometheus.MustRegister(addGoalCounter)
	prometheus.MustRegister(removeGoalCounter)
	prometheus.MustRegister(httpRequestsCounter)

}

func createConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err

	}
	return db, nil
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob(os.Getenv("KD_DATA_PATH") + "/*")
	db, err := createConnection()
	if err != nil {
		log.Println("Error connecting to postgres sql", err)
		return
	}
	defer db.Close()

	router.GET("/", func(c *gin.Context) {
		rows, err := db.Query("SELECT * from goals")
		if err != nil {
			log.Println("Error querying database", err)
			c.String(http.StatusInternalServerError, "Error querying databasse")
			return
		}
		defer rows.Close()

		var goals []struct {
			ID   int
			Name string
		}
		for rows.Next() {
			var goal struct {
				ID   int
				Name string
			}
		}
	})
}
