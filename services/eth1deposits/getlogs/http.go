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

package getlogs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// post sends an HTTP post request and returns the body.
func (s *Service) post(ctx context.Context, endpoint string, body io.Reader) (io.Reader, error) {
	// #nosec G404
	log := log.With().Str("id", fmt.Sprintf("%02x", rand.Int31())).Logger()
	if e := log.Trace(); e.Enabled() {
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			return nil, errors.New("failed to read request body")
		}
		body = bytes.NewReader(bodyBytes)

		e.Str("endpoint", endpoint).Str("body", string(bodyBytes)).Msg("POST request")
	}

	reference, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}
	url := s.base.ResolveReference(reference).String()

	opCtx, cancel := context.WithTimeout(ctx, s.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodPost, url, body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create POST request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to call POST endpoint")
	}
	// skipcq:GO-S2307
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read POST response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		return nil, fmt.Errorf("POST failed with status %d: %s", resp.StatusCode, string(data))
	}
	cancel()

	log.Trace().Str("response", string(data)).Msg("POST response")

	return bytes.NewReader(data), nil
}
