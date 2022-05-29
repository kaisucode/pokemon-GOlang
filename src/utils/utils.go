package utils

// import "fmt"

// func main() {
//   fmt.Println("vim-go")
// }

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func use(...interface{}) {}
