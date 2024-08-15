// Copyright © 2022 Weald Technology Trading.
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

package chaindb

import (
	"context"
	"time"

	api "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/wealdtech/chaind/services/chaindb"
)

type service struct{}

// New creates a new mock chain database.
func New() chaindb.Service {
	return &service{}
}

// AttestationsForBlock fetches all attestations made for the given block.
func (s *service) AttestationsForBlock(_ context.Context, _ phase0.Root) ([]*chaindb.Attestation, error) {
	return nil, nil
}

// AttestationsInBlock fetches all attestations contained in the given block.
func (s *service) AttestationsInBlock(_ context.Context, _ phase0.Root) ([]*chaindb.Attestation, error) {
	return nil, nil
}

// AttestationsForSlotRange fetches all attestations made for the given slot range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startSlot 2 and endSlot 4 will provide
// attestations for slots 2 and 3.
func (s *service) AttestationsForSlotRange(_ context.Context,
	_ phase0.Slot,
	_ phase0.Slot,
) (
	[]*chaindb.Attestation,
	error,
) {
	return nil, nil
}

// AttestationsInSlotRange fetches all attestations made in the given slot range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startSlot 2 and endSlot 4 will provide
// attestations in slots 2 and 3.
func (s *service) AttestationsInSlotRange(_ context.Context,
	_ phase0.Slot,
	_ phase0.Slot,
) (
	[]*chaindb.Attestation,
	error,
) {
	return nil, nil
}

// IndeterminateAttestationSlots fetches the slots in the given range with attestations that do not have a canonical status.
func (s *service) IndeterminateAttestationSlots(_ context.Context,
	_ phase0.Slot,
	_ phase0.Slot,
) (
	[]phase0.Slot,
	error,
) {
	return nil, nil
}

// SetAttestation sets an attestation.
func (s *service) SetAttestation(_ context.Context, _ *chaindb.Attestation) error {
	return nil
}

// AttesterSlashingsForSlotRange fetches all attester slashings made for the given slot range.
// It will return slashings from blocks that are canonical or undefined, but not from non-canonical blocks.
func (s *service) AttesterSlashingsForSlotRange(_ context.Context,
	_ phase0.Slot,
	_ phase0.Slot,
) (
	[]*chaindb.AttesterSlashing,
	error,
) {
	return nil, nil
}

// AttesterSlashingsForValidator fetches all attester slashings made for the given validator.
// It will return slashings from blocks that are canonical or undefined, but not from non-canonical blocks.
func (s *service) AttesterSlashingsForValidator(_ context.Context,
	_ phase0.ValidatorIndex,
) (
	[]*chaindb.AttesterSlashing,
	error,
) {
	return nil, nil
}

// SetAttesterSlashing sets an attester slashing.
func (s *service) SetAttesterSlashing(_ context.Context, _ *chaindb.AttesterSlashing) error {
	return nil
}

// BeaconCommitteeBySlotAndIndex fetches the beacon committee with the given slot and index.
func (s *service) BeaconCommitteeBySlotAndIndex(_ context.Context,
	_ phase0.Slot,
	_ phase0.CommitteeIndex,
) (
	*chaindb.BeaconCommittee,
	error,
) {
	return nil, nil
}

// AttesterDuties fetches the attester duties at the given slot range for the given validator indices.
func (s *service) AttesterDuties(_ context.Context,
	_ phase0.Slot,
	_ phase0.Slot,
	_ []phase0.ValidatorIndex,
) (
	[]*chaindb.AttesterDuty,
	error,
) {
	return nil, nil
}

// SetBeaconCommittee sets a beacon committee.
func (s *service) SetBeaconCommittee(_ context.Context, _ *chaindb.BeaconCommittee) error {
	return nil
}

// Blocks provides blocks according to the filter.
func (s *service) Blocks(_ context.Context, _ *chaindb.BlockFilter) ([]*chaindb.Block, error) {
	return []*chaindb.Block{}, nil
}

// BlocksBySlot fetches all blocks with the given slot.
func (s *service) BlocksBySlot(_ context.Context, _ phase0.Slot) ([]*chaindb.Block, error) {
	return nil, nil
}

// BlocksForSlotRange fetches all blocks with the given slot range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startSlot 2 and endSlot 4 will provide
// blocks duties for slots 2 and 3.
func (s *service) BlocksForSlotRange(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]*chaindb.Block, error) {
	return nil, nil
}

