package dbu

import "fmt"

type Key string

// V return string value
func (k Key) V() string {
	return string(k)
}

func (k Key) Eq() string {
	return fmt.Sprintf("%s = ?", k)
}

func (k Key) Max() string {
	return fmt.Sprintf("max(%s)", k)
}

func (k Key) Min() string {
	return fmt.Sprintf("min(%s)", k)
}

func (k Key) Inc() string {
	return fmt.Sprintf("%s = %s + ?", k, k)
}

func (k Key) Dec() string {
	return fmt.Sprintf("%s = %s - ?", k, k)
}

func (k Key) Ne() string {
	return fmt.Sprintf("%s != ?", k)
}

func (k Key) Is() string {
	return fmt.Sprintf("%s IS ?", k)
}

func (k Key) IsNot() string {
	return fmt.Sprintf("%s IS NOT ?", k)
}

// Gt return `greater than` format
func (k Key) Gt() string {
	return fmt.Sprintf("%s > ?", k)
}

func (k Key) Gte() string {
	return fmt.Sprintf("%s >= ?", k)
}

func (k Key) Lt() string {
	return fmt.Sprintf("%s < ?", k)
}

func (k Key) Lte() string {
	return fmt.Sprintf("%s <= ?", k)
}

func (k Key) In() string {
	return fmt.Sprintf("%s IN (?)", k)
}

func (k Key) ArrayContains() string {
	return fmt.Sprintf("%s @> (?)", k)
}

func (k Key) ArrayOverlap() string {
	return fmt.Sprintf("%s && (?)", k)
}

func (k Key) Nin() string {
	return fmt.Sprintf("%s NOT IN (?)", k)
}

func (k Key) Desc() string {
	return fmt.Sprintf("%s DESC", k)
}

func (k Key) Asc() string {
	return fmt.Sprintf("%s ASC", k)
}

func (k Key) Like() string {
	return fmt.Sprintf("%s LIKE ?", k)
}

func (k Key) Between() string {
	return fmt.Sprintf("%s BETWEEN ? AND ?", k)
}

func (k Key) IsNULL() string {
	return fmt.Sprintf("%s IS NULL", k)
}

func (k Key) IsNotNULL() string {
	return fmt.Sprintf("%s IS NOT NULL", k)
}

func (k Key) Any() string {
	return fmt.Sprintf("%s.*", k)
}

func (k Key) SetDefault() string {
	return fmt.Sprintf("%s = DEFAULT", k)
}

func (k Key) SetNull() string {
	return fmt.Sprintf("%s = NULL", k)
}

func (k Key) Escape() Key {
	return Key(fmt.Sprintf("\"%s\"", k))
}

func Concat(a, b string) Key {
	return Key(fmt.Sprintf("%s.%s", a, b))
}

func (k Key) DescNULLSLAST() string {
	return fmt.Sprintf("%s DESC NULLS LAST", k)
}

func (k Key) AscNullsFIRST() string {
	return fmt.Sprintf("%s ASC NULLS FIRST", k)
}

func (k Key) Distinct() string {
	return fmt.Sprintf("DISTINCT %s ", k)
}

func (k Key) Excluded() string {
	return fmt.Sprintf("%s = EXCLUDED.%s", k, k)
}
