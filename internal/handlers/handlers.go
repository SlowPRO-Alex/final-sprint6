package handlers

import (
	"fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/SlowPRO-Alex/final-sprint6/tree/main/internal/service" 

)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
        return
    }
    err := r.ParseMultipartForm(10 << 20) // 10MB лимит
    if err != nil {
        http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
        return
    }
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Ошибка при извлечении файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    fileData, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
        return
    }

    converted, err := service.typeDetector(string(fileData)) 
    if err != nil {
        http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
        return
    }

    ext := filepath.Ext(handler.Filename)
    filename := fmt.Sprintf("%s%s", time.Now().UTC().Format("20060102T150405"), ext)
    outFile, err := os.Create(filename)
    if err != nil {
        http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    _, err = outFile.WriteString(converted)
    if err != nil {
        http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(converted))
}