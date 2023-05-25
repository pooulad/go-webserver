package main


import(
	"fmt"
	"log"
	"net/http"
)

func helloHanler(){
	
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.Handle("/hello",helloHanler)


	fmt.PrintF("starting server ar port 8080 \n")
	err := http.ListenAndServe("8080",nil)
	if err != nil {
		log.
	}
}