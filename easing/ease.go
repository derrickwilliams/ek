// Package easing contains easing functions
package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DoublePi is 2 * Pi
const DoublePi = math.Pi * 2

// ////////////////////////////////////////////////////////////////////////////////// //

// Easing is type for any easing func
type Easing func(t, b, c, d float64) float64

// ////////////////////////////////////////////////////////////////////////////////// //
