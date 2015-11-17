package mynumber

import "testing"

func TestB2n(t *testing.T) {
	if b2n('0') != 0 {
		t.Error("b2n('0') should returns 0")
	}
	if b2n('1') != 1 {
		t.Error("b2n('1') should returns 1")
	}
	if b2n('2') != 2 {
		t.Error("b2n('2') should returns 2")
	}
	if b2n('3') != 3 {
		t.Error("b2n('3') should returns 3")
	}
	if b2n('4') != 4 {
		t.Error("b2n('4') should returns 4")
	}
	if b2n('5') != 5 {
		t.Error("b2n('5') should returns 5")
	}
	if b2n('6') != 6 {
		t.Error("b2n('6') should returns 6")
	}
	if b2n('7') != 7 {
		t.Error("b2n('7') should returns 7")
	}
	if b2n('8') != 8 {
		t.Error("b2n('8') should returns 8")
	}
	if b2n('9') != 9 {
		t.Error("b2n('9') should returns 9")
	}
}

func TestValidate(t *testing.T) {
	if ValidateStr("123456789018") != true {
		t.Error("\"123456789018\" should be valid")
	}
	if ValidateStr("123456789010") != false {
		t.Error("\"123456789010\" must not be valid")
	}
	if ValidateStr("123456789011") != false {
		t.Error("\"123456789011\" must not be valid")
	}
	if ValidateStr("123456789012") != false {
		t.Error("\"123456789012\" must not be valid")
	}
	if ValidateStr("123456789013") != false {
		t.Error("\"123456789013\" must not be valid")
	}
	if ValidateStr("123456789014") != false {
		t.Error("\"123456789014\" must not be valid")
	}
	if ValidateStr("123456789015") != false {
		t.Error("\"123456789015\" must not be valid")
	}
	if ValidateStr("123456789016") != false {
		t.Error("\"123456789016\" must not be valid")
	}
	if ValidateStr("123456789019") != false {
		t.Error("\"123456789019\" must not be valid")
	}
}
