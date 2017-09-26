// https://github.com/cytoscape-ci/image-queue/blob/master/main.go
package gin

/*
import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func AuthRequiredHandler(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// func setBucketRoutes(r *gin.Engine, d *bolt.DB) {

		r.POST("/network/:network_id", func(c *gin.Context) {
			id := c.Param("network_id")
			d.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("queue"))
				q := b.Get([]byte("network_ids"))
				append(*q, id)
				err = b.Put([]byte("network_ids"), q)
				return err
			})
			c.JSON(200, gin.H{"message": "Network submission successful."})
		})
		r.GET("/network", func(c *gin.Context) {
			d.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("queue"))
				q := b.Get([]byte("network_ids"))
				len := q.Len() - 1
				network_id = (*q)[len]
				*q = (*q)[:x]
				err = b.Put([]byte("network_ids"), q)
				return err
			})
			c.JSON(200, gin.H{"network_id": network_id})
		})

	}

}
*/

/*
func main() {
	database := getDatabase()
	defer database.Close()
	router := getGinRouter(database)
	setRoutes(router, database)
	router.Run(":8080")
}
*/
