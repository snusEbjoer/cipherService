package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/snusEbjoer/cipherService/internal/models"
	"github.com/snusEbjoer/cipherService/internal/utils"
	encode "github.com/snusEbjoer/cipherService/internal/utils/encodeutils"
)

const TTL = 3 * time.Minute
const RTTL = 10 * time.Second

type Handler = func(w http.ResponseWriter, r *http.Request)

type Service struct {
	cache *redis.Client
	ctx   context.Context
}

func NewEncodeService(cache *redis.Client, ctx context.Context) *Service {
	return &Service{
		cache: cache,
		ctx:   ctx,
	}
}

func (s *Service) EncodeSHA256() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path[1:]
		str := r.PathValue("string")

		val, err := s.cache.Get(s.ctx, path).Result()

		if err == nil { // some times this happens...
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(val))
			return
		}

		encodedString := encode.NewSHA256(str).ToString()

		resp, err := json.Marshal(models.EncodedString{
			EncodedString: encodedString,
		})

		if err != nil {
			utils.NewInternalServerError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (s *Service) EncodeMD5() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path[1:]
		str := r.PathValue("string")

		val, err := s.cache.Get(s.ctx, path).Result()

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(val))
			return
		}

		encodedString := encode.NewMD5(str).ToString()

		resp, err := json.Marshal(models.EncodedString{
			EncodedString: encodedString,
		})

		if err := s.cache.Set(s.ctx, path, resp, TTL).Err(); err != nil {
			log.Println(err.Error())
		}

		if err != nil {
			utils.NewInternalServerError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
