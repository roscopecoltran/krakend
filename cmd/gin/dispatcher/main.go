package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

var cnt int

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/fire-and-forget", fireAndForget)
	router.GET("/load-concurrently", loadConcurrently)
	router.GET("/waterfall", waterfall)
	router.Run("127.0.0.1:8080")
}

func fireAndForget(c *gin.Context) {
	start := time.Now()
	var results []*http.Response
	go hGet("http://www.gog.com")
	go hGet("http://www.google.com")
	go hGet("http://www.bing.com")
	timer := time.Since(start).Seconds()
	fmt.Println("request done in %f \n", timer)
	c.IndentedJSON(http.StatusOK, results)
	return
}

func loadConcurrently(c *gin.Context) {
	start := time.Now()
	var results []*http.Response
	ch := make(chan *http.Response)
	go hGetC("http://www.gog.com", ch)
	go hGetC("http://www.google.com", ch)
	go hGetC("http://www.bing.com", ch)
	results = append(results, <-ch)
	results = append(results, <-ch)
	results = append(results, <-ch)
	log.Printf("Codes: %d, %d, %d \n", results[0].StatusCode, results[1].StatusCode, results[2].StatusCode)
	log.Printf("Body[0]: %#v \n", results[0].Body)
	log.Printf("Body[1]: %#v \n", results[1].Body)
	log.Printf("Body[2]: %#v \n", results[2].Body)
	timer := time.Since(start).Seconds()
	fmt.Println("request done in %f \n", timer)
	pp.Print(results)
	c.IndentedJSON(http.StatusOK, results)
}

func waterfall(c *gin.Context) {
	start := time.Now()
	var results []*http.Response
	// var bodies []string
	results = append(results, hGet("http://www.gog.com"))
	results = append(results, hGet("http://www.google.com"))
	results = append(results, hGet("http://www.bing.com"))
	timer := time.Since(start).Seconds()
	fmt.Println("request done in %f \n", timer)
	pp.Print(results)
	c.IndentedJSON(http.StatusOK, results)
	return
}

func hGet(url string) *http.Response {
	start := time.Now()
	cnt++
	resp, err := http.Get(url)
	timer := time.Since(start).Seconds()
	if err != nil {
		log.Printf("get %s failed after %f: %s (%d) \n", url, timer, err, cnt)
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	} /* else {
		log.Printf("get %s done in %f (%d) \n", url, timer, cnt)
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}*/
	//fmt.Printf("hGet NON channeled")
	//pp.Print(resp)
	return resp
}

func hGetC(url string, c chan *http.Response) {
	fmt.Printf("hGet channeled")
	c <- hGet(url)
}
