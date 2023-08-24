package lengthconv

func MToFt(m Metre) Foot { return Foot(m * 3.2808) }
func FtToM(m Foot) Metre { return Metre(m / 3.2808) }
