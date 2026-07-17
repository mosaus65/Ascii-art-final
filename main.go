package main


import(
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . <filename>")
	}

	input := os.Args[1]
    filename := os.Args[2]

	banner, err := loadBanner(filename)
	if err != nil {
		fmt.Println(err)
	}
	result := Generate(input, banner)
	fmt.Println(result)
}

