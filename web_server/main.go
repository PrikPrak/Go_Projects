/* Simple web server, example straight from Mat Ryer's book 
   Go Programming Blueprints:
   https://www.packtpub.com/application-development/go-programming-blueprints */
   

package main

import (
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
                  <html>
                    <head>
                      <title>Chat</title>
                    </head>
                    <body>
                      Let's chat!
                    </body>
                  </html>
                `))
	})

	// Start the web server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
