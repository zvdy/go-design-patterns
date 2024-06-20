package main

import "fmt"

type USB interface {
	InsertIntoUSBPort()
}

type MicroUSB interface {
	InsertIntoMicroUSBPort()
}

type USBTypeC interface {
	InsertIntoUSBTypeCPort()
}

type MicroUSBDevice struct{}

func (m *MicroUSBDevice) InsertIntoMicroUSBPort() {
	fmt.Println("MicroUSB device connected.")
}

type USBTypeCDevice struct{}

func (u *USBTypeCDevice) InsertIntoUSBTypeCPort() {
	fmt.Println("USB Type-C device connected.")
}

type MicroUSBToUSBAdapter struct {
	MicroUSBDevice MicroUSB
}

func (m *MicroUSBToUSBAdapter) InsertIntoUSBPort() {
	m.MicroUSBDevice.InsertIntoMicroUSBPort()
}

type USBTypeCToUSBAdapter struct {
	USBTypeCDevice USBTypeC
}

func (u *USBTypeCToUSBAdapter) InsertIntoUSBPort() {
	u.USBTypeCDevice.InsertIntoUSBTypeCPort()
}

type Device struct {
	USBPort USB
}

func (d *Device) Connect() {
	d.USBPort.InsertIntoUSBPort()
}

func main() {
	device := &Device{}

	microUSBDevice := &MicroUSBDevice{}
	microUSBAdapter := &MicroUSBToUSBAdapter{MicroUSBDevice: microUSBDevice}
	device.USBPort = microUSBAdapter
	device.Connect()

	usbTypeCDevice := &USBTypeCDevice{}
	usbTypeCAdapter := &USBTypeCToUSBAdapter{USBTypeCDevice: usbTypeCDevice}
	device.USBPort = usbTypeCAdapter
	device.Connect()
}
