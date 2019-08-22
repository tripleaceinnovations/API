package models

import (
	"testing"
	//"github.com/tripleaceinnovations/learngo/models"
)

<<<<<<< HEAD
=======
//Test IsPalindrome function
func TestIsPalindromeEmptyArg(t *testing.T) {
	//test for empty argument
	//emptyResult := models.IsPalindrome("")
	emptyResult := IsPalindrome("")

	if emptyResult != false {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "false", emptyResult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "false", emptyResult)
	}
}

>>>>>>> 7ba05b6a2ae41480890fe2ff061ada0edcb722e7
//test for valid postive arguments
func TestIsPalindromeValidPosArg(t *testing.T) {
	//postiveresult := models.IsPalindrome("level")
	postiveresult := IsPalindrome("level")

	if postiveresult != true {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "true", postiveresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "true", postiveresult)
	}

}

//test for valid negative arguments
func TestIsPalindromeValidNegArg(t *testing.T) {
	//negativeresult := models.IsPalindrome("test")
	negativeresult := IsPalindrome("test")

	if negativeresult != false {
		t.Errorf("IsPalindrome function failed, expected `%v`, got `%v`", "false", negativeresult)
	} else {
		t.Logf("IsPalindrome function success, expected `%v`, got `%v`", "false", negativeresult)
	}

}
