package util

func DieOnError(err error) {
	if err != nil {
		panic(err)
	}
}
