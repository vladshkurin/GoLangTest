package main

import (
    "net/http"
    "strings"
    "gopkg.in/gin-gonic/gin.v1"
)

type Request struct {
    Site       []string
    SearchText string
}

type Response struct {
    FoundAtSite string
}

func main() {
    router := gin.Default()

    router.POST("/checkText", func(c *gin.Context) {
        var json Request
        var json_res Response
        c.BindJSON(&json)
        site := json.Site
        search := strings.ToLower(json.SearchText)
        search_status := false

        for i := 0; i < len(site); i++ {
            domain := strings.Split(site[i], "//")
            site_name := strings.Split(domain[1], ".")
            if site_name[0] == search {
                search_status = true
                json_res.FoundAtSite = site[i]
            }
        }
        if search_status == true {
            c.JSON(200, json_res)
        } else {
            c.Writer.WriteHeader(http.StatusNoContent)
        }
    })
    router.Run(":8080")
