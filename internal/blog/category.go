package blog

const (
	Golang Category = "golang"
	Cpp    Category = "cpp"
	Rust   Category = "rust"
)

type Category string

func (c Category) String() string {
	if c == "" {
		return "unknown"
	}
	return string(c)
}
