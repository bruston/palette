package main

import (
	"encoding/json"
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	listen := flag.String("listen", ":8080", "host:port to use")
	maxSize := flag.Int64("max_size", 1000000, "maximum size of image to process in bytes")
	readTimeout := flag.Duration("r_timeout", time.Second*60, "read timeout in seconds")
	writeTimeout := flag.Duration("w_timeout", time.Second*60, "write timeout in seconds")
	flag.Parse()
	srv := &http.Server{
		ReadTimeout:  time.Second * *readTimeout,
		WriteTimeout: time.Second * *writeTimeout,
		Addr:         *listen,
	}
	srv.SetKeepAlivesEnabled(true)
	http.Handle("/", palHandler{*maxSize})
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type rgb struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

type palHandler struct {
	maxSize int64
}

func (p palHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	minPixels, _ := strconv.ParseInt(r.URL.Query().Get("min"), 10, 64)
	prettyPrint, _ := strconv.ParseBool(r.URL.Query().Get("pretty"))
	lr := &io.LimitedReader{R: r.Body, N: p.maxSize}
	img, _, err := image.Decode(lr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seen := make(map[rgb]int64)
	bounds := img.Bounds()
	var colours []rgb
	for i := 0; i <= bounds.Max.X; i++ {
		for j := 0; j <= bounds.Max.Y; j++ {
			pixel := img.At(i, j)
			red, green, blue, alpha := pixel.RGBA()
			if alpha > 0 {
				seen[rgb{int(red / 257), int(green / 257), int(blue / 257)}]++
			}
		}
	}
	for v, count := range seen {
		if count > minPixels {
			colours = append(colours, v)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if prettyPrint {
		enc.SetIndent("", "\t")
	}
	enc.Encode(colours)
}
