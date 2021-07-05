package libraries

import (
	"github.com/spf13/viper"
	"log"
)

func NewConfig(log *log.Logger) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigName("config.local")
	v.SetConfigType("yaml")
	v.AddConfigPath("fx_gin/configs")
	err = v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	return
}
