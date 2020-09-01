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
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return number, nil
}
