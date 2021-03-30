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

type upload string // all scalars are treated as strings
func NewUpload(val string) upload {
	return upload(val)
}
