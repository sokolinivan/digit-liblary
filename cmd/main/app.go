package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"digit-liblary/internal/config"
	"digit-liblary/internal/user"
	"digit-liblary/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}
 
func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	configuration := config.GetConfig()
	
	logger.Info("rigister handlers")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, logger, configuration)
}

func start(router *httprouter.Router, logger *logging.Logger, cfg *config.Config) {
	logger.Info("start server")
	
	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		socketPah := path.Join(appDir, "app.sock")
		logger.Debugf("socket path %s", socketPah)

		logger.Info("listen unix socket")
		listener, listenErr =  net.Listen("unix", socketPah)
		logger.Infof("server is listening unix socket %s", socketPah)
	} else {
		logger.Info("listen tcp")
			
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}