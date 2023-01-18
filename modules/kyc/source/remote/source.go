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
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
)

type Source struct {
	*remote.Source
	q kyctypes.QueryClient
}

func NewSource(source *remote.Source, querier kyctypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      querier,
	}
}

func (s Source) GetParams(height int64) (kyctypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &kyctypes.QueryParamsRequest{})
	if err != nil {
		return kyctypes.Params{}, fmt.Errorf("error while getting kyc params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetKycStatus(height int64, req kyctypes.QueryGetKycStatusRequest) (kyctypes.QueryGetKycStatusResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.KycStatus(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetKycStatusResponse{}, fmt.Errorf("error while getting kyc status: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllKycStatus(height int64, req kyctypes.QueryAllKycStatusRequest) (kyctypes.QueryAllKycStatusResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.KycStatusAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllKycStatusResponse{}, fmt.Errorf("error while getting all kyc status: %s", err)
	}
	return *res, nil
}

func (s Source) GetDetailedAccount(height int64, req kyctypes.QueryGetDetailedAccountRequest) (kyctypes.QueryGetDetailedAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.DetailedAccount(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetDetailedAccountResponse{}, fmt.Errorf("error while getting detailed account: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllDetailedAccount(height int64, req kyctypes.QueryAllDetailedAccountsRequest) (kyctypes.QueryAllDetailedAccountsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.DetailedAccountsAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllDetailedAccountsResponse{}, fmt.Errorf("error while getting all detailed accounts: %s", err)
	}
	return *res, nil
}

func (s Source) GetInvite(height int64, req kyctypes.QueryGetInviteRequest) (kyctypes.QueryGetInviteResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Invite(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetInviteResponse{}, fmt.Errorf("error while getting invite: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllInvite(height int64, req kyctypes.QueryAllInviteRequest) (kyctypes.QueryAllInviteResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.InviteAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllInviteResponse{}, fmt.Errorf("error while getting all invites: %s", err)
	}
	return *res, nil
}

func (s Source) GetNbInvitePerDay(height int64, req kyctypes.QueryGetNbInvitePerDayRequest) (kyctypes.QueryGetNbInvitePerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbInvitePerDay(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetNbInvitePerDayResponse{}, fmt.Errorf("error while getting nb invite per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNbInvitePerDay(height int64, req kyctypes.QueryAllNbInvitePerDayRequest) (kyctypes.QueryAllNbInvitePerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbInvitePerDayAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllNbInvitePerDayResponse{}, fmt.Errorf("error while getting all nb invite per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetNetworkKyb(height int64, req kyctypes.QueryGetNetworkKybRequest) (kyctypes.QueryGetNetworkKybResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkKyb(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetNetworkKybResponse{}, fmt.Errorf("error while getting network kyb: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNetworkKyb(height int64, req kyctypes.QueryAllNetworkKybRequest) (kyctypes.QueryAllNetworkKybResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkKybAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllNetworkKybResponse{}, fmt.Errorf("error while getting all network kyb: %s", err)
	}
	return *res, nil
}

func (s Source) GetHuman(height int64, req kyctypes.QueryGetHumanRequest) (kyctypes.QueryGetHumanResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Human(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetHumanResponse{}, fmt.Errorf("error while getting human: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllHuman(height int64, req kyctypes.QueryAllHumanRequest) (kyctypes.QueryAllHumanResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.HumanAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllHumanResponse{}, fmt.Errorf("error while getting all humans: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccount(height int64, req kyctypes.QueryGetAccountRequest) (kyctypes.QueryGetAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Account(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetAccountResponse{}, fmt.Errorf("error while getting account: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccount(height int64, req kyctypes.QueryAllAccountRequest) (kyctypes.QueryAllAccountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllAccountResponse{}, fmt.Errorf("error while getting all accounts: %s", err)
	}
	return *res, nil
}

func (s Source) GetIdentityProvider(height int64, req kyctypes.QueryGetIdentityProviderRequest) (kyctypes.QueryGetIdentityProviderResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.IdentityProvider(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetIdentityProviderResponse{}, fmt.Errorf("error while getting identity provider: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllIdentityProvider(height int64, req kyctypes.QueryAllIdentityProviderRequest) (kyctypes.QueryAllIdentityProviderResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.IdentityProviderAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllIdentityProviderResponse{}, fmt.Errorf("error while getting all identity providers: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccountLinkProposal(height int64, req kyctypes.QueryGetAccountLinkProposalRequest) (kyctypes.QueryGetAccountLinkProposalResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposal(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetAccountLinkProposalResponse{}, fmt.Errorf("error while getting account link proposal: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccountLinkProposal(height int64, req kyctypes.QueryAllAccountLinkProposalRequest) (kyctypes.QueryAllAccountLinkProposalResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllAccountLinkProposalResponse{}, fmt.Errorf("error while getting all account link proposals: %s", err)
	}
	return *res, nil
}

func (s Source) GetAccountLinkProposalsForHumanId(height int64, req kyctypes.QueryGetAccountLinkProposalsForHumanIdRequest) (kyctypes.QueryGetAccountLinkProposalsForHumanIdResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalsForHumanId(ctx, &req)
	if err != nil {
		return kyctypes.QueryGetAccountLinkProposalsForHumanIdResponse{}, fmt.Errorf("error while getting account link proposals for human id: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllAccountLinkProposalsForHumanId(height int64, req kyctypes.QueryAllAccountLinkProposalsForHumanIdRequest) (kyctypes.QueryAllAccountLinkProposalsForHumanIdResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.AccountLinkProposalsForHumanIdAll(ctx, &req)
	if err != nil {
		return kyctypes.QueryAllAccountLinkProposalsForHumanIdResponse{}, fmt.Errorf("error while getting all account link proposals for human id: %s", err)
	}
	return *res, nil
}
