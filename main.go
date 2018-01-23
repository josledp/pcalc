package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/ALTree/bigfloat"
)

type pile struct {
	elements []*big.Float
	top      int
}

func (p *pile) Push(f *big.Float) error {
	if p.top == len(p.elements) {
		p.elements = append(p.elements, f)
	} else {
		p.elements[p.top] = f
	}
	p.top++
	return nil
}

func (p *pile) Pop() (*big.Float, error) {
	if p.top == 0 {
		return nil, fmt.Errorf("pile is empty")
	}
	p.top--
	return p.elements[p.top], nil
}

var p *pile

func main() {
	var precission uint
	var pprecission int
	flag.UintVar(&precission, "precission", 512, "Setup float precission")
	flag.IntVar(&pprecission, "print-precission", 64, "Setup float precission for printing")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	p = new(pile)

MainFor:
	for scanner.Scan() {
		entries := strings.Split(scanner.Text(), " ")
		for _, e := range entries {
			if e == "" {
				continue
			}
			match, err := regexp.MatchString("^[0-9\\.,]*$", e)
			if err != nil {
				log.Printf("error matching: %v", err)
				continue MainFor
			}
			if match {
				f, _, err := big.ParseFloat(e, 10, precission, big.ToNearestEven)
				if err != nil {
					log.Printf("error parsing bigfloat: %v", err)
					continue MainFor
				}
				p.Push(f)
				continue
			}
			switch e {
			case "p":
				f, err := p.Pop()
				if err != nil {
					log.Println(err)
					continue MainFor
				}
				fmt.Println(f.Text('f', pprecission))
				p.Push(f)
			case "n":
				f, err := p.Pop()
				if err != nil {
					log.Println(err)
					continue MainFor
				}
				fmt.Println(f.Text('f', pprecission))
			case "pp":
				f, err := p.Pop()
				if err != nil {
					log.Println(err)
					continue MainFor
				}
				tmp, _ := f.Int64()
				pprecission = int(tmp)
			case "+":
				err = dualOp(new(big.Float).Add)
			case "*":
				err = dualOp(new(big.Float).Mul)
			case "/":
				err = dualOp(new(big.Float).Quo)
			case "-":
				err = dualOp(new(big.Float).Sub)
			case "^":
				err = dualOp(bigfloat.Pow)
			case "v":
				err = singleOp(bigfloat.Sqrt)
			case "q":
				os.Exit(0)
			default:
				log.Printf("unkown command %s", e)
			}
			if err != nil {
				log.Println(err)
				continue MainFor
			}
		}
	}
}

func singleOp(operation func(f *big.Float) *big.Float) error {
	f, err := p.Pop()
	if err != nil {
		return err
	}
	err = p.Push(operation(f))
	if err != nil {
		return err
	}
	return nil
}
func dualOp(operation func(f1, f2 *big.Float) *big.Float) error {
	f1, err := p.Pop()
	if err != nil {
		return err
	}
	f2, err := p.Pop()
	if err != nil {
		return err
	}
	err = p.Push(operation(f2, f1))
	if err != nil {
		return err
	}
	return nil
}
