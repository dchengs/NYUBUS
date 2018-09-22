package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Input reads user input to stdin and convert it to string. It also strips \n from the end
func Input(target string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s :", target)
	output, _ := reader.ReadString('\n')
	output = strings.Replace(output, "\n", "", -1)
	return output
}
