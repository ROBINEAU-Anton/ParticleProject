package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewSystem() System {
	/*
	NewSystem is a function that initializes a particle system
	And adds the chosen number of particles (InitNumParticles)
	And returns it to the main function of the project, which will display it
	*/
	rand.Seed(time.Now().UnixNano())

	//Change the config file for each gamemod 
	if config.General.Gamemod == "circle"{
		circle()
	}else if config.General.Gamemod == "torch"{
		torch()
	}else if config.General.Gamemod == "buble"{
		buble()
	}else if config.General.Gamemod == "paint"{
		paint()
	}

	s := System{Content: list.New()}
	if config.General.Gamemod == "torch"{
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		t := newTorch()
		s.Torch = &t
	}else if config.General.Gamemod == "paint"{
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		c := newCrayon()
		s.Crayon = &c
	}
	s.add_number(config.General.InitNumParticles)
	return s
}