// BlockByRoot fetches the block with the given root.
func (s *service) BlockByRoot(_ context.Context, _ phase0.Root) (*chaindb.Block, error) {
	return nil, nil
}

// BlocksByParentRoot fetches the blocks with the given parent root.
func (s *service) BlocksByParentRoot(_ context.Context, _ phase0.Root) ([]*chaindb.Block, error) {
	return nil, nil
}

// EmptySlots fetches the slots in the given range without a block in the database.
func (s *service) EmptySlots(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]phase0.Slot, error) {
	return nil, nil
}

// LatestBlocks fetches the blocks with the highest slot number in the database.
func (s *service) LatestBlocks(_ context.Context) ([]*chaindb.Block, error) {
	return nil, nil
}

// IndeterminateBlocks fetches the blocks in the given range that do not have a canonical status.
func (s *service) IndeterminateBlocks(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]phase0.Root, error) {
	return nil, nil
}

// CanonicalBlockPresenceForSlotRange returns a boolean for each slot in the range for the presence
// of a canonical block.
// Ranges are inclusive of start and exclusive of end i.e. a request with startSlot 2 and endSlot 4 will provide
// presence duties for slots 2 and 3.
func (s *service) CanonicalBlockPresenceForSlotRange(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]bool, error) {
	return nil, nil
}

// LatestCanonicalBlock returns the slot of the latest canonical block known in the database.
func (s *service) LatestCanonicalBlock(_ context.Context) (phase0.Slot, error) {
	return 0, nil
}

// SetBlock sets a block.
func (s *service) SetBlock(_ context.Context, _ *chaindb.Block) error {
	return nil
}

// Spec provides the spec information of the chain.
func (s *service) Spec(ctx context.Context) (map[string]any, error) {
	return s.ChainSpec(ctx)
}

