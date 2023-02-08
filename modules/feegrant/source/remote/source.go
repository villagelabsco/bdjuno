/*
 * Copyright 2022 LimeChain Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remote

import (
	"fmt"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/villagelabsco/juno/v4/node/remote"
)

type Source struct {
	*remote.Source
	q feegranttypes.QueryClient
}

func NewSource(source *remote.Source, q feegranttypes.QueryClient) *Source {
	return &Source{Source: source, q: q}
}

func (s Source) GetAllowances(height int64, req feegranttypes.QueryAllowancesRequest) (feegranttypes.QueryAllowancesResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Allowances(ctx, &req)
	if err != nil {
		return feegranttypes.QueryAllowancesResponse{}, fmt.Errorf("error while fetching allowances: %v", err)
	}
	return *res, nil
}