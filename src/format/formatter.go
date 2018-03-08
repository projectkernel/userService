package format

type Formatter interface {
	Format(interface{}) string
}
