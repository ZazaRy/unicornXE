package utils

import (
	"github.com/ZazaRy/unicornXE/pkg/characters"
	"math/rand"
)


func All[T any](teams []T, condition func(T) bool) bool {
	for i := 0; i < len(teams); i++ {
		if !condition(teams[i]) {
			return false
		}
	}
	return true
}

func Any[T any](teams []T, condition func(T) bool) bool {
	for i := 0; i < len(teams); i++ {
		if condition(teams[i]) {
			return true
		}
	}
	return false
}

func Min[T any](values []T, method func(T) int) T {
	minIndex := values[0]
	for i := 1; i < len(values); i++ {
		if method(values[i]) < method(minIndex) {
			minIndex = values[i]
		}
	}
	return minIndex
}

func Max(values []int, method func(int) int) int {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if method(values[i]) > method(max) {
			max = values[i]
		}
	}
	return max
}

func Sum[T int | float64](slice []T) T {
	var res T
	for i := 0; i < len(slice); i++ {
		res += slice[i]
	}
	return res
}

func partition(combatants []characters.BaseCombatant, low, high int) int {
	pivotIndex := rand.Intn(high-low+1) + low
	combatants[pivotIndex], combatants[high] = combatants[high], combatants[pivotIndex]
	pivot := combatants[high].GetInitiative()
	i := low

	for j := low; j < high; j++ {
		if combatants[j].GetInitiative() > pivot {
			combatants[i], combatants[j] = combatants[j], combatants[i]
			i++
		}
	}
	combatants[i], combatants[high] = combatants[high], combatants[i]
	return i
}

func QuickSort(combatants []characters.BaseCombatant, low, high int) {
	for low < high {
		pi := partition(combatants, low, high)

		if pi-low < high-pi {
			QuickSort(combatants, low, pi-1)
			low = pi + 1
		} else {
			QuickSort(combatants, pi+1, high)
			high = pi - 1
		}
	}
}

