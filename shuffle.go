package shuffle

import (
	"time"
	"reflect"
	"math/rand"
)

// NewRandSource creates a new random source using UnixNano
func NewRandSource() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

// Suffle reorders in-place the values of an slice
func Shuffle(slice interface{}, source rand.Source) {
    rv     := reflect.ValueOf(slice)
    swap   := reflect.Swapper(slice)
	random := rand.New(source)
    length := rv.Len()
    for i := length - 1; i > 0; i-- {
            j := random.Intn(i + 1)
            swap(i, j)
    }
}
