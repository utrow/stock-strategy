package strategys

type Strategy struct {
	Reserve        Reserve
	HeikenFood     HeikenFood
	HeikenSmoothed HeikenSmoothed
}

func MakeStrategy() Strategy {
	return Strategy{}
}
