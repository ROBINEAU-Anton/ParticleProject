package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
	"project-particles/config"
)

func (s *System) Update() {

	//update each alive particle of the system and delete the others
	for e := s.Content.Front(); e != nil;  {
		p, ok :=e.Value.(*Particle)
		next := e.Next()
		if !ok {
			continue
		}
		if p.Alive {
			p.update()
		}else {
			s.Content.Remove(e)
		}
		e = next 
	}

	//Generate more particles
	//by click or automaticaly according to the options choosen in the config.json
	s.spawn()
	
	
	//to close the window with the escape key
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape){
		os.Exit(0)
	}

	//Update the torch or pen according to the gamemod 
	if config.General.Gamemod == "torch" {
		s.update_torch()
	}else if config.General.Gamemod == "paint"{
		s.update_crayon()
	}
}
