package app

import(
	"github.com/asim/go-micro/v3/config"
	"os"
	yaml "github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	"github.com/asim/go-micro/v3/config/reader"
	"github.com/asim/go-micro/v3/config/reader/json"
	"github.com/asim/go-micro/v3/config/source/file"
	"fmt"
	"strings"
	"github.com/micro/micro/v3/service/logger"
)

var g map[string]interface{}

func Setconfig(){
	profile := "dev"
	profiles := map[string]bool{"dev":true, "prod":true};
	if( profiles[os.Getenv("GO_MICRO_PROFILE")] ){
		profile = os.Getenv("GO_MICRO_PROFILE");
	}
	enc := yaml.NewEncoder()
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader( 
				reader.WithEncoder(enc),
			),
		),
	)
	if err := c.Load(file.NewSource(
		file.WithPath(("config/app.yaml")),
	)); err != nil {
		logger.Errorf(err.Error())
		return
	}
	g=c.Map()
	if err := c.Load(file.NewSource(
		file.WithPath(("config/"+profile+"-config.yaml")),
	)); err != nil {
		logger.Errorf(err.Error())
		return
	}
	for k,v :=range c.Map() {
		g[k]=v
	}
	LoadEnv()	
}

func LoadEnv(){
	for _,e :=range os.Environ(){
		pair := strings.SplitN(e,"=",2)
		if(strings.HasPrefix(pair[0],"GO_MICRO")){
			g[pair[0]]=pair[1]
			logger.Infof(pair[0]+" loaded")
		}
	}
}

func GetVal(key string) string{
	return fmt.Sprintf("%v",g[key])
}
