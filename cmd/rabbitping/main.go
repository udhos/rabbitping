// Package main implements the rabbitping tool.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const version = "1.2.0"

type application struct {
	me            string
	conf          config
	serverMetrics *http.Server
	serverHealth  *http.Server
	met           *metrics
}

func longVersion(me string) string {
	return fmt.Sprintf("%s runtime=%s GOOS=%s GOARCH=%s GOMAXPROCS=%d",
		me, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0))
}

func main() {

	//
	// parse cmd line
	//

	var showVersion bool
	flag.BoolVar(&showVersion, "version", showVersion, "show version")
	flag.Parse()

	//
	// show version
	//

	me := filepath.Base(os.Args[0])

	{
		v := longVersion(me + " version=" + version)
		if showVersion {
			fmt.Println(v)
			return
		}
		log.Print(v)
	}

	app := &application{
		me:   me,
		conf: getConfig(me),
	}

	if app.conf.restartDeploy == "" {
		log.Fatalf("missing required env var RESTART_DEPLOY")
	}

	//
	// start metrics server
	//

	{
		app.met = newMetrics(app.conf.metricsNamespace, app.conf.metricsLatencyBuckets)

		mux := http.NewServeMux()
		app.serverMetrics = &http.Server{
			Addr:    app.conf.metricsAddr,
			Handler: mux,
		}

		mux.Handle(app.conf.metricsPath, promhttp.Handler())

		go func() {
			log.Printf("metrics server: listening on %s %s",
				app.conf.metricsAddr, app.conf.metricsPath)
			err := app.serverMetrics.ListenAndServe()
			log.Fatalf("metrics server: exited: %v", err)
		}()
	}

	//
	// start health server
	//

	{
		mux := http.NewServeMux()
		app.serverHealth = &http.Server{
			Addr:    app.conf.healthAddr,
			Handler: mux,
		}

		mux.HandleFunc(app.conf.healthPath, func(w http.ResponseWriter,
			_ /*r*/ *http.Request) {
			http.Error(w, "health ok", 200)
		})

		go func() {
			log.Printf("health server: listening on %s %s",
				app.conf.healthAddr, app.conf.healthPath)
			err := app.serverHealth.ListenAndServe()
			log.Fatalf("health server: exited: %v", err)
		}()
	}

	//
	// start pinger
	//

	go pinger(app)

	<-make(chan struct{}) // wait forever
}
