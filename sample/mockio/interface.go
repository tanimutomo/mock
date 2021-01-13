package foo

type Repository interface {
	Get(foo int, bar string) (Model, error)
	List(foos []int, bars []string) ([]Model, error)
}
