package test

import (
"fmt"
"os"
"util"
"testing"
)

func Test_Config(t *testing.T) {
	myConfig := new(util.Config)
	path, _ := os.Getwd()
	fmt.Println(path)
	path += "/src/util/config.txt"
	fmt.Println(path)
	myConfig.InitConfig(path)
	fmt.Println(myConfig.Read("default", "path"))
	fmt.Printf("%v", myConfig.Mymap)
}
