package gin

import (
	"encoding/json"
	_ "fmt"
	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	"io/ioutil"

	"github.com/roscopecoltran/krakend/logging"
)

// DebugHandler creates a dummy handler function, useful for quick integration tests
func IndexerHandler(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug("Method:", c.Request.Method)
		logger.Debug("URL:", c.Request.RequestURI)
		logger.Debug("Query:", c.Request.URL.Query())
		logger.Debug("Params:", c.Params)
		logger.Debug("Headers:", c.Request.Header)
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		logger.Debug("Body:", string(body))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (r ginRouter) registerIndexerEndpoints() {
	// handler := IndexerHandler(r.cfg.Logger)
	// r.cfg.Engine.GET("/__debug/*param", handler)
	// r.cfg.Engine.POST("/__debug/*param", handler)
	// r.cfg.Engine.PUT("/__debug/*param", handler)
	r.cfg.Engine.POST("/api/search/:index", Search)
	r.cfg.Engine.POST("/api/index/:index/:docId", Index)
	r.cfg.Engine.PUT("/api/update/:index/:docId", Index)
	r.cfg.Engine.DELETE("/api/delete/:index/:docId", Delete)
}

func Index(c *gin.Context) {
	var index bleve.Index
	indexName := c.Params.ByName("index")
	index, err := bleve.Open(indexName)

	if err != nil {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexName, mapping)
		if err != nil {
			c.JSON(400, gin.H{"status": "Opening index error"})
			return
		}
	}

	defer func() {
		if index != nil {
			index.Close()
		}
	}()

	docId := c.Params.ByName("docId")
	if docId == "" {
		c.JSON(400, gin.H{"status": "Missing id"})
		return
	}

	var form map[string]interface{}

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"status": "Error reading body"})
		return
	}

	err = json.Unmarshal(requestBody, &form)

	if err != nil {
		c.JSON(400, gin.H{"status": "Malformed Payload JSON"})
		return
	}

	err = index.Index(docId, form)
	if err != nil {
		c.JSON(400, gin.H{"status": "Error indexing document"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func Search(c *gin.Context) {
	indexName := c.Params.ByName("index")
	index, err := bleve.Open(indexName)

	if err != nil {
		c.JSON(400, gin.H{"status": "Error opening index"})
		return
	}

	defer func() {
		if index != nil {
			index.Close()
		}
	}()

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"status": "Error reading body"})
		return
	}

	var searchRequest bleve.SearchRequest
	err = json.Unmarshal(requestBody, &searchRequest)

	if err != nil {
		c.JSON(400, gin.H{"status": "Error parsing query"})
		return
	}

	err = searchRequest.Query.Validate()
	if err != nil {
		c.JSON(400, gin.H{"status": "Error validating query"})
		return
	}

	searchResponse, err := index.Search(&searchRequest)
	if err != nil {
		c.JSON(400, gin.H{"status": "Error executing the query"})
		return
	}

	reply := make([]string, 0)

	//query := bleve.NewFuzzyQuery("blve")
	//searchRequest := bleve.NewSearchRequest(query)
	//searchResponse, _ := index.Search(searchRequest)

	for _, r := range searchResponse.Hits {
		reply = append(reply, r.ID)
	}

	response, err := json.Marshal(&reply)
	c.JSON(200, gin.H{"status": string(response[:])})
}

func Delete(c *gin.Context) {
	indexName := c.Params.ByName("index")
	index, err := bleve.Open(indexName)

	if err != nil {
		c.JSON(400, gin.H{"status": "Error opening index"})
		return
	}

	defer func() {
		if index != nil {
			index.Close()
		}
	}()

	docId := c.Params.ByName("docId")
	if docId == "" {
		c.JSON(400, gin.H{"status": "Missing id"})
		return
	}

	err = index.Delete(docId)
	if err != nil {
		c.JSON(400, gin.H{"status": "Error deleting document"})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
