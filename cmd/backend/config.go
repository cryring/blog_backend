package main

type Config struct {
	Address string `arg:"--addr"`
	LogPath string `arg:"--log"`
	Config  string `arg:"--config"`
}
