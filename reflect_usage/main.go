// DeepEqual(A,B):
//		if A and B are pointer, compare with the values of pointer
//		if A is value, B is pointer, return false -> TypeOf(A) != TypeOf(B), return false
package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	ValInt    int
	ValString string
}

func (v1 *MyType) Less(v2 *MyType) bool {
	return v1.ValInt < v2.ValInt
}

func main() {
	//////////////////////////// Slice Int ///////////////////////////////////////////////
	sliceInt := make([]int, 0, 5)
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 89)
	sliceInt = append(sliceInt, 5)
	sliceInt = append(sliceInt, 45)
	sliceInt = append(sliceInt, 19)
	// same type: int
	contains, ind := SliceContains(sliceInt, 5)
	fmt.Println("Contains Slice: same type: int -> ", contains, ind)
	// int8 vs int
	contains, ind = SliceContains(sliceInt, int8(5))
	fmt.Println("Contains Slice: int8 vs int -> ", contains, ind)
	// max
	maxSliceInt, maxInd := SliceMax(sliceInt)
	fmt.Println("Max Slice: int -> ", maxSliceInt, maxInd)
	fmt.Println()

	//////////////////////////// Slice Float ///////////////////////////////////////////////
	sliceFloat := make([]float32, 0, 5)
	sliceFloat = append(sliceFloat, 1.1)
	sliceFloat = append(sliceFloat, 89.9)
	sliceFloat = append(sliceFloat, 5.5)
	sliceFloat = append(sliceFloat, 45.5)
	sliceFloat = append(sliceFloat, 8.8)
	// same type: float32
	contains, ind = SliceContains(sliceFloat, float32(5.5))
	fmt.Println("Contains Slice: same type: float32 -> ", contains, ind)
	// float64 vs float32
	contains, ind = SliceContains(sliceFloat, float64(5.5))
	fmt.Println("Contains Slice: float64 vs float32 -> ", contains, ind)
	// max
	maxSliceFloat, maxInd := SliceMax(sliceFloat)
	fmt.Println("Max Slice: float -> ", maxSliceFloat, maxInd)
	fmt.Println()

	//////////////////////////// Slice String ///////////////////////////////////////////////
	sliceString := make([]string, 0, 5)
	sliceString = append(sliceString, "a")
	sliceString = append(sliceString, "c")
	sliceString = append(sliceString, "z")
	sliceString = append(sliceString, "b")
	sliceString = append(sliceString, "i")
	contains, ind = SliceContains(sliceFloat, "c")
	fmt.Println("Contains Slice: string -> ", contains, ind)
	// max
	maxSliceString, maxInd := SliceMax(sliceString)
	fmt.Println("Max Slice: string -> ", maxSliceString, maxInd)
	fmt.Println()

	//////////////////////////// Array Int ///////////////////////////////////////////////
	arrayInt := [5]int{1, 89, 5, 45, 19}
	// same type: int
	contains, ind = SliceContains(arrayInt, 5)
	fmt.Println("Contains Array: same type: int -> ", contains, ind)
	// int8 vs int
	contains, ind = SliceContains(arrayInt, int8(5))
	fmt.Println("Contains Array: int8 vs int -> ", contains, ind)
	// max
	maxArrayInt, maxInd := SliceMax(arrayInt)
	fmt.Println("Max Array: int -> ", maxArrayInt, maxInd)
	fmt.Println()

	//////////////////////////// Slice Struct ///////////////////////////////////////////////
	sliceStruct := make([]MyType, 0, 5)
	sliceStruct = append(sliceStruct, MyType{
		ValInt:    1,
		ValString: "a",
	})
	sliceStruct = append(sliceStruct, MyType{
		ValInt:    19,
		ValString: "g",
	})
	sliceStruct = append(sliceStruct, MyType{
		ValInt:    8,
		ValString: "c",
	})
	sliceStruct = append(sliceStruct, MyType{
		ValInt:    45,
		ValString: "y",
	})
	sliceStruct = append(sliceStruct, MyType{
		ValInt:    12,
		ValString: "b",
	})
	contains, ind = SliceContains(sliceStruct, MyType{
		ValInt:    8,
		ValString: "c",
	})
	fmt.Println("Contains Slice: struct -> ", contains, ind)
	// max
	maxSliceStruct, maxInd := SliceMax(sliceStruct)
	fmt.Println("Max Slice: struct -> ", maxSliceStruct, maxInd)
	fmt.Println()

	//////////////////////////// Slice Struct Pointer ///////////////////////////////////////////////
	sliceStructPointer := make([]*MyType, 0, 5)
	sliceStructPointer = append(sliceStructPointer, &MyType{
		ValInt:    1,
		ValString: "a",
	})
	sliceStructPointer = append(sliceStructPointer, &MyType{
		ValInt:    19,
		ValString: "g",
	})
	sliceStructPointer = append(sliceStructPointer, &MyType{
		ValInt:    8,
		ValString: "c",
	})
	sliceStructPointer = append(sliceStructPointer, &MyType{
		ValInt:    45,
		ValString: "y",
	})
	sliceStructPointer = append(sliceStructPointer, &MyType{
		ValInt:    12,
		ValString: "b",
	})
	contains, ind = SliceContains(sliceStructPointer, &MyType{
		ValInt:    8,
		ValString: "c",
	})
	fmt.Println("Contains Slice: struct pointer -> ", contains, ind)
	contains, ind = SliceContains(sliceStructPointer, MyType{
		ValInt:    8,
		ValString: "c",
	})
	fmt.Println("Contains Slice: struct pointer&value -> ", contains, ind)
	// max
	maxSliceStructPointer, maxInd := SliceMax(sliceStructPointer)
	fmt.Println("Max Slice: struct pointer -> ", maxSliceStructPointer, maxInd)

	return

	/*
		list := make([]interface{}, 0, 10)
		list = append(list, 1)
		list = append(list, 156484516458)
		list = append(list, -156484516458)
		list = append(list, 1.23)
		list = append(list, true)

		var a int8
		a = 22
		list = append(list, a)

		var b float32
		b = 22.1
		list = append(list, &b)
		toto(list)

		return

		listContains := make([]interface{}, 0, 10)
		listContains = append(listContains, 1)
		listContains = append(listContains, 89)
		listContains = append(listContains, 94)
		listContains = append(listContains, 5)
		listContains = append(listContains, 89)

		fmt.Println(Contains(listContains, 5))

		listContains = listContains[:0]
		listContains = append(listContains, "abc")
		listContains = append(listContains, "123")
		listContains = append(listContains, "gogogo")

		fmt.Println(Contains(listContains, "123"))
		return
	*/
}

