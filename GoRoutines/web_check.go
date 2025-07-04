package main

import (
	"fmt"
	"net/http"
	"sync"
)


func main(){
	websites := []string{
		"https://google.com",
		"https://github.com",
		"https://nonexistent123.com",
	}

	var wg sync.WaitGroup

	for i:=0;i<3;i++{
		wg.Add(1)
		go GetUrl(websites[i])
	}
	wg.Wait()

}

func GetUrl(url string){
	resp,err:= http.Get(url)
	if err != nil {
		fmt.Printf("Error Fetching Data for : %q\n\n", url)
		return 
	}
	fmt.Printf("%q is Running : UP\n\n", url)
	defer resp.Body.Close()
}



