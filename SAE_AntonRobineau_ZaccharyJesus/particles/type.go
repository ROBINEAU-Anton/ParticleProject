package particles

import "container/list"

// System defines a particle system.
/* 
At the moment it only contains a list of particles, but this can
evolve during your project. 
*/
type System struct {
	Content *list.List
	SpawnValue float64 
	Torch      *Torch
	Crayon     *Crayon
}

// Particle defines a particle.
/* 
It has a position, rotation, size, color, and opacity. 
You will certainly add other characteristics to the particles
during the project. 
*/
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX,SpeedY 					float64
	Alive 							bool 
	LiveSpan 						float64
}

//Torch define a torch
/*
It has only a positon x and  and sclae x and y*/
type Torch struct{
	PositionX, PositionY            float64
	ScaleX, ScaleY                  float64
}

//Torch define a crayon
/*
It has only a positon x and  and sclae x and y*/
type Crayon struct{
	PositionX, PositionY            float64
	ScaleX, ScaleY                  float64
}
