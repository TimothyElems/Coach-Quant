package main

/*
Coin Flip Quartet

"We flip a fair coin 9 times and aim to find the probability of getting exactly 4 heads among all tosses."

This is a Binomial Distribution problem.
The goal is to write the functions needed, from scratch, and come to a solution of the original question.
Functions needed:
	- Factorial
	- Binomial Coefficient
	- Binomial Distribution
*/
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("4! is %v\n", factorial(4))
	fmt.Printf("5! is %v\n", factorial(5))
	fmt.Printf("The Binomial Coefficient for 7 of 10 is %v\n", binomialCoefficient(7, 10))
	fmt.Printf("The probability that a 60%% free throw shooter goes 7 for 10, is %v\n", binomialDistribution(10, 7, 0.6))
	fmt.Println("And now for the Grand Finale!")
	fmt.Printf("The probability of getting 4 heads on 9 coin tosses is %v", binomialDistribution(9, 4, 0.5))
}

// Factorial
func factorial(factor float64) float64 {
	if factor == 0 || factor == 1 {
		return 1
	} else {
		return factor * factorial((factor - 1))
	}
}

// Binomial Coefficient: Doesn't use PascalCase cause it won't be exported
func binomialCoefficient(k, n float64) float64 {
	return (factorial(n) / (factorial(k) * factorial((n - k))))
}

// Binomial Distribution
func binomialDistribution(n, k float64, p float64) float64 {
	bc := binomialCoefficient(k, n)
	// The power should be bootstrapped, but I'm running out of time
	bd := (bc * math.Pow(p, k) * math.Pow((1-p), (n-k))) * 100
	return math.Round(bd*100) / 100
}
