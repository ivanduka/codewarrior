package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"strconv"
)

type arguments struct {
	start   int
	end     int
	padding int
	target  string
	results chan<- string
	done    <-chan struct{}
}

func crackPart(args arguments) {
	for i := args.start; i <= args.end; i += 1 {
		select {
		case <-args.done:
			args.results <- ""
			return
		default:
			paddedString := fmt.Sprintf("%0*d", args.padding, i)
			currentHash := md5.Sum([]byte(paddedString))
			currentHashString := hex.EncodeToString(currentHash[:])
			if currentHashString == args.target {
				args.results <- paddedString
				<-args.done
				return
			}
		}
	}
	args.results <- ""
	<-args.done
}

const maximum = 99999

func Crack(hash string) string {
	workUnits := divideIntegers()
	results := make(chan string)
	defer close(results)
	done := make(chan struct{})
	defer close(done)
	padding := len(strconv.Itoa(maximum))

	for _, pair := range workUnits {
		args := arguments{
			start:   pair[0],
			end:     pair[1],
			padding: padding,
			target:  hash,
			results: results,
			done:    done,
		}
		go crackPart(args)
	}

	var correctAnswer string

	for range workUnits {
		result := <-results
		if result != "" {
			correctAnswer = result
			for range workUnits {
				done <- struct{}{}
			}
		}
	}

	return correctAnswer
}

func md5Hash(s string) string {
	currentHash := md5.Sum([]byte(s))
	return hex.EncodeToString(currentHash[:])
}

func divideIntegers() [][2]int {
	numCPU := runtime.NumCPU()
	slice := make([][2]int, numCPU)
	division := maximum / numCPU
	lastStart := -1
	for index := range slice {
		slice[index][0] = lastStart + 1
		if index == len(slice)-1 {
			slice[index][1] = maximum
		} else {
			slice[index][1] = lastStart + 1 + division
			lastStart += division + 1
		}
	}
	return slice
}

func main() {
	s := md5Hash("9999999")
	fmt.Println(s)
}
