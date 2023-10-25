package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/png"
	"os"
	"strconv"

	Network "github.com/jcatterton/GoNeat/GoNeat"
)

func Start() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Driving Simulator",
		Bounds: pixel.R(0, 0, 1500, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.Black

	carImage, err := loadPicture("./car.png")
	if err != nil {
		panic(err)
	}

	walls := []Wall{
		{x1: 50, y1: 950, x2: 800, y2: 950},
		{x1: 800, y1: 950, x2: 875, y2: 925},
		{x1: 875, y1: 925, x2: 925, y2: 875},
		{x1: 925, y1: 875, x2: 950, y2: 800},
		{x1: 950, y1: 800, x2: 950, y2: 600},
		{x1: 950, y1: 800, x2: 950, y2: 600},
		{x1: 950, y1: 600, x2: 975, y2: 525},
		{x1: 975, y1: 525, x2: 1025, y2: 475},
		{x1: 1025, y1: 475, x2: 1100, y2: 450},
		{x1: 1100, y1: 450, x2: 1300, y2: 450},
		{x1: 1300, y1: 450, x2: 1400, y2: 350},
		{x1: 1400, y1: 350, x2: 1400, y2: 250},
		{x1: 1400, y1: 250, x2: 1300, y2: 150},
		{x1: 1300, y1: 150, x2: 800, y2: 150},
		{x1: 800, y1: 150, x2: 700, y2: 225},
		{x1: 700, y1: 225, x2: 625, y2: 325},
		{x1: 625, y1: 325, x2: 625, y2: 500},
		{x1: 625, y1: 325, x2: 550, y2: 225},
		{x1: 550, y1: 225, x2: 450, y2: 150},
		{x1: 450, y1: 150, x2: 50, y2: 150},
		{x1: 50, y1: 150, x2: 50, y2: 950},
		{x1: 150, y1: 850, x2: 800, y2: 850},
		{x1: 800, y1: 850, x2: 850, y2: 800},
		{x1: 850, y1: 800, x2: 850, y2: 600},
		{x1: 850, y1: 600, x2: 900, y2: 475},
		{x1: 900, y1: 475, x2: 975, y2: 400},
		{x1: 975, y1: 400, x2: 1100, y2: 350},
		{x1: 1100, y1: 350, x2: 1300, y2: 350},
		{x1: 1300, y1: 350, x2: 1300, y2: 250},
		{x1: 1300, y1: 250, x2: 800, y2: 250},
		{x1: 800, y1: 250, x2: 725, y2: 325},
		{x1: 525, y1: 325, x2: 525, y2: 600},
		{x1: 725, y1: 325, x2: 725, y2: 600},
		{x1: 525, y1: 325, x2: 450, y2: 250},
		{x1: 450, y1: 250, x2: 150, y2: 250},
		{x1: 150, y1: 250, x2: 150, y2: 850},
		{x1: 525, y1: 600, x2: 725, y2: 600},
	}

	fitnessGate := []Wall{
		{x1: 250, x2: 250, y1: 850, y2: 950},
		{x1: 350, x2: 350, y1: 850, y2: 950},
		{x1: 450, x2: 450, y1: 850, y2: 950},
		{x1: 550, x2: 550, y1: 850, y2: 950},
		{x1: 650, x2: 650, y1: 850, y2: 950},
		{x1: 750, x2: 750, y1: 850, y2: 950},
		{x1: 800, x2: 875, y1: 850, y2: 925},
		{x1: 850, x2: 925, y1: 800, y2: 875},
		{x1: 850, x2: 950, y1: 750, y2: 750},
		{x1: 850, x2: 950, y1: 600, y2: 600},
		{x1: 900, x2: 975, y1: 475, y2: 525},
		{x1: 975, x2: 1025, y1: 400, y2: 475},
		{x1: 1100, x2: 1100, y1: 350, y2: 450},
		{x1: 1250, x2: 1250, y1: 350, y2: 450},
		{x1: 1300, x2: 1400, y1: 350, y2: 350},
		{x1: 1300, x2: 1400, y1: 250, y2: 250},
		{x1: 1250, x2: 1250, y1: 250, y2: 150},
		{x1: 1150, x2: 1150, y1: 250, y2: 150},
		{x1: 1050, x2: 1050, y1: 250, y2: 150},
		{x1: 950, x2: 950, y1: 250, y2: 150},
		{x1: 800, x2: 800, y1: 250, y2: 150},
		{x1: 700, x2: 763, y1: 225, y2: 283},
		{x1: 625, x2: 725, y1: 325, y2: 325},
		{x1: 625, x2: 725, y1: 450, y2: 450},
		{x1: 625, x2: 625, y1: 500, y2: 600},
		{x1: 625, x2: 525, y1: 450, y2: 450},
		{x1: 525, x2: 625, y1: 325, y2: 325},
		{x1: 550, x2: 483, y1: 225, y2: 283},
		{x1: 450, x2: 450, y1: 150, y2: 250},
		{x1: 350, x2: 350, y1: 150, y2: 250},
		{x1: 250, x2: 250, y1: 150, y2: 250},
		{x1: 50, x2: 150, y1: 150, y2: 250},
		{x1: 50, x2: 150, y1: 950, y2: 850},
		{x1: 50, x2: 150, y1: 300, y2: 300},
		{x1: 50, x2: 150, y1: 400, y2: 400},
		{x1: 50, x2: 150, y1: 500, y2: 500},
		{x1: 50, x2: 150, y1: 600, y2: 600},
		{x1: 50, x2: 150, y1: 700, y2: 700},
		{x1: 50, x2: 150, y1: 800, y2: 800},
	}

	for i := range fitnessGate {
		fitnessGate[i].Draw(win, *imd, false)
	}

	pop := Network.InitPopulation(3, 2, 600)

	linesVisible := false
	brainVisible := true
	gatesVisible := false

	cars := make([]Car, len(pop.GetAllGenomes()))
	for i := range cars {
		cars[i].xPos = 500
		cars[i].yPos = 900
		cars[i].angle = 0
		cars[i].angularVel = 0
		cars[i].linearVel = 0
		cars[i].dead = false
		cars[i].sprite = *pixel.NewSprite(carImage, carImage.Bounds())
		cars[i].score = 0
	}

	livingCars := 600

	for !win.Closed() {
		win.Clear(colornames.White)

		if win.JustPressed(pixelgl.Key1) {
			linesVisible = !linesVisible
		}

		if win.JustPressed(pixelgl.Key2) {
			brainVisible = !brainVisible
		}

		if win.JustPressed(pixelgl.Key3) {
			gatesVisible = !gatesVisible
		}

		for i := range walls {
			walls[i].Draw(win, *imd, true)
		}

		for x := range pop.GetAllGenomes() {
			if win.JustPressed(pixelgl.KeyX) {
				for i := range cars {
					cars[i].dead = true
				}
			}

			if !cars[x].dead {
				distanceVecs := cars[x].Look(walls)

				if err := pop.GetAllGenomes()[x].TakeInput(
					[]float64{
						distanceVecs[0],
						distanceVecs[1],
						distanceVecs[2],
					}); err != nil {
					panic(err)
				}
				pop.GetAllGenomes()[x].FeedForward()
				outputs := pop.GetAllGenomes()[x].GetOutputs()

				cars[x].MoveForward()

				if outputs[0] > 0.5 {
					cars[x].TurnRight()
				}
				if outputs[1] > 0.5 {
					cars[x].TurnLeft()
				}
				if outputs[0] <= 0.5 && outputs[1] <= 0.5 {
					cars[x].SlowAngularVel()
				}

				cars[x].Move()

				wallCollision, gateCollision := checkForCollisions(cars[x], walls, fitnessGate)
				if wallCollision {
					livingCars--
					pop.GetAllGenomes()[x].SetFitness(cars[x].score)
					cars[x].score = 0.0
					cars[x].dead = true
				}

				if gateCollision {
					if !cars[x].inFitnessGate {
						cars[x].inFitnessGate = true
						cars[x].score = cars[x].score + 100
					}
				} else {
					cars[x].inFitnessGate = false
				}

				cars[x].Draw(win, pixel.IM.Moved(pixel.V(cars[x].xPos, cars[x].yPos)).Rotated(pixel.V(cars[x].xPos, cars[x].yPos), cars[x].angle))
			}
		}

		if allCarsDead(cars) {

			pop.NaturalSelection()
			win.SetTitle("Driving Simulator - " +
				"Generation: " + strconv.Itoa(pop.GetGeneration()) + " - " +
				"Best Fitness: " + strconv.FormatFloat(pop.GetGrandChampion().GetFitness(), 'f', 0, 64) + " - " +
				"Stagnation: " + strconv.Itoa(pop.GetSpecies()[0].GetStagnation()) + " - " +
				"Species: " + strconv.Itoa(len(pop.GetSpecies())))

			for i := range cars {
				cars[i].xPos = 500
				cars[i].yPos = 900
				cars[i].angle = 0
				cars[i].angularVel = 0
				cars[i].linearVel = 0
				cars[i].dead = false
				cars[i].sprite = *pixel.NewSprite(carImage, carImage.Bounds())
			}

			livingCars = len(cars)
		}

		if brainVisible {
			for i := range pop.GetAllGenomes() {
				if !cars[i].dead {
					drawGenome(pop.GetAllGenomes()[i], win)
					break
				}
			}
		}
		if linesVisible {
			for car := range cars {
				if !cars[car].dead {
					for i := range cars[car].sightVecs {
						imd.Push(cars[car].shortSightVecs[i].A)
						imd.Push(cars[car].shortSightVecs[i].B)
						imd.Line(1)
					}
				}
			}
		}
		if gatesVisible {
			for i := range fitnessGate {
				fitnessGate[i].Draw(win, *imd, false)
			}
		}

		imd.Draw(win)

		imd.Clear()

		win.Update()
	}
}

