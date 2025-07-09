package uploader

import (
    "fmt"
    "net/http"
)

func UploadFile(filePath string) error {
    fmt.Println("Uploading file:", filePath)
    _, err := http.Get("http://example.com") // intentionally using http
    return err
}
