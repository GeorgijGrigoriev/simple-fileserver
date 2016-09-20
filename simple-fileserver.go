package main

import (
        "flag"
        "net/http"
        )
func main()  {
    
    port := flag.String("p", "8080", "Порт для запуска сервера")
    directory := flag.String("d", ".", "Директория для отображения файлов")
    flag.Parse()
    http.Handle("/", http.FileServer(http.Dir(*directory)))
    http.ListenAndServe(":"+*port, nil)
}
