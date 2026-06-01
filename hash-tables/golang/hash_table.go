package main

import "fmt"

type Element struct {
	key        int
	next, prev *Element
}

type HashTable struct {
	bucket     []*Element
	bucketSize int
}

func (ht *HashTable) put(key int) {
	hashValue := ht.calculateHashValue(key)
	z := &Element{key, nil, nil}

	z.next = ht.bucket[hashValue]
	if ht.bucket[hashValue] != nil {
		ht.bucket[hashValue].prev = z
	}
	ht.bucket[hashValue] = z
}

func (ht *HashTable) remove(key int) {
	z := ht.contains(key)
	if z == nil {
		fmt.Printf("%d is not found in the hash table\n", key)

		return
	}

	hashVal := ht.calculateHashValue(z.key)
	if z.prev != nil {
		z.prev.next = z.next
	} else {
		ht.bucket[hashVal] = z.next
	}
	if z.next != nil {
		z.next.prev = z.prev
	}
}

func (ht *HashTable) contains(key int) *Element {
	hashVal := ht.calculateHashValue(key)
	for p := ht.bucket[hashVal]; p != nil; p = p.next {
		if p.key == key {
			return p
		}
	}

	return nil
}

func (ht *HashTable) printHashTable() {
	for i := 0; i < ht.bucketSize; i++ {
		fmt.Printf("%d : ", i)

		for p := ht.bucket[i]; p != nil; p = p.next {
			fmt.Printf("%d", p.key)
			if p.next != nil {
				fmt.Print(" -> ")
			} else {
				fmt.Print("\n")
			}
		}
	}
}

func (ht *HashTable) calculateHashValue(key int) int {
	return key % ht.bucketSize
}

func main() {
	bucketSize := 13
	ht := &HashTable{make([](*Element), bucketSize), bucketSize}

	for i := range 100 {
		ht.put(i)
	}
	ht.printHashTable()

	for {
		fmt.Print("1:put 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
			ht.put(key)
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
			ht.remove(key)
		case 3:
			ht.printHashTable()
		default:
			fmt.Println("invalid operation")
		}
	}
}
