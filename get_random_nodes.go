package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"reflect"
)

func GetRandomIndices(indiceCount, highestIndex int) []int {
	if highestIndex+1 < indiceCount {
		return nil
	}

	indices := make([]int, 0)
	if highestIndex+1 == indiceCount {
		for i := 0; i < indiceCount; i++ {
			indices = append(indices, i)
		}
		return indices
	}

	for len(indices) < indiceCount {
		n := RandomInt(highestIndex + 1)
		if IndexInSlice(indices, n, numbericEqual) != -1 {
			continue
		}
		indices = append(indices, n)
	}
	return indices
}
func numbericEqual(a interface{}, b interface{}) bool {
	return a.(int) == b.(int)
}

type Equals func(a interface{}, b interface{}) bool

// IndexInSlice returns the index of given object o in array
func IndexInSlice(array interface{}, o interface{}, equals Equals) int {
	arr := reflect.ValueOf(array)
	for i := 0; i < arr.Len(); i++ {
		if equals(arr.Index(i).Interface(), o) {
			return i
		}
	}
	return -1
}

func RandomInt(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("Got invalid (non positive) value: %d", n))
	}
	m := int(RandomUInt64()) % n
	if m < 0 {
		return n + m
	}
	return m
}
func RandomUInt64() uint64 {
	b := make([]byte, 8)
	_, err := io.ReadFull(cryptorand.Reader, b)
	if err == nil {
		n := new(big.Int)
		return n.SetBytes(b).Uint64()
	}
	rand.Seed(rand.Int63())
	return uint64(rand.Int63())
}

func main() {

	fmt.Println(RandomUInt64())
	fmt.Println(RandomInt(5))
	fmt.Println(GetRandomIndices(4, 9))

	nodes := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5"}
	indices := GetRandomIndices(2, len(nodes)-1)
	nodeSelects := make([]string, 0)
	for _, index := range indices {
		nodeSelects = append(nodeSelects, nodes[index])

	}

	fmt.Printf("nodeSelects:%v\n", nodeSelects)

}
