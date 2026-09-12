package main

import (
	"errors"
	"fmt"
)

func average(values []int) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("невозможно вычислить среднее для пустого списка")
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return float64(sum) / float64(len(values)), nil
}

func main() {
	durations := []int{
		120, 95, 110, 140, 85,
	}
	avg, err := average(durations)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Среднее время обработки: %.2f мс\n", avg)
	}
}
