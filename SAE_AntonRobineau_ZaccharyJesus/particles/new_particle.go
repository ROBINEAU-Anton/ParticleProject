package particles 

import(
	"project-particles/config"
	"math/rand"
	//"math"
) 


func newParticle() (Particle){

	//Creation of a BasicParticle 
	p := Particle{
		PositionX: float64(config.General.WindowSizeX/2), //the particle spawn in the center of the screen if the randomspawn if not activacted
		PositionY: float64(config.General.WindowSizeY/2) ,
		Rotation: 0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		SpeedX : (rand.NormFloat64()-rand.NormFloat64()),
		SpeedY : (rand.NormFloat64()-rand.NormFloat64()),
		Alive : true,
		LiveSpan : config.General.LiveSpan,
	}
	//Particle configuration
	p.color() // choose random color for the particle if the option is set on true
	p.opacity() //choose a random opacity between 30% and 100% if the option is set on true
	p.scale() //choose a random scale x and y between 75% and 100% if the option is set on true
	p.speed() //change the speedtype depending on gamemod setting 
	p.spawn() // Spawn choose randomly if the option is set on true
	p.rotation() //choose a random Rotation if the option is set on true

	//Return the particle
	return p
}

func newTorch() (Torch){
	t := Torch{
		PositionX : 0,
		PositionY : 0,
		ScaleX : 1,ScaleY : 1,
	}
	return t 
}

func newCrayon() (Crayon){
	c := Crayon{
		PositionX : 0,
		PositionY : 0,
		ScaleX : 1,ScaleY : 1,
	}
	return c
}