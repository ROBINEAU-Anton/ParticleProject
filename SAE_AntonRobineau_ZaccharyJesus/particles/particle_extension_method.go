package particles

import(
	"project-particles/config"
	"math/rand"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
) 

//5.1
func (p *Particle) gravity(){
	/*
	This function is a method of a particle
	It updates its verticale speed if the gravity option is true
	*/
	if config.General.Gravity{
		if p.SpeedY < 0 {
			p.SpeedY += 0.035
		}else{
			p.SpeedY += 0.045
		}
	}
}

//5.3
func (p *Particle) updateLiveSpan(){
	//Update the particle LiveSpan 
	p.LiveSpan--
	//kill the particles if their LiveSpan is finish and the LiveSpanActivator is set on true
	if config.General.LiveSpanActivator && p.LiveSpan < 0 {
		p.Alive = false
	}
}
//5.4
//s to change the specs of a particle at its creation
func (p *Particle) color(){
	/*
	This fonction change the color of the particle before drawing it according to the setting
	*/
	if config.General.RandomColor{
		p.ColorRed,p.ColorBlue,p.ColorGreen  = rand.Float64(),rand.Float64(),rand.Float64()
	}else if config.General.Gamemod == "torch" || config.General.Gamemod == "paint"{
		p.ColorRed,p.ColorBlue,p.ColorGreen = 1,0,0
	}else if config.General.Gamemod == "buble"{
		p.ColorRed,p.ColorBlue,p.ColorGreen = rand.Float64()*0.4,0.8+rand.Float64()*0.2,rand.Float64()*0.8
	}
}

func (p *Particle) scale(){
	/*
	This fonction change the scale of the particle between 75% & 125% before drawing it according to the setting
	*/
	if config.General.RandomScale{
		p.ScaleX = 0.75 + rand.Float64() * (0.5)
		p.ScaleY = 0.75 + rand.Float64() * (0.5)
	}else if config.General.Gamemod == "paint" {
		p.ScaleX,p.ScaleY = 1,1
	}
}
func (p *Particle) opacity(){
	/*
	This fonction change the opacity of the particle between 30% & 100% before drawing it according to the setting
	*/	
	if config.General.RandomOpacity{
		p.Opacity = 0.3 + rand.Float64() * 0.7
	}
}

func (p *Particle) spawn(){
	/*
	This fonction change the spawn coordinates of the particle before drawing it according to the option 
	(randomSpawn or Gamemode)
	*/
	if config.General.RandomSpawn{
		p.PositionX = float64(rand.Intn(config.General.WindowSizeX - 10))
		p.PositionY = float64(rand.Intn(config.General.WindowSizeY -10))
	}else if config.General.Gamemod == "circle"{
		x,y := ebiten.CursorPosition()
		p.PositionX,p.PositionY = float64(x),float64(y)
	}else if config.General.Gamemod == "torch"{
		x,y := ebiten.CursorPosition()
		p.PositionX,p.PositionY = float64(x)-5,float64(y)-55
	}else if config.General.Gamemod == "paint"{
		x,y := ebiten.CursorPosition()
		p.PositionX,p.PositionY = float64(x),float64(y) + 100
	}else if config.General.Gamemod == "buble"{
		x,y := ebiten.CursorPosition()
		p.PositionX,p.PositionY = float64(x),float64(y)
	}
}

func (p *Particle) rotation(){
	/*
	This fonction change the axis of rotation of the particle before drawing it according to the setting
	*/
	//Rotation configuration 
	if  config.General.RandomRotation{
		p.Rotation = rand.Float64() 
	}
}
func (p *Particle) speed(){
	/*
	This fonction change the speed of the particle before drawing it according to the Gamemode
	*/
	if config.General.Gamemod == "circle"{
		a :=rand.Float64() * 2 * math.Pi
		p.SpeedX = math.Cos(a)  * 8
		p.SpeedY = math.Sin(a)  * 8
	}else if config.General.Gamemod == "torch"{
		p.SpeedX = rand.Float64()*0.2 - rand.Float64()*0.2
		p.SpeedY = -rand.Float64()
	}else if config.General.Gamemod == "buble"{
		a :=rand.Float64() * 2 * math.Pi
		p.SpeedX = math.Cos(a)  * 2
		p.SpeedY = math.Sin(a)  * 2
	}else if config.General.Gamemod == "paint"{
		p.SpeedX,p.SpeedY = 0,0
	}
}

//5.4
//Methods to uptade the specs of a particle during its livespan
func (p *Particle) updatePosition(){
	/*
	Update the position of a particle by adding the speed 
	And BlowerMode+BounceMode if they are activacted 
	*/ 

	//Blower mode developement (To spire the particles with RightClick)
	if config.General.RightClickBlower{
		if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0{

			x,y := ebiten.CursorPosition()

			gapX,gapY := float64(x)-p.PositionX,float64(y)-p.PositionY
			gap := math.Sqrt(gapX*gapX + gapY*gapY)
			gapX,gapY = gapX/gap,gapY/gap

			p.SpeedX,p.SpeedY = gapX*5,gapY*5

			if p.PositionX == float64(x) {
					p.SpeedX = 0 
			}
			if p.PositionY == float64(y){
				p.SpeedY = 0 
			}

		}
	}

	//Update the position of a particle by adding the speed
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY

	// Inverse the speed of the particles if it touch the border if the option is set on
	if config.General.Bounce{
		// Inverse the speed of the particles if it touch the border 
		if p.PositionX > float64(config.General.WindowSizeX) - 10 * p.ScaleX || p.PositionX < 10 * p.ScaleX{
			p.SpeedX = -p.SpeedX
		}
		if p.PositionY > float64(config.General.WindowSizeY) - 10 * p.ScaleY || p.PositionY < 10 * p.ScaleY{
			p.SpeedY = -p.SpeedY
		}
	}
}

func (p *Particle) updateRotation(){
	/*
	Update the rotation of the particle by adding a value
	*/
	p.Rotation += 0.08
}

func (p *Particle) updateOpacity(){
	/*
	Update the opacity of the particle: if the opacity of the particle is less than 1 it decreases it 
	and if it is greater than 0 it increases it by adding/subtracting a value. The particle dies 
	if the opacity goes below 0
	*/
	if config.General.OpacityManagementMode == 1 {
		//increase Opacity
		if p.Opacity < 1 {
			p.Opacity +=0.001
		}
	}
	if config.General.OpacityManagementMode == 2 {
		//decrease Opacity
		if p.Opacity > 0 {
			p.Opacity -=0.007
		}else{
			p.Alive = false
		}
	}
}

func (p *Particle) updateColor(){
	//change the color of the particles for a beatiful render in torch gamemod
	p.ColorGreen += 0.008
	
}

//5.6
func (p *Particle) revive(){
	//this fonction was create for a futur extension in order to revive the dead particles 
	p.color() // choose random color for the particle if the option is set on true
	p.opacity() //choose a random opacity between 30% and 100% if the option is set on true
	p.scale() //choose a random scale x and y between 75% and 100% if the option is set on true
	p.spawn() // Spawn choose randomly if the option is set on true
	p.rotation() //choose a random Rotation if the option is set on true
	p.speed() //change the speedtype depending on gamemod setting
	p.Alive = true
}




