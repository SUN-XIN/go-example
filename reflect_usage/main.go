package main

import (
	"fmt"
	"reflect"
)

func main() {
	listV1 := make([]int, 0, 10)
	listV1 = append(listV1, 1)
	listV1 = append(listV1, 89)
	listV1 = append(listV1, 5)
	listV1 = append(listV1, 45)
	fmt.Println(ContainsV1(listV1, 5))

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

	m1 := 5.0
	m2 := 15.0
	m, err := Max(m1, m2)
	if err != nil {
		fmt.Printf("Max: %+v \n", err)
		return
	}

	fmt.Printf("Max is: %v \n", m)
	return
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

func ContainsV1(list interface{}, val interface{}) (int, bool) {
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

	//vTarget := reflect.ValueOf(val)
	for i := 0; i < l; i++ {
		fmt.Println(i, v.Index(i))
		if reflect.DeepEqual(v.Index(i).Interface(), val) {
			return i, true
		}
	}

	return -1, false
}
