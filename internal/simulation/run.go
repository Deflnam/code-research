package simulation

type Config struct {
	P      float64
	N      int
	Length int
}

type Result struct {
	Total    int
	Missed   int
	Detected int
}
