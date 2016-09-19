package main

import ("fmt"
        "flag"
        )
func main()  {
    port := flag.String("p", "8080", "Порт для запуска сервера")
    flag.Parse()

    fmt.Println(*port)
}