// Copyright © 2020 Weald Technology Trading.
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
	LatestEpoch         phase0.Epoch   `json:"latest_epoch"`
	LatestBalancesEpoch phase0.Epoch   `json:"latest_balances_epoch"`
	MissedEpochs        []phase0.Epoch `json:"missed_epochs,omitempty"`
}

// metadataKey is the key for the metadata.
var metadataKey = "validators.standard"

// getMetadata gets metadata for this service.
func (s *Service) getMetadata(ctx context.Context) (*metadata, error) {
	md := &metadata{}
	mdJSON, err := s.chainDB.Metadata(ctx, metadataKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain metadata")
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
		md.LatestEpoch = phase0.Epoch(omd.LatestEpoch)
		md.LatestBalancesEpoch = phase0.Epoch(omd.LatestBalancesEpoch)
		md.MissedEpochs = make([]phase0.Epoch, len(omd.MissedEpochs))
		for i := range omd.MissedEpochs {
			md.MissedEpochs[i] = phase0.Epoch(omd.MissedEpochs[i])
		}
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
	LatestEpoch         uint64   `json:"latest_epoch"`
	LatestBalancesEpoch uint64   `json:"latest_balances_epoch"`
	MissedEpochs        []uint64 `json:"missed_epochs,omitempty"`
}
