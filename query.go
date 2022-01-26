package seedwork

type Query interface {
	Prepare()
	Validate() error
}

func ProcessQuery(q Query) error {
	q.Prepare()

	return Validate(q, "query")
}
