package node

type Source interface {
	BuildResource(d Device, index int)
}

func (s *SourceGeneric) BuildResource(d Device, index int) {
	label := getResourceLabel("TestGenericSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Generic", d)
	s.Format = VideoFormat
}

func (s *SourceAudio) BuildResource(d Device, index int) {
	label := getResourceLabel("TestAudioSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Audio", d)
	s.Format = AudioFormat
}

func (s *SourceData) BuildResource(d Device, index int) {
	label := getResourceLabel("TestDataSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Data", d)
	s.Format = DataFormat
}
