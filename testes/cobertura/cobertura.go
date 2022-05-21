package cobertura

import "fmt"

func main() {
	fmt.Println("go test -cover")
	fmt.Println("go test -coverprofile=resultado.out")
	fmt.Println("go tool cover -func=resultado.out")

	fmt.Println(`go test ./... -coverprofile="resultado.out"`)
	fmt.Println(`go tool cover -func resultado.out`)
	fmt.Println(`go tool cover -html resultado.out`)

}
