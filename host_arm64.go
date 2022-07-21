// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package host

import (
	// Make sure CPU and board drivers are registered.
	_ "adev73/x/host/v3/allwinner"
	_ "adev73/x/host/v3/bcm283x"
	_ "adev73/x/host/v3/pine64"
	_ "adev73/x/host/v3/rpi"
)