func SliceMax(list interface{}) (interface{}, int) {
	if reflect.TypeOf(list).Kind() != reflect.Slice &&
		reflect.TypeOf(list).Kind() != reflect.Array {
		fmt.Println("NOT Slice/Array")
		return nil, -1
	}

	v := reflect.ValueOf(list)
	l := v.Len()
	if l == 0 {
		fmt.Println("len 0")
		return nil, -1
	}

	max := v.Index(0)
	maxInd := 0

	k := max.Kind()
	if k == reflect.Ptr {
		max = reflect.Indirect(max)
		k = max.Kind()
	}
	for i := 1; i < l; i++ {
		cur := v.Index(i)

		if cur.Kind() == reflect.Ptr {
			cur = reflect.Indirect(cur)
		}

		if cur.Kind() != k {
			fmt.Println("Must be the same kind")
			return nil, -1
		}

		switch k {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i1 := max.Int()
			i2 := cur.Int()
			if i1 < i2 {
				max = v.Index(i)
				maxInd = i
			}
		case reflect.String:
			s1 := max.String()
			s2 := cur.String()
			if s1 < s2 {
				max = v.Index(i)
				maxInd = i
			}
		case reflect.Float32, reflect.Float64:
			f1 := max.Float()
			f2 := cur.Float()
			if f1 < f2 {
				max = v.Index(i)
				maxInd = i
			}
		case reflect.Struct:
			t := max.Type()
			switch t.String() {
			case "main.MyType":
				v1 := max.Interface().(MyType)
				v2 := cur.Interface().(MyType)
				if v1.Less(&v2) {
					max = cur
					maxInd = i
				}
			default:
				fmt.Println("Unknown Type", t.String())
			}
		default:
			fmt.Println("Unknown kind", k)
			return nil, -1
		}
	}

	return max, maxInd
}

