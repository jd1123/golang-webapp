package vapelog

type ModType int
type BatteryType int

const (
	builtin BatteryType = iota
	removeable
)

const (
	Box ModType = iota
	Tube
	EGO
	Custom
)

type Mod struct {
	Manufacturer             string
	Name                     string
	MinV, MaxV               float64
	MinW, MaxW               float64
	VapeWattage, VapeVoltage float64
	Batt                     BatteryType
	Type                     ModType
	Meta                     interface{}
}

type Atomizer struct {
	BaseType     string
	Clone        bool
	Name         string
	Manufacturer string
	Rebuilable   bool
	NumCoils     int
	Resistence   float64
	Meta         interface{}
}

type Battery struct {
	Name         string
	Manufacturer string
	Mah          int
	Amps         float64
	Meta         interface{}
}

type VapeSetup struct {
	VapeMod      Mod
	VapeAtomizer Atomizer
	VapeBattery  Battery
	Meta         interface{}
}

func NewMod(manufacturer, name string) Mod {
	return Mod{Manufacturer: manufacturer, Name: name}
}
