package combat


import (
    _"fmt"
    "math"
    "reflect"
)


func All(slice interface{}, condition func(interface{}) bool) bool {
    v := reflect.ValueOf(slice)

    if v.Kind() != reflect.Slice {
        panic("All: not a slice")
    }

    for i := 0; i < v.Len(); i++ {
        if !condition(v.Index(i).Interface()) {
            return false
        }
    }

    return true
}


func Any(slice []int, condition func(int) bool) bool{

    for _, v := range slice{
        if condition(v){
            return true
        }
    }
    return false
}


func Min(slice []float64) float64{
    min := math.Inf(1)
    for _, v := range slice{
        if v < min {
            min = v
        }
    }
    return min
}

func Max(slice []float64) float64{
    max := math.Inf(-1)
    for _, v := range slice{
        if v > max{
            max = v
        }
    }
    return max
}
