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

	"github.com/forbole/juno/v4/node/remote"
	rbactypes "github.com/villagelabsco/village/x/rbac/types"
)

type Source struct {
	*remote.Source
	q rbactypes.QueryClient
}

func NewSource(source *remote.Source, q rbactypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (rbactypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &rbactypes.QueryParamsRequest{})
	if err != nil {
		return rbactypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetAuthorizations(height int64, req rbactypes.QueryGetAuthorizationsRequest) (rbactypes.QueryGetAuthorizationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Authorizations(ctx, &req)
	if err != nil {
		return rbactypes.QueryGetAuthorizationsResponse{}, fmt.Errorf("error while getting authorizations: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAuthorizations(height int64, req rbactypes.QueryAllAuthorizationsRequest) (rbactypes.QueryAllAuthorizationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AuthorizationsAll(ctx, &req)
	if err != nil {
		return rbactypes.QueryAllAuthorizationsResponse{}, fmt.Errorf("error while getting all authorizations: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAuthorizationsByNetwork(height int64, req rbactypes.QueryAllAuthorizationsByNetworkRequest) (rbactypes.QueryAllAuthorizationsByNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AuthorizationsByNetworkAll(ctx, &req)
	if err != nil {
		return rbactypes.QueryAllAuthorizationsByNetworkResponse{}, fmt.Errorf("error while getting all authorizations by network: %s", err)
	}
	return *res, nil
}

func (s Source) GetAuthorizationsForAccount(height int64, req rbactypes.QueryListAuthorizationsForAccountRequest) (rbactypes.QueryListAuthorizationsForAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ListAuthorizationsForAccount(ctx, &req)
	if err != nil {
		return rbactypes.QueryListAuthorizationsForAccountResponse{}, fmt.Errorf("error while getting authorizations for account: %s", err)
	}
	return *res, nil
}
