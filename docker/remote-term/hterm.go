package main

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"golang.org/x/net/websocket"
)

//go:embed assets/hterm.html
//go:embed assets/hterm.js
var folder embed.FS

type termMsg struct {
	Args   *string `json:"args"`
	Data   *string `json:"data"`
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
}

func Handle(w http.ResponseWriter, r *http.Request, pty func(string) *Pty) {
	switch {
	case strings.HasSuffix(r.URL.Path, "/hterm.js"):
		log.Println("hterm handling js:", r.URL.Path)
		HandleAsset(w, r, "assets/hterm.js")
	case strings.HasSuffix(r.URL.Path, "/hterm"):
		log.Println("hterm handling socket:", r.URL.Path)
		HandleSocket(w, r, pty)
	default:
		log.Println("hterm handling index:", r.URL.Path)
		HandleAsset(w, r, "assets/hterm.html")
	}
}

func HandleAsset(w http.ResponseWriter, r *http.Request, asset string) {
	data, err := folder.ReadFile(asset)
	if err != nil {
		log.Println("hterm:", err)
		http.NotFound(w, r)
	}
	if _, err := w.Write(data); err != nil {
		log.Println("hterm:", err)
	}
}

func HandleSocket(w http.ResponseWriter, r *http.Request, ptyFunc func(string) *Pty) {
	websocket.Handler(func(conn *websocket.Conn) {
		var obj termMsg
		dec := json.NewDecoder(conn)
		err := dec.Decode(&obj)
		if err != nil {
			log.Println("hterm:", err)
			return
		}
		if obj.Args == nil || obj.Width == nil || obj.Height == nil {
			log.Println("hterm: no args")
			return
		}
		pty := ptyFunc(*obj.Args)
		pty.Size(*obj.Width, *obj.Height)
		go func() {
			if _, err := io.Copy(conn, pty); err != nil {
				log.Println("hterm:", err)
			}
		}()
		for {
			var obj termMsg
			err := dec.Decode(&obj)
			if err != nil {
				log.Println("hterm:", err)
				break
			}

			if obj.Width != nil && obj.Height != nil {
				pty.Size(*obj.Width, *obj.Height)
				continue
			}
			if obj.Data != nil {
				_, err = io.WriteString(pty, *obj.Data)
				if err != nil {
					log.Println("hterm:", err)
					break
				}
			}
		}
	}).ServeHTTP(w, r)
}

type Pty struct {
	*os.File
}

// winsize stores the Height and Width of a terminal.
type winsize struct {
	Height uint16
	Width  uint16
}

func (pty *Pty) Size(width int, height int) {
	ws := &winsize{Width: uint16(width), Height: uint16(height)}
	syscall.Syscall(syscall.SYS_IOCTL, pty.Fd(), uintptr(syscall.TIOCSWINSZ), uintptr(unsafe.Pointer(ws)))

}

func NewPty(cmd *exec.Cmd) (*Pty, error) {
	cmdPty, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}
	return &Pty{cmdPty}, nil
}
