package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// Load .env file -> PORT=3000 -> nowhere in the code we have mentioned the variable port, then how did we get the port number?
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// os.Getenv("PORT") -> this could have also worked
}
