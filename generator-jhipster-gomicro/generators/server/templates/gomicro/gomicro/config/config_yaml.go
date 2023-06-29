package app

import(
	"github.com/asim/go-micro/v3/config"
	"os"
	yaml "github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	"github.com/asim/go-micro/v3/config/reader"
	"github.com/asim/go-micro/v3/config/reader/json"
	"github.com/asim/go-micro/v3/config/source/file"
	"fmt"
	"github.com/micro/micro/v3/service/logger"
)

var g map[string]interface{}

func Setconfig(){
	var environment =os.Getenv("GO_MICRO_ENV")
	enc := yaml.NewEncoder()
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader( 
				reader.WithEncoder(enc),
			),
		),
	)
	if err := c.Load(file.NewSource(
		file.WithPath(("./app.yaml")),
	)); err != nil {
		logger.Errorf(err.Error())
		return
	}
	g=c.Map()
	if err := c.Load(file.NewSource(
		file.WithPath(("./"+environment+"-config.yaml")),
	)); err != nil {
		logger.Errorf(err.Error())
		return
	}
	for k,v :=range c.Map() {
		g[k]=v
	}	
}

func Getval(key string) string{
	return fmt.Sprintf("%v",g[key])
}
