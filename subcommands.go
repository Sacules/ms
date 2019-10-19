package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	ms "gitlab.com/sacules/ms/schedule"
)

// Sometimes when creating a new weekly block,
// I forget that I may have not listened or rated something in the last block.
// Check that before creating a new block.

func checkLastBlock() {

}

func newBlock() {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		q       = new(ms.Queue)
	)

	err := q.Load()
	if err != nil {
		fmt.Printf("Error reading queue: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Name for new block: ")
	scanner.Scan()
	name := scanner.Text()

records:
	fmt.Printf("Amount of records for this week: ")
	scanner.Scan()
	num := scanner.Text()

	n, err := strconv.Atoi(num)
	if n <= 0 {
		err = fmt.Errorf("amount of records must be an integer greater than 0")
	}
	if err != nil {
		fmt.Printf("Error with amount of record: %s\n", err)
		fmt.Println("Pleae try again.")
		goto records
	}

	albums := make([]ms.Album, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Album %d: ", i+1)
		scanner.Scan()
		al := scanner.Text()
		if al == "" {
			break
		}

		albums[i].Name = al
	}

	b := &ms.Block{name, albums}
	q.Add(b)

	err = q.Save()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
