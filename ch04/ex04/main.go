package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	rotate(s, 1)
	fmt.Println(s)
}

func rotate(s []int, cnt int) {
	if cnt < 0 || len(s) < cnt+1 {
		return
	}

	bk := s
	target := make([]int, len(s[:cnt]))
	copy(target, s[:cnt])

	fmt.Println(target)
	for i := range s {
		// そのあとは末尾に移動する対象のindexから入れる
		if len(s) <= i+cnt {
			fmt.Println(cnt + i - len(s))
			s[i] = target[cnt+i-len(s)]
			continue
		}

		// 末尾に行くまではcnt+iを入れる
		s[i] = bk[cnt+i]
	}

}
