package utils

import (
	_"fmt"
	_ "math"
	_ "reflect"

	"github.com/ZazaRy/unicornXE/pkg/characters"
)


func All[T any](slice []T, condition func(T) bool) bool{
    for i:=0; i<len(slice); i++{
        if !condition(slice[i]){
            return false
        }
    }
    return true
}

func Any[T any](slice []T, condition func(T) bool) bool{
    for i:=0; i<len(slice); i++{
        if condition(slice[i]){
            return true
        }
    }
    return false
}


func Min[T any](values []T, method func(T) int) T{
    minIndex := values[0]
    for i:=1; i < len(values); i++{
        if method(values[i]) < method(minIndex){
            minIndex = values[i]
        }
    }
    return minIndex
}

func Max[T any](values []int, method func(int) int) int{
    max := values[0]
    for i:=1; i < len(values); i++{
        if method(values[i]) > method(max){
            max = values[i]
        }
    }
    return max
}



func Sum[T int | float64](slice []T) T{
    var res T
    for i:=0; i < len(slice);i++{
        res += slice[i]
    }
    return res
}


func partition(combatants []characters.BaseCombatant, low, high int) int{
    pivot := combatants[high].GetInitiative()
    i := low

    for j:= low; j < high; j++{
        if combatants[j].GetInitiative() > pivot{
            combatants[i], combatants[j] = combatants[j], combatants[i]
            i++
        }
    }
    combatants[i], combatants[high] = combatants[high], combatants[i]
    return i
}


func QuickSort(combatants []characters.BaseCombatant, low, high int){
    if low < high {
        pivot := partition(combatants, low, high)

        QuickSort(combatants, low, pivot-1)
        QuickSort(combatants, pivot+1, high)
    }

}
