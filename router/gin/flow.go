package gin

/*
	Refs:
	- https://github.com/mackristof/react-webapp/blob/master/main/react-webapp.go
*/

/*
import (
	"fmt"
	// b64 "encoding/base64"
	// "os"
	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/trustmaster/goflow"
	"github.com/roscopecoltran/goflow"

	"github.com/roscopecoltran/krakend/logging"
)

// struct Greeter used communication
type FlowRequest struct {
	flow.Component
	Node <-chan string
	Res  chan<- string
}

func (g *FlowRequest) OnName(name string) {
	response := fmt.Sprintf("Hello, %s!", name)
	g.Res <- response
}

type FlowLogger struct {
	flow.Component
	Line <-chan string
	Res  chan<- string
}

// graphql gin
func (p *FlowLogger) OnLine(line string) {
	fmt.Println(line)
	p.Res <- line
}

type EndpointFlow struct {
	flow.Graph
}

func NewRequestFlow() *FlowRequest {
	n := new(EndpointFlow)
	n.InitGraphState()
	n.Add(new(FlowRequest), "request")
	n.Add(new(FlowLogger), "logger")
	n.Connect("request", "Res", "logger", "Line")
	n.MapInPort("In", "request", "Name")
	n.MapOutPort("Out", "logger", "Res")
	return n
}

func FlowRoutes() {
	net := NewGreetingFlow()
	in := make(chan string)
	out := make(chan string)
	net.SetInPort("In", in)
	net.SetOutPort("Out", out)
	flow.RunNet(net)

	router := gin.Default()
	router.GET("/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		in <- name
		var resp struct {
			Title string `json:"greeter"`
			Value string
		}
		resp.Title = "golang"
		resp.Value = <-out
		c.JSON(200, resp)
	})
	router.Run(":8080")
	close(in)
	close(out)
	<-net.Wait()
}

// AuthRequired require Authorization with token Basic64 encoded
func FlowHandler(logger logging.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		logger.Debug("Method:", c.Request.Method)
		logger.Debug("URL:", c.Request.RequestURI)
		logger.Debug("Query:", c.Request.URL.Query())
		logger.Debug("Params:", c.Params)
		logger.Debug("Headers:", c.Request.Header)
	}
}

*/
