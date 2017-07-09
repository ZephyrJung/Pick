package main
import (
	"net/http"
	"log"
)

const redis_host = "127.0.0.1:6379"

func main() {
	http.HandleFunc("/", router) //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func router(w http.ResponseWriter,r *http.Request){
	listLabels()
}