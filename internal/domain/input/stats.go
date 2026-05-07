package input

type Stats struct {
	FirstServe  *Serve
	SecondServe *Serve
	ServeReturn *Serve

	Aces         int
	DoubleFaults int

	Winners        int
	UnforcedErrors int

	Surface    *Surface
	MatchLevel *MatchLevel
}

func NewStats(
	firstServe, secondServe, serveReturn *Serve,
	aces, doubleFaults, winners, unforcedErrors int,
	surface *Surface,
	matchLevel *MatchLevel,
) *Stats {
	return &Stats{
		FirstServe:     firstServe,
		SecondServe:    secondServe,
		ServeReturn:    serveReturn,
		Aces:           aces,
		DoubleFaults:   doubleFaults,
		Winners:        winners,
		UnforcedErrors: unforcedErrors,
		Surface:        surface,
		MatchLevel:     matchLevel,
	}
}

func (s Stats) Validate() error {
	return nil
}
