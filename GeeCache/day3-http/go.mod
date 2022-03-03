module day3

go 1.17

require rsc.io/quote v1.5.2
require geecache v0.0.0

replace geecache => ./geecache
require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
