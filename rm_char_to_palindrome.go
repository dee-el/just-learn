package main


// Solution is your solution code.
func Solution (a string) bool {
  r := ValidPalindrome{
    Str: a,
    Low : 0,
    High : len(a) - 1,
  }

  res := isPossiblePalindromeByRemoveChar(r) 
  return res
}

type ValidPalindrome struct {
  Str string
  Low int
  High int
}

func isPalindrome(str string, low int, high int) bool {
  for low < high {
    if (str[low] != str[high]) {
      return false
    }

    low++
    high--
  }

  return true
}

func isPossiblePalindromeByRemoveChar(r ValidPalindrome) bool {
  low := r.Low
  high := r.High

  for low < high {
    if (r.Str[low] != r.Str[high]) {
      return isPalindrome(r.Str, low + 1, high) || isPalindrome(r.Str, low, high - 1) 
    } 

    low++
    high--
  }

  return true
}
