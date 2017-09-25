package gin

import (
	//"io"
	"io/ioutil"
	"os"

	"github.com/dogtools/dog"
	"github.com/gin-gonic/gin"
	"github.com/roscopecoltran/krakend/logging"
)

var (
	dogfile dog.Dogfile // Dogfile object
	// Define two tasks in the Dogfile format using YAML
	dogfileYAMLTest string = `
		- task: hello-dog
		  description: Say Hello
		  post: hello-world
		  code: echo "Hello Dog!"
		- task: hello-world
		  description: Say Hello Again
		  code: echo "Hello World!"`
)

/*
	Refs:
	- https://github.com/dogtools/dog/blob/master/examples/hello/hello.go
*/

func TaskHandler(logger logging.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		logger.Debug("Method:", c.Request.Method)
		logger.Debug("URL:", c.Request.RequestURI)
		logger.Debug("Query:", c.Request.URL.Query())
		logger.Debug("Params:", c.Params)
		logger.Debug("Headers:", c.Request.Header)

		// Get task name from path
		taskName := c.Request.URL.Path[1:]

		// Parse Dogfile
		dogfile, err := dog.Parse([]byte(dogfileYAMLTest))
		if err != nil {
			logger.Debug("error: ", err)
			c.AbortWithStatus(500)
		}

		// Generate task chain for the task named as the URL path
		taskChain, err := dog.NewTaskChain(dogfile, taskName)
		if err != nil {
			logger.Debug("task chain generation failed: ", err)
			c.AbortWithStatus(500)
		}

		// Run task chain, HTTP client receives info about how task finished
		err = taskChain.Run(os.Stdout, os.Stderr)
		if err != nil {
			logger.Debug("TaskName:", taskName, "error:", err)
			c.AbortWithStatus(500)
		}
		logger.Debug("TaskName: ", taskName, " finished.")

		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()

		logger.Debug("Body:", string(body))
		c.JSON(200, gin.H{
			"message": "task",
		})

	}

}
