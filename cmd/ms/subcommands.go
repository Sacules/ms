package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gitlab.com/sacules/ms"
)

func newblock() error {
	scanner := bufio.NewScanner(os.Stdin)
	q := new(ms.Queue)

	err := q.Load()
	if err != nil {
		return fmt.Errorf("newblock: %v", err)
	}

	fmt.Printf("Name for new block: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Amount of records for this week: ")
	scanner.Scan()
	num := scanner.Text()

	n, err := strconv.Atoi(num)
	if err != nil {
		return fmt.Errorf("newblock: %v", err)
	}

	if n <= 0 {
		return fmt.Errorf("newblock: amount of records must be greater than 0")
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

	b := &ms.Block{Name: name, Albums: albums}
	q.Add(b)

	err = q.Save()
	if err != nil {
		return fmt.Errorf("newblock: %v", err)
	}

	return nil
}

func status() error {
	q := new(ms.Queue)
	err := q.Load()
	if err != nil {
		return fmt.Errorf("status: %v", err)
	}

	var total, listened int
	for i, block := range q {
		for _, album := range block.Albums {
			if i == 2 {
				continue
			}

			if (i == 0 && album.FirstListen) ||
				(i == 1 && album.SecondListen) ||
				(i == 3 && album.ThirdListen) {
				listened++
			}

			total++
		}
	}

	fmt.Println("Total records:", total)
	fmt.Println("Total listened:", listened)

	return nil
}
