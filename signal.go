// Copyright 2015 Matthew Holt and The Caddy Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"net/http"
)

func getTime() string {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	return timestamp
}

func trapSignalsPosix(d chan bool) {
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

		for sig := range sigchan {
			switch sig {
			case syscall.SIGINT:
				fmt.Println(getTime(), "interupting process immediately signal: SIGINT")
				os.Exit(30)

			case syscall.SIGQUIT:
				fmt.Println(getTime(), "quitting process immediately signal: SIGQUIT")
				os.Exit(20)

			case syscall.SIGTERM:
				fmt.Println(getTime(), "shutting down apps then terminating signal: SIGTERM")

			case syscall.SIGUSR1:
				fmt.Println(getTime(), "not implemented signal: SIGUSR1")

			case syscall.SIGUSR2:
				fmt.Println(getTime(), "not implemented signal: SIGUSR2")

			case syscall.SIGHUP:
				// ignore; this signal is sometimes sent outside of the user's control
				fmt.Println(getTime(), "not implemented signal: SIGHUP")
			}
		}
		d <- true
	}()
}

func main() {
	done := make(chan bool, 1)

	fmt.Println(getTime(), "Starting signal trapper...")

	trapSignalsPosix(done)

	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":8080", nil)
	<-done

	os.Exit(10)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<html><head></head><body><p1>Hello, time: %s, url: %s, node: %s !<p1></body></html>", getTime(), r.URL.Path, name)
}
