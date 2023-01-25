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
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/forbole/juno/v3/node/local"
)

type Source struct {
	*local.Source
	q nfttypes.QueryServer
}

func NewSource(source *local.Source, q nfttypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) Balance(height int64, req nfttypes.QueryBalanceRequest) (nfttypes.QueryBalanceResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryBalanceResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Balance(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryBalanceResponse{}, fmt.Errorf("error while getting nft balance: %s", err)
	}

	return *res, nil
}

func (s Source) Owner(height int64, req nfttypes.QueryOwnerRequest) (nfttypes.QueryOwnerResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryOwnerResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Owner(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryOwnerResponse{}, fmt.Errorf("error while getting nft owner: %s", err)
	}

	return *res, nil
}

func (s Source) Supply(height int64, req nfttypes.QuerySupplyRequest) (nfttypes.QuerySupplyResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QuerySupplyResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Supply(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QuerySupplyResponse{}, fmt.Errorf("error while getting nft supply: %s", err)
	}

	return *res, nil
}

func (s Source) NFTs(height int64, req nfttypes.QueryNFTsRequest) (nfttypes.QueryNFTsResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryNFTsResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.NFTs(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryNFTsResponse{}, fmt.Errorf("error while getting nfts: %s", err)
	}

	return *res, nil
}

func (s Source) NFT(height int64, req nfttypes.QueryNFTRequest) (nfttypes.QueryNFTResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryNFTResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.NFT(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryNFTResponse{}, fmt.Errorf("error while getting nft: %s", err)
	}

	return *res, nil
}

func (s Source) Class(height int64, req nfttypes.QueryClassRequest) (nfttypes.QueryClassResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryClassResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Class(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryClassResponse{}, fmt.Errorf("error while getting class: %s", err)
	}

	return *res, nil
}

func (s Source) Classes(height int64, req nfttypes.QueryClassesRequest) (nfttypes.QueryClassesResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nfttypes.QueryClassesResponse{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Classes(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nfttypes.QueryClassesResponse{}, fmt.Errorf("error while getting classes: %s", err)
	}

	return *res, nil
}
