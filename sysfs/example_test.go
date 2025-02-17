// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package sysfs_test

import (
	"fmt"
	"log"

	"github.com/adev73/host/v3"
	"github.com/adev73/host/v3/sysfs"

	"periph.io/x/conn/v3/gpio"
)

func ExampleLEDByName() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for _, led := range sysfs.LEDs {
		fmt.Printf("- %s: %s\n", led, led.Func())
	}
	led, err := sysfs.LEDByName("LED0")
	if err != nil {
		log.Fatalf("failed to find LED: %v", err)
	}
	if err := led.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
}
