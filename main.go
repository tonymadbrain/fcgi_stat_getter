package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tomasen/fcgi_client"
)

var (
	host	= flag.String("h", "", "Host of PHP server")
	port	= flag.Int("p", 3000, "Port of PHP server")
	check	= flag.String("c", "", "Check type: ping, stat or clear")
	key	= flag.String("k", "", `Key for check type 'stat': pool, accepted_conn,
		listen_queue, max_listen_queue, listen_queue_len, idle_processes,
		active_processes, total_processes, max_active_processes,
		max_children_reached, slow_requests, latency`)
	timeout	= int(1000)
	s	[]string
)

func main() {
	flag.Parse()

	if flag.NFlag() < 3 {
		os.Exit(1)
	}

	switch *check {
	case "ping":
		out := makeRequest("ping")

		fmt.Println(strings.TrimSpace(string(out)))
	case "stat":
		res := make(map[string]string)
		out := makeRequest("status")
		s = strings.Split(strings.TrimSpace(string(out)), "\n")

		for _, pair := range s {
			z := strings.Split(pair, ":")
			res[z[0]] = z[1]
		}

		fmt.Println(strings.TrimSpace(res[*key]))
	case "clear":
		out := makeRequest("status")

		fmt.Println(strings.TrimSpace(string(out)))
	case "latency":
		start := time.Now()
		makeRequest("status")
		elapsed := float64(time.Since(start).Nanoseconds()) / float64(int64(time.Millisecond))

		fmt.Println(elapsed)
	default:
		log.Fatal("Bad check type")
	}
}

func makeRequest(script string) []byte {
	connect := fmt.Sprintf("%s:%d", *host, *port)
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = "/" + script
	env["SCRIPT_NAME"] = "/" + script
	env["SERVER_SOFTWARE"] = "go / fcgiclient / imhozbx "
	env["REMOTE_ADDR"] = connect

	fcgi, err := fcgiclient.DialTimeout("tcp", connect, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := fcgi.Get(env)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return content
}
