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
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/forbole/juno/v3/node/remote"
)

type Source struct {
	*remote.Source
	q nfttypes.QueryClient
}

func NewSource(source *remote.Source, q nfttypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) Balance(height int64, req nfttypes.QueryBalanceRequest) (nfttypes.QueryBalanceResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Balance(ctx, &req)
	if err != nil {
		return nfttypes.QueryBalanceResponse{}, fmt.Errorf("error while getting nft balance: %s", err)
	}
	return *res, nil
}

func (s Source) Owner(height int64, req nfttypes.QueryOwnerRequest) (nfttypes.QueryOwnerResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Owner(ctx, &req)
	if err != nil {
		return nfttypes.QueryOwnerResponse{}, fmt.Errorf("error while getting nft owner: %s", err)
	}
	return *res, nil
}

func (s Source) Supply(height int64, req nfttypes.QuerySupplyRequest) (nfttypes.QuerySupplyResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Supply(ctx, &req)
	if err != nil {
		return nfttypes.QuerySupplyResponse{}, fmt.Errorf("error while getting nft supply: %s", err)
	}
	return *res, nil
}

func (s Source) NFTs(height int64, req nfttypes.QueryNFTsRequest) (nfttypes.QueryNFTsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NFTs(ctx, &req)
	if err != nil {
		return nfttypes.QueryNFTsResponse{}, fmt.Errorf("error while getting nfts: %s", err)
	}
	return *res, nil
}

func (s Source) NFT(height int64, req nfttypes.QueryNFTRequest) (nfttypes.QueryNFTResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NFT(ctx, &req)
	if err != nil {
		return nfttypes.QueryNFTResponse{}, fmt.Errorf("error while getting nft: %s", err)
	}
	return *res, nil
}

func (s Source) Class(height int64, req nfttypes.QueryClassRequest) (nfttypes.QueryClassResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Class(ctx, &req)
	if err != nil {
		return nfttypes.QueryClassResponse{}, fmt.Errorf("error while getting nft class: %s", err)
	}
	return *res, nil
}

func (s Source) Classes(height int64, req nfttypes.QueryClassesRequest) (nfttypes.QueryClassesResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Classes(ctx, &req)
	if err != nil {
		return nfttypes.QueryClassesResponse{}, fmt.Errorf("error while getting nft classes: %s", err)
	}
	return *res, nil
}
