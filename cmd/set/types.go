package set

type Options struct {
	Subreddits []string
	Index      int
	Permanent  bool
}

type logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
