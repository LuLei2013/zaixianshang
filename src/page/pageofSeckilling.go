package page

import(
	"fmt"
	"html/template"
	"net/http"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("page.html")
		t.Execute(w, nil)
	}
}