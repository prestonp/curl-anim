package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/prestonp/curl-anim/service"
)

var (
	host       = flag.String("host", "0.0.0.0", "server host")
	port       = flag.String("port", "8081", "server port")
	fps        = flag.Int("fps", 15, "frames per second")
	framesPath = flag.String("frames-path", "./frames", "dir to frames")
  repo       = flag.String("repo", "https://github.com/prestonp/curl-anim", "source repo")
)

func main() {
	flag.Parse()

	addr := net.JoinHostPort(*host, *port)
	svc := service.New(*fps, *framesPath, *repo)

	fmt.Println("listening on", addr)
	panic(http.ListenAndServe(addr, svc))
}
