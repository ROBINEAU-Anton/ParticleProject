# R1.01.SAE.eq09_JESUS-Zacchary_ROBINEAU-Anton

## Name
Project_Particle

## Description
Creation of a particle generator in the go language using a 2D engine named Ebiten

## Pre-requisites
- Golang - https://go.dev

## Setup
Install the source code on gitlab directly or through the terminal with the following command:<br>
```
git clone https://gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae/projets/groupe2/r1.01.sae.eq09_jesus-zacchary_robineau-anton.git
```
Once installed, you will have to go to the terminal and go to the installer folder before and execute the following commands:
```
go install
go build
```

## Run

 If you want to launch a particle generation and you have already completed the previous steps, you just have to launch the file 'project-particles.exe'

## Test 

 If you want to run the test set from the particles folder in the console run the command go test
 The test set will evolve during the development of this project 

## Config.json
> ### <ins>Window configuration</ins>
> - WindowTitle | Choose the name of the window.
> - Debug | Activate the information display at the top left of the screen (true/false)
> - Fullscreen | Choose if the application launches in full screen (true/false) coming soon
> #### Window size
> - WindowSizeX | Window length (in pixels)
> - WindowSizeY | Window widht (in pixels)
> ### <ins>System configuration</ins>
> - OnClick |to create particles with the LeftClick (true/false)
> - SpawnRate | Number of particles at each call of the update function(60/s)
> - InitNumParticles | The number of particles on the screen at startup
> - NumMaxParticles | The number of particles maximum
> - Gravity | Adds a gravity effect to the particles (true/false)
> - Bounce | To activate the bounces on the screen border (true/false)
> - RightClickBlower | to strive the particles with RightCLick (true/false)
> - LiveSpanActivator | Activates a lifetime for the particles (true/false)
> - LiveSpan | Defines the lifetime of the particles (60 pour 1s)
> ### <ins>Particle configuration</ins>
> - RandomSpawn | Appearance of particles at a random location (true/false)
> - RandomRotation | Choose a random inital rotation for particles true/false)
> - RandomColor | Choose a random color for each particles (true/false)
> - RandomOpacity | Choose a random opacity between 30% and 100% for each particles (true/false)
> - RandomScale | Choose a random scale for each particles (true/false)
> - OpacityManagementMode | 0 to no changement | 1 to increase opacity | 2 to decrease opacity 
> ### <ins>Image configuration</ins>
> - ParticleImage | Image display for particles (path)
> - TorchImage | Image display for the torch in torch gamemod  (path)
> - CrayonImage | Image display for the pen in paint gamemod  (path)
> ### <ins>Gamemod</ins>
> - "Basic" to have your personal configuration
> - "buble" to create a buble with LeftClick and strive with RightClick
> - "circle" to create circle with LeftClick and strive with RightClick
> - "torch" Torch mod | SpaceKey to create fire 
> - "paint" Draw mod | LeftClick to draw and C to clear 

## Visual

### <ins>Basic gamemod</ins>

![](/SAE_AntonRobineau_ZaccharyJesus/assets/screen5.png)

### <ins>Buble gamemod</ins>

![](/SAE_AntonRobineau_ZaccharyJesus/assets/screen1.png)

### <ins>Circle gamemod</ins>
![](/SAE_AntonRobineau_ZaccharyJesus/assets/screen2.png)

### <ins>Torch gamemod</ins>
![](/SAE_AntonRobineau_ZaccharyJesus/assets/screen3.png)

### <ins>Paint gamemod</ins>
![](/SAE_AntonRobineau_ZaccharyJesus/assets/screen4.png)


## Authors
Zacchary JESUS - Anton ROBINEAU

## Project status
This project is still under development.