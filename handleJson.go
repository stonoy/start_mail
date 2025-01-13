package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respWithError(w http.ResponseWriter, code int, msg string) {
	reply := struct {
		Msg string
	}{
		Msg: msg,
	}

	if code > 499 {
		log.Printf("error code -> %v, msg -> %v", code, msg)
	}

	respWithJson(w, code, reply)
}

func respWithJson(w http.ResponseWriter, code int, reply interface{}) {
	// marshal/convert reply to byte
	dat, err := json.Marshal(reply)
	if err != nil {
		log.Printf("can not marshal -> %v, err -> %v", reply, err)

		respWithError(w, 500, fmt.Sprintf("can not marshal err -> %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
