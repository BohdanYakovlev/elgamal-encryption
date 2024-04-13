package main

import (
	"fmt"
	"math/rand"
)

type openKey struct {
	y int
	g int
	p int
}

func numbers() (res map[int]int) {
	res = map[int]int{}
	res[1] = 257
	res[2] = 263
	res[3] = 269
	res[4] = 271
	res[5] = 277
	res[6] = 281
	res[7] = 283
	res[8] = 293
	res[9] = 307
	res[10] = 311
	return
}

func getRandomNumber() int {
	return numbers()[rand.Intn(10)+1]
}

func setKey() (openKey, int) {
	key := openKey{p: getRandomNumber()}
	key.g = getPrimitiveRoot(key.p)
	closeKey := rand.Intn(key.p-3) + 2
	key.y = pow(key.g, closeKey, key.p)
	return key, closeKey
}

func mod(n, m int) int {
	return n % m
}

func areCoprime(n1, n2 int) bool {
	for n2 == 0 {
		n1, n2 = n2, n1%n2
	}
	return n1 == 1
}

func pow(n, p, m int) int {
	res := n
	for i := 2; i <= p; i++ {
		res *= n
		res = mod(res, m)
	}
	return res
}

func getSimpleFactors(n int) []int {
	factors := []int{}
	index := 2
	for n > 1 {
		if mod(n, index) == 0 {
			factors = append(factors, index)
			for mod(n, index) == 0 {
				n /= index
			}
		}
		index++
	}
	return factors
}

func isPrimitiveRoot(g, n int) bool {
	e := n - 1
	factors := getSimpleFactors(e)
	for f := range factors {
		if pow(g, e/factors[f], n) == 1 {
			return false
		}
	}
	return true
}

func getPrimitiveRoot(n int) int {
	for i := 2; i < n; i++ {
		if isPrimitiveRoot(i, n) {
			return i
		}
	}
	return 0
}

type cipText struct {
	a int
	b int
}

func elGamalCyp(key openKey, message int) cipText {
	tempKey := rand.Intn(key.p-2) + 1
	return cipText{pow(key.g, tempKey, key.p), mod(pow(key.y, tempKey, key.p)*message, key.p)}
}

func decodeElGamal(hash cipText, oKey openKey, cKey int) int {
	return mod(pow(hash.a, oKey.p-1-cKey, oKey.p)*hash.b, oKey.p)
}

func main() {
	m1, m2 := 20, 10
	oKey, cKey := setKey()
	cyp1 := elGamalCyp(oKey, m1)
	fmt.Println(cyp1, m1)
	cip2 := elGamalCyp(oKey, m2)
	fmt.Println(cip2, m2)
	temp := cipText{mod(cyp1.a*cip2.a, oKey.p), mod(cyp1.b*cip2.b, oKey.p)}
	fmt.Println(temp, decodeElGamal(temp, oKey, cKey))
}
