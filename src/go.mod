module main

go 1.13

require (
	ch01 v0.0.0 //indirect
	ch02 v0.0.0 //indirect
)

replace (
	ch01 v0.0.0 => ./ch01
	ch02 v0.0.0 => ./ch02
)
