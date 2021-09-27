package constant

// Channel
const (
	CatOne   = "one"
	CatTwo   = "two"
	CatThree = "three"
	CatFour  = "four"
	CatFive  = "five"
)

type Person struct {
	Name string `json:"name"`
	City string `json:"city"`
}

const (
	AccessLogPath = "./"
)
