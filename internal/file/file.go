package file

import (
	"net/http"
)

func FileDownload(dirPath string) {
	fs := http.Dir(dirPath)
	fileServer := http.FileServer(fs)

	http.Handle("/download/", http.StripPrefix("/download", fileServer))
}
