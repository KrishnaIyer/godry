// Copyright © 2020 Krishna Iyer Easwaran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"context"

	"go.uber.org/zap"
)

type loggerKeyType string

var loggerKey loggerKeyType = "logger"

// FromContext retrieves a logger from a context and panics if there isn't one.
func FromContext(ctx context.Context) *zap.Logger {
	val := ctx.Value(loggerKey)
	logger, ok := val.(*zap.Logger)
	if !ok {
		panic("No logger in context")
	}
	return logger
}

// New returns a new zap logger.
func New() (*zap.Logger, error) {
	cfg := zap.Config{
		DisableCaller:     true,
		DisableStacktrace: true,
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

// NewContext returns a new context with a logger and panics if a nil value is passed.
func NewContext(ctx context.Context, logger *zap.Logger) context.Context {
	if logger == nil {
		panic("Nil Logger")
	}
	return context.WithValue(ctx, loggerKey, logger)
}
