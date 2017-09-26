package gin

/*
	Refs:
	- https://github.com/kanocz/goginjsonrpc

	rpc := TestRPC{}
	router := gin.Default()
	router.POST("/", func(c *gin.Context) { goginjsonrpc.ProcessJsonRPC(c, &rpc); })
	router.Run("127.0.0.1:8000")

*/
