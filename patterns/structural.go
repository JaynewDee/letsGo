package main

import (
	"fmt"
)

/*
	ADAPTER
*/
// interface to adapt to
type remote interface {
	on()
	off()
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
}

// TV implements remote directly
type TV struct {
	vol     int
	channel int
	isOn    bool
}

func (tv *TV) on() {
	fmt.Println("TV on.")
	tv.isOn = true
}
func (tv *TV) off() {
	fmt.Println("TV off.")
	tv.isOn = false
}
func (tv *TV) volumeUp() int {
	fmt.Println("Volume turned up.")
	tv.vol++
	return tv.vol
}
func (tv *TV) volumeDown() int {
	fmt.Println("Volume turned down.")
	tv.vol--
	return tv.vol
}
func (tv *TV) channelUp() int {
	fmt.Println("Channel up.")
	tv.channel++
	return tv.channel
}
func (tv *TV) channelDown() int {
	fmt.Println("Channel down.")
	tv.channel--
	return tv.channel
}

// Radio has the same behaviors but the API is defined differently
type Radio struct {
	currentChannel int
	currentVolume  int
	pluggedIn      bool
}

func (r *Radio) plugIn() {
	r.pluggedIn = true
}
func (r *Radio) unplug() {
	r.pluggedIn = false
}
func (r *Radio) seekUp() int {
	r.currentChannel++
	return r.currentChannel
}
func (r *Radio) seekDown() int {
	r.currentChannel--
	return r.currentChannel
}
func (r *Radio) decreaseVolume() int {
	r.currentVolume--
	return r.currentVolume
}
func (r *Radio) increaseVolume() int {
	r.currentVolume++
	return r.currentVolume
}

// The adapter gets all remote interface methods, which call the corresponding radio methods
type radioAdapter struct {
	radio *Radio
}

func (rd *radioAdapter) on() {
	rd.radio.plugIn()
}
func (rd *radioAdapter) off() {
	rd.radio.unplug()
}
func (rd *radioAdapter) volumeUp() int {
	return rd.radio.decreaseVolume()
}
func (rd *radioAdapter) volumeDown() int {
	return rd.radio.increaseVolume()
}
func (rd *radioAdapter) channelUp() int {
	return rd.radio.seekUp()
}
func (rd *radioAdapter) channelDown() int {
	return rd.radio.seekDown()
}

/*
FACADE
*/
type MilkKind string
type Size int

const (
	WholeMilk   MilkKind = "whole"
	CoconutMilk MilkKind = "coconut"
	AlmondMilk  MilkKind = "almond"
)
const (
	Small  Size = 3
	Medium Size = 5
	Large  Size = 8
)

func milkMatcher(kind MilkKind, to string) (MilkKind, error) {
	if to == "enum" {
		switch kind {
		case WholeMilk:
			return "whole", nil
		case CoconutMilk:
			return "coconut", nil
		case AlmondMilk:
			return "almond", nil
		default:
			return "", fmt.Errorf("failed to get string value of enum")
		}
	} else {
		switch kind {
		case "whole":
			return WholeMilk, nil
		case "coconut":
			return CoconutMilk, nil
		case "almond":
			return AlmondMilk, nil
		default:
			return "", fmt.Errorf("failed to get enum value of string")
		}
	}
}

type coffeeOps interface {
	useBeanVariety(variety string)
	useDrinkSize(size int)
	useFoam(units int)
	grindBeans(amount float32)
	useHotWater()
	useColdWater()
	addSyrup(flavor string)
	addMilk(kind milkKind)
}

type coffeeDrinks interface {
	Drip()
	Iced()
	Espresso()
	CaramelLatte()
}

type CoffeeMachine struct {
	operator coffeeOps
}

func (cm *CoffeeMachine) Drip() {
	cm.operator.useDrinkSize(Small)
	cm.operator.grindBeans()
}
