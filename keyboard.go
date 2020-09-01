//reads from the standard I/P i.e. keybaoard as string and then converts the entry into float 64 number
package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) //the buffer input output method reads reads the I/P
	input, err := reader.ReadString('\n') //the entered data is read till newline char and saved in the variable input and the error if saved in err
	if err != nil {
		log.Fatal(err)
	}//checks if there is an error in the data read
	input = strings.TrimSpace(input) //removes all white spaces
	number, err := strconv.ParseFloat(input, 64) // converts the input into the float 64  
	if err != nil {
		log.Fatal(err)
	} // checks if there is an error in the input
	return number, nil //returns the converted input and nil for error 
}
