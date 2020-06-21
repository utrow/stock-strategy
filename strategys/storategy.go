package strategys

type Strategy struct {
	Reserve        Reserve
	HeikenFood     HeikenFood
	HeikenSmoothed HeikenSmoothed
	StockChange    StockChange
	Weeks          Weeks
}

func MakeStrategy() Strategy {
	return Strategy{}
}
