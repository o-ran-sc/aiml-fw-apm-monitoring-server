package monitoring

//go:generate mockgen -source=monitoring.go -destination=./mock/monitoring_mock.go -package=mock

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/api/monitoring/scheme"
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/common/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/controller/agent"
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-server/pkg/controller/subscribe"
)

const (
	// REGISTER_URL is Register URL
	REGISTER_URL = "/register"
	// SUBSCRIBE_URL is Subscribe URL
	SUBSCRIBE_URL = "/subscribe"
)

type Command interface {
	Register(e *gin.RouterGroup)
}

// Server is server struct
type Server struct {
	Command
}

var agentExecutor agent.Command
var subscribeExecutor subscribe.Command

func init() {
	agentExecutor = agent.Executor{}
	subscribeExecutor = subscribe.Executor{}
}

// NewServer is Server constructor
func NewServer() (server Server) {
	return
}

// Register is for adding routing handler of Server
func (a Server) Register(e *gin.RouterGroup) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	e.POST(REGISTER_URL, a.registerMonitoringAgent)
	e.GET(REGISTER_URL, a.getMonitoringAgentList)

	e.GET(SUBSCRIBE_URL, a.subscribeList)
	e.POST(SUBSCRIBE_URL, a.subscribeMLApp)
	e.DELETE(SUBSCRIBE_URL, a.unSubscribeMLApp)
}

// Register godoc
// @Summary Register Monitoring Agent
// @Description register monitoring agent with info
// @Produce json
// @Param body body scheme.AgentScheme true "Agent info"
// @Success 201 {string} string "created"
// @Router /v1/monitoring/register [POST]
// @Tags Register
func (a Server) registerMonitoringAgent(c *gin.Context) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	var agentScheme scheme.AgentScheme

	if err := c.ShouldBind(&agentScheme); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request Input : "+err.Error())
		return
	}

	if err := agentExecutor.Register(agentScheme.Name, agentScheme.Endpoint); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error : "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register monitoring agent completed"})
}

// GetMonitoringAgentList godoc
// @Summary Get Monitoring Agent List
// @Description Returns information about registered agents
// @Produce json
// @Success 200 {string} string "OK"
// @Router /v1/monitoring/register [GET]
// @Tags Register
func (a Server) getMonitoringAgentList(c *gin.Context) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	list, err := agentExecutor.GetAgentList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error : "+err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// Subscribe godoc
// @Summary Subscribe ML App
// @Description subscribe ML App with ML App Name
// @Produce json
// @Param body body scheme.SubscribeScheme true "ML App info"
// @Success 201 {string} string "created"
// @Router /v1/monitoring/subscribe [POST]
// @Tags Subscribe
func (a Server) subscribeMLApp(c *gin.Context) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	var subscribeScheme scheme.SubscribeScheme

	if err := c.ShouldBind(&subscribeScheme); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request Input : "+err.Error())
		return
	}

	err := subscribeExecutor.Subscribe(subscribeScheme.Agent, subscribeScheme.Name, subscribeScheme.Data.Type, subscribeScheme.Data.Interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "subscribe completed"})
}

// Subscribe godoc
// @Summary Subscribe list of ML App
// @Description subscribe list of ML App
// @Produce json
// @Success 200 {string} string "OK"
// @Router /v1/monitoring/subscribe [GET]
// @Tags Subscribe
func (a Server) subscribeList(c *gin.Context) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	list, err := subscribeExecutor.GetSubscribeList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error : "+err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// Subscribe godoc
// @Summary Unsubscribe ML App
// @Description unsubscribe ML App
// @Produce json
// @Param body body scheme.UnsubscribeScheme true "ML App info"
// @Success 204 {string} string "deleted"
// @Router /v1/monitoring/subscribe [DELETE]
// @Tags Subscribe
func (a Server) unSubscribeMLApp(c *gin.Context) {
	logger.Debug("[IN]")
	defer logger.Debug("[OUT]")

	var unSubscribeScheme scheme.UnsubscribeScheme

	if err := c.ShouldBind(&unSubscribeScheme); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request Input : "+err.Error())
		return
	}

	err := subscribeExecutor.UnSubscribe(unSubscribeScheme.Agent, unSubscribeScheme.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Deleted"})
}
