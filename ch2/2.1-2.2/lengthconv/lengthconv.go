package lengthconv

import "fmt"

type Metre float64
type Foot float64

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (m Foot) String() string  { return fmt.Sprintf("%gft", m) }
