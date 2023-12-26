/*
* --- 1071. Greatest Common Divisor of Strings ---
* https://leetcode.com/problems/greatest-common-divisor-of-strings
* For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
* Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
*/
package leetCode75
import (
    "fmt"
    "testing"
)

func ttt(){
    fmt.Println(gcdOfStrings("OBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNO","OBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNO"))
    fmt.Println(gcdOfStrings("CXTXNCXTXNCXTXNCXTXNCXTXN","CXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXN"))
    fmt.Println(gcdOfStrings("ABABABAB","ABAB"))
    fmt.Println(gcdOfStrings("ABCDEF","ABC"))
    fmt.Println(gcdOfStrings("ABCABC","ABC"))
    fmt.Println(gcdOfStrings("ABABAB","ABAB"))
    fmt.Println(gcdOfStrings("LEET","CODE"))
}

func TestGcdOfStrings(t *testing.T){
    fmt.Println(gcdOfStrings("OBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNO","OBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNOOBCNO"))
}

func GcdOfStrings(str1 string, str2 string) string {
    var largestDivisibleNumber int = calculateLargestDivisibleNumber( len(str1), len(str2))

    intersectionString := str1[0:largestDivisibleNumber]
    
    //check string 1 is dividable
    if !isStringDividable(str1, intersectionString) {
        return ""
    }

    //check string 2 is dividable
    if !isStringDividable(str2, intersectionString) {
        return ""
    }

    return intersectionString
}


func isStringDividable(main string, devideTo string) bool {
    for i := 0 ; i < len(main); i = i + len(devideTo) {
        if i + len(devideTo) > len(main) {
            return false
        }

        if devideTo != string(main[i: i + len(devideTo)]) {
            return false
        }
    }
    return true
}

func calculateLargestDivisibleNumber(num1 int, num2 int) int {
    //find larger & smaller number
    var larger int = 0
    var smaller int = 0
    if num1 > num2 {
        larger = num1
        smaller = num2
    }else{
        larger = num2
        smaller = num1
    }

    //find largest divisible number
    largestDivisibleNumber := 1
    for i := 1 ; i <= smaller ; i++ {
        if smaller % i == 0 {
            if larger % i == 0 {
                largestDivisibleNumber = i
            }
        }
    }
    return largestDivisibleNumber
}
