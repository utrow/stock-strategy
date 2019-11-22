package services

type Service struct  {
	CsvReader CsvReader
}

func MakeService() Service {
	return Service{}
}
