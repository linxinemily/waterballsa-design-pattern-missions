package domain

type Card[C any] interface {
	CompareTo(C) int
}

func Compare[T Card[T]](c1, c2 T) int {
	return c1.CompareTo(c2)
}