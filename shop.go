package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

var (
	Products = map[string]int{
		"apple":     20,
		"orange":    30,
		"banana":    40,
		"pineapple": 50,
	}
	Users = map[string]int{
		"вася": 300,
		"петя": 30000000,
		"анна": 5000,
	}

	Orders = map[int]string{
		1: "вася",
		2: "петя",
		3: "вася",
	}

	addList = map[string]int{
		"pear":   60,
		"carrot": 80,
	}
	updList = map[string]int{
		"banana":    100,
		"pineapple": 100,
	}

	updNameList = map[string]string{
		"pineapple":    "pineapplePen",
	}
	sumList = [] string {"banana", "orange"}

	superSumList = map[int][]string{
		1: sumList,
		2: {"banana", "orange"},
		3: {"pineapple", "apple", "pineapple"},
	}
)
//метод добавления
func addProduct(in map[string]int, products map[string]int) {
	for key, value := range in {
		products[key] = value
	}
}

func TestAdd(t *testing.T) {
	addProduct(addList,Products)
	println(Products)
}

//метод удаления
func delProduct(arrayForAdd []string, products map[string]int) {
	for _,product := range arrayForAdd {
		delete(products, product)
	}
}
func TestDelete(t *testing.T) {
	delList := []string {"apple"}
	delProduct(delList,Products)
	println(Products)
}
//метод изменения имени
func updateName(names map[string]string, products map[string]int) {
	for oldName, newName := range names {
		price := products[oldName]
		delete(names, oldName)
		products[newName] = price
	}
}
func TestUpdateName(t *testing.T) {
	updateName(updNameList,Products)
	println(Products)
}
//изменения цены
func updatePrice(in map[string]int, products map[string]int) {
	for name, price := range in {
		_, ok := products[name]
		if ok {
			products[name] = price
		}
	}
}
func TestUpdatePrice(t *testing.T) {
	updatePrice(updList,Products)
	println(Products)
}
//сумма заказа
func basicSum(in []string, products map[string]int) (out int) {
	for _, name := range in {
		price, ok := products[name]
		if ok {
			out += price
		}
	}
	return
}
func TestBasicSum(t *testing.T) {
	basicSum(sumList,Products)
	println(Products)
}

//сумма заказов
func superSum(in map[int][]string, users map[string]int, products map[string]int, orders map[int]string) map[int]int {
	out := make(map[int]int)
	history := make(map[string]int)

	for id, order := range in {
		// ключ - отсортированные лексиграфически продукты
		sort.Strings(order)

		key := strings.Join(order, "")
		//история заказов
		sum, exist := history[key]
		if exist {
			out[id] = sum
		} else {
			new_sum := basicSum(order, products)
			history[key] = new_sum
			out[id] = new_sum
		}

		getMoney(id, out[id], users, orders)
	}
	return out
}

func TestSuperSum(t *testing.T) {
	superSum(superSumList,Users,Products,Orders)
	println(Products)
	println(Users)
}

// Снять деньги у пользователя
func getMoney(in_order_id int, in_sum int, users map[string]int, orders map[int]string) {
	money, exist := users[orders[in_order_id]]
	if exist && money >= in_sum {
		users[orders[in_order_id]] -= in_sum
	}
}

// сортировка имено
func sortByKey(mode bool,users map[string]int) {
	var keys []string
	for login, _ := range users {
		keys = append(keys, login)
	}

	if mode {
		sort.Strings(keys)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	for _, id := range keys {
		fmt.Println(id, users[id])
	}
}

func sortByValue(mode bool, users map[string]int) {
	var (
		keys   []string
		values []int
	)

	for login, money := range users {
		keys = append(keys, login)
		values = append(values, money)
	}

	keys, values = sortByValue_(keys, values, mode)

	for i := 0; i < len(values); i++ {
		fmt.Println(keys[i], values[i])
	}
}

func sortByValue_(keys []string, values []int, mode bool) (out_keys []string, out_values []int) {
	if len(values) < 2 {
		return keys, values
	}

	left, right := 0, len(values)-1
	pvt := rand.Int() % len(values)

	values[pvt], values[right] = values[right], values[pvt]
	keys[pvt], keys[right] = keys[right], keys[pvt]

	for i, _ := range values {
		if mode {
			if values[i] < values[right] {
				values[left], values[i] = values[i], values[left]
				keys[left], keys[i] = keys[i], keys[left]

				left++
			}
		} else {
			if values[i] > values[right] {
				values[left], values[i] = values[i], values[left]
				keys[left], keys[i] = keys[i], keys[left]

				left++
			}
		}
	}

	values[left], values[right] = values[right], values[left]
	keys[left], keys[right] = keys[right], keys[left]

	sortByValue_(keys[:left], values[:left], mode)
	sortByValue_(keys[left+1:], values[left+1:], mode)

	return keys, values
}

func TestSort(t *testing.T) {
	sortByKey(true,Users)
	sortByKey(false,Users)
	sortByValue(true,Users)
	sortByValue(false,Users)
	println(Users)
}
