package subscribe

//go:generate mockgen -source=subscribe.go -destination=./mock/subscribe_mock.go -package=mock

import (
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/common/logger"
)

type subscriber struct {
}

type Command interface {
	Subscribe(host string, name string, dataType []string, interval int) error
	UnSubscribe(host string, name string) error
	GetSubscribeList() (info []SubscribeInfo, err error)
}

type Executor struct {
	Command
}

var subscribeList subscriber

func init() {
}

func (Executor) Subscribe(host string, name string, dataType []string, interval int) (err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}

func (Executor) UnSubscribe(host string, name string) (err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}

func (Executor) GetSubscribeList() (info []SubscribeInfo, err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}
