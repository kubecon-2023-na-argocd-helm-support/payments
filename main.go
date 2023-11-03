package main

import (
	_ "embed"
	"net/http"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

//go:embed index.html
var indexHTML []byte

func main() {
	redisAddress := "localhost:6379"
	if val := os.Getenv("REDIS_ADDRESS"); val != "" {
		redisAddress = val
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		_, _ = writer.Write(indexHTML)
	})
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		redisClient := redis.NewClient(&redis.Options{
			Addr: redisAddress,
		})
		defer redisClient.Close()
		var cnt int64
		var err error
		if r.Method == http.MethodPost {
			cnt, err = redisClient.Incr(r.Context(), "count").Result()
		} else {
			cnt, err = redisClient.Get(r.Context(), "count").Int64()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			_, _ = w.Write([]byte(strconv.FormatInt(cnt, 10)))
		}
	})
	_ = http.ListenAndServe(":8080", nil)
}
