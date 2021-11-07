module github.com/dragovorn/go-pong

go 1.17

replace github.com/isshoni-soft/sakura => ./engine/
replace github.com/isshoni-soft/roxxy => ./engine/roxxy/
replace github.com/isshoni-soft/kirito => ./engine/kirito/

require (
	github.com/go-gl/gl v0.0.0-20210905235341-f7a045908259 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20210727001814-0db043d8d5be // indirect
	github.com/isshoni-soft/kirito v0.0.0-20211106044012-849c4cdb0637 // indirect
	github.com/isshoni-soft/roxxy v0.0.0-20211106044041-28a6add6336b // indirect
	github.com/isshoni-soft/sakura v0.0.0-20211022193349-d711737eafa0 // indirect
)
