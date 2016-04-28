package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	//"strings"
	//"io"
	//"os"
	"fmt"
	"sync"
	"time"
)

var config map[string]interface{}
var baseTime time.Time

func main() {

	config = loadConfig()
	//baseurl := config["url"].(string)

	//baseurl := "http://192.168.2.114:9090/zaixianshang/seckilling"
	baseurl := "http://127.0.0.1:9090/zaixianshang/seckilling"
	concurrency, _ := strconv.Atoi(config["count"].(string))
	totalSuccess := 0
	totalFail := 0
	var lock sync.Mutex
	sem := make(chan int, concurrency)

	for i := 0; i < concurrency; i++ {

		go func(index int) {

			//tnow := time.Now()
			//starttimeint := (int)(tnow.Sub(baseTime).Seconds() * 1000000)
			url := baseurl + "?" + "userid=" + strconv.Itoa(index) + "&productid=111"

			//fmt.Println(url)
			resp, err := http.Get(url)

			if err != nil {
				//fmt.Println(err)
				//fmt.Println(err)
				lock.Lock()
				totalFail = totalFail + 1
				lock.Unlock()
				return

			}

			defer resp.Body.Close()

			//body, _ := ioutil.ReadAll(resp.Body)

			//fmt.Println("index:", index)
			//fmt.Println(string(body))

			//if string(body) == "uid=123" {
			lock.Lock()
			totalSuccess = totalSuccess + 1
			lock.Unlock()
			//} else {
			//	totalFail++
			//}

			sem <- index

		}(i)
	}

	//timeout := make(chan bool, 1)
	//启动timeout协程，由于缓存为1，不可能泄露
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	timeout <- true
	//}()

	//select {
	//case <-sem:
	//	// a read from ch has occurred
	//case <-timeout:
	// the read from ch has timed out
	//}

	for m := 0; m < concurrency; m++ {

		<-sem
		//fmt.Println("read squence:", tmp)
	}

	fmt.Println("totalSuccess:", totalSuccess)
	fmt.Println("totalFail:", totalFail)

}

func readFile(filename string) (map[string]interface{}, error) {
	var FdMap map[string]interface{}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &FdMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return FdMap, nil
}

func loadConfig() map[string]interface{} {

	CFTablesMap, err := readFile("./config/download.config")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		return nil
	}
	return CFTablesMap
}
