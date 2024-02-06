package main
import (
	"fmt"
	
	"cmp"
)

func main() {
	v1 := []float64{1.3, 2.5, 3.7, 4.9}
	v2 := []int64{1, 2, 3, 4}

	fmt.Println(Sum01(v1))
	fmt.Println(Sum01(v2))
	anyType(1,"asd")
	comparableType(1,2)
	cstSlice1 := CustomSlice[int]{1,2,3,4}
	cstSlice2 := CustomSlice[string]{"a","b","c","d"}
	fmt.Println(cstSlice1)
	fmt.Println(cstSlice2)
	
	x,y:=Point(5),Point(2)
	fmt.Println(MinNumber(x,y))

	fd :=MyGenericStruct[MyFirstData]{StrValue:"Test",Data:MyFirstData{}}
	fd.Data.PrintOne()
	sd :=MyGenericStruct[MySecondData]{StrValue:"Test",Data:MySecondData{}}
	sd.Data.PrintTwo()

}

func Sum01[N int | int32 |  float32 | int64| float64](v1 []N) N {
	var result N
	for _, v := range v1 {
		result += v
	}
	return result
}
type Number interface {
	int | int32 |  float32 | int64| float64
}

func Sum02[N Number](v []N) N {
	var result N
	for _, v := range v {
		result += v
	}
	return result
}

func anyType[N1 any, N2 any](v1 N1,v2 N2)  {
	
	fmt.Println(v1,v2)
}

func comparableType[N cmp.Ordered](v1,v2 N){

	fmt.Println(v1,v2)
	fmt.Println(v1>=v2)
}

type CustomSlice[V int | int32 | int64 | float32 | float64 | string] []V

func MinNumber[T N1](x,y T)T{
	if x<y{
		return x
	}
	return y
}

type Point int

type N1 interface {
	~int | ~int32 |  ~float32 | ~int64| ~float64
}

type(
	MyGenericStruct[D any] struct {
		StrValue string
		Data D
	}
	MyFirstData struct {}
	MySecondData struct {}
)

func (d MyFirstData) PrintOne(){
	fmt.Println("Print First Data")
}

func (d MySecondData) PrintTwo(){
	fmt.Println("Print Second Data")
}