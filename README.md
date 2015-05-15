# go-ternary

Non-idiomatic attempt to create ternary operator and even something more than it.

```golang

import( . github.com/rshmelev/go-ternary/if )

// such oneliner. very golang. much idiomatic.
a := If(a > b).Then(10).Else("20").Str()

// first non-default value is returned
b := Or(a).Or(b).Or(c).Str() 

// init zero-value variables
IfDefaultValue(&a).ThenSetTo(10)

```

automatic conversion is supported with [github.com/rshmelev/go-uniconv]
so, as a bonus you can convert result to any of basic types

__WARNING__: code like `If(a != nil).Then(a.Field).Else(b)` __will cause panic with access to nil pointer__.   
You should clearly understand that such "ternary" thing __is not using lazy evaluation__.

author: rshmelev@gmail.com

