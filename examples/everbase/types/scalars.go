package types

type boolean string // all scalars are treated as strings
func NewBoolean(val string) boolean {
	return boolean(val)
}

type float string // all scalars are treated as strings
func NewFloat(val string) float {
	return float(val)
}

type id string // all scalars are treated as strings
func NewId(val string) id {
	return id(val)
}

type _int string // all scalars are treated as strings
func NewInt(val string) _int {
	return _int(val)
}

type _string string // all scalars are treated as strings
func NewString(val string) _string {
	return _string(val)
}
