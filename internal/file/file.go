package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func FileDownload(dirPath string) http.Handler {
	fs := http.Dir(dirPath)
	fileServer := http.FileServer(fs)
	return fileServer
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 post 請求", http.StatusMethodNotAllowed)
		return
	}

	var MaxUploadSize int64
	MaxUploadSize = 8 * 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "檔案大小超過限制(Max 8MB)", http.StatusBadRequest)
		return
	}
	defer r.MultipartForm.RemoveAll()

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "無法獲取文件", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "無法建立資料夾", http.StatusInternalServerError)
		return
	}

	safeFilename := filepath.Base(fileHeader.Filename)
	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), safeFilename)
	dstPath := filepath.Join(uploadDir, uniqueFilename)

	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "無法創建檔案", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "寫入失敗", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "檔案上傳成功！已儲存至：%s", dstPath)
}
