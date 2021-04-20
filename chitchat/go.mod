module chitchat

go 1.16

require (
	github.com/lib/pq v1.10.0 // indirect
	local.packages/data v1.0.0
)

replace local.packages/data => ./data
