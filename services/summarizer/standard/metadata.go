// Copyright © 2021, 2023 Weald Technology Limited.
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
	"encoding/json"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// metadata stored about this service.
type metadata struct {
	LastValidatorEpoch       phase0.Epoch `json:"latest_validator_epoch"`
	LastBlockEpoch           phase0.Epoch `json:"latest_block_epoch"`
	LastEpoch                phase0.Epoch `json:"latest_epoch"`
	LastValidatorDay         int64        `json:"last_validator_day"`
	PeriodicValidatorRollups bool         `json:"periodic_validator_rollups"`
}

// metadataKey is the key for the metadata.
var metadataKey = "summarizer.standard"

// getMetadata gets metadata for this service.
func (s *Service) getMetadata(ctx context.Context) (*metadata, error) {
	md := &metadata{
		LastValidatorDay: -1,
	}
	mdJSON, err := s.chainDB.Metadata(ctx, metadataKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch metadata")
	}
	if mdJSON == nil {
		return md, nil
	}
	if err := json.Unmarshal(mdJSON, md); err != nil {
		// Assume this is the old format.
		omd := &oldmetadata{}
		if err := json.Unmarshal(mdJSON, omd); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal metadata")
		}
		// Convert to new format.
		md.LastValidatorEpoch = phase0.Epoch(omd.LastValidatorEpoch)
		md.LastBlockEpoch = phase0.Epoch(omd.LastBlockEpoch)
		md.LastEpoch = phase0.Epoch(omd.LastEpoch)
		md.LastValidatorDay = omd.LastValidatorDay
		md.PeriodicValidatorRollups = omd.PeriodicValidatorRollups
	}

	return md, nil
}

// setMetadata sets metadata for this service.
func (s *Service) setMetadata(ctx context.Context, md *metadata) error {
	mdJSON, err := json.Marshal(md)
	if err != nil {
		return errors.Wrap(err, "failed to marshal metadata")
	}
	if err := s.chainDB.SetMetadata(ctx, metadataKey, mdJSON); err != nil {
		return errors.Wrap(err, "failed to update metadata")
	}
	return nil
}

// oldmetadata is pre-0.8.8 metadata using unquoted strings for epochs.
type oldmetadata struct {
	LastValidatorEpoch       uint64 `json:"latest_validator_epoch"`
	LastBlockEpoch           uint64 `json:"latest_block_epoch"`
	LastEpoch                uint64 `json:"latest_epoch"`
	LastValidatorDay         int64  `json:"last_validator_day"`
	PeriodicValidatorRollups bool   `json:"periodic_validator_rollups"`
}
