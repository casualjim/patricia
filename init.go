package patricia

func init() {
	initBuildLeftMasks()
}

type Value struct {
	ID uint32
	Addr []byte
}