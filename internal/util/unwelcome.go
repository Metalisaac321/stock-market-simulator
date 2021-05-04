package util
import "fmt"

func Unwelcome(name string) string {
	return fmt.Sprintf("Bye %s", name)
}