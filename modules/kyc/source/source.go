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

import kyctypes "github.com/villagelabs/villaged/x/kyc/types"

type Source interface {
	GetParams(height int64) (kyctypes.Params, error)
	GetKycStatus(height int64, req kyctypes.QueryGetKycStatusRequest) (kyctypes.QueryGetKycStatusResponse, error)
	GetAllKycStatus(height int64, req kyctypes.QueryAllKycStatusRequest) (kyctypes.QueryAllKycStatusResponse, error)
	GetDetailedAccount(height int64, req kyctypes.QueryGetDetailedAccountRequest) (kyctypes.QueryGetDetailedAccountResponse, error)
	GetAllDetailedAccount(height int64, req kyctypes.QueryAllDetailedAccountsRequest) (kyctypes.QueryAllDetailedAccountsResponse, error)
	GetInvite(height int64, req kyctypes.QueryGetInviteRequest) (kyctypes.QueryGetInviteResponse, error)
	GetAllInvite(height int64, req kyctypes.QueryAllInviteRequest) (kyctypes.QueryAllInviteResponse, error)
	GetNbInvitePerDay(height int64, req kyctypes.QueryGetNbInvitePerDayRequest) (kyctypes.QueryGetNbInvitePerDayResponse, error)
	GetAllNbInvitePerDay(height int64, req kyctypes.QueryAllNbInvitePerDayRequest) (kyctypes.QueryAllNbInvitePerDayResponse, error)
	GetNetworkKyb(height int64, req kyctypes.QueryGetNetworkKybRequest) (kyctypes.QueryGetNetworkKybResponse, error)
	GetAllNetworkKyb(height int64, req kyctypes.QueryAllNetworkKybRequest) (kyctypes.QueryAllNetworkKybResponse, error)
	GetHuman(height int64, req kyctypes.QueryGetHumanRequest) (kyctypes.QueryGetHumanResponse, error)
	GetAllHuman(height int64, req kyctypes.QueryAllHumanRequest) (kyctypes.QueryAllHumanResponse, error)
	GetAccount(height int64, req kyctypes.QueryGetAccountRequest) (kyctypes.QueryGetAccountResponse, error)
	GetAllAccount(height int64, req kyctypes.QueryAllAccountRequest) (kyctypes.QueryAllAccountResponse, error)
	GetIdentityProvider(height int64, req kyctypes.QueryGetIdentityProviderRequest) (kyctypes.QueryGetIdentityProviderResponse, error)
	GetAllIdentityProvider(height int64, req kyctypes.QueryAllIdentityProviderRequest) (kyctypes.QueryAllIdentityProviderResponse, error)
	GetAccountLinkProposal(height int64, req kyctypes.QueryGetAccountLinkProposalRequest) (kyctypes.QueryGetAccountLinkProposalResponse, error)
	GetAllAccountLinkProposal(height int64, req kyctypes.QueryAllAccountLinkProposalRequest) (kyctypes.QueryAllAccountLinkProposalResponse, error)
	GetAccountLinkProposalsForHumanId(height int64, req kyctypes.QueryGetAccountLinkProposalsForHumanIdRequest) (kyctypes.QueryGetAccountLinkProposalsForHumanIdResponse, error)
	GetAllAccountLinkProposalsForHumanId(height int64, req kyctypes.QueryAllAccountLinkProposalsForHumanIdRequest) (kyctypes.QueryAllAccountLinkProposalsForHumanIdResponse, error)
}
