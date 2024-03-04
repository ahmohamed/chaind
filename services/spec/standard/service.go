// Copyright © 2021 Weald Technology Trading.
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
	"context"
	"time"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
	"github.com/wealdtech/chaind/services/chaindb"
)

// Service is a spec service.
type Service struct {
	eth2Client         eth2client.Service
	chainDB            chaindb.Service
	chainSpecSetter    chaindb.ChainSpecSetter
	genesisSetter      chaindb.GenesisSetter
	forkScheduleSetter chaindb.ForkScheduleSetter
}

// module-wide log.
var log zerolog.Logger

// New creates a new service.
func New(ctx context.Context, params ...Parameter) (*Service, error) {
	parameters, err := parseAndCheckParameters(params...)
	if err != nil {
		return nil, errors.Wrap(err, "problem with parameters")
	}

	// Set logging.
	log = zerologger.With().Str("service", "spec").Str("impl", "standard").Logger()
	if parameters.logLevel != log.GetLevel() {
		log = log.Level(parameters.logLevel)
	}

	chainSpecSetter, isChainSpecSetter := parameters.chainDB.(chaindb.ChainSpecSetter)
	if !isChainSpecSetter {
		return nil, errors.New("chain DB does not support chain spec setting")
	}

	genesisSetter, isGenesisSetter := parameters.chainDB.(chaindb.GenesisSetter)
	if !isGenesisSetter {
		return nil, errors.New("chain DB does not support genesis setting")
	}

	forkScheduleSetter, isForkScheduleSetter := parameters.chainDB.(chaindb.ForkScheduleSetter)
	if !isForkScheduleSetter {
		return nil, errors.New("chain DB does not support fork schedule setting")
	}

	s := &Service{
		eth2Client:         parameters.eth2Client,
		chainDB:            parameters.chainDB,
		chainSpecSetter:    chainSpecSetter,
		genesisSetter:      genesisSetter,
		forkScheduleSetter: forkScheduleSetter,
	}

	// Update spec in the _foreground_.  This ensures that spec information
	// is available to other modules when they start.
	s.updateSpec(ctx)

	// Set up a periodic refresh of the spec information.
	runtimeFunc := func(_ context.Context, _ any) (time.Time, error) {
		// Run daily.
		return time.Now().AddDate(0, 0, 1), nil
	}
	jobFunc := func(ctx context.Context, data any) {
		log.Trace().Msg("Updating spec")
		s := data.(*Service)
		s.updateSpec(ctx)
	}
	if err := parameters.scheduler.SchedulePeriodicJob(ctx, "spec", "update spec",
		runtimeFunc,
		nil,
		jobFunc,
		s,
	); err != nil {
		return nil, errors.Wrap(err, "failed to set up periodic refresh of spec")
	}

	return s, nil
}

func (s *Service) updateSpec(ctx context.Context) {
	ctx, cancel, err := s.chainDB.BeginTx(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to begin transaction")
		return
	}

	if err := s.updateChainSpec(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to update spec")
	}

	if err := s.updateGenesis(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to update genesis")
	}

	if err := s.updateForkSchedule(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to update fork schedule")
	}

	if err := s.chainDB.CommitTx(ctx); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func (s *Service) updateChainSpec(ctx context.Context) error {
	// Fetch the chain spec.
	specResponse, err := s.eth2Client.(eth2client.SpecProvider).Spec(ctx, &api.SpecOpts{})
	if err != nil {
		return errors.Wrap(err, "failed to obtain chain spec")
	}
	spec := specResponse.Data

	// Update the database.
	for k, v := range spec {
		if err := s.chainSpecSetter.SetChainSpecValue(ctx, k, v); err != nil {
			return errors.Wrap(err, "failed to set chain spec value")
		}
	}
	return nil
}

func (s *Service) updateGenesis(ctx context.Context) error {
	// Fetch genesis parameters.
	genesisResponse, err := s.eth2Client.(eth2client.GenesisProvider).Genesis(ctx, &api.GenesisOpts{})
	if err != nil {
		return errors.Wrap(err, "failed to obtain genesis")
	}
	genesis := genesisResponse.Data

	// Update the database.
	if err := s.genesisSetter.SetGenesis(ctx, genesis); err != nil {
		return errors.Wrap(err, "failed to set genesis")
	}

	return nil
}

func (s *Service) updateForkSchedule(ctx context.Context) error {
	// Fetch fork schedule.
	scheduleResponse, err := s.eth2Client.(eth2client.ForkScheduleProvider).ForkSchedule(ctx, &api.ForkScheduleOpts{})
	if err != nil {
		return errors.Wrap(err, "failed to obtain fork schedule")
	}
	schedule := scheduleResponse.Data

	// Remove any fork schedules that are in the far future.
	realSchedule := make([]*phase0.Fork, 0, len(schedule))
	for i := range schedule {
		if schedule[i].Epoch != 0xffffffffffffffff {
			realSchedule = append(realSchedule, schedule[i])
		}
	}

	// Update the database.
	if err := s.forkScheduleSetter.SetForkSchedule(ctx, realSchedule); err != nil {
		return errors.Wrap(err, "failed to set fork schedule")
	}

	return nil
}
