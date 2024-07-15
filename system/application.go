package system

import (
	"fmt"
	"github.com/LiFeAiR/ozon_sea_fight/dto"
	"strconv"
	"strings"
)

type App interface {
	Clear()
	CreateMatrix(maxIndex int)
	CreateShips(coordinates string) error
	MakeShot(coordinates string) (*dto.Shot, error)
	State() dto.State
	ShipsCreated() bool
	GetFightMatrix() map[string]map[string]bool
}

type Application struct {
	FightMatrix map[string]map[string]bool
	ShipMatrix  map[string]map[string]int
	Ships       []*dto.Ship
	state       dto.State
}

func NewApplication() App {
	app := &Application{}
	app.init()

	return app
}

func (app *Application) init() {
	app.FightMatrix = make(map[string]map[string]bool)
	app.ShipMatrix = make(map[string]map[string]int)
	app.state = dto.State{}
	app.state.KnockedSlice = make([]int, 0)
	app.Ships = make([]*dto.Ship, 0)
}

func (app *Application) Clear() {
	app.init()
}

func (app *Application) CreateMatrix(maxIndex int) {
	for i := 1; i <= maxIndex; i++ {
		s := strconv.Itoa(i)
		app.FightMatrix[s] = make(map[string]bool)
		app.ShipMatrix[s] = make(map[string]int)

		for j := 0; j < maxIndex; j++ {
			s1 := fmt.Sprint(j + 65)
			app.FightMatrix[s][s1] = false
			app.ShipMatrix[s][s1] = -1
		}
	}
}

func (app *Application) CreateShips(coordinates string) error {
	shipsSlice := strings.Split(coordinates, ",")
	app.Ships = make([]*dto.Ship, len(shipsSlice))

	for k, coords := range shipsSlice {
		coordsSlice := strings.Split(coords, " ")
		ship := &dto.Ship{
			Coordinates: dto.Coordinates{
				LeftTop:     coordsSlice[0],
				RightBottom: coordsSlice[1],
			},
			Knock: 0,
		}

		knocks, err := app.putShipToMatrix(ship, k)
		if err != nil {
			app.Ships = make([]*dto.Ship, 0)
			return err
		}

		ship.Knock = knocks
		app.Ships[k] = ship
	}

	app.state = dto.State{ShipCount: len(app.Ships)}
	app.state.KnockedSlice = make([]int, 0)

	return nil
}

func (app *Application) putShipToMatrix(ship *dto.Ship, idx int) (int, error) {
	var (
		row map[string]int
		ok  bool
	)
	rowMin := string(ship.Coordinates.LeftTop[0])
	columnMin := string(ship.Coordinates.LeftTop[1])

	rowMax := string(ship.Coordinates.RightBottom[0])
	columnMax := string(ship.Coordinates.RightBottom[1])

	row, ok = app.ShipMatrix[rowMin]
	if !ok {
		return 0, fmt.Errorf("error: shipMatrix => row")
	}

	if _, ok = row[columnMin]; !ok {
		return 0, fmt.Errorf("error: shipMatrix => row => column")
	}

	row, ok = app.ShipMatrix[rowMax]
	if !ok {
		return 0, fmt.Errorf("error: shipMatrix => row")
	}

	if _, ok = row[columnMax]; !ok {
		return 0, fmt.Errorf("error: shipMatrix => row => column")
	}

	iMin, err := strconv.Atoi(rowMin)
	if err != nil {
		return 0, err
	}

	iMax, err := strconv.Atoi(rowMax)
	if err != nil {
		return 0, err
	}

	jMin := int(columnMin[0]) - 65
	jMax := int(columnMax[0]) - 65
	knock := 0

	for i := iMin; i <= iMax; i++ {
		s := strconv.Itoa(i)

		for j := jMin; j <= jMax; j++ {
			s1 := fmt.Sprint(j + 65)
			app.ShipMatrix[s][s1] = idx
			knock++
		}
	}

	return knock, nil
}

func (app *Application) MakeShot(coordinates string) (*dto.Shot, error) {
	var (
		row     map[string]int
		ok      bool
		shipIdx int
	)
	longitude := string(coordinates[0])
	row, ok = app.ShipMatrix[longitude]
	if !ok {
		return nil, fmt.Errorf("shipMatrix => row")
	}

	latitude := string(coordinates[1])
	if shipIdx, ok = row[latitude]; !ok {
		return nil, fmt.Errorf("shipMatrix => row => column")
	}

	if app.FightMatrix[longitude][latitude] {
		return nil, fmt.Errorf("shot repeated")
	}

	app.FightMatrix[longitude][latitude] = true
	response := &dto.Shot{}
	if shipIdx > -1 {
		response.Knock = true
		ship := app.Ships[shipIdx]
		if ship.Knock > 0 {
			ship.Knock--
		}
		if ship.Knock == 0 {
			response.Destroy = true
			app.state.Destroyed++
			if app.state.Knocked > 0 {
				app.state.Knocked--
			}
		}
		find := false
		for _, v := range app.state.KnockedSlice {
			if v == shipIdx {
				find = true
				break
			}
		}
		if !find {
			app.state.Knocked++
			app.state.KnockedSlice = append(app.state.KnockedSlice, shipIdx)
		}
	}

	if app.state.Destroyed == app.state.ShipCount {
		response.End = true
	}

	app.state.ShotCount++
	return response, nil
}

func (app *Application) State() dto.State {
	return app.state
}

func (app *Application) GetFightMatrix() map[string]map[string]bool {
	return app.FightMatrix
}

func (app *Application) ShipsCreated() bool {
	return len(app.Ships) > 0
}
