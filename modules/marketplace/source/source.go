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

import marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"

type Source interface {
	GetParams(height int64) (marketplacetypes.Params, error)
	GetListing(height int64, req marketplacetypes.QueryGetListingRequest) (marketplacetypes.QueryGetListingResponse, error)
	GetAllListing(height int64, req marketplacetypes.QueryAllListingRequest) (marketplacetypes.QueryAllListingResponse, error)
	GetOrder(height int64, req marketplacetypes.QueryGetOrderRequest) (marketplacetypes.QueryGetOrderResponse, error)
	GetAllOrder(height int64, req marketplacetypes.QueryAllOrderRequest) (marketplacetypes.QueryAllOrderResponse, error)
}
