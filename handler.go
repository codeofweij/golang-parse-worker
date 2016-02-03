package main

import (
	"fmt"
	"net/http"
	"os"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func ParseRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Now, we retrieve the person's name from the request.
	filepath := r.FormValue("filepath")

	exists, err := exists(filepath)
	if exists && err == nil {
		// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
		work := WorkRequest{filepath: filepath}

		// Push the work onto the queue.
		WorkQueue <- work
		fmt.Println("request queued")
		// And let the user know their work request was created.
		w.WriteHeader(http.StatusCreated)
	}
	return
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
