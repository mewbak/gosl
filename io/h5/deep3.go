// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows,!darwin

package h5

/*
#include "hdf5.h"
#include "hdf5_hl.h"
#include "stdlib.h"
*/
import "C"

import "github.com/cpmech/gosl/utl"

// Deep3 ///////////////////////////////////////////////////////////////////////////////////////

// Deep3Put puts a deep slice with 3 levels and name described in path into HDF5 file
//  NOTE: path = "/mydeep3"  or   path = "/group/mydeep3"
func (o *File) Deep3Put(path string, a [][][]float64) {
	I, P, S := utl.SerializeDeep3(a)
	o.putArray(path+"/S", []int{len(S)}, S)
	o.putArrayIntNoGroups(path+"/I", I)
	o.putArrayIntNoGroups(path+"/P", P)
}

// Deep3Read reads a deep slice with 3 levels from file
func (o *File) Deep3Read(path string) (a [][][]float64) {
	_, S := o.getArray(path+"/S", false) // ismat=false
	_, I := o.getArrayInt(path+"/I", false)
	_, P := o.getArrayInt(path+"/P", false)
	a = utl.DeserializeDeep3(I, P, S, false)
	return
}
