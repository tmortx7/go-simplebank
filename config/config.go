package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		DBDriver string
		DBSource string
	}
	Security struct {
		TokenSymmetricKey    string
		AccessTokenDuration  time.Duration
		RefreshTokenDuration time.Duration
	}
	Server struct {
		Address string
	}
}

// C is config variable
var C config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) {
	Config := &C

	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config")

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
