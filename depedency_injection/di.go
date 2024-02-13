package depedencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, input string) {
	fmt.Fprintf(w, "Hello %s", input)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
