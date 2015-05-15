package ifff

import . "github.com/rshmelev/go-uniconv"

type IThisOrThisBase interface {
	Or(a interface{}) IThisOrThis
}

type IThisOrThis interface {
	UniversalConverter
	IThisOrThisBase
}

type ThisOrThis struct {
	*SoftConverter
	foundNotDefaultValue bool
}

func Or(a interface{}) IThisOrThisBase {
	return either(a)
}

func either(a interface{}) IThisOrThisBase {
	x := &ThisOrThis{&SoftConverter{}, false}
	x.Or(a)
	return x
}

func (x *ThisOrThis) Or(a interface{}) IThisOrThis {
	if !IsDefaultValue(a) && !x.foundNotDefaultValue {
		x.foundNotDefaultValue = true
		x.Value = a
	}
	return x
}
