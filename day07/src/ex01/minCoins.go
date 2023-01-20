package main

import (
	"sort"
)

func minCoins(val int, coins []int) []int {
    res := make([]int, 0)
    i := len(coins) - 1
    for i >= 0 {
        for val >= coins[i] {
            val -= coins[i]
            res = append(res, coins[i])
        }
        i -= 1
    }
    return res
}

func makeResult(storage *[]int, sum int) []int {

	result := make([]int, 0)
	sort.Ints(*storage)
	i := len(*storage) - 1
	if (i < 0) {
		return []int{}
	}
	for (sum > 0 && i > -1) {
		if (sum >= (*storage)[i]) {
			sum -= (*storage)[i]
			result = append(result, (*storage)[i])
		} else {
			i -= 1
		}
	}
	if (sum != 0) {
		return []int{}
	}
	return result
}

func minCoins2(val int, coins []int) []int {
	mapValue := make(map[int]int)
	storage := make([]int, 0)

	for i := 0; i < len(coins); i++ {
		if (coins[i] <= 0) {
			return []int{}
		}
		if mapValue[coins[i]] == 0 {
			mapValue[coins[i]] = coins[i]
			storage = append(storage, coins[i])
		}
	}

	return makeResult(&storage, val)
}

// func main() {
// 	coins := []int{10, 10, 5, 100}
// 	nbr := 116
// 	fmt.Println(minCoins2(nbr, coins))

// 	tmp := []int{1,5,10}
// 	tn := 13
// 	fmt.Println(minCoins2(tn, tmp))
// }