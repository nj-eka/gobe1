package env

import (
	"github.com/joho/godotenv"
	"os"
	fp "path/filepath"
)

func lookupDotEnvPath() (path string){
	path = fp.Join(fp.Dir(os.Args[0]), ".env")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		wd, _ := os.Getwd()
		path = fp.Join(wd, ".env")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			path = ""
		}
	}
	return
}

func init(){
	if dotEnvPath := lookupDotEnvPath(); dotEnvPath != ""{
		_ = godotenv.Load(dotEnvPath)
	}
}

func GetEnv(key string, defaultValue string) string{
	if value, ok := os.LookupEnv(key); ok{
		return value
	}
	return defaultValue
}
