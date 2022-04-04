package main

import "fmt"

func main (){
	// urls:= []string{
	// 	"fb.com",
	// 	"google.com",
	// 	"linkedin.com",
	// }

	// c := make(chan string)
	// defer close(c)
	// for _, url:= range(urls){
	// 	 go checkURL(url,c )
	// }

	// var p []string
	// for i:=0 ; i<len(urls); i++{
	// 	p = append(p, <-c)
	// }

	// fmt.Println(p)

	test := map[string]string{
		"a":"nepal",
		"b":"india",
	}
	fmt.Println(test)
	callMap(test)
	fmt.Println(test)

}

func callMap(testMap map[string]string){
	testMap["a"]="check"
}


func checkURL(a string , c chan string ){
	c <- a
}