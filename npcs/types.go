package npcs

type Power struct {
	Attack  int
	Defense int
}
type Location struct {
	X float64
	Y float64
	Z float64
}
type NonPlayerCahracter struct {
	Name  string
	Speed int
	HP    int
	Power Power
	Loc   Location
}
