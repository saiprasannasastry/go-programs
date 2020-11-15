// Given a binary string s (a string consisting only of '0's and '1's), we can split s into 3 non-empty strings s1, s2, s3 (s1+ s2+ s3 = s).

// Return the number of ways s can be split such that the number of characters '1' is the same in s1, s2, and s3.

// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:

// Input: s = "10101"
// Output: 4
// Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
// "1|010|1"
// "1|01|01"
// "10|10|1"
// "10|1|01"
// Example 2:

// Input: s = "1001"
// Output: 0
// Example 3:

// Input: s = "0000"
// Output: 3
// Explanation: There are three ways to split s in 3 parts.
// "0|0|00"
// "0|00|0"
// "00|0|0"
// Example 4:

// Input: s = "100100010100110"
// Output: 12
package main
import "fmt"
func main(){
	fmt.Println(splitString("0000"))
	fmt.Println(splitString("1001"))
	fmt.Println(splitString("10101"))
}

func splitString(s string) int{
	n := len(s)
	count1 :=0
	mod := 1_000_000_007
	for _,v := range s{
		if v =='1'{
			count1 +=1
		}
	}
	if count1 ==0{
		return (((n-1)*(n-2))/2)%mod
	}
	if count1 %3 !=0{
		return 0
	}
	count1b := 0
	n1 ,n2 := 0,0
	oneThird := count1/3

	for _,v := range s{
		if v =='1'{
			count1b +=1
		}

	if count1b == oneThird{
		n1++
	}
	if count1b ==2*oneThird{
		n2++
	}
}
	return (n1*n2 )%mod


}