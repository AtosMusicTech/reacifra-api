package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfigFile() {
	dotEnvFilename := getFilenameDotEnv("./.env", "../../.env")
	err := godotenv.Load(dotEnvFilename)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error ao ler o arquivo de configuraçao")
	}

	viper.WatchConfig()
}

func getFilenameDotEnv(filenames ...string) string {
	for _, filename := range filenames {
		if existFilename(filename) {
			return filename
		}
	}

	return ""
}

func existFilename(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}
