package main

type Item struct {
	Name string
	Value int32
}

func initStorage() ([]Item) {

	var itemStorage []Item

	itemStorage = append(itemStorage, Item{"CE", 10})
	itemStorage = append(itemStorage, Item{"NT", 15})
	itemStorage = append(itemStorage, Item{"DE", 21})
	itemStorage = append(itemStorage, Item{"YR", 23})
	itemStorage = append(itemStorage, Item{"AA", 15})

	return (itemStorage)
}

func fetchCandy(candyToFind string) (Item, bool) {
	for _, item := range candyStorage {
		if (item.Name == candyToFind) {
			return item, true
		}
	}
	return Item{}, false
}