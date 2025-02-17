// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package sysfssmoketest

import (
	"errors"
	"flag"
	"fmt"
	"sort"

	"github.com/adev73/host/v3/sysfs"

	"periph.io/x/conn/v3/gpio"
)

// Benchmark is imported by periph-smoketest.
type Benchmark struct {
	short bool
	p     *sysfs.Pin
	pull  gpio.Pull
}

// Name implements the SmokeTest interface.
func (s *Benchmark) Name() string {
	return "sysfs-benchmark"
}

// Description implements the SmokeTest interface.
func (s *Benchmark) Description() string {
	return "Benchmarks sysfs gpio functionality"
}

// Run implements the SmokeTest interface.
func (s *Benchmark) Run(f *flag.FlagSet, args []string) error {
	num := f.Int("p", -1, "Pin number to use")
	f.BoolVar(&s.short, "short", false, "Skip many partially redundant benchmarks")
	if err := f.Parse(args); err != nil {
		return err
	}

	if f.NArg() != 0 {
		f.Usage()
		return errors.New("unsupported flags")
	}
	if *num == -1 {
		f.Usage()
		return errors.New("-p is required")
	}
	if s.p = sysfs.Pins[*num]; s.p == nil {
		list := make([]int, 0, len(sysfs.Pins))
		for i := range sysfs.Pins {
			list = append(list, i)
		}
		sort.Ints(list)
		valid := ""
		for i, v := range list {
			if i == 0 {
				valid += fmt.Sprintf("%d", v)
			} else {
				valid += fmt.Sprintf(", %d", v)
			}
		}
		return fmt.Errorf("invalid pin %d; valid: %s", *num, valid)
	}
	s.pull = gpio.PullNoChange
	s.runGPIOBenchmark()
	return nil
}
