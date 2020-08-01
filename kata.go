package main

import (
	"context"
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
	ch      chan<- string
	ctx     context.Context
}

func crackPart(args arguments) {
	for i := args.start; i <= args.end; i += 1 {
		select {
		case <-args.ctx.Done():
			fmt.Println("Returning early!")
			return
		default:
			paddedString := fmt.Sprintf("%0*d", args.padding, i)
			currentHash := md5.Sum([]byte(paddedString))
			currentHashString := hex.EncodeToString(currentHash[:])
			if currentHashString == args.target {
				fmt.Println("Found!")
				args.ch <- paddedString
				return
			}
		}
	}
	fmt.Println("Not found!")
}

const maximum = 99999999

func Crack(hash string) string {
	workUnits := divideIntegers()
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	padding := len(strconv.Itoa(maximum))

	for _, pair := range workUnits {
		args := arguments{
			start:   pair[0],
			end:     pair[1],
			padding: padding,
			target:  hash,
			ch:      ch,
			ctx:     ctx,
		}
		go crackPart(args)
	}

	return <-ch
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
