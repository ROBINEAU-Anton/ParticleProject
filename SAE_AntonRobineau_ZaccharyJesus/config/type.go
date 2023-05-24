package config

/*
Config defines the fields that can be found in a config file.
In the file the fields must have the same name as in the type if underneath, 
including capital letters.  All fields must be required start with capital letters, 
otherwise it will not be possible to recover their values from the config file. 
You can add fields and they will be automatically read in the config file. 
You will need to do this several times during the project. 
*/
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	TorchImage 				 string
	CrayonImage 			 string
	Debug                    bool
	OnClick 				 bool
	InitNumParticles         int
	NumMaxParticles 	     int
	RandomSpawn              bool
	SpawnRate                float64
	Gamemod  			     string //Circle , Basic , 
	FullScreen 				 bool
	Gravity 				 bool
	LiveSpanActivator		 bool
	LiveSpan 				 float64
	RandomColor 			 bool
	RandomOpacity 			 bool
	RandomScale  			 bool
	OpacityManagementMode    int
	RandomRotation			 bool
	Bounce 					 bool
	RightClickBlower 					 bool
}

var General Config
