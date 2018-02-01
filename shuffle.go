package shuffle

import (
	"time"
	"reflect"
	"math/rand"
)

func NewRandSource() rand.Source {
	return rand.NewSource(time.Now().Unix())
}


// Suffle slice/array in place
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
