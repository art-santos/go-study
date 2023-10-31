//Write a comment describing what this code does, as well as any other information you think is relevant. Include any function names, variable names, or other identifiers that you think are important. You can also include any other information that you think is relevant, such as the purpose of the code, the context in which it is used, or any other information that you think is relevant.

package fibonacci

import (
	"math/big"
)

func Fibonacci(n int) []*big.Int {
	seq := make([]*big.Int, n)
	seq[0], seq[1] = big.NewInt(1), big.NewInt(1)
	for i := 2; i < n; i++ {
		seq[i] = new(big.Int).Set(seq[i-1])
		seq[i].Add(seq[i], seq[i-2])
	}
	return seq
}
