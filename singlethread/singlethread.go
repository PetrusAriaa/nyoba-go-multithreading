package singlethread

import (
	fio "bufio"
	"fmt"
	"os"
	"time"
)

func readLines(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("an error occured")
		return
	}

	defer file.Close()

	s := fio.NewScanner(file)
	_temp := 0
	for s.Scan() {
		_temp++
		// line := s.Text()
		// fmt.Sprint(line)
		// fmt.Printf("Read: %s\n", line)
	}

	if err := s.Err(); err != nil {
		fmt.Println("an error occured")
	}

}

func SingleThread() {
	// var wg sync.WaitGroup
	filename := []string{
		"./source/dummy_1.txt",
		"./source/dummy_2.txt",
		"./source/dummy_3.txt",
		"./source/dummy_4.txt",
		"./source/dummy_5.txt",
		"./source/dummy_6.txt",
		"./source/dummy_7.txt",
		"./source/dummy_8.txt",
		"./source/dummy_9.txt",
		"./source/dummy_10.txt",
	}
	startTime := time.Now()

	for _, file := range filename {
		readLines(file)
	}
	end := time.Now()
	diff := end.Sub(startTime)

	fmt.Printf("Singlethread Time: %v\n", diff)
}
