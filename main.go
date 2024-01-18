package main

import (
    "fmt"
    "log"
    "net/http" 

    "secondhandmarketapp/handler" 
	"secondhandmarketapp/backend"
    "secondhandmarketapp/util"
  
)
func main() {
    fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
    if err != nil {
        panic(err)
    }

    // log.Printf("Elasticsearch Address: %s", config.ElasticsearchConfig.Address)
    // log.Printf("Elasticsearch Username: %s", config.ElasticsearchConfig.Username)
    // log.Printf("Elasticsearch Password: %s", config.ElasticsearchConfig.Password)


    backend.InitElasticsearchBackend(config.ElasticsearchConfig)
    backend.InitGCSBackend(config.GCSConfig)


    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}