module github.com/mayank0802/crud

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	last.com/controller v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	last.com/config v0.0.0-00010101000000-000000000000 // indirect
	last.com/model v0.0.0-00010101000000-000000000000 // indirect
)

replace last.com/config => ./config

replace last.com/controller => ./controller

replace last.com/model => ./model
