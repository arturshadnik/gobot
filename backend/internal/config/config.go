// configuration, .env parse
package config

import (
	"os"

	"github.com/joho/godotenv"
)

var _ error = godotenv.Load()

var OpenAIApiKey string = os.Getenv("OPENAI_API_KEY")
var OpenAIUrl string = os.Getenv("OPENAI_URL")
