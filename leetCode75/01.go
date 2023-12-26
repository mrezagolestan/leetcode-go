/*
* --- 1768. Merge Strings Alternately ---
* https://leetcode.com/problems/merge-strings-alternately
* You are given two strings word1 and word2.
* Merge the strings by adding letters in alternating order, starting with word1. 
* If a string is longer than the other, append the additional letters onto the end of the merged string.

*/
package LeetCode75
import ("fmt")

func main2(){
    fmt.Println(mergeAlternately("abc","pqr"));
    fmt.Println(mergeAlternately("ab","pqrs"));
    fmt.Println(mergeAlternately("abcd","pq"));
    
}
func mergeAlternately(word1 string, word2 string) string {
    var final string
    lenWord1 := len(word1)
    lenWord2 := len(word2)
    iteration := max(lenWord1, lenWord2)


    for i := 0; i < iteration; i++ {
        if i < lenWord1 {
            final += string(word1[i])
        }
        if i < lenWord2 {
            final += string(word2[i])
        }
    }
    return final
}

func max(number1 int, number2 int) int {
    if number1 > number2 {
        return number1;
    }else{
        return number2;
    }
}