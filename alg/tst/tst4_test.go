package tst

import (
	"fmt"
	"math"
	"testing"
)

//禁止搜索答案，可以搜索类库的使用
//请在备注里面写下搜索的关键词以及查看的网页的链接

// 1.Prime number
// 1.1 Given a positive integer number n. Determine whether n is a prime number or not. (You may implement your program in any programming language.)

//   Example: 2 is a prime number. 6 is not a prime number
//     (1). input：n = 2   output: True
//     (2). input: n = 6   output: False

// 1.2 Please write down the time complexity of your program in terms of big O notation. Is it possible to improve the program?
func IsPrime(number int) bool {
	for i:=2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return true
		}
	}

	return true
}


//禁止搜索答案，可以搜索类库的使用
//请在备注里面写下搜索的关键词以及查看的网页的链接
// 关键词：go fmt 方法
// 链接： https://www.google.com/search?q=go+fmt+%E6%96%B9%E6%B3%95&oq=go+fmt+%E6%96%B9%E6%B3%95&aqs=chrome..69i57j0i546l2j69i64.9482j0j7&sourceid=chrome&ie=UTF-8

// 2.Largest number
// 2.1 Given a list of non negative integers, arrange them such that they form the largest number. (You may implement your program in pseudocode or any programming language.)

//   Example:
//     (1). input: [10,2]          output: "210"
//     (2). input: [3,30,34,5,9]   output: "9534330"

// 2.2 Is it possible to improve the program?

func IsFirstNLager(n,m int) bool {
	strN := fmt.Sprintf("%d%d",n,m)
	strM := fmt.Sprintf("%d%d",m,n)

	return strN > strM
}

func QuickSortNumber(nums []int) {
	if len(nums) <= 1 {
		return
	}
	head := nums[0]
	i,p := 0,len(nums)-1
	for i < p {
		if IsFirstNLager(nums[i],head) {
			i++
		} else {
			nums[i],nums[p] = nums[p],nums[i]
			p--
		}
	}

	if IsFirstNLager(head,nums[i]) {
		nums[i],nums[0] = nums[0],nums[i]
	}

	QuickSortNumber(nums[:i])
	QuickSortNumber(nums[i:])
}


func LargestNumber(nums []int) string {
	QuickSortNumber(nums)
	//sort.Slice(nums, func(i,j int)bool{
	//	s1 := strconv.Itoa(nums[i])
	//	s2 := strconv.Itoa(nums[j])
	//	return s1+s2 > s2+s1
	//})

	var s string
	for _,n:= range nums {
		s = fmt.Sprintf("%s%d",s,n)
	}

	return s
}

func TestMethod(t *testing.T) {
	fmt.Println("is prime 2 = ", IsPrime(2))

	fmt.Println("IsFirstNLager is: ", IsFirstNLager(5,30))

	fmt.Println("largest number 10, 2 =", LargestNumber([]int{3,30}))
	fmt.Println("largest number =", LargestNumber([]int{3,30,34,5,9}))
}
