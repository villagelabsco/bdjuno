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
	"github.com/villagelabsco/juno/v4/node/local"
	marketplacetypes "github.com/villagelabsco/villaged/x/marketplace/types"
)

type Source struct {
	*local.Source
	q marketplacetypes.QueryClient
}

func NewSource(source *local.Source, q marketplacetypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (marketplacetypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return marketplacetypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(
		sdk.WrapSDKContext(ctx),
		&marketplacetypes.QueryParamsRequest{},
	)
	if err != nil {
		return marketplacetypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}

	return res.Params, nil
}

func (s Source) GetListing(height int64, req marketplacetypes.QueryGetListingRequest) (marketplacetypes.QueryGetListingResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return marketplacetypes.QueryGetListingResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Listing(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return marketplacetypes.QueryGetListingResponse{}, fmt.Errorf("error while getting listing: %s", err)
	}

	return *res, nil
}

func (s Source) GetAllListing(height int64, req marketplacetypes.QueryAllListingRequest) (marketplacetypes.QueryAllListingResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return marketplacetypes.QueryAllListingResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.ListingAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return marketplacetypes.QueryAllListingResponse{}, fmt.Errorf("error while getting all listings: %s", err)
	}

	return *res, nil
}

func (s Source) GetOrder(height int64, req marketplacetypes.QueryGetOrderRequest) (marketplacetypes.QueryGetOrderResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return marketplacetypes.QueryGetOrderResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Order(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return marketplacetypes.QueryGetOrderResponse{}, fmt.Errorf("error while getting order: %s", err)
	}

	return *res, nil
}

func (s Source) GetAllOrder(height int64, req marketplacetypes.QueryAllOrderRequest) (marketplacetypes.QueryAllOrderResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return marketplacetypes.QueryAllOrderResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.OrderAll(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return marketplacetypes.QueryAllOrderResponse{}, fmt.Errorf("error while getting all orders: %s", err)
	}

	return *res, nil
}
