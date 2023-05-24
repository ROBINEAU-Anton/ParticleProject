package particles

import (
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*This next fonction allows to modifie the config.json file according to the gamemod */
func circle(){
	if config.General.Gamemod == "circle"{
		config.General.LiveSpanActivator = false
		config.General.RandomRotation = false
		config.General.Gravity = false
		config.General.InitNumParticles = 0
		config.General.NumMaxParticles = 40000
		config.General.SpawnRate = 0
		config.General.OnClick = true
		config.General.RightClickBlower = true
		config.General.Bounce = false
	}
}

func buble(){
	if config.General.Gamemod == "buble"{
		config.General.LiveSpanActivator = false
		config.General.RandomRotation = false
		config.General.Gravity = false
		config.General.InitNumParticles = 0
		config.General.NumMaxParticles = 100000
		config.General.SpawnRate = 0
		config.General.OnClick = true
		config.General.RandomColor = false
		config.General.Bounce = false
	}
}

func torch(){
	if config.General.Gamemod == "torch"{
		config.General.LiveSpanActivator = false
		config.General.LiveSpan = 90
		config.General.RandomRotation = false
		config.General.Gravity = false
		config.General.InitNumParticles = 0
		config.General.NumMaxParticles = 10000
		config.General.SpawnRate = 0
		config.General.OnClick = false
		config.General.RandomColor = false
		config.General.OpacityManagementMode = 2 
		config.General.RightClickBlower = false
		config.General.Bounce = false
	}
}

func paint(){
	if config.General.Gamemod == "paint"{
		config.General.LiveSpanActivator = false
		config.General.LiveSpan = 0
		config.General.RandomRotation = false
		config.General.Gravity = false
		config.General.InitNumParticles = 0
		config.General.NumMaxParticles = 30000
		config.General.SpawnRate = 0
		config.General.OnClick = false
		config.General.RandomColor = false
		config.General.OpacityManagementMode = 0
		config.General.RandomScale = false
		config.General.RandomOpacity = false
		config.General.RightClickBlower = false
		config.General.Bounce = false
	}
}


//This next fonctions allows to update object for some gamemod 
func (s *System) update_torch(){
	x,y := ebiten.CursorPosition()
		s.Torch.PositionX,s.Torch.PositionY = float64(x)-15,float64(y)-50
		if inpututil.IsKeyJustReleased(ebiten.KeySpace){
			if config.General.SpawnRate == 0{
				config.General.SpawnRate = 1 
			}else{
				config.General.SpawnRate = 0
			}
	}
}



func (s *System) update_crayon(){
	if config.General.LiveSpanActivator{
		config.General.LiveSpanActivator = false
	}
	x,y := ebiten.CursorPosition()
	s.Crayon.PositionX,s.Crayon.PositionY = float64(x),float64(y)+15
	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		s.add()
	}
	//to clear the window with the C key 
	if inpututil.IsKeyJustReleased(ebiten.KeyC){
		config.General.LiveSpanActivator = true
	}
}