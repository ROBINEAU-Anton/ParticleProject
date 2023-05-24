package particles


//This function allows to add a Particle to the system
func (s *System) add(){
		p := newParticle()
		s.Content.PushFront(&p)
	}

//This function allows you to add a finite number of particles to the system
func (s *System) add_number(number int){
	for i:=0;i<number;i++{
		s.add()
	}
}
