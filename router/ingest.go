package routes

import (
	"encoding/json"
	"fmt"
	"loghawk/parser"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var hitCounts = make(map[string]int)
var mu sync.Mutex

func IncrementHitCount() {
	currentHour := time.Now().Format("2006-01-02 15:00:00") // Use the hour as the key
	mu.Lock()
	hitCounts[currentHour]++
	mu.Unlock()
}

const poolSize = 5

var workerPool chan struct{}

func init() {
	workerPool = make(chan struct{}, poolSize)
}

func worker(tag, log string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire a worker from the pool
	workerPool <- struct{}{}

	// Release the worker when done
	defer func() {
		<-workerPool
	}()

	// Your parsing logic here
	parser.ParseLogs(tag, log)
}

func GetIngestRoutes(db *gorm.DB, router *gin.Engine) {

	router.POST("/ingest", func(c *gin.Context) {
		var input json.RawMessage
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data map[string]string
		err := json.Unmarshal(input, &data)
		if err != nil {
			fmt.Println("Err > ", err)
		}

		fmt.Println("Data > ", data)
		fmt.Println("Recieved logs from tag : ", data["tag"], data["log"])

		log, dataOk := data["log"]
		if dataOk {
			IncrementHitCount()
			// parser.ParseLogs(data["tag"], log)
			var wg sync.WaitGroup

			wg.Add(1)
			go worker(data["tag"], log, &wg)

			// Wait for all workers to finish
			wg.Wait()

		}

		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
}
