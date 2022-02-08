package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	//postgres
	host     = "PSQL_HOST"
	port     = "PSQL_PORT"
	username = "PSQL_USERNAME"
	dbName   = "PSQL_DBNAME"
	sslMode  = "PSQL_SSLMODE"
	password = "PSQL_PASSWORD"
	//tg
	token = "TG_TOKEN"
	chat  = "TG_CHAT"
	//google
	sheetDB     = "SHEET_DB"
	sheetMsg    = "SHEET_MSG"
	sheetAdmins = "SHEET_ADMINS"
	//cache
	cacheAddr = "CACHE_ADDR"
	keepTime  = "CACHE_KEEPTIME"
)

type (
	Conf struct {
		Postgres PostgresConfig
		Tg       TgConfig
		Sheets   SheetsConfig
		Redis    RedisConfig
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		DBName   string
		SSLMode  string
		Password string
	}
	TgConfig struct {
		Token string
		Chat  int64
	}

	SheetsConfig struct {
		DB     string
		Msg    string
		Admins string
	}

	RedisConfig struct {
		KeepTime int64
		Addr     string
	}
)

func InitConf() (*Conf, error) {
	var test bool
	flag.BoolVar(&test, "test", false, "off for docker")
	flag.Parse()
	return envVar(test)
}

func envVar(test bool) (*Conf, error) {

	if test {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, errors.Wrap(err, "envVar load")
		}
	}

	chatInt, err := strconv.Atoi(os.Getenv(chat))
	if err != nil {
		return nil, errors.Wrap(err, "chat")
	}
	keepTimeInt, err := strconv.Atoi(os.Getenv(keepTime))
	if err != nil {
		return nil, errors.Wrap(err, "keepTime")
	}

	return &Conf{
		Postgres: PostgresConfig{
			Host:     os.Getenv(host),
			Port:     os.Getenv(port),
			Username: os.Getenv(username),
			DBName:   os.Getenv(dbName),
			SSLMode:  os.Getenv(sslMode),
			Password: os.Getenv(password),
		},
		Tg: TgConfig{
			Token: os.Getenv(token),
			Chat:  int64(chatInt),
		},
		Sheets: SheetsConfig{
			DB:     os.Getenv(sheetDB),
			Msg:    os.Getenv(sheetMsg),
			Admins: os.Getenv(sheetAdmins),
		},
		Redis: RedisConfig{
			KeepTime: int64(keepTimeInt),
			Addr:     os.Getenv(cacheAddr),
		},
	}, nil
}