// ChainSpec fetches all chain specification values.
func (s *service) ChainSpec(_ context.Context) (map[string]any, error) {
	return map[string]any{
		"ALTAIR_FORK_EPOCH":                        uint64(74240),
		"ALTAIR_FORK_VERSION":                      phase0.Version{0x01, 0x00, 0x00, 0x00},
		"BASE_REWARD_FACTOR":                       uint64(64),
		"BELLATRIX_FORK_EPOCH":                     uint64(18446744073709551615),
		"BELLATRIX_FORK_VERSION":                   phase0.Version{0x02, 0x00, 0x00, 0x00},
		"BLS_WITHDRAWAL_PREFIX":                    []byte{0x00},
		"CHURN_LIMIT_QUOTIENT":                     uint64(65536),
		"CONFIG_NAME":                              "mainnet",
		"DEPOSIT_CHAIN_ID":                         1,
		"DEPOSIT_CONTRACT_ADDRESS":                 []byte{0x00, 0x00, 0x00, 0x00, 0x21, 0x9a, 0xb5, 0x40, 0x35, 0x6c, 0xBB, 0x83, 0x9C, 0xbe, 0x05, 0x30, 0x3d, 0x77, 0x05, 0xFa},
		"DEPOSIT_NETWORK_ID":                       1,
		"DOMAIN_AGGREGATE_AND_PROOF":               phase0.DomainType{0x06, 0x00, 0x00, 0x00},
		"DOMAIN_BEACON_ATTESTER":                   phase0.DomainType{0x01, 0x00, 0x00, 0x00},
		"DOMAIN_BEACON_PROPOSER":                   phase0.DomainType{0x00, 0x00, 0x00, 0x00},
		"DOMAIN_CONTRIBUTION_AND_PROOF":            phase0.DomainType{0x09, 0x00, 0x00, 0x00},
		"DOMAIN_DEPOSIT":                           phase0.DomainType{0x03, 0x00, 0x00, 0x00},
		"DOMAIN_RANDAO":                            phase0.DomainType{0x02, 0x00, 0x00, 0x00},
		"DOMAIN_SELECTION_PROOF":                   phase0.DomainType{0x05, 0x00, 0x00, 0x00},
		"DOMAIN_SYNC_COMMITTEE":                    phase0.DomainType{0x07, 0x00, 0x00, 0x00},
		"DOMAIN_SYNC_COMMITTEE_SELECTION_PROOF":    phase0.DomainType{0x08, 0x00, 0x00, 0x00},
		"DOMAIN_VOLUNTARY_EXIT":                    phase0.DomainType{0x04, 0x00, 0x00, 0x00},
		"EFFECTIVE_BALANCE_INCREMENT":              uint64(1000000000),
		"EJECTION_BALANCE":                         uint64(16000000000),
		"EPOCHS_PER_ETH1_VOTING_PERIOD":            uint64(64),
		"EPOCHS_PER_HISTORICAL_VECTOR":             uint64(65536),
		"EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION":    uint64(256),
		"EPOCHS_PER_SLASHINGS_VECTOR":              uint64(8192),
		"EPOCHS_PER_SYNC_COMMITTEE_PERIOD":         uint64(256),
		"ETH1_FOLLOW_DISTANCE":                     uint64(2048),
		"GENESIS_DELAY":                            604800 * time.Second,
		"GENESIS_FORK_VERSION":                     phase0.Version{0x00, 0x00, 0x00, 0x00},
		"HISTORICAL_ROOTS_LIMIT":                   uint64(16777216),
		"HYSTERESIS_DOWNWARD_MULTIPLIER":           uint64(1),
		"HYSTERESIS_QUOTIENT":                      uint64(4),
		"HYSTERESIS_UPWARD_MULTIPLIER":             uint64(5),
		"INACTIVITY_PENALTY_QUOTIENT":              uint64(67108864),
		"INACTIVITY_PENALTY_QUOTIENT_ALTAIR":       uint64(50331648),
		"INACTIVITY_PENALTY_QUOTIENT_MERGE":        uint64(16777216),
		"INACTIVITY_SCORE_BIAS":                    uint64(4),
		"INACTIVITY_SCORE_RECOVERY_RATE":           uint64(16),
		"MAX_ATTESTATIONS":                         uint64(128),
		"MAX_ATTESTER_SLASHINGS":                   uint64(2),
		"MAX_COMMITTEES_PER_SLOT":                  uint64(64),
		"MAX_DEPOSITS":                             uint64(16),
		"MAX_EFFECTIVE_BALANCE":                    uint64(32000000000),
		"MAX_PROPOSER_SLASHINGS":                   uint64(16),
		"MAX_SEED_LOOKAHEAD":                       uint64(4),
		"MAX_VALIDATORS_PER_COMMITTEE":             uint64(2048),
		"MAX_VOLUNTARY_EXITS":                      uint64(16),
		"MIN_ANCHOR_POW_BLOCK_DIFFICULTY":          uint64(4294967296),
		"MIN_ATTESTATION_INCLUSION_DELAY":          uint64(1),
		"MIN_DEPOSIT_AMOUNT":                       uint64(1000000000),
		"MIN_EPOCHS_TO_INACTIVITY_PENALTY":         uint64(4),
		"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT":       uint64(16384),
		"MIN_GENESIS_TIME":                         time.Unix(1606824000, 0),
		"MIN_PER_EPOCH_CHURN_LIMIT":                uint64(4),
		"MIN_SEED_LOOKAHEAD":                       uint64(1),
		"MIN_SLASHING_PENALTY_QUOTIENT":            uint64(128),
		"MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR":     uint64(64),
		"MIN_SLASHING_PENALTY_QUOTIENT_MERGE":      uint64(32),
		"MIN_SYNC_COMMITTEE_PARTICIPANTS":          uint64(1),
		"MIN_VALIDATOR_WITHDRAWABILITY_DELAY":      uint64(256),
		"PRESET_BASE":                              "mainnet",
		"PROPORTIONAL_SLASHING_MULTIPLIER":         uint64(1),
		"PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR":  uint64(2),
		"PROPORTIONAL_SLASHING_MULTIPLIER_MERGE":   uint64(3),
		"PROPOSER_REWARD_QUOTIENT":                 uint64(8),
		"PROPOSER_WEIGHT":                          uint64(8),
		"RANDOM_SUBNETS_PER_VALIDATOR":             uint64(1),
		"SAFE_SLOTS_TO_UPDATE_JUSTIFIED":           uint64(8),
		"SECONDS_PER_ETH1_BLOCK":                   14 * time.Second,
		"SECONDS_PER_SLOT":                         12 * time.Second,
		"SHARDING_FORK_EPOCH":                      uint64(18446744073709551615),
		"SHARDING_FORK_VERSION":                    phase0.Version{0x03, 0x00, 0x00, 0x00},
		"SHARD_COMMITTEE_PERIOD":                   uint64(256),
		"SHUFFLE_ROUND_COUNT":                      uint64(90),
		"SLOTS_PER_EPOCH":                          uint64(32),
		"SLOTS_PER_HISTORICAL_ROOT":                uint64(8192),
		"SYNC_COMMITTEE_SIZE":                      uint64(512),
		"SYNC_COMMITTEE_SUBNET_COUNT":              uint64(4),
		"SYNC_REWARD_WEIGHT":                       uint64(2),
		"TARGET_AGGREGATORS_PER_COMMITTEE":         uint64(16),
		"TARGET_AGGREGATORS_PER_SYNC_SUBCOMMITTEE": uint64(16),
		"TARGET_COMMITTEE_SIZE":                    uint64(128),
		"TERMINAL_BLOCK_HASH":                      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		"TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH":     uint64(18446744073709551615),
		"TERMINAL_TOTAL_DIFFICULTY":                uint64(0),
		"TIMELY_HEAD_FLAG_INDEX":                   []byte{0x02},
		"TIMELY_HEAD_WEIGHT":                       uint64(14),
		"TIMELY_SOURCE_FLAG_INDEX":                 []byte{0x00},
		"TIMELY_SOURCE_WEIGHT":                     uint64(14),
		"TIMELY_TARGET_FLAG_INDEX":                 []byte{0x01},
		"TIMELY_TARGET_WEIGHT":                     uint64(26),
		"TRANSITION_TOTAL_DIFFICULTY":              uint64(0),
		"VALIDATOR_REGISTRY_LIMIT":                 uint64(1099511627776),
		"WEIGHT_DENOMINATOR":                       uint64(64),
		"WHISTLEBLOWER_REWARD_QUOTIENT":            uint64(512),
	}, nil
}

