
package main

import (
	"github.com/ledongthuc/pdf"
	"fmt"
)



func main() {
	f,r,err := pdf.Open("./Bhutan-Penal-Code.pdf")
	if err != nil {
		fmt.Println("Error opening PDF:", err)
		return
	}
	defer f.Close()
	text,err := r.GetPlainText()
	fmt.Println("Text:", text)

}