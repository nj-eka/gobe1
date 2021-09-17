package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type Secret interface{
	Question() string
	Solution() string
	Test(answer string) bool
}

var ops = [...]string{"+", "-", "*"}

type smath struct{
	a, b int
	op string
	res string
}

func (s *smath) Question() string {
	return fmt.Sprintf("%d %s %d = ? ", s.a, s.op, s.b)
}

func (s *smath) Solution() string {
	return fmt.Sprintf("%d %s %d = %s ", s.a, s.op, s.b, s.res)
}

func (s *smath) Test(answer string) bool {
	return s.res == strings.ToLower(strings.Trim(answer, "\t\n "))
}

func NewSMGame() Secret{
	smgame := smath{
		a:   rand.Intn(10),
		b:   rand.Intn(10),
		op:  ops[rand.Intn(len(ops))],
	}
	smgame.resolve()
	return &smgame
}

func (s *smath) resolve(){
	var res int
	switch s.op{
	case "+":
		res = s.a + s.b
	case "-":
		res = s.a - s.b
	case "*":
		res = s.a * s.b
	default:
		panic("unknown operator")
	}
	s.res = strconv.Itoa(res)
}


type scores map[string]int

type pair struct{
	key string
	value int
}

func (p pair) String() string{
	return fmt.Sprintf("%s = %d", p.key, p.value)
}

func (s scores) Sorted() []pair{
	pairs := make([]pair, 0, len(s))
	for key, value := range s{
		pairs = append(pairs, pair{key, value})
	}
	sort.SliceStable(pairs, func(i, j int) bool{
		return pairs[i].value > pairs[j].value
	})
	return pairs
}
