package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func promethes() {
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
			if err := rows.Scan(&goal.ID, &goal.Name); err != nil {
				log.Println("Error scanning row", err)
				continue
			}

			goals = append(goals, goal)
		}
		httpRequestsCounter.WithLabelValues("/").Inc()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"goals": goals,
		})
	})
	router.POST("/add_goal", func(c *gin.Context) {
		goalName := c.PostForm("goal_name")
		if goalName != "" {
			_, err = db.Exec("INSERT INTO goals (goal_name) VALUES($1)", goalName)
			if err != nil {
				log.Println("error inserting goal", err)
				c.String(http.StatusInternalServerError, "Error inserting values")
				return
			}
			// Increment the add goal counter
			addGoalCounter.Inc()
			httpRequestsCounter.WithLabelValues("/add_goal").Inc()

		}
		c.Redirect(http.StatusFound, "/")
	})
	router.POST("/remove_goal", func(c *gin.Context) {
		goalID := c.PostForm("goal_id")
		if goalID != "" {
			_, err = db.Exec("DELETE FROM goals WHERE id = $1", goalID) // Increment the remove goal counter
			if err != nil {
				log.Println("Error deleting goal", err)
				c.String(http.StatusInternalServerError, "Error deleting goal from the database")
				return
			}
			removeGoalCounter.Inc()
			httpRequestsCounter.WithLabelValues("/remove_goal").Inc()
		}
		c.Redirect(http.StatusFound, "/")
	})
	router.GET("/health", func(c *gin.Context) {
		httpRequestsCounter.WithLabelValues("/health").Inc()
		c.String(http.StatusOK, "OK")
	})
	// Expose metric endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":8080")
}
