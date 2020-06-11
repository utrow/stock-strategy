package strategys

type Strategy struct {
	Reserve        Reserve
	HeikenFood     HeikenFood
	HeikenSmoothed HeikenSmoothed
	StockChange    StockChange
}

func MakeStrategy() Strategy {
	return Strategy{}
}
