package main

import (
  "sort"
  "strconv"
)

// Solution is your solution code.
func Solution (a int) int {
  digits := splitDigits(a)
  arrangedDigits := arrangement(digits)
  if (len(arrangedDigits) == 0) {
    return -1
  }

  return concatDigits(arrangedDigits)
}

func arrangement(digits []int) []int {
  i := len(digits) - 2
  for i >= 0 && digits[i] >= digits[i+1] {
    i--
  }

  if (i == -1) {
    return []int{}
  }

  j := len(digits) - 1
  for (digits[j] <= digits[i]) {
    j--
  }

  digits[i], digits[j] = digits[j], digits[i]
  sort.Ints(digits[i+1:])

  return digits
}

func concatDigits(digits []int) int {
  str := ""
  for _, digit := range digits {
    str += strconv.Itoa(digit)
  }

  res, _ := strconv.Atoi(str)
  return res
}

func splitDigits(num int) []int {
  str := strconv.Itoa(num)
  digits := make([]int, len(str))

  for k, v := range str {
    digits[k], _ = strconv.Atoi(string(v))
  }

  return digits
}
