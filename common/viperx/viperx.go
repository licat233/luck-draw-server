package viperx

import (
	"fmt"
	"luck_draw-server/common/utils"
	"path"

	"github.com/spf13/viper"
)

func NewViper() (viperx *viper.Viper) {
	viperx = viper.New()
	viperx.SetConfigName("config")
	viperx.SetConfigType("yaml")
	rootpath := utils.CurrentProjectRootPath()
	fpath := path.Join(rootpath, "deploy/config")
	viperx.AddConfigPath(fpath)
	err := viperx.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return
}
