import (
	"fmt"
	"os"
	"crypto/rand"
	"math/big"
	"github.com/fatih/color"
)

func generateRandNum(seed int64) int64 {
	num, err := rand.Int(rand.Reader, big.NewInt(seed))
	if err != nil {
                panic("KABOOM!")
    }
	n := num.Int64()
	return n
}

func operation(n int64, d int64, opr string) int64{
	var col = color.New(color.FgWhite).SprintFunc()
	mycol := []string{"red","blue","green","yellow"}
	mcol := mycol[generateRandNum(4)]
	switch mcol {
	case "red":
		col = color.New(color.FgRed).SprintFunc()
		color.Red("To solve captcha use %s in red.", "numbers")
	case "blue":
		col = color.New(color.FgBlue).SprintFunc()
		color.Blue("To solve captcha use %s in blue.", "numbers")
	case "green":
		col = color.New(color.FgGreen).SprintFunc()
		color.Green("To solve captcha use %s in green.", "numbers")
	case "yellow":
		col = color.New(color.FgYellow).SprintFunc()
		color.Yellow("To solve captcha use %s in yellow.", "numbers")
	}
	switch opr {
	case "+":
		fmt.Printf("Please provide the result for %v%v%v plus %v%v%v\n",generateRandNum(1000), col(n),generateRandNum(1000), generateRandNum(1000),col(d),generateRandNum(1000))
		return n+d
	case "-":
		fmt.Printf("Please provide the result for %v%v%v minus %v%v%v\n",generateRandNum(1000), col(n),generateRandNum(1000), generateRandNum(1000),col(d),generateRandNum(1000))
		return n-d
	case "%":
		fmt.Printf("Please provide the remainder for %v%v%v divided by %v%v%v\n",generateRandNum(1000), col(n),generateRandNum(1000), generateRandNum(1000),col(d),generateRandNum(1000))
		return n%d
	case "*":
		fmt.Printf("Please provide the result for %v%v%v multiply %v%v%v\n",generateRandNum(1000), col(n),generateRandNum(1000), generateRandNum(1000),col(d),generateRandNum(1000))
		return n*d
	}
	return -1
}

func captcha() {
		operators := []string{"+","-","%","*"}
		opr := operators[generateRandNum(4)]
 		n := generateRandNum(10) + 1	//add 1 to generate number between 1 to 10 rather 0 to 9
		d := generateRandNum(10) + 1	//add 1 to generate number between 1 to 10 rather 0 to 9
        r := operation(n,d,opr)

        var ans int64
        fmt.Scanln(&ans)
        if ans != r {
                fmt.Printf("Sorry! was expecting %v\n", r)
                os.Exit(1)
        }
}

func main(){
	captcha()
}
