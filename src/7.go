
import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.New(rand.NewSource(42))
	num1 := random.Intn(50)
	num2 := random.Intn(30)
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}