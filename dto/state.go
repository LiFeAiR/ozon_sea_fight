package dto

type State struct {
	ShipCount    int   `json:"ship_count"` // всего кораблей
	Destroyed    int   `json:"destroyed"`  // потоплено
	Knocked      int   `json:"knocked"`    // подбито
	KnockedSlice []int `json:"-"`
	ShotCount    int   `json:"shot_count"` // сделано выстрелов
}
