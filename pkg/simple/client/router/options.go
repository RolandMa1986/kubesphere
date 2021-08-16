/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package router

import (
	"runtime"
	"time"

	"github.com/spf13/pflag"

	"kubesphere.io/kubesphere/pkg/utils/reflectutils"
)

// Options contains configuration of the default router
type Options struct {
	WatchesPath             string `json:"watchesPath,omitempty" yaml:"watchesPath"`
	Namespace               string `json:"namespace,omitempty" yaml:"namespace"`
	ReconcilePeriod         time.Duration
	MaxConcurrentReconciles int
}

// NewRouterOptions creates a default router
func NewRouterOptions() *Options {
	return &Options{
		WatchesPath:             "",
		Namespace:               "",
		ReconcilePeriod:         time.Minute,
		MaxConcurrentReconciles: runtime.NumCPU(),
	}
}

func (s *Options) IsEmpty() bool {
	return s.WatchesPath == ""
}

// Validate check options values
func (s *Options) Validate() []error {
	var errors []error

	return errors
}

// ApplyTo overrides options if it's valid, which watchesPath is not empty
func (s *Options) ApplyTo(options *Options) {
	if s.WatchesPath != "" {
		reflectutils.Override(options, s)
	}
}

// AddFlags add options flags to command line flags,
// if watchesPath if left empty, following options will be ignored
func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.WatchesPath, "watchesPath", c.WatchesPath, "Path to the watches file to use.")
	fs.DurationVar(&c.ReconcilePeriod, "reconcile-period", c.ReconcilePeriod, "Default reconcile period for controllers")
	fs.IntVar(&c.MaxConcurrentReconciles, "max-concurrent-reconciles", c.MaxConcurrentReconciles, "Maximum number of concurrent reconciles for controllers.")
}
