package models

import (
	"testing"

	"github.com/tripleaceinnovations/api/models"
)

//Test IsPalindrome function
func testIsPalindromeEmptyArg(t *testing.T) {
	//test for empty argument
	emptyResult := models.IsPalindrome("")

	if emptyResult != false {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "false", emptyResult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "false", emptyResult)
	}
}

//test for valid postive arguments
func testIsPalindromeValidPosArg(t *testing.T) {
	postiveresult := models.IsPalindrome("level")

	if postiveresult != true {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "true", postiveresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "true", postiveresult)
	}

}

//test for valid negative arguments
func testIsPalindromeValidNegArg(t *testing.T) {
	negativeresult := models.IsPalindrome("bash")

	if negativeresult != false {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "false", negativeresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "false", negativeresult)
	}

}
