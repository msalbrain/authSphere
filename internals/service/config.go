package service


import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"strconv"
)

type Server struct {
	APPLICATION_NAME string
	APPLICATION_URL string
	Port string
	Host string
	PasswordCost int64
}

type Database struct {
	Dbname string
	Mode string
}

type Mail struct {
	Username string
	Password string
	Host string
	Port string
}

type Config struct {
	Server
	Database
	Mail
	JWTSecret string
	ACCESSTOKENLIFE int64
	REFRESHTOKENLIFE int64
}


func GetEnvirometConfig() Config {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	passwordcost, err := strconv.Atoi(os.Getenv("PASSWORD_COST"))
	if err != nil {
		panic("password settings invalid")
	}
	accessTokenLife, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE"))
	if err != nil {
		panic("accessTokenLife setting invalid")
	}
	refreshTokenLife, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE"))
	if err != nil {
		panic("refreshTokenLife setting invalid")
	}


	newConfig := Config{
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
			PasswordCost: int64(passwordcost),
		},
		Database: Database{
			Dbname: os.Getenv("DATABASE_NAME"),
			Mode: os.Getenv("MODE"),
		},
		Mail: Mail{
			Username: os.Getenv("EMAIL_USERNAME"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Port: os.Getenv("EMAIL_PORT"),
			Host: os.Getenv("EMAIL_HOST"),
		},
		JWTSecret: os.Getenv("SECRET"),
		ACCESSTOKENLIFE: int64(accessTokenLife),
		REFRESHTOKENLIFE: int64(refreshTokenLife),
	}

	return newConfig
}

