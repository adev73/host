// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package host

import (
	// Make sure CPU and board drivers are registered.
	_ "github.com/adev73/host/v3/allwinner"
	_ "github.com/adev73/host/v3/bcm283x"
	_ "github.com/adev73/host/v3/pine64"
	_ "github.com/adev73/host/v3/rpi"
)