// ChainSpecValue fetches a chain specification value given its key.
func (s *service) ChainSpecValue(_ context.Context, _ string) (any, error) {
	return nil, nil
}

// SetChainSpecValue sets the value of the provided key.
func (s *service) SetChainSpecValue(_ context.Context, _ string, _ any) error {
	return nil
}

// ForkSchedule provides details of past and future changes in the chain's fork version.
func (s *service) ForkSchedule(_ context.Context) ([]*phase0.Fork, error) {
	return nil, nil
}

// SetForkSchedule sets the fork schedule.
func (s *service) SetForkSchedule(_ context.Context, _ []*phase0.Fork) error {
	return nil
}

// Genesis fetches genesis values.
func (s *service) Genesis(_ context.Context) (*api.Genesis, error) {
	return nil, nil
}

// SetGenesis sets the genesis information.
func (s *service) SetGenesis(_ context.Context, _ *api.Genesis) error {
	return nil
}

// ETH1DepositsByPublicKey fetches Ethereum 1 deposits for a given set of validator public keys.
func (s *service) ETH1DepositsByPublicKey(_ context.Context, _ []phase0.BLSPubKey) ([]*chaindb.ETH1Deposit, error) {
	return nil, nil
}

// SetETH1Deposit sets an Ethereum 1 deposit.
func (s *service) SetETH1Deposit(_ context.Context, _ *chaindb.ETH1Deposit) error {
	return nil
}

// ProposerDuties provides proposer duties according to the filter.
func (*service) ProposerDuties(_ context.Context, _ *chaindb.ProposerDutyFilter) ([]*chaindb.ProposerDuty, error) {
	return nil, nil
}

// ProposerDutiesForSlotRange fetches all proposer duties for the given slot range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startSlot 2 and endSlot 4 will provide
// proposer duties for slots 2 and 3.
func (s *service) ProposerDutiesForSlotRange(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]*chaindb.ProposerDuty, error) {
	return nil, nil
}

// ProposerDutiesForValidator provides all proposer duties for the given validator index.
func (s *service) ProposerDutiesForValidator(_ context.Context, _ phase0.ValidatorIndex) ([]*chaindb.ProposerDuty, error) {
	return nil, nil
}

// SetProposerDuty sets a proposer duty.
func (s *service) SetProposerDuty(_ context.Context, _ *chaindb.ProposerDuty) error {
	return nil
}

// ProposerSlashingsForSlotRange fetches all proposer slashings made for the given slot range.
// It will return slashings from blocks that are canonical or undefined, but not from non-canonical blocks.
func (s *service) ProposerSlashingsForSlotRange(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]*chaindb.ProposerSlashing, error) {
	return nil, nil
}

// ProposerSlashingsForValidator fetches all proposer slashings made for the given validator.
// It will return slashings from blocks that are canonical or undefined, but not from non-canonical blocks.
func (s *service) ProposerSlashingsForValidator(_ context.Context, _ phase0.ValidatorIndex) ([]*chaindb.ProposerSlashing, error) {
	return nil, nil
}

