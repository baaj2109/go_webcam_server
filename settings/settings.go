package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	Name         string `mapsturcture:"name"`
	Version      string `mapsturcture:"version"`
	StartTime    string `mapsturcture:"start_time"`
	MachineID    string `mapsturcture:"machine_id"`
	*LogConfig   `mapsturcture:"log"`
	*MySqlConfig `mapsturcture:"mysql"`
	*RedisConfig `mapsturcture:"redis"`
	*JWTConfig   `mapsturcture:"jwt"`
}

type LogConfig struct {
	Level      string `mapsturcture:"level"`
	Filename   string `mapsturcture:"filename"`
	MaxSize    int    `mapsturcture:"max_size"`
	MaxAge     int    `mapsturcture:"max_age"`
	MaxBackups int    `mapsturcture:"max_backups"`
}

type MySqlConfig struct {
	Host         string `mapsturcture:"host"`
	Port         int    `mapsturcture:"port"`
	User         string `mapsturcture:"user"`
	Password     string `mapsturcture:"password"`
	Database     string `mapsturcture:"database"`
	MaxOpenConns int    `mapsturcture:"max_open_conns"`
	MaxIdleConns int    `mapsturcture:"max_idle_conns"`
	MaxLifeTime  int    `mapsturcture:"max_life_time"`
}

type RedisConfig struct {
	Host         string `mapsturcture:"host"`
	Port         int    `mapsturcture:"port"`
	Password     string `mapsturcture:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type JWTConfig struct {
	Secret string `mapsturcture:"secret"`
	Issuer string `mapsturcture:"issuer"`
}

func Init() error {
	// read config
	viper.SetConfigName("config/app.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
		}
		viper.Unmarshal(&Conf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return nil
}
