// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package host

import (
	// Make sure CPU and board drivers are registered.
	_ "github.com/adev73/host/v3/allwinner"
	_ "github.com/adev73/host/v3/am335x"
	_ "github.com/adev73/host/v3/bcm283x"
	_ "github.com/adev73/host/v3/beagle/bone"
	_ "github.com/adev73/host/v3/beagle/green"
	_ "github.com/adev73/host/v3/chip"
	_ "github.com/adev73/host/v3/odroidc1"

	// While this board is ARM64, it may run ARM 32 bits binaries so load it on
	// 32 bits builds too.
	_ "github.com/adev73/host/v3/pine64"
	_ "github.com/adev73/host/v3/rpi"
)
