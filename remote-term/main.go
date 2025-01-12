package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func getShell() string {
	shell := os.Getenv("SHELL")
	if shell != "" {
		return shell
	}

	validShells := []string{"/bin/zsh", "/bin/bash", "/bin/ash", "/bin/sh"}
	for _, shell := range validShells {
		if _, err := os.Stat(shell); err == nil {
			return shell
		}
	}

	return "/bin/sh"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// todo: add support for basic auth

		w.Header().Set("Hterm-Title", "Brick Terminal")
		Handle(w, r, func(args string) *Pty {
			cmd := exec.Command(getShell(), "-i")
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, "USER=root")
			pty, err := NewPty(cmd)
			if err != nil {
				log.Fatal(err)
			}
			return pty
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// check if the port is a number
	for _, c := range port {
		if c < '0' || c > '9' {
			log.Fatal("Invalid port number:", port)
		}
	}

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
