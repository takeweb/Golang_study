package util

// うるう年かどうか判定する
func IsLeapYear(year int) bool {
	if year%400 == 0 { // 400で割り切れたらうるう年
		return true
	} else if year%100 == 0 { // 100で割り切れたらうるう年じゃない
		return false
	} else if year%4 == 0 { // 4で割り切れたらうるう年
		return true
	} else {
		return false
	}
}
