package common

import (
	"echo-production-example/model"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const ymlPrefix = "config" //config.yml

var yml model.Yaml

func init() {
	viper.SetConfigName(ymlPrefix)
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Error reading config file : ")
	}

	err := viper.Unmarshal(&yml)
	if err != nil {
		fmt.Errorf("unable to decode into struct : ")
	}
}

func TestMssql(t *testing.T) {

	t.Run("", func(t *testing.T) {
		_, err := SqlxInit(yml)
		if err != nil {
			panic(err)
		}

		assert.NoError(t, err, "result is no error")

		SqlxClose()
	})

}
