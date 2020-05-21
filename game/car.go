package game

import (
	"github.com/faiface/pixel"
	"math"
)

type Car struct {
	xPos           float64
	yPos           float64
	angularVel     float64
	angle          float64
	linearVel      float64
	sprite         pixel.Sprite
	dead           bool
	sightVecs      []pixel.Line
	shortSightVecs []pixel.Line
	score          float64
}

func (c *Car) Draw(t pixel.Target, m pixel.Matrix) {
	c.sprite.Draw(t, m)
}

func (c *Car) Bounds() []pixel.Line {
	x1 := (c.sprite.Frame().Min.X-c.sprite.Frame().Center().X)*math.Cos(c.angle) - (c.sprite.Frame().Min.Y-c.sprite.Frame().Center().Y)*math.Sin(c.angle)
	y1 := (c.sprite.Frame().Min.X-c.sprite.Frame().Center().X)*math.Sin(c.angle) + (c.sprite.Frame().Min.Y-c.sprite.Frame().Center().Y)*math.Cos(c.angle)

	x2 := (c.sprite.Frame().Min.X-c.sprite.Frame().Center().X)*math.Cos(c.angle) - (c.sprite.Frame().Max.Y-c.sprite.Frame().Center().Y)*math.Sin(c.angle)
	y2 := (c.sprite.Frame().Min.X-c.sprite.Frame().Center().X)*math.Sin(c.angle) + (c.sprite.Frame().Max.Y-c.sprite.Frame().Center().Y)*math.Cos(c.angle)

	x3 := (c.sprite.Frame().Max.X-c.sprite.Frame().Center().X)*math.Cos(c.angle) - (c.sprite.Frame().Min.Y-c.sprite.Frame().Center().Y)*math.Sin(c.angle)
	y3 := (c.sprite.Frame().Max.X-c.sprite.Frame().Center().X)*math.Sin(c.angle) + (c.sprite.Frame().Min.Y-c.sprite.Frame().Center().Y)*math.Cos(c.angle)

	x4 := (c.sprite.Frame().Max.X-c.sprite.Frame().Center().X)*math.Cos(c.angle) - (c.sprite.Frame().Max.Y-c.sprite.Frame().Center().Y)*math.Sin(c.angle)
	y4 := (c.sprite.Frame().Max.X-c.sprite.Frame().Center().X)*math.Sin(c.angle) + (c.sprite.Frame().Max.Y-c.sprite.Frame().Center().Y)*math.Cos(c.angle)

	return []pixel.Line{
		pixel.L(pixel.V(x1+c.xPos, y1+c.yPos), pixel.V(x2+c.xPos, y2+c.yPos)),
		pixel.L(pixel.V(x2+c.xPos, y2+c.yPos), pixel.V(x4+c.xPos, y4+c.yPos)),
		pixel.L(pixel.V(x1+c.xPos, y1+c.yPos), pixel.V(x3+c.xPos, y3+c.yPos)),
		pixel.L(pixel.V(x3+c.xPos, y3+c.yPos), pixel.V(x4+c.xPos, y4+c.yPos)),
	}
}

func (c *Car) Move() {
	c.angle = c.angle + c.angularVel
	c.xPos = c.xPos + c.linearVel*math.Cos(c.angle)
	c.yPos = c.yPos + c.linearVel*math.Sin(c.angle)
}

func (c *Car) TurnLeft() {
	if c.angularVel <= math.Pi/24 {
		c.angularVel = c.angularVel + 0.005
	} else {
		c.angularVel = math.Pi / 24
	}
}

func (c *Car) TurnRight() {
	if c.angularVel >= -math.Pi/24 {
		c.angularVel = c.angularVel - 0.005
	} else {
		c.angularVel = -math.Pi / 24
	}
}

func (c *Car) MoveForward() {
	if c.linearVel <= 5 {
		c.linearVel = c.linearVel + 0.2
	} else {
		c.linearVel = 5
	}
}

func (c *Car) MoveBackwards() {
	if c.linearVel >= -5 {
		c.linearVel = c.linearVel - 0.2
	} else {
		c.linearVel = -5
	}
}

func (c *Car) SlowLinearVel() {
	for c.linearVel > 0 {
		c.linearVel = c.linearVel - 0.2
		if c.linearVel < 0.2 && c.linearVel > -0.2 {
			c.linearVel = 0
		}
	}
	for c.linearVel < 0 {
		c.linearVel = c.linearVel + 0.2
		if c.linearVel < 0.2 && c.linearVel > -0.2 {
			c.linearVel = 0
		}
	}
}

func (c *Car) SlowAngularVel() {
	if c.angularVel > -math.Pi/20 && c.angularVel < math.Pi/20 {
		c.angularVel = 0
	}
	for c.angularVel > 0 {
		c.angularVel = c.angularVel - math.Pi/20
	}
	for c.angularVel < 0 {
		c.angularVel = c.angularVel + math.Pi/20
	}
}

func (c *Car) Look(walls []Wall) []float64 {
	c.sightVecs = make([]pixel.Line, 3)
	c.sightVecs[0] = pixel.L(pixel.V(c.xPos, c.yPos), pixel.V(c.xPos+1000*math.Cos(c.angle), c.yPos+1000*math.Sin(c.angle)))
	c.sightVecs[1] = pixel.L(pixel.V(c.xPos, c.yPos), pixel.V(c.xPos+1000*math.Cos(c.angle+math.Pi/4), c.yPos+1000*math.Sin(c.angle+math.Pi/4)))
	c.sightVecs[2] = pixel.L(pixel.V(c.xPos, c.yPos), pixel.V(c.xPos+1000*math.Cos(c.angle-math.Pi/4), c.yPos+1000*math.Sin(c.angle-math.Pi/4)))

	distanceVecs := []float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

	c.shortSightVecs = make([]pixel.Line, 3)

	for i := range c.sightVecs {
		for j := range walls {
			tempVec, intersect := c.sightVecs[i].Intersect(walls[j].Line())
			if intersect {
				tempLine := pixel.L(pixel.V(c.xPos, c.yPos), tempVec)
				if tempLine.Len() < distanceVecs[i] {
					c.shortSightVecs[i] = pixel.L(pixel.V(c.xPos, c.yPos), tempVec)
					distanceVecs[i] = tempLine.Len()
				}
			}
		}
	}

	distanceVecs[0] = distanceVecs[0] - c.sprite.Frame().W()/2
	distanceVecs[1] = distanceVecs[1] - math.Sqrt(((c.sprite.Frame().W()/2)*(c.sprite.Frame().W()/2))+((c.sprite.Frame().H()/2)*(c.sprite.Frame().H()/2)))
	distanceVecs[2] = distanceVecs[2] - math.Sqrt(((c.sprite.Frame().W()/2)*(c.sprite.Frame().W()/2))+((c.sprite.Frame().H()/2)*(c.sprite.Frame().H()/2)))

	return distanceVecs
}
