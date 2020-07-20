package main

import "fmt"

var (
	count   = 50
	surplus = 0
	users   = []string{
		"Matthew",
		"Sarah",
		"Augustus",
		"Heidi",
		"Emilie",
		"Peter",
		"Giana",
		"Adriano",
		"Aaron",
		"Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

/**
解法一
*/
func dispatchCoin() (map[string]int, int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'A', 'a', 'e', 'E':
				distribution[name]++
				count--
			case 'i', 'I':
				distribution[name] += 2
				count -= 2
			case 'o', 'O':
				distribution[name] += 3
				count -= 3
			case 'U', 'u':
				distribution[name] += 5
				count -= 5
			}
		}
	}
	return distribution, count
}

/**
计算金币
*/
func calc(username string) int {
	var sum int
	for _, value := range username {
		switch value {
		case 'A', 'a', 'e', 'E':
			sum = sum + 1
		case 'i', 'I':
			sum = sum + 2
		case 'o', 'O':
			sum = sum + 3
		case 'U', 'u':
			sum = sum + 5
		}
	}
	return sum
}

func case2() int {
	surplus = count
	for _, username := range users {
		allCoin := calc(username)
		surplus -= allCoin
		value, ok := distribution[username]
		// 检测map是否存在  如果不存在直接就是等于返回的金币  如果存在则为返回的金币+之前的金币
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return surplus
}

func main() {
	//coin, count := dispatchCoin()
	//fmt.Println(coin)
	//fmt.Println(count)

	coins := case2()
	fmt.Println(coins)
}
