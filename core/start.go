package main

import (
	"net/http"
	"strconv"
	"strings"
	"work7z-go/core/handlers"
	"work7z-go/core/log"
	"work7z-go/core/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var IsDevMode bool = true

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Info("recover just now")
		}
	}()
	log.Ref().Info("Service is launching...")
	if IsDevMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	R_Engine := gin.Default()

	handlers.SetupRoutes(R_Engine)

	// set begin port
	port := 32017

	var host string = ""
	if tools.IsDockerMode() {
		host = "0.0.0.0:" + strconv.Itoa(port)
	} else {
		host = "127.0.0.1:" + strconv.Itoa(port)
	}

	actualServerPath := "127.0.0.1"
	if tools.IsDockerMode() {
		actualServerPath = "127.0.0.1" // TODO: consider to use other domain
	}
	fullURL := "http://" + strings.ReplaceAll(host, "0.0.0.0", actualServerPath) + "" + "/entry?t=" + ""

	println("")
	println("-----------------------------------------------")
	println("PLEASE ACCESS THE LINK BELOW IN BROWSER.")
	println("请复制下方链接并在浏览器端打开(for zh_CN users)")
	println("請復製下方鏈接並在瀏覽器端打開(for zh_HK users)")
	println("" + fullURL + "  ")
	println("-----------------------------------------------")
	println("")

	if !tools.IsDevMode {
		// global.OpenInBrowser(fullURL)
	}

	Srv := &http.Server{
		Addr:    host,
		Handler: R_Engine,
	}

	go func() {
		if err := Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			tools.ShouldNoErr(err, "Unable to launch the service")
		}
	}()

}
