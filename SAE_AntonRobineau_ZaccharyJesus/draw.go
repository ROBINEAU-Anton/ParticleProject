package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"math"
)

/*
Draw is responsible for displaying on the screen the current state of the particle system
g.system. It is called automatically about 60 times per second per second the Ebiten 
library. This function may be slightly modified when It is specified in the subject.
*/
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok && p.Alive{
			//Draw the particle if it is Alive 
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	//Draw the torch in torch mode and print more info
	if config.General.Gamemod == "torch"{
		top := ebiten.DrawImageOptions{}
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Push SpaceKey to turn on the torch"),config.General.WindowSizeX/2-100,20,)
		top.GeoM.Translate(g.system.Torch.PositionX, g.system.Torch.PositionY)
		top.GeoM.Scale(g.system.Torch.ScaleX, g.system.Torch.ScaleY)
		screen.DrawImage(assets.TorchImage, &top)
	}

	//Draw the pen in paint mode and print more info
	if config.General.Gamemod == "paint"{
		cop := ebiten.DrawImageOptions{}
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Click to draw"),config.General.WindowSizeX/2-30,20,)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Push C To Clear"),config.General.WindowSizeX - 150,20,)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"),0,90,)
		cop.GeoM.Translate(g.system.Crayon.PositionX, g.system.Crayon.PositionY)
		cop.GeoM.Scale(g.system.Crayon.ScaleX, g.system.Crayon.ScaleY)
		screen.DrawImage(assets.CrayonImage, &cop)
	}

	//Print the debug info FPS/Gamemod/Number of particles 
	if config.General.Debug {
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("FPS : ",math.Round(ebiten.ActualTPS()),),0,0)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Gamemod : ",config.General.Gamemod),config.General.WindowSizeX/2-30,0)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Particles : ",g.system.Content.Len()),0,20,)
		if config.General.RightClickBlower {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint("Right Click to Strive"),200,20,)
		}
	}

}
