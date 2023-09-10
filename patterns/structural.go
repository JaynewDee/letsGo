package main

import "fmt"

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
	vol int
	channel int
	isOn bool
}
func(tv *TV) on() {
	fmt.Println("TV on.")
	tv.isOn = true
}
func(tv *TV) off() {
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
	currentVolume int
	pluggedIn bool
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
// The adapter has all remote interface methods, which call the corresponding radio methods
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