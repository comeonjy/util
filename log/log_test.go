// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/10 5:59 下午
package log_test

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/elastic"
	"github.com/comeonjy/util/email"
	"github.com/comeonjy/util/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func init()  {
	config.LoadConfig()
}

func TestInit(t *testing.T) {
	elastic.Init(config.GetConfig().Elastic)
	email.Init(config.GetConfig().Email)
	log.Init(config.GetConfig().Log)
	logrus.Info("this is info")
	logrus.Error("this is error")
	logrus.Debug("this is debug")
	logrus.Fatal("this is fatal")
}
