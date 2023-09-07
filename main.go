package main

import (
	m "me.example/coba/multithread"
	s "me.example/coba/singlethread"
)

func main() {
	s.SingleThread()
	m.Thread()
}
