package learn

import (
	"fmt"
	"testing"
)

type Device struct {
	Serial     string
	hardwareID int
}

//function linked to struct
func (d *Device) SetHardwareID(h int) {
	d.hardwareID = h
}

func NewDevice(hw int) *Device {
	return &Device{
		Serial:     "abc0",
		hardwareID: hw,
	}
}

func (d *Device) String() string {
	return fmt.Sprintf("A big fat device with serial number %q", d.Serial)
}

func TestDevice(t *testing.T) {
	d := NewDevice(123)
	d.SetHardwareID(55)
	t.Logf("%v", d)
	t.Log(d)
	t.Fail()
}
