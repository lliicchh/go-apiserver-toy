package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lliicchh/apiserver/config"
	"github.com/lliicchh/apiserver/router"
	"github.com/lliicchh/apiserver/model"
	"github.com/lliicchh/apiserver/router/middleware"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// get gin mode
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	router.Load(
		g,
		middleware.RequestId(),
		middleware.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Error("the router has no reponse, or it might took too long to start up : ", err)
		}
		log.Infof("the router has been deployed successfully")
	}()

	log.Infof("start to listening on the http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to router")

}
