package main

import "fmt"

const Pi = 3.14
//書き換え不可型宣言いらない

const(
		Username = "test_user"
		Password = "test_pass"
)

func main(){
		fmt.Println(Pi, Username, Password)
}