package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
)

type pile struct {
	elements []*big.Float
	top      int
}

func (p *pile) Push(f *big.Float) error {
	if p.top == cap(p.elements) {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	p := new(pile)

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
				f, _, err := big.ParseFloat(e, 10, 512, big.ToNearestEven)
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
				fmt.Println(f.Text('f', 512))
			}
		}
	}
}