// SetProposerSlashing sets an proposer slashing.
func (s *service) SetProposerSlashing(_ context.Context, _ *chaindb.ProposerSlashing) error {
	return nil
}

// SyncAggregateForBlock provides the sync aggregate for the supplied block root.
func (s *service) SyncAggregateForBlock(_ context.Context, _ phase0.Root) (*chaindb.SyncAggregate, error) {
	return nil, nil
}

// SetSyncAggregate sets the sync aggregate.
func (s *service) SetSyncAggregate(_ context.Context, _ *chaindb.SyncAggregate) error {
	return nil
}

// Validators fetches all validators.
func (s *service) Validators(_ context.Context) ([]*chaindb.Validator, error) {
	return nil, nil
}

// ValidatorsByPublicKey fetches all validators matching the given public keys.
// This is a common starting point for external entities to query specific validators, as they should
// always have the public key at a minimum, hence the return map keyed by public key.
func (s *service) ValidatorsByPublicKey(_ context.Context,
	_ []phase0.BLSPubKey,
) (
	map[phase0.BLSPubKey]*chaindb.Validator,
	error,
) {
	return nil, nil
}

// ValidatorsByIndex fetches all validators matching the given indices.
func (s *service) ValidatorsByIndex(_ context.Context,
	_ []phase0.ValidatorIndex,
) (
	map[phase0.ValidatorIndex]*chaindb.Validator,
	error,
) {
	return nil, nil
}

// ValidatorBalancesByEpoch fetches all validator balances for the given epoch.
func (s *service) ValidatorBalancesByEpoch(
	_ context.Context,
	_ phase0.Epoch,
) (
	[]*chaindb.ValidatorBalance,
	error,
) {
	return nil, nil
}

// ValidatorBalancesByIndexAndEpoch fetches the validator balances for the given validators and epoch.
func (s *service) ValidatorBalancesByIndexAndEpoch(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ phase0.Epoch,
) (
	map[phase0.ValidatorIndex]*chaindb.ValidatorBalance,
	error,
) {
	return nil, nil
}

// ValidatorBalancesByIndexAndEpochRange fetches the validator balances for the given validators and epoch range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startEpoch 2 and endEpoch 4 will provide
// balances for epochs 2 and 3.
func (s *service) ValidatorBalancesByIndexAndEpochRange(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ phase0.Epoch,
	_ phase0.Epoch,
) (
	map[phase0.ValidatorIndex][]*chaindb.ValidatorBalance,
	error,
) {
	return nil, nil
}

// ValidatorBalancesByIndexAndEpochs fetches the validator balances for the given validators at the specified epochs.
func (s *service) ValidatorBalancesByIndexAndEpochs(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ []phase0.Epoch,
) (
	map[phase0.ValidatorIndex][]*chaindb.ValidatorBalance,
	error,
) {
	return nil, nil
}

// AggregateValidatorBalancesByIndexAndEpoch fetches the aggregate validator balances for the given validators and epoch.
func (s *service) AggregateValidatorBalancesByIndexAndEpoch(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ phase0.Epoch,
) (
	*chaindb.AggregateValidatorBalance,
	error,
) {
	return nil, nil
}

// AggregateValidatorBalancesByIndexAndEpochRange fetches the aggregate validator balances for the given validators and
// epoch range.
// Ranges are inclusive of start and exclusive of end i.e. a request with startEpoch 2 and endEpoch 4 will provide
// balances for epochs 2 and 3.
func (s *service) AggregateValidatorBalancesByIndexAndEpochRange(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ phase0.Epoch,
	_ phase0.Epoch,
) (
	[]*chaindb.AggregateValidatorBalance,
	error,
) {
	return nil, nil
}

// AggregateValidatorBalancesByIndexAndEpochs fetches the validator balances for the given validators at the specified epochs.
func (s *service) AggregateValidatorBalancesByIndexAndEpochs(_ context.Context,
	_ []phase0.ValidatorIndex,
	_ []phase0.Epoch,
) (
	[]*chaindb.AggregateValidatorBalance,
	error,
) {
	return nil, nil
}

// SetValidator sets a validator.
func (s *service) SetValidator(_ context.Context, _ *chaindb.Validator) error {
	return nil
}

// SetValidatorBalance sets a validator balance.
func (s *service) SetValidatorBalance(_ context.Context, _ *chaindb.ValidatorBalance) error {
	return nil
}

