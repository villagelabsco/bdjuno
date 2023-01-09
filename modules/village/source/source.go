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

import villagetypes "github.com/villagelabs/villaged/x/village/types"

type Source interface {
	GetParams(height int64) (villagetypes.Params, error)
	GetNetwork(height int64, req villagetypes.QueryGetNetworkRequest) (villagetypes.QueryGetNetworkResponse, error)
	GetAllNetwork(height int64, req villagetypes.QueryAllNetworkRequest) (villagetypes.QueryAllNetworkResponse, error)
	GetUserNetworks(height int64, req villagetypes.QueryGetUserNetworksRequest) (villagetypes.QueryGetUserNetworksResponse, error)
	GetAllUserNetworks(height int64, req villagetypes.QueryAllUserNetworksRequest) (villagetypes.QueryAllUserNetworksResponse, error)
	GetNbNetworkCreationPerDay(height int64, req villagetypes.QueryGetNbNetworkCreationPerDayRequest) (villagetypes.QueryGetNbNetworkCreationPerDayResponse, error)
	GetAllNbNetworkCreationPerDay(height int64, req villagetypes.QueryAllNbNetworkCreationPerDayRequest) (villagetypes.QueryAllNbNetworkCreationPerDayResponse, error)
	GetAccountsInNetwork(height int64, req villagetypes.QueryGetAccountsInNetworkRequest) (villagetypes.QueryGetAccountsInNetworkResponse, error)
	GetAllAccountsInNetwork(height int64, req villagetypes.QueryAllAccountsInNetworkRequest) (villagetypes.QueryAllAccountsInNetworkResponse, error)
}
