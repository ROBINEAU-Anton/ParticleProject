package particles 

import(
	"math"
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (s *System) spawn() {
	/*
	spawnUpdate is a system method that updates the number of particles in real time based on the spawnrate
	For this it adds the spawrate to the spawnvalue which is an attribute of the system
	If the spawnvalue is greater than one then we create particles and we keep only the decimal part
	This method is called in the update function which is run 60 times per second
	This fonction work with two mode automaticaly or on click according to the options
	If the onclick option is on , we create 65 particles by update fonction call during a click
	*/
	if s.Content.Len()< config.General.NumMaxParticles{
		if config.General.OnClick {
				if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0{
					s.add_number(65)
				}
		}else{
			s.SpawnValue += config.General.SpawnRate
			if s.SpawnValue >= 1{
				s.add_number(int(s.SpawnValue))							
				_, s.SpawnValue = math.Modf(s.SpawnValue)
			}
		}
	}
}
