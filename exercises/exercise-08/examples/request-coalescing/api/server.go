package api

import (
	"crypto/rand"
	"echo-server/database"
	"echo-server/database/message"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"golang.org/x/sync/singleflight"
)

type server struct {
	messageRepo message.Repository
	g           *singleflight.Group
}

func NewServer(
	messageRepo message.Repository,
) *server {
	g := &singleflight.Group{}
	return &server{messageRepo, g}
}

func (serv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/generate":
		for i := 0; i < 100; i++ {
			serv.messageRepo.Create(&database.Message{
				ID:   int64(i),
				User: randomString(16),
				Text: randomString(100),
			})
		}
	case "/benchmark/with-request-coalescing":
		randomId, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		msg, err, _ := serv.g.Do(fmt.Sprint(randomId.Int64()), func() (interface{}, error) {
			return serv.messageRepo.FindById(randomId.Int64())
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(msg)
	case "/benchmark/without-request-coalescing":
		randomId, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		msg, err := serv.messageRepo.FindById(randomId.Int64())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(msg)
	}
}

const alphabeth = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabeth))))
		result[i] = alphabeth[randomIndex.Int64()]
	}

	return string(result)
}