func checkForCollisions(c Car, w []Wall, g []Wall) (bool, bool) {
	wallCollision := false
	gateCollision := false

	for i := range w {
		_, intersect1 := w[i].Line().Intersect(c.Bounds()[0])
		_, intersect2 := w[i].Line().Intersect(c.Bounds()[1])
		_, intersect3 := w[i].Line().Intersect(c.Bounds()[2])
		_, intersect4 := w[i].Line().Intersect(c.Bounds()[3])
		if intersect1 || intersect2 || intersect3 || intersect4 {
			wallCollision = true
		}
	}

	for i := range g {
		_, intersect1 := g[i].Line().Intersect(c.Bounds()[0])
		_, intersect2 := g[i].Line().Intersect(c.Bounds()[1])
		_, intersect3 := g[i].Line().Intersect(c.Bounds()[2])
		_, intersect4 := g[i].Line().Intersect(c.Bounds()[3])
		if intersect1 || intersect2 || intersect3 || intersect4 {
			gateCollision = true
		}
	}
	return wallCollision, gateCollision
}

func allCarsDead(cars []Car) bool {
	for i := range cars {
		if cars[i].dead == false {
			return false
		}
	}
	return true
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func drawGenome(g *Network.Genome, win *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)

	imd := imdraw.New(nil)

	w := win.Bounds().W() / 2
	h := win.Bounds().H() / 3

	for i := 0; i < g.GetLayers(); i++ {
		for j := range g.GetNodesWithLayer(i + 1) {
			if g.GetNodesWithLayer(i + 1)[j].IsActivated() {
				imd.Color = pixel.RGB(0, 1, 0)
			} else {
				imd.Color = pixel.RGB(1, 0, 0)
			}
			imd.Push(pixel.V(
				(float64(i)+0.5)*(w/float64(g.GetLayers())),
				(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))))
			imd.Circle(5, 20)

			for k := range g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections() {
				imd.Color = pixel.RGB(0, 0, 0)
				imd.Push(
					pixel.V(
						(float64(i)+0.5)*(w/float64(g.GetLayers()))+10,
						(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))),
					pixel.V(
						(float64(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer())-0.5)*(w/float64(g.GetLayers()))-10,
						(float64(Network.NodeIndex(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()),
							g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB()))+0.5)*(h/float64(len(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()))))))
				imd.Line(2)
			}

			basicTxt.Color = colornames.White
			_, err := fmt.Fprintf(basicTxt, strconv.Itoa(g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber()))
			if err != nil {
				panic(err)
			}
			basicTxt.Draw(win, pixel.IM.Moved(pixel.V(
				(float64(i)+0.5)*(w/float64(g.GetLayers()))-1,
				(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))+20)))
			basicTxt.Clear()
		}
	}
	imd.Draw(win)
}
