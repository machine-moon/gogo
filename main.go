package main // this is similar java

// similar to C / python

// 2 styles can do () for multi line or manual
import "fmt"
import "github.com/google/uuid"

//func keywordis defines a function
// main() is reserved for entry point

// use var keyword followed by variableName and TYPE
//var myvar int = 10 // global variable

//ORRRR

// :=
// u see that ':'
// that is like python dynamic typing
// compiler will INFER type
// very nice syntax lolz
//mySecondVar := 10

func main() {

	// so now lets go over this
	// for loop
	// see this for (init ; condition ; increment)

	// SAME AS IN C BRUH

	//copilot annoying lolz
	//

	// same as C right? k.

	for i := 0; i < 1000; i++ {
		// here we give the variable named uuid something
		uuid := uuid.New()
		fmt.Println("Generated UUID:", uuid)
		// println == print line
	}
}
