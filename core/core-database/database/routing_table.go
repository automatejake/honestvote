package database

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ExistsInTable(ipaddr string) (bool, string) {
	// data, err := ioutil.ReadFile("routingtable.txt")
	file, err := os.Open("routingtable.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if words[0] == ipaddr {
			fmt.Println("Same")
		} else {
			fmt.Println(words[0], "\n", ipaddr)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false, "ipaddress"
}
