module package_test

go 1.16

require (
	local.packages/workhello2 v1.0.0
    local.packages/workhello3 v1.0.0
)

replace local.packages/workhello2 => ./package2/workhello2
replace local.packages/workhello3 => ../package3/workhello3