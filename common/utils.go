package common

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

func MustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	CheckErr(err, "Failed to convert to integer")
	return result
}

func SetIntersection[T constraints.Ordered](pS ...[]T) []T {
	hash := make(map[T]*int) // value, counter
	result := make([]T, 0)
	for _, slice := range pS {
		duplicationHash := make(map[T]bool) // duplication checking for individual slice
		for _, value := range slice {
			if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
				if counter := hash[value]; counter != nil { // is found in hash counter map
					if *counter++; *counter >= len(pS) { // is found in every slice
						result = append(result, value)
					}
				} else { // not found in hash counter map
					i := 1
					hash[value] = &i
				}
				duplicationHash[value] = true
			}
		}
	}
	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Remove ith element from slice
func RemoveIndex[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

// Insert element into slice at index
func Insert[T any](slice []T, index int, element T) []T {
	slice = append(slice, element)
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}

// Integer exponentiation
func IntPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

// Concatenate two ints
func ConcatInts(a, b int) int {
	strValue := fmt.Sprintf("%d%d", a, b)
	return MustAtoi(strValue)
}
