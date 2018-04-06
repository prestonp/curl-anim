package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prestonp/curl-anim/color"
)

func readFrames(path string) []string {
	frames := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		frame, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		frames = append(frames, string(frame))
	}
	return frames
}

func New(fps int, framesPath, repo string) http.Handler {
	frames := readFrames(framesPath)
	defaultFpsDuration := 1000 / fps

	svc := http.NewServeMux()
	svc.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("user-agent"), "curl") {
			w.Header().Set("location", repo)
			w.WriteHeader(http.StatusFound)
			return
		}

		f, ok := w.(http.Flusher)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "text/event-stream")
		w.Header().Set("connection", "keep-alive")

		fpsQuery := r.URL.Query().Get("fps")
		var duration = defaultFpsDuration
		if fpsQuery != "" {
			fps, err := strconv.Atoi(fpsQuery)
			if err != nil || fps > 60 || fps < 1 {
				http.Error(w, "invalid fps query (must be integer between 1 - 60)", http.StatusBadRequest)
				return
			}
			duration = 1000 / fps
		}

		ticker := time.NewTicker(time.Duration(duration) * time.Millisecond)
		printer := color.New()
		frameIdx := 0
		reqClose := r.Context().Done()
		for {
			select {
			case <-ticker.C:
				fmt.Fprint(w, printer.Print(frames[frameIdx]))
				frameIdx++
				if frameIdx >= len(frames) {
					frameIdx = 0
				}
				f.Flush()
			case <-reqClose:
				return
			}
		}
	})

	return svc
}
