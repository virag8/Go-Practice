package main
import(
	"fmt"
)
func main() {
	yamlFile, err := ioutil.ReadFile("../data.yaml")
    if err != nil {
        fmt.Printf("yamlFile.Get err   #%v ", err)
    }
}