// +build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var dirStack []string

// ////////////////////////////////////////////////////////////////////////////////// //

// Push change current working directory and add previous working directory to stack
func Push(dir string) string {
	if dirStack == nil {
		dirStack = append(dirStack, Current())
	}

	err := os.Chdir(dir)

	if err != nil {
		return ""
	}

	wd, _ := os.Getwd()

	dirStack = append(dirStack, wd)

	return wd
}

// Pop change current working directory to previous in stack
func Pop() string {
	if dirStack == nil {
		dirStack = append(dirStack, Current())
	}

	dl := len(dirStack)

	switch dl {

	case 0, 1:
		// nop

	default:
		err := os.Chdir(dirStack[dl-2])

		if err != nil {
			return ""
		}

		dirStack = dirStack[0 : dl-1]
	}

	wd, _ := os.Getwd()

	return wd
}

// Current return current working directory
func Current() string {
	wd, err := os.Getwd()

	if err != nil {
		return ""
	}

	return wd
}
