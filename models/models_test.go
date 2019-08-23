package models

import (
	"testing"
)

//test for valid postive argument
func TestIsPalindromeValidPosArg(t *testing.T) {
	postiveresult := IsPalindrome("level")

	if postiveresult != true {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "true", postiveresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "true", postiveresult)
	}

}

//test for valid negative argument
func TestIsPalindromeValidNegArg(t *testing.T) {
	negativeresult := IsPalindrome("test")

	if negativeresult != false {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "false", negativeresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "false", negativeresult)
	}

}