// SetValidatorBalances sets multiple validator balances.
func (s *service) SetValidatorBalances(_ context.Context, _ []*chaindb.ValidatorBalance) error {
	return nil
}

// DepositsByPublicKey fetches deposits for a given set of validator public keys.
func (s *service) DepositsByPublicKey(_ context.Context, _ []phase0.BLSPubKey) (map[phase0.BLSPubKey][]*chaindb.Deposit, error) {
	return nil, nil
}

// DepositsForSlotRange fetches all deposits made in the given slot range.
// It will return deposits from blocks that are canonical or undefined, but not from non-canonical blocks.
func (s *service) DepositsForSlotRange(_ context.Context, _ phase0.Slot, _ phase0.Slot) ([]*chaindb.Deposit, error) {
	return nil, nil
}

// SetDeposit sets a deposit.
func (s *service) SetDeposit(_ context.Context, _ *chaindb.Deposit) error {
	return nil
}

// SetVoluntaryExit sets a voluntary exit.
func (s *service) SetVoluntaryExit(_ context.Context, _ *chaindb.VoluntaryExit) error {
	return nil
}

// SetValidatorEpochSummary sets a validator epoch summary.
func (s *service) SetValidatorEpochSummary(_ context.Context, _ *chaindb.ValidatorEpochSummary) error {
	return nil
}

// SetValidatorEpochSummaries sets multiple validator epoch summaries.
func (s *service) SetValidatorEpochSummaries(_ context.Context, _ []*chaindb.ValidatorEpochSummary) error {
	return nil
}

// BlockSummaryForSlot obtains the summary of a block for a given slot.
func (s *service) BlockSummaryForSlot(_ context.Context, _ phase0.Slot) (*chaindb.BlockSummary, error) {
	return nil, nil
}

// ValidatorSummaries provides summaries according to the filter.
func (s *service) ValidatorSummaries(_ context.Context, _ *chaindb.ValidatorSummaryFilter) ([]*chaindb.ValidatorEpochSummary, error) {
	return nil, nil
}

// ValidatorSummariesForEpoch obtains all summaries for a given epoch.
func (s *service) ValidatorSummariesForEpoch(_ context.Context, _ phase0.Epoch) ([]*chaindb.ValidatorEpochSummary, error) {
	return nil, nil
}

// ValidatorSummaryForEpoch obtains the summary of a validator for a given epoch.
func (s *service) ValidatorSummaryForEpoch(_ context.Context, _ phase0.ValidatorIndex, _ phase0.Epoch) (*chaindb.ValidatorEpochSummary, error) {
	return nil, nil
}

// SetBlockSummary sets a block summary.
func (s *service) SetBlockSummary(_ context.Context, _ *chaindb.BlockSummary) error {
	return nil
}

// SetEpochSummary sets an epoch summary.
func (s *service) SetEpochSummary(_ context.Context, _ *chaindb.EpochSummary) error {
	return nil
}

// SyncCommittee provides a sync committee for the given sync committee period.
func (s *service) SyncCommittee(_ context.Context, _ uint64) (*chaindb.SyncCommittee, error) {
	return nil, nil
}

// SetSyncCommittee sets a sync committee.
func (s *service) SetSyncCommittee(_ context.Context, _ *chaindb.SyncCommittee) error {
	return nil
}

// Withdrawals provides withdrawals according to the filter.
func (s *service) Withdrawals(_ context.Context, _ *chaindb.WithdrawalFilter) ([]*chaindb.Withdrawal, error) {
	return []*chaindb.Withdrawal{}, nil
}

// BeginTx begins a transaction.
func (s *service) BeginTx(_ context.Context) (context.Context, context.CancelFunc, error) {
	return nil, nil, nil
}

// CommitTx commits a transaction.
func (s *service) CommitTx(_ context.Context) error {
	return nil
}

// BeginROTx begins a read-only transaction.
func (s *service) BeginROTx(_ context.Context) (context.Context, error) {
	return nil, nil
}

// CommitROTx commits a read-only transaction.
func (s *service) CommitROTx(_ context.Context) {}

// SetMetadata sets a metadata key to a JSON value.
func (s *service) SetMetadata(_ context.Context, _ string, _ []byte) error {
	return nil
}

// Metadata obtains the JSON value from a metadata key.
func (s *service) Metadata(_ context.Context, _ string) ([]byte, error) {
	return nil, nil
}
