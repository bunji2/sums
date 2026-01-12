// sums.exe 合計値になる加算因数を見つけるためのプログラム
// Usage sums.exe 合計値 数列
// 例： sums.exe 14 9,1,5,6,5,4
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum, nums, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sum:", sum)
	fmt.Println("nums:", nums)

	results, counter := findOut(sum, nums)
	for idx, result := range results {
		fmt.Printf("%d:%d=>%d%% (%d/%d)\n", idx, nums[idx], result*100/counter, result, counter)
	}

}

// 引数のから sum と nums をパースする
func parseArgs() (sum int, nums []int, err error) {
	sum, err = strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	for _, s := range strings.Split(os.Args[2], ",") {
		n, e := strconv.Atoi(s)
		if e != nil {
			err = e
			return
		}
		nums = append(nums, n)
	}

	return
}

// 合計が sum となる数列を見つけ出す
// 返り値:
// r : nums の要素ごとの合計の算出に使われた回数
// counter : 合計が sums となる数列の個数
func findOut(sum int, nums []int) (r []int, counter int) {
	r = make([]int, len(nums))
	a := arrayOf(len(nums))

	for _, p := range perm(a) {
		s := 0
		for i, idx := range p {
			s += nums[idx]
			if s == sum {
				// DEBUG // disp(nums, p[:i+1])
				for j := 0; j <= i; j++ {
					r[p[j]] += 1
				}
				counter += 1
				break
			}
		}
	}
	return
}

// 数列を表示
func disp(nums, poss []int) {
	a := []string{fmt.Sprintf("[%d]%d", poss[0], nums[poss[0]])}
	for _, pos := range poss[1:] {
		a = append(a, fmt.Sprintf("[%d]%d", pos, nums[pos]))
	}
	fmt.Println(strings.Join(a, "+"))
}

func perm(a []int) (r [][]int) {
	//fmt.Println("a =", a)

	// a が空のときは空を返す
	if len(a) < 1 {
		return
	}

	// a の要素数が1のときは a だけの配列を返す
	if len(a) == 1 {
		r = [][]int{a}
		return
	}

	// a のそれぞれの要素を除外してできる順列をもとめ、
	// それぞれの先頭に除外したaの要素を追加したものを結果に加えていく
	for i := 0; i < len(a); i++ {
		for _, b := range perm(except(i, a)) {
			//fmt.Println("b =", b)
			if len(b) > 0 {
				r = append(r, append([]int{a[i]}, b...))
			}
		}
	}
	//fmt.Println(a, "=>", r)
	return
}

// 配列 a の i 番目の要素を除外した配列を返す
// 例：except(0,[0,1,2]) => [1,2]
// 例：except(1,[0,1,2]) => [0,2]
// 例：except(2,[0,1,2]) => [0,1]
func except(i int, a []int) (r []int) {
	r = []int{}
	for j := 0; j < len(a); j++ {
		if i != j {
			r = append(r, a[j])
		}
	}
	//fmt.Println("except:r =", r)
	return
}

// n までの(nを含まない)整数を要素とする配列を返す
// 例：arrayOf(3) => [0,1,2]
func arrayOf(n int) (r []int) {
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return
}