func SliceContains(list interface{}, val interface{}) (int, bool) {
	if reflect.TypeOf(list).Kind() != reflect.Slice &&
		reflect.TypeOf(list).Kind() != reflect.Array {
		fmt.Println("NOT Slice/Array")
		return -1, false
	}

	v := reflect.ValueOf(list)
	l := v.Len()
	if l == 0 {
		fmt.Println("len 0")
		return -1, false
	}

	for i := 0; i < l; i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), val) {
			return i, true
		}
	}

	return -1, false
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
func toto(list []interface{}) {
	for i, intf := range list {
		fmt.Printf("%d: type %s \n", i, reflect.TypeOf(intf).String())
		fmt.Printf("%d: kind %s \n", i, reflect.TypeOf(intf).Kind().String())
		fmt.Printf("%d: val %v \n", i, reflect.ValueOf(intf))
	}
}

func Contains(list []interface{}, val interface{}) (int, bool) {
	if len(list) == 0 {
		fmt.Println("empty list")
		return -1, false
	}

	k := reflect.TypeOf(val).Kind()

	for i, intf := range list {
		if reflect.TypeOf(intf).Kind() != k {
			fmt.Println("not same kind")
			return -1, false
		}

		if reflect.DeepEqual(intf, val) {
			return i, true
		}

		//fmt.Printf("%d: val %v %v \n", i, reflect.ValueOf(intf), reflect.ValueOf(val))

		//fmt.Printf("%d: type %s \n", i, reflect.TypeOf(intf).String())
		//fmt.Printf("%d: kind %s \n", i, reflect.TypeOf(intf).Kind().String())
		//fmt.Printf("%d: val %v \n", i, reflect.ValueOf(intf))
	}

	return -1, false
}

func Max(a, b interface{}) (interface{}, error) {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	k1 := t1.Kind()
	k2 := t2.Kind()

	if k1 != k2 {
		return nil, fmt.Errorf("Should be the same kind")
	}

	var v1, v2 reflect.Value
	switch k1 {
	case reflect.Bool,
		reflect.Array,
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Slice,
		reflect.Struct:
		return nil, fmt.Errorf("Not supported kind")
	case reflect.Ptr:
		v1 = reflect.Indirect(reflect.ValueOf(a))
		v2 = reflect.Indirect(reflect.ValueOf(b))
	default:
		v1 = reflect.ValueOf(a)
		v2 = reflect.ValueOf(b)
	}

	switch k1 {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("Int kind (%v) \n", k1.String())
		i1 := v1.Int()
		i2 := v2.Int()
		if i1 > i2 {
			return i1, nil
		}

		return i2, nil
	case reflect.String:
		fmt.Printf("String kind (%v) \n", k1.String())
		s1 := v1.String()
		s2 := v2.String()
		if s1 > s2 {
			return s1, nil
		}

		return s2, nil
	case reflect.Float32, reflect.Float64:
		fmt.Printf("Float kind (%v) \n", k1.String())
		f1 := v1.Float()
		f2 := v2.Float()
		if f1 > f2 {
			return f1, nil
		}

		return f2, nil
	}

	return nil, nil
}
