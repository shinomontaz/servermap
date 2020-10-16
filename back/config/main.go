package config

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "net/http/pprof"
)

// Config struct
type Config struct {
	ListenPort string
	TestFlag   bool

	HostsFile string
	VmsFile   string
}

type Env struct {
	Config *Config
}

func NewEnv(path string) *Env {
	v := viper.New()

	v.AddConfigPath(path)

	v.SetConfigName("conf.tpl") // всегда должен присутсвовать
	v.SetConfigType("json")
	err := v.ReadInConfig()
	checkErr(err)

	v.SetConfigName("conf") // переопределяет шаблонный конфиг и если нет переменных окружения то он и используется
	v.MergeInConfig()

	v.AutomaticEnv()

	var cfg Config
	err = v.Unmarshal(&cfg)
	checkErr(err)

	if cfg.TestFlag {
		log.SetLevel(log.DebugLevel)
	}

	return &Env{
		Config: &cfg,
	}
}

func checkErr(err error) {
	if err != nil {
		_, filename, lineno, ok := runtime.Caller(1)
		message := ""
		if ok {
			message = fmt.Sprintf("%v:%v: %v\n", filename, lineno, err)
		}
		log.Panic(message, err)
	}
}
