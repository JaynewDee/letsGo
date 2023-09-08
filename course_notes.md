### PUBLIC
- defining a value with an initial capital letter gives the value public visibility!

### PASS BY VALUE
Primitive Scalar types like ints, floats, strings, and structs are PASS BY VALUE
- "Pass by reference" can be accomplished by passing a pointer!

### PASS BY REFERENCE
Collection types like slices and maps are PASS BY REFERENCE


### "Zero Values"
- numeric: `0`
- complexNumeric: `(0+0i)`

- bool: `false`
- string: `""`
- pointer(any): `nil`
- array: `[]0`
- slice: `nil`
- map: `nil`
- channel: `nil`
- struct: each field gets respective "zero value"
- function: `nil`
- interface: `nil`