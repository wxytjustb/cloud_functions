package main

import (
	"net/http"
	"cloud_functions"
)

func main()  {
	http.HandleFunc("/api", p.Api)
	http.ListenAndServe("0.0.0.0:6789", nil)
}
