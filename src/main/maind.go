package main
import(

	"log"
	"net/http"
	"page"
)
func main() {
	http.HandleFunc("/page", page.GetPage)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
