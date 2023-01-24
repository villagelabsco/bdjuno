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
	"github.com/forbole/juno/v3/node/remote"
	identitytypes "github.com/villagelabs/villaged/x/identity/types"
)

type Source struct {
	*remote.Source
	q identitytypes.QueryClient
}

func (s Source) GetParams(height int64) (identitytypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &identitytypes.QueryParamsRequest{})
	if err != nil {
		return identitytypes.Params{}, fmt.Errorf("error while getting kyc params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetKycStatus(height int64, req identitytypes.QueryGetKycStatusRequest) (identitytypes.QueryGetKycStatusResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.KycStatus(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetKycStatusResponse{}, fmt.Errorf("error while getting kyc status: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllKycStatus(height int64, req identitytypes.QueryAllKycStatusRequest) (identitytypes.QueryAllKycStatusResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.KycStatusAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllKycStatusResponse{}, fmt.Errorf("error while getting all kyc status: %s", err)
	}
	return *res, nil
}

func (s Source) GetDetailedAccount(height int64, req identitytypes.QueryGetDetailedAccountRequest) (identitytypes.QueryGetDetailedAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.DetailedAccount(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetDetailedAccountResponse{}, fmt.Errorf("error while getting detailed account: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllDetailedAccount(height int64, req identitytypes.QueryAllDetailedAccountsRequest) (identitytypes.QueryAllDetailedAccountsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.DetailedAccountsAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllDetailedAccountsResponse{}, fmt.Errorf("error while getting all detailed accounts: %s", err)
	}
	return *res, nil
}

func (s Source) GetInvite(height int64, req identitytypes.QueryGetInviteRequest) (identitytypes.QueryGetInviteResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Invite(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetInviteResponse{}, fmt.Errorf("error while getting invite: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllInvite(height int64, req identitytypes.QueryAllInviteRequest) (identitytypes.QueryAllInviteResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.InviteAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllInviteResponse{}, fmt.Errorf("error while getting all invites: %s", err)
	}
	return *res, nil
}

func (s Source) GetNbInvitePerDay(height int64, req identitytypes.QueryGetNbInvitePerDayRequest) (identitytypes.QueryGetNbInvitePerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbInvitePerDay(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetNbInvitePerDayResponse{}, fmt.Errorf("error while getting nb invite per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNbInvitePerDay(height int64, req identitytypes.QueryAllNbInvitePerDayRequest) (identitytypes.QueryAllNbInvitePerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbInvitePerDayAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllNbInvitePerDayResponse{}, fmt.Errorf("error while getting all nb invite per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetNetworkKyb(height int64, req identitytypes.QueryGetNetworkKybRequest) (identitytypes.QueryGetNetworkKybResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkKyb(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetNetworkKybResponse{}, fmt.Errorf("error while getting network kyb: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNetworkKyb(height int64, req identitytypes.QueryAllNetworkKybRequest) (identitytypes.QueryAllNetworkKybResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkKybAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllNetworkKybResponse{}, fmt.Errorf("error while getting all network kyb: %s", err)
	}
	return *res, nil
}

func (s Source) GetHuman(height int64, req identitytypes.QueryGetHumanRequest) (identitytypes.QueryGetHumanResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Human(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetHumanResponse{}, fmt.Errorf("error while getting human: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllHuman(height int64, req identitytypes.QueryAllHumanRequest) (identitytypes.QueryAllHumanResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.HumanAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllHumanResponse{}, fmt.Errorf("error while getting all humans: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccount(height int64, req identitytypes.QueryGetAccountRequest) (identitytypes.QueryGetAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Account(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetAccountResponse{}, fmt.Errorf("error while getting account: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccount(height int64, req identitytypes.QueryAllAccountRequest) (identitytypes.QueryAllAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllAccountResponse{}, fmt.Errorf("error while getting all accounts: %s", err)
	}
	return *res, nil
}

func (s Source) GetIdentityProvider(height int64, req identitytypes.QueryGetIdentityProviderRequest) (identitytypes.QueryGetIdentityProviderResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.IdentityProvider(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetIdentityProviderResponse{}, fmt.Errorf("error while getting identity provider: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllIdentityProvider(height int64, req identitytypes.QueryAllIdentityProviderRequest) (identitytypes.QueryAllIdentityProviderResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.IdentityProviderAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllIdentityProviderResponse{}, fmt.Errorf("error while getting all identity providers: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccountLinkProposal(height int64, req identitytypes.QueryGetAccountLinkProposalRequest) (identitytypes.QueryGetAccountLinkProposalResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposal(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetAccountLinkProposalResponse{}, fmt.Errorf("error while getting account link proposal: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccountLinkProposal(height int64, req identitytypes.QueryAllAccountLinkProposalRequest) (identitytypes.QueryAllAccountLinkProposalResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllAccountLinkProposalResponse{}, fmt.Errorf("error while getting all account link proposals: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccountLinkProposalsForHumanId(height int64, req identitytypes.QueryGetAccountLinkProposalsForHumanIdRequest) (identitytypes.QueryGetAccountLinkProposalsForHumanIdResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalsForHumanId(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetAccountLinkProposalsForHumanIdResponse{}, fmt.Errorf("error while getting account link proposals for human id: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccountLinkProposalsForHumanId(height int64, req identitytypes.QueryAllAccountLinkProposalsForHumanIdRequest) (identitytypes.QueryAllAccountLinkProposalsForHumanIdResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalsForHumanIdAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllAccountLinkProposalsForHumanIdResponse{}, fmt.Errorf("error while getting all account link proposals for human id: %s", err)
	}
	return *res, nil
}

func (s Source) GetNetwork(height int64, req identitytypes.QueryGetNetworkRequest) (identitytypes.QueryGetNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Network(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetNetworkResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllNetwork(height int64, req identitytypes.QueryAllNetworkRequest) (identitytypes.QueryAllNetworkResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllNetworkResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAccountNetworks(height int64, req identitytypes.QueryGetAccountNetworksRequest) (identitytypes.QueryGetAccountNetworksResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountNetworks(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetAccountNetworksResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllAccountNetworks(height int64, req identitytypes.QueryAllAccountNetworksRequest) (identitytypes.QueryAllAccountNetworksResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountNetworksAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllAccountNetworksResponse{}, err
	}
	return *res, nil
}

func (s Source) GetNetworkAccounts(height int64, req identitytypes.QueryGetNetworkAccountsRequest) (identitytypes.QueryGetNetworkAccountsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkAccounts(ctx, &req)
	if err != nil {
		return identitytypes.QueryGetNetworkAccountsResponse{}, err
	}
	return *res, nil
}

func (s Source) GetAllNetworkAccounts(height int64, req identitytypes.QueryAllNetworkAccountsRequest) (identitytypes.QueryAllNetworkAccountsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkAccountsAll(ctx, &req)
	if err != nil {
		return identitytypes.QueryAllNetworkAccountsResponse{}, err
	}
	return *res, nil
}
