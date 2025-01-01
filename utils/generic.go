package utils

import "golang.org/x/exp/constraints"

// Numeric represents all numeric types
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Ordered represents all ordered types
type Ordered interface {
	constraints.Ordered
}

// Comparable represents all comparable types
type Comparable interface {
	comparable
}

// Any represents any type
type Any interface {
	any
}
