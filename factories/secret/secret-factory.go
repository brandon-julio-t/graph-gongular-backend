package secret

import (
	"crypto/rand"
	"log"
	"os"
)

type Factory struct{}

func (*Factory) Create() []byte {
	if secret := []byte(os.Getenv("APP_KEY")); len(secret) > 0 {
		return secret
	}
	return makeSecret()
}

func makeSecret() []byte {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Error while generating APP_KEY")
	}

	return key
}
