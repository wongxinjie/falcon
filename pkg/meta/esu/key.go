package esu

type Key string

// V return string value
func (k Key) V() string {
	return string(k)
}

