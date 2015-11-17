package mynumber

func b2n(b byte) int {
	if b < '0' || b > '9' {
		return -1
	}
	return (int)(b - '0')
}

// Validate validates nums as MyNumber.  If it is valid returns true, otherwise
// false.
func Validate(nums []byte) bool {
	if len(nums) != 12 {
		return false
	}
	s := 0
	for i, b := range nums[0:11] {
		p := 11 - i
		if p <= 6 {
			p++
		} else {
			p -= 5
		}
		q := b2n(b)
		if q < 0 {
			return false
		}
		s += p * q
	}
	s %= 11
	if s <= 1 {
		s = 0
	} else {
		s = 11 - s
	}
	return b2n(nums[11]) == s
}

// ValidateStr validates nums as MyNumber.  If it is valid returns true,
// otherwise false.
func ValidateStr(nums string) bool {
	return Validate([]byte(nums))
}
