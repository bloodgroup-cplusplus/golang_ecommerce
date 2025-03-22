package main

import (
	"database/sql"
	"fmt"
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
}
