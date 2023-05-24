package particles

import (
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//each fonctions uses in this one are coded et explained in the particle_extension_methode.go file 
func (p *Particle) update(){
	/*
	This function is a method of a particle
	It updates its position according to its speed and its livespan
	If the particle is Alive 
	And kill the particle if it go out of the screen
	*/
	if p.Alive{
		//update the position of a particle by adding the speed 
		p.updatePosition()
 		//The vertical speed of the particle changes with gravity. 
		p.gravity()
		//Update of the livespan of the particle
		p.updateLiveSpan()
		//Opacity Management 
		p.updateOpacity()
		if config.General.Gamemod == "torch"{
			p.updateColor()
		}else if config.General.Gamemod == "buble"{
			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){
				p.SpeedY,p.SpeedX = 2,0	
			}
		}

		//Verify if the particle go out of the screen 
		if p.PositionX > float64(config.General.WindowSizeX) +10 || p.PositionX < -10  ||p.PositionY > float64(config.General.WindowSizeY) +10  || p.PositionY < -10 {
			p.Alive = false // And set the Alive parameter to false
		}
	}
}
