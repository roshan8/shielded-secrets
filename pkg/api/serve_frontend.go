package api

import (
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	app "shielded-secrets"
	"strings"
)

var frontendFs fs.FS

const defaultHtml = "index.html"

func ServeFrontendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	if path == "/" { // Add other paths that you route on the UI side here
		sendFileToClient(w, defaultHtml)
		return
	}

	path = strings.TrimPrefix(path, "/")

	file, err := frontendFs.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			sendFileToClient(w, defaultHtml)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(path))
	w.Header().Set("Content-Type", contentType)

	io.Copy(w, file)
}

func init() {
	var err error
	frontendFs, err = fs.Sub(app.Frontend, "fe/build")
	if err != nil {
		log.Fatal("failed to get frontend fs", err)
	}
}

func sendFileToClient(w http.ResponseWriter, path string) {
	pathTrimmed := strings.TrimPrefix(path, "/")
	file, err := frontendFs.Open(pathTrimmed)

	if err != nil {
		log.Fatalln(err)
	}

	contentType := mime.TypeByExtension(filepath.Ext(path))
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, file)
}
