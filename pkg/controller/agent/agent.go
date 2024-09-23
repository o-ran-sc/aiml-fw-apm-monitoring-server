package agent

//go:generate mockgen -source=agent.go -destination=./mock/agent_mock.go -package=mock

import (
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/common/logger"
)

type Command interface {
	Register(name string, endpoint string) error
	UnRegister(name string) error
	GetAgentList() (list []string, err error)
}

type Executor struct {
	Command
}


func init() {
}

func (Executor) Register(name string, endpoint string) (err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}

func (Executor) UnRegister(name string) (err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}

func (Executor) GetAgentList() (list []string, err error) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	return
}
