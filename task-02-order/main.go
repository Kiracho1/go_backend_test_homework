package main

import "fmt"

func calculateOrder(catalog map[string]int, order []string) (int, []string) {
    total := 0
    catalogNotFound := []string{}

    for _, item := range order {
        price, ok := catalog[item]
        if ok {
            total += price
        } else {
            catalogNotFound = append(catalogNotFound, item)
        }
    }
    return total, catalogNotFound
}

func main() {
    catalog := map[string]int{
        "Клавиатура": 3500,
        "Мышь":       1500,
        "Монитор":    24000,
        "Наушники":   5200,
        "Веб-камера": 6900,
    }

    order := []string{
        "Мышь",
        "Монитор",
        "Коврик",
        "Наушники",
        "Клавиатура",
        "Принтер",
    }
    
    total, catalogNotFound := calculateOrder(catalog, order)
    fmt.Printf(
        "Стоимость заказа: %d\n"+
        "Товары не найдены: %v\n",
        total, catalogNotFound,
    )
}