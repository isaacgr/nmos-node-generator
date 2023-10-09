package node

type Source interface {
	BuildResource(d Device, index int, useRandomResource bool)
	getId() string
}

func (s *SourceGeneric) BuildResource(d Device, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"GenericSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Generic", d, useRandomResource)
	s.Format = VideoFormat
}

func (s *SourceAudio) BuildResource(d Device, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"AudioSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Audio", d, useRandomResource)
	s.Format = AudioFormat
	c1 := SourceChannels{
		"Audio 1",
		SourceChannelSymbols["L"],
	}
	c2 := SourceChannels{
		"Audio 2",
		SourceChannelSymbols["R"],
	}
	s.Channels = append(s.Channels, c1, c2)
}

func (s *SourceData) BuildResource(d Device, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"DataSource", index)
	s.BaseSource = SetBaseSourceProperties(label, "NMOS Test Source Data", d, useRandomResource)
	s.Format = DataFormat
}

func (s *SourceGeneric) getId() string {
	return s.ID
}
func (s *SourceAudio) getId() string {
	return s.ID
}
func (s *SourceData) getId() string {
	return s.ID
}
