// ADOBE CONFIDENTIAL
// ___________________
//
// Copyright 2025 Adobe
// All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains
// the property of Adobe and its suppliers, if any. The intellectual
// and technical concepts contained herein are proprietary to Adobe
// and its suppliers and are protected by all applicable intellectual
// property laws, including trade secret and copyright laws.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Adobe.

package linters

import "fmt"

// ==============================================
// Default linters
// ==============================================

func callErr() { // also unused
	errorFunc() // tigger errcheck
}

//nolint:unused // demo
func errorFunc() error {
	return fmt.Errorf("an error")
}

//nolint:unused // demo
func foo() {
	x := 1 // tigger ineffassign
	x = 2
	println(x)
}

//nolint:unused // demo
func unreachableCode() {
	return
	println("This will never run") // trigger govet
}

type bar string // trigger unused

// ==============================================
// nolintlint
// ==============================================

// trigger nolintlint
// nolint
func trueFn() bool {
	return true
}

// ==============================================
// gocyclo
// ==============================================
//
//nolint:unused // demo
func checkConditions(a, b, c, d, e int) string { // trigger gocyclo
	if a > 0 {
		if b > 0 {
			a = 1
		} else {
			if c < 0 {
				a = 1
			} else if c > 0 {
				a = 2
			} else {
				a = 0
			}
		}
	} else if a < 0 {
		if c < 0 {
			if d < 0 {
				a = 1
			} else if d > 0 {
				a = 2
			} else {
				a = 0
			}
		}
	} else {
		if c < 0 {
			a = 1
		} else if c > 0 {
			a = 2
		} else {
			a = 0
		}
	}
	switch a {
	case 1:
		return "a is one"
	case 2:
		return "a is two"
	default:
		return "unknown"
	}
}

// ==============================================
// nestif
// ==============================================

//nolint:unused // demo
func nestedIfs(a, b, c, d int) bool { // trigger nestif
	if a > 0 {
		if b > 1 { // +1
			if c > 2 { // +2
				if d > 3 { // +3
					return false
				}
				return false

			}
			return true
		}
		return false
	}
	return true
}

// ==============================================
// gosec
// ==============================================

//nolint:unused // demo
func outOfBoundAccess() {
	data := []int{1, 2, 3}
	index := 5

	fmt.Println("Value:", data[index]) // trigger gosec
}
