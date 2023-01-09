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

package local

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v3/node/local"
	villagetypes "github.com/villagelabs/villaged/x/village/types"
)

type Source struct {
	*local.Source
	q villagetypes.QueryClient
}

func NewSource(source *local.Source, q villagetypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (villagetypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(
		sdk.WrapSDKContext(ctx),
		&villagetypes.QueryParamsRequest{},
	)
	if err != nil {
		return villagetypes.Params{}, fmt.Errorf("error while querying params: %s", err)
	}

	return res.Params, nil
}

func (s Source) GetNetwork(height int64, req villagetypes.QueryGetNetworkRequest) (villagetypes.QueryGetNetworkResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryGetNetworkResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Network(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryGetNetworkResponse{}, fmt.Errorf("error while querying network: %s", err)
	}

	return *res, nil
}

func (s Source) GetAllNetwork(height int64, req villagetypes.QueryAllNetworkRequest) (villagetypes.QueryAllNetworkResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryAllNetworkResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.NetworkAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryAllNetworkResponse{}, fmt.Errorf("error while querying network: %s", err)
	}

	return *res, nil
}

func (s Source) GetUserNetworks(height int64, req villagetypes.QueryGetUserNetworksRequest) (villagetypes.QueryGetUserNetworksResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryGetUserNetworksResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.UserNetworks(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryGetUserNetworksResponse{}, fmt.Errorf("error while querying user networks: %s", err)
	}

	return *res, nil
}

func (s Source) GetAllUserNetworks(height int64, req villagetypes.QueryAllUserNetworksRequest) (villagetypes.QueryAllUserNetworksResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryAllUserNetworksResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.UserNetworksAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryAllUserNetworksResponse{}, fmt.Errorf("error while querying user networks: %s", err)
	}

	return *res, nil
}

func (s Source) GetNbNetworkCreationPerDay(height int64, req villagetypes.QueryGetNbNetworkCreationPerDayRequest) (villagetypes.QueryGetNbNetworkCreationPerDayResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryGetNbNetworkCreationPerDayResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.NbNetworkCreationPerDay(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryGetNbNetworkCreationPerDayResponse{}, fmt.Errorf("error while querying nb network creation per day: %s", err)
	}

	return *res, nil
}

func (s Source) GetAllNbNetworkCreationPerDay(height int64, req villagetypes.QueryAllNbNetworkCreationPerDayRequest) (villagetypes.QueryAllNbNetworkCreationPerDayResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryAllNbNetworkCreationPerDayResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.NbNetworkCreationPerDayAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryAllNbNetworkCreationPerDayResponse{}, fmt.Errorf("error while querying nb network creation per day: %s", err)
	}

	return *res, nil
}

func (s Source) GetAccountsInNetwork(height int64, req villagetypes.QueryGetAccountsInNetworkRequest) (villagetypes.QueryGetAccountsInNetworkResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryGetAccountsInNetworkResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.AccountsInNetwork(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryGetAccountsInNetworkResponse{}, err
	}

	return *res, nil
}

func (s Source) GetAllAccountsInNetwork(height int64, req villagetypes.QueryAllAccountsInNetworkRequest) (villagetypes.QueryAllAccountsInNetworkResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return villagetypes.QueryAllAccountsInNetworkResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.AccountsInNetworkAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return villagetypes.QueryAllAccountsInNetworkResponse{}, err
	}

	return *res, nil
}
