package reflection

import "reflect"

func walkInternal(x interface{}, fn func(input string), seen map[uintptr]struct{}) {
	val := reflect.ValueOf(x)

	walkValue := func(value reflect.Value) {
		walkInternal(value.Interface(), fn, seen)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, result := range valFnResult {
			walkValue(result)
		}
	case reflect.Pointer:
		addr := val.Pointer()
		if _, ok := seen[addr]; !ok {
			seen[addr] = struct{}{}
			walkValue(val.Elem())
		}
	}
}

func walk(x interface{}, fn func(input string)) {
	seen := map[uintptr]struct{}{}
	walkInternal(x, fn, seen)
}
