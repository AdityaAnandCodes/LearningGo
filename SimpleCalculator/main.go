package main


import(
	"fmt"
)
func main(){
	var result int64 = 0;
	var a int64;
	var b int64;
	fmt.Println("Welcome To A Calculator Built in Go ! ");
	for {	
	fmt.Println("Please Enter The First Number");
	fmt.Scanln(&a);
	fmt.Println("Please Enter The Second Number");
	fmt.Scanln(&b);
	fmt.Println("Enter The Operation You Want to perform (+,-,*,/)");
	var operation string;
	fmt.Scanln(&operation);
	switch operation{
		case "+":
			result = a + b;
		case "-":
			result = a - b;
		case "/":
			if b != 0{
				result = a/b ;
			}else{
				fmt.Println("Error! Division By Zero is not allowed");
			}
		case "*":
			result = a * b ;
		default:
			fmt.Println("Error! Invalid Operation");
	}
	fmt.Println("The Result is ",result);
	}
}