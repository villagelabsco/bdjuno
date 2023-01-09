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
	"github.com/forbole/juno/v3/node/remote"
	villagetypes "github.com/villagelabs/villaged/x/village/types"
)

type Source struct {
	*remote.Source
	q villagetypes.QueryClient
}

func NewSource(source *remote.Source, querier villagetypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      querier,
	}
}

func (s Source) GetParams(height int64) (villagetypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &villagetypes.QueryParamsRequest{})
	if err != nil {
		return villagetypes.Params{}, err
	}
	return res.Params, nil
}

func (s Source) GetNetwork(height int64, req villagetypes.QueryGetNetworkRequest) (villagetypes.QueryGetNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Network(ctx, &req)
	if err != nil {
		return villagetypes.QueryGetNetworkResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllNetwork(height int64, req villagetypes.QueryAllNetworkRequest) (villagetypes.QueryAllNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkAll(ctx, &req)
	if err != nil {
		return villagetypes.QueryAllNetworkResponse{}, err
	}
	return *res, nil
}

func (s Source) GetUserNetworks(height int64, req villagetypes.QueryGetUserNetworksRequest) (villagetypes.QueryGetUserNetworksResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.UserNetworks(ctx, &req)
	if err != nil {
		return villagetypes.QueryGetUserNetworksResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllUserNetworks(height int64, req villagetypes.QueryAllUserNetworksRequest) (villagetypes.QueryAllUserNetworksResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.UserNetworksAll(ctx, &req)
	if err != nil {
		return villagetypes.QueryAllUserNetworksResponse{}, err
	}
	return *res, nil
}

func (s Source) GetNbNetworkCreationPerDay(height int64, req villagetypes.QueryGetNbNetworkCreationPerDayRequest) (villagetypes.QueryGetNbNetworkCreationPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbNetworkCreationPerDay(ctx, &req)
	if err != nil {
		return villagetypes.QueryGetNbNetworkCreationPerDayResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllNbNetworkCreationPerDay(height int64, req villagetypes.QueryAllNbNetworkCreationPerDayRequest) (villagetypes.QueryAllNbNetworkCreationPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbNetworkCreationPerDayAll(ctx, &req)
	if err != nil {
		return villagetypes.QueryAllNbNetworkCreationPerDayResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAccountsInNetwork(height int64, req villagetypes.QueryGetAccountsInNetworkRequest) (villagetypes.QueryGetAccountsInNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountsInNetwork(ctx, &req)
	if err != nil {
		return villagetypes.QueryGetAccountsInNetworkResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllAccountsInNetwork(height int64, req villagetypes.QueryAllAccountsInNetworkRequest) (villagetypes.QueryAllAccountsInNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountsInNetworkAll(ctx, &req)
	if err != nil {
		return villagetypes.QueryAllAccountsInNetworkResponse{}, err
	}
	return *res, nil
}
