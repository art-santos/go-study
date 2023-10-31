// NumberToWords converts a number to a string
// Example: 123 => "one hundred twenty-three"

package naming

import (
    "github.com/divan/num2words"
    "math/big"
    "fmt"
)

func NumberToWords(number interface{}) (string, error) {
    var num int
    switch n := number.(type) {
    case int:
        num = n
    case *big.Int:
        num64 := n.Int64()
        if num64 > (1<<31 - 1) || num64 < -(1<<31) {
            return "", fmt.Errorf("Number is out of int range")
        }
        num = int(num64)
    default:
        return "", fmt.Errorf("Unsupported number type")
    }

    words := num2words.Convert(num)

    return words, nil
}
