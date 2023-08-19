package rank

type Rank struct {
	name   string
	id     int
	format string
}

func (r *Rank) Name() string {
	return r.name
}

func (r *Rank) ID() int {
	return r.id
}

func (r *Rank) Format() string {
	return r.format
}
