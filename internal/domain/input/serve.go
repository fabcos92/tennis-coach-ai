package input

type Serve struct {
	In  *Percent
	Won *Percent
}

func NewServe(in, won *Percent) *Serve {
	return &Serve{in, won}
}
