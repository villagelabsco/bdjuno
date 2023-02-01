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

import identitytypes "github.com/villagelabsco/villaged/x/identity/types"

type Source interface {
	GetParams(height int64) (identitytypes.Params, error)
	GetKycStatus(height int64, req identitytypes.QueryGetKycStatusRequest) (identitytypes.QueryGetKycStatusResponse, error)
	GetAllKycStatus(height int64, req identitytypes.QueryAllKycStatusRequest) (identitytypes.QueryAllKycStatusResponse, error)
	GetDetailedAccount(height int64, req identitytypes.QueryGetDetailedAccountRequest) (identitytypes.QueryGetDetailedAccountResponse, error)
	GetAllDetailedAccount(height int64, req identitytypes.QueryAllDetailedAccountsRequest) (identitytypes.QueryAllDetailedAccountsResponse, error)
	GetInvite(height int64, req identitytypes.QueryGetInviteRequest) (identitytypes.QueryGetInviteResponse, error)
	GetAllInvite(height int64, req identitytypes.QueryAllInviteRequest) (identitytypes.QueryAllInviteResponse, error)
	GetNbInvitePerDay(height int64, req identitytypes.QueryGetNbInvitePerDayRequest) (identitytypes.QueryGetNbInvitePerDayResponse, error)
	GetAllNbInvitePerDay(height int64, req identitytypes.QueryAllNbInvitePerDayRequest) (identitytypes.QueryAllNbInvitePerDayResponse, error)
	GetNetworkKyb(height int64, req identitytypes.QueryGetNetworkKybRequest) (identitytypes.QueryGetNetworkKybResponse, error)
	GetAllNetworkKyb(height int64, req identitytypes.QueryAllNetworkKybRequest) (identitytypes.QueryAllNetworkKybResponse, error)
	GetHuman(height int64, req identitytypes.QueryGetHumanRequest) (identitytypes.QueryGetHumanResponse, error)
	GetAllHuman(height int64, req identitytypes.QueryAllHumanRequest) (identitytypes.QueryAllHumanResponse, error)
	GetAccount(height int64, req identitytypes.QueryGetAccountRequest) (identitytypes.QueryGetAccountResponse, error)
	GetAllAccount(height int64, req identitytypes.QueryAllAccountRequest) (identitytypes.QueryAllAccountResponse, error)
	GetIdentityProvider(height int64, req identitytypes.QueryGetIdentityProviderRequest) (identitytypes.QueryGetIdentityProviderResponse, error)
	GetAllIdentityProvider(height int64, req identitytypes.QueryAllIdentityProviderRequest) (identitytypes.QueryAllIdentityProviderResponse, error)
	GetAccountLinkProposal(height int64, req identitytypes.QueryGetAccountLinkProposalRequest) (identitytypes.QueryGetAccountLinkProposalResponse, error)
	GetAllAccountLinkProposal(height int64, req identitytypes.QueryAllAccountLinkProposalRequest) (identitytypes.QueryAllAccountLinkProposalResponse, error)
	GetAccountLinkProposalsForHumanId(height int64, req identitytypes.QueryGetAccountLinkProposalsForHumanIdRequest) (identitytypes.QueryGetAccountLinkProposalsForHumanIdResponse, error)
	GetAllAccountLinkProposalsForHumanId(height int64, req identitytypes.QueryAllAccountLinkProposalsForHumanIdRequest) (identitytypes.QueryAllAccountLinkProposalsForHumanIdResponse, error)
	GetNetwork(height int64, req identitytypes.QueryGetNetworkRequest) (identitytypes.QueryGetNetworkResponse, error)
	GetAllNetwork(height int64, req identitytypes.QueryAllNetworkRequest) (identitytypes.QueryAllNetworkResponse, error)
	GetAccountNetworks(height int64, req identitytypes.QueryGetAccountNetworksRequest) (identitytypes.QueryGetAccountNetworksResponse, error)
	GetAllAccountNetworks(height int64, req identitytypes.QueryAllAccountNetworksRequest) (identitytypes.QueryAllAccountNetworksResponse, error)
	GetNetworkAccounts(height int64, req identitytypes.QueryGetNetworkAccountsRequest) (identitytypes.QueryGetNetworkAccountsResponse, error)
	GetAllNetworkAccounts(height int64, req identitytypes.QueryAllNetworkAccountsRequest) (identitytypes.QueryAllNetworkAccountsResponse, error)
}
