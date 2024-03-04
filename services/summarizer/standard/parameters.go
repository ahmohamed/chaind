// Copyright © 2021 Weald Technology Limited.
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

package standard

import (
	"errors"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/rs/zerolog"
	"github.com/wealdtech/chaind/services/chaindb"
	"github.com/wealdtech/chaind/services/chaintime"
	"github.com/wealdtech/chaind/services/metrics"
)

type parameters struct {
	logLevel                  zerolog.Level
	monitor                   metrics.Service
	eth2Client                eth2client.Service
	chainDB                   chaindb.Service
	chainTime                 chaintime.Service
	epochSummaries            bool
	blockSummaries            bool
	validatorSummaries        bool
	validatorRetain           []phase0.BLSPubKey
	validatorEpochRetention   string
	maxDaysPerRun             uint64
	validatorBalanceRetention string
}

// Parameter is the interface for service parameters.
type Parameter interface {
	apply(p *parameters)
}

type parameterFunc func(*parameters)

func (f parameterFunc) apply(p *parameters) {
	f(p)
}

// WithLogLevel sets the log level for the module.
func WithLogLevel(logLevel zerolog.Level) Parameter {
	return parameterFunc(func(p *parameters) {
		p.logLevel = logLevel
	})
}

// WithMonitor sets the monitor for the module.
func WithMonitor(monitor metrics.Service) Parameter {
	return parameterFunc(func(p *parameters) {
		p.monitor = monitor
	})
}

// WithETH2Client sets the Ethereum 2 client for this module.
func WithETH2Client(eth2Client eth2client.Service) Parameter {
	return parameterFunc(func(p *parameters) {
		p.eth2Client = eth2Client
	})
}

// WithChainDB sets the chain database for this module.
func WithChainDB(chainDB chaindb.Service) Parameter {
	return parameterFunc(func(p *parameters) {
		p.chainDB = chainDB
	})
}

// WithChainTime sets the chain time service for this module.
func WithChainTime(chainTime chaintime.Service) Parameter {
	return parameterFunc(func(p *parameters) {
		p.chainTime = chainTime
	})
}

// WithEpochSummaries states if the module should generate epoch summaries.
func WithEpochSummaries(enabled bool) Parameter {
	return parameterFunc(func(p *parameters) {
		p.epochSummaries = enabled
	})
}

// WithBlockSummaries states if the module should generate block summaries.
func WithBlockSummaries(enabled bool) Parameter {
	return parameterFunc(func(p *parameters) {
		p.blockSummaries = enabled
	})
}

// WithValidatorSummaries states if the module should generate validator summaries.
func WithValidatorSummaries(enabled bool) Parameter {
	return parameterFunc(func(p *parameters) {
		p.validatorSummaries = enabled
	})
}

// WithValidatorRetain states if the module should retain balance and epoch summaries for a subset of validator.
func WithValidatorRetain(validatorRetain []phase0.BLSPubKey) Parameter {
	return parameterFunc(func(p *parameters) {
		p.validatorRetain = validatorRetain
	})
}

// WithMaxDaysPerRun provides the maximum number of days to process in a single run of the summarizer.
func WithMaxDaysPerRun(maxDaysPerRun uint64) Parameter {
	return parameterFunc(func(p *parameters) {
		p.maxDaysPerRun = maxDaysPerRun
	})
}

// WithValidatorEpochRetention provides the amount of validator epoch data to retain.
func WithValidatorEpochRetention(retention string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.validatorEpochRetention = retention
	})
}

// WithValidatorBalanceRetention provides the amount of validator balance data to retain.
func WithValidatorBalanceRetention(retention string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.validatorBalanceRetention = retention
	})
}

// parseAndCheckParameters parses and checks parameters to ensure that mandatory parameters are present and correct.
func parseAndCheckParameters(params ...Parameter) (*parameters, error) {
	parameters := parameters{
		logLevel: zerolog.GlobalLevel(),
	}
	for _, p := range params {
		if params != nil {
			p.apply(&parameters)
		}
	}

	if parameters.eth2Client == nil {
		return nil, errors.New("no consensus client specified")
	}
	if parameters.chainDB == nil {
		return nil, errors.New("no chain database specified")
	}
	if parameters.chainTime == nil {
		return nil, errors.New("no chain time specified")
	}
	if parameters.maxDaysPerRun == 0 {
		return nil, errors.New("no max days per run specified")
	}

	return &parameters, nil
}
