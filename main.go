package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type responseobject struct{
// 	status infloat32 `status`

// }
func pingURL(url string) (Duration time.Duration,status any, error error) {
    start := time.Now()

    resp, err := http.Get(url)
    if err != nil {
    
        return 0,400, err
    }
  //  defer resp.Body.Close()

    latency := time.Since(start)
    return latency, resp.StatusCode,nil
}
 func v1EndpointHandler(c *gin.Context) {
	url:=c.Request.URL.Query()["url"][0]
	fmt.Print(url)
    latency,status,err:= pingURL(url)
	// fmt.Print("Response:",resp,err)
	if err!=nil{
		 c.JSON(http.StatusBadRequest,gin.H{"error": err.Error(),"url-status":status,"latency":latency})
         return
	 }
	c.JSON(http.StatusOK, gin.H{"latency": latency,"url-status":status})
}


func main() {
    router := gin.Default()

    v1 := router.Group("/v1")

    v1.GET("/get", v1EndpointHandler)


  
 

    router.Run(":5000")
}