package seedwork

type Command interface {
	Prepare()
	Validate() error
}

func ProcessCommand(q Query) error {
	q.Prepare()

	return Validate(q, "command")
}
