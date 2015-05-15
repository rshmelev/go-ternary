package ifff

import (
	"reflect"

	. "github.com/rshmelev/go-uniconv"
)

type ThenStruct struct {
	Condition bool
	ThenValue interface{}
	ElseValue interface{}

	*SoftConverter
}

type PreThenInterface interface {
	Then(a interface{}) *ThenStruct
	ThenElse(a interface{}, b interface{}) *ThenStruct
}
type PreThenSetInterface interface {
	PreThenInterface
	ThenSetTo(a interface{})
	Or(cond bool) PreThenSetInterface
	And(cond bool) PreThenSetInterface
}

func If(condition bool) PreThenInterface {
	t := &ThenStruct{condition, nil, nil, &SoftConverter{}}
	return t
}

// pointers are used here to forward vars to change them, so
// don't send pointers to check if they're nil or not.
// or send pointers to pointers.
func IsDefaultValue(a interface{}) bool {
	rvoa := reflect.ValueOf(a)
	if rvoa.Kind() == reflect.Ptr {
		rvoa = rvoa.Elem()
	}
	isnil := false
	switch rvoa.Kind() {
	case reflect.Chan, reflect.Interface, reflect.Slice, reflect.Func, reflect.Map, reflect.Ptr:
		isnil = rvoa.IsNil()
	default:
		isnil = rvoa.Interface() == reflect.New(rvoa.Type()).Elem().Interface()
	}
	return isnil
}
func IfDefaultValue(a interface{}) PreThenSetInterface {
	t := &ThenStruct{IsDefaultValue(a), nil, a, &SoftConverter{}}
	return t
}

func (t *ThenStruct) Or(cond bool) PreThenSetInterface {
	t.Condition = t.Condition || cond
	return t
}
func (t *ThenStruct) And(cond bool) PreThenSetInterface {
	t.Condition = t.Condition && cond
	return t
}

func (t *ThenStruct) ThenSetTo(a interface{}) {
	v := reflect.ValueOf(t.ElseValue)
	if v.Kind() != reflect.Ptr {
		panic("IfHasDefaultValue - ThenSet got value param, not a pointer")
	}
	if !t.Condition {
		return
	}
	v = v.Elem()
	v.Set(reflect.ValueOf(a))
}

func (t *ThenStruct) Then(a interface{}) *ThenStruct {
	t.ThenValue = a
	if t.Condition {
		t.Value = a
	}
	return t
}
func (t *ThenStruct) ThenElse(a interface{}, b interface{}) *ThenStruct {
	t.ThenValue = a
	t.ElseValue = b
	t.Value = t.res()
	return t
}

func (t *ThenStruct) Else(a interface{}) *ThenStruct {
	if !t.Condition {
		t.Value = a
	}
	t.ElseValue = a
	return t
}

// for verbosity
func (t *ThenStruct) ElseDefaultValue() *ThenStruct {
	return t
}

func (t *ThenStruct) res() interface{} {
	if t.Condition {
		return t.ThenValue
	}
	return t.ElseValue
}

//func (t *ThenStruct) Str() string {
//	v := t.res()
//	return fmt.Sprint(v)
//}
