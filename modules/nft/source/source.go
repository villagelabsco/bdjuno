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

package source

import (
	"github.com/cosmos/cosmos-sdk/x/nft"
)

type Source interface {
	Balance(height int64, req nft.QueryBalanceRequest) (nft.QueryBalanceResponse, error)
	Owner(height int64, req nft.QueryOwnerRequest) (nft.QueryOwnerResponse, error)
	Supply(height int64, req nft.QuerySupplyRequest) (nft.QuerySupplyResponse, error)
	NFTs(height int64, req nft.QueryNFTsRequest) (nft.QueryNFTsResponse, error)
	NFT(height int64, req nft.QueryNFTRequest) (nft.QueryNFTResponse, error)
	Class(height int64, req nft.QueryClassRequest) (nft.QueryClassResponse, error)
	Classes(height int64, req nft.QueryClassesRequest) (nft.QueryClassesResponse, error)
}
