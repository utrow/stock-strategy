package strategys

type Strategy struct  {
	Reserve Reserve
}

func MakeStrategy() Strategy {
	return Strategy{}
}
