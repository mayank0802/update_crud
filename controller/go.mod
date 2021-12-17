module last.com/controller

go 1.17

replace last.com/config => ../config

replace last.com/model => ../model

require (
	last.com/config v0.0.0-00010101000000-000000000000
	last.com/model v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/mux v1.8.0 // indirect
