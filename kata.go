package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"strconv"
)

type crackingUnit struct {
	start   int
	end     int
	padding int
	target  string
	result  chan<- string
	done    <-chan struct{}
}

func crackPart(cu crackingUnit) {
	for i := cu.start; i <= cu.end; i += 1 {
		select {
		case <-cu.done:
			fmt.Println("Woot!!!")
			cu.result <- ""
			return
		default:
			paddedString := fmt.Sprintf("%0*d", cu.padding, i)
			currentHash := md5.Sum([]byte(paddedString))
			currentHashString := hex.EncodeToString(currentHash[:])
			if currentHashString == cu.target {
				fmt.Println("Found!!!")
				cu.result <- paddedString
				<-cu.done
				return
			}
		}
	}
	fmt.Println("Wow!!!")
	cu.result <- ""
	<-cu.done
}

const maximum = 99999999

func Crack(hash string) string {
	workUnits := divideIntegers()
	result := make(chan string)
	defer close(result)
	done := make(chan struct{}, len(workUnits))
	defer close(done)
	padding := len(strconv.Itoa(maximum))

	for _, pair := range workUnits {
		cu := crackingUnit{
			start:   pair[0],
			end:     pair[1],
			padding: padding,
			target:  hash,
			result:  result,
			done:    done,
		}
		go crackPart(cu)
	}

	var correctAnswer string
	for range workUnits {
		data := <-result
		if data != "" {
			correctAnswer = data
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
