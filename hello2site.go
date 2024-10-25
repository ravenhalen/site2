package main
import (
         "fmt"
         "net/http"
     )

     func handler(w http.ResponseWriter, r *http.Request) {
         fmt.Fprintln(w, "Hello from Vercel!")
     }

     func main() {}
