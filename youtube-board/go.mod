module youtube-board

go 1.16

require (
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	local.packages/my v0.0.0-00010101000000-000000000000
)

replace local.packages/my => ./my
