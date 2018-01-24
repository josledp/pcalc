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

const (
	e   = "2.71828182845904523536028747135266249775724709369995957496696762772407663035354759457138217852516642742746"
	pi  = "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214"
	phi = "1.61803398874989484820458683436563811772030917980576286213544862270526046281890244970720720418939113748475"
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

var precission uint
var pprecission int

func main() {

	flag.UintVar(&precission, "precission", 512, "Setup float precission")
	flag.IntVar(&pprecission, "print-precission", 64, "Setup float precission for printing")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	p = new(pile)

	for scanner.Scan() {
		entries := strings.Split(scanner.Text(), " ")
		for _, e := range entries {
			if e == "" {
				continue
			}
			str, err := parseEntry(e)
			if err != nil {
				log.Printf("error parsing %s: %v", e, err)
				break
			}
			if str != "" {
				fmt.Println(str)
			}
		}
	}
}

func parseEntry(s string) (string, error) {
	if isFloat(s) {
		f, err := parseFloat(s)
		if err != nil {
			return "", fmt.Errorf("error parsing bigfloat: %v", err)
		}
		return "", p.Push(f)
	}
	switch s {
	case "p":
		f, err := p.Pop()
		if err != nil {
			return "", err
		}
		return f.Text('f', pprecission), p.Push(f)
	case "n":
		f, err := p.Pop()
		if err != nil {
			return "", err
		}
		return f.Text('f', pprecission), nil
	case "pp":
		f, err := p.Pop()
		if err != nil {
			return "", err
		}
		tmp, _ := f.Int64()
		pprecission = int(tmp)
	case "pi":
		return parseEntry(pi)
	case "e":
		return parseEntry(e)
	case "phi":
		return parseEntry(phi)
	case "+":
		return "", dualOp(new(big.Float).Add)
	case "*":
		return "", dualOp(new(big.Float).Mul)
	case "/":
		return "", dualOp(new(big.Float).Quo)
	case "-":
		return "", dualOp(new(big.Float).Sub)
	case "^":
		return "", dualOp(bigfloat.Pow)
	case "v":
		return "", singleOp(bigfloat.Sqrt)
	case "q":
		//FIXME: we should not do an os.Exit from here
		os.Exit(0)
	default:
		return "", fmt.Errorf("unkown command")
	}
	return "", nil
}

func isFloat(s string) bool {
	//this should never return an error
	match, err := regexp.MatchString("^-?[0-9\\.]+(e-?[0-9]+)?$", s)
	if err != nil {
		log.Printf("error matching: %v", err)
	}
	return match
}

func parseFloat(s string) (*big.Float, error) {
	f, _, err := big.ParseFloat(s, 10, precission, big.ToNearestEven)
	return f, err
}

func singleOp(operation func(f *big.Float) *big.Float) error {
	f, err := p.Pop()
	if err != nil {
		return err
	}
	return p.Push(operation(f))
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
	return p.Push(operation(f2, f1))
}
