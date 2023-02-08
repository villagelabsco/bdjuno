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

package identity

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/villagelabsco/bdjuno/v3/types"
	"github.com/villagelabsco/bdjuno/v3/utils"
	juno "github.com/villagelabsco/juno/v4/types"
	identitytypes "github.com/villagelabsco/villaged/x/identity/types"
	rbactypes "github.com/villagelabsco/villaged/x/rbac/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *identitytypes.MsgVerifyAccount:
		return m.handleMsgVerifyAccount(index, tx, cosmosMsg)
	case *identitytypes.MsgRevokeAccount:
		return m.handleMsgRevokeAccount(index, tx, cosmosMsg)
	case *identitytypes.MsgCreateInvite:
		return m.handleMsgCreateInvite(index, tx, cosmosMsg)
	case *identitytypes.MsgClaimInvite:
		return m.handleMsgClaimInvite(index, tx, cosmosMsg)
	case *identitytypes.MsgRescindInvite:
		return m.handleMsgRescindInvite(index, tx, cosmosMsg)
	case *identitytypes.MsgCreateMultipleInvites:
		return m.handleMsgCreateMultipleInvites(index, tx, cosmosMsg)
	case *identitytypes.MsgVerifyNetwork:
		return m.handleMsgVerifyNetwork(index, tx, cosmosMsg)
	case *identitytypes.MsgRevokeNetwork:
		return m.handleMsgRevokeNetwork(index, tx, cosmosMsg)
	case *identitytypes.MsgCreateHumanId:
		return m.handleMsgCreateHumanId(index, tx, cosmosMsg)
	case *identitytypes.MsgRegisterIdentityProvider:
		return m.handleMsgRegisterIdentityProvider(index, tx, cosmosMsg)
	case *identitytypes.MsgSetIdentityProviderAdminAccounts:
		return m.handleMsgSetIdentityProviderAdminAccounts(index, tx, cosmosMsg)
	case *identitytypes.MsgSetIdentityProviderProviderAccounts:
		return m.handleMsgSetIdentityProviderProviderAccounts(index, tx, cosmosMsg)
	case *identitytypes.MsgJoinNetwork:
		return m.handleMsgJoinNetwork(index, tx, cosmosMsg)
	case *identitytypes.MsgCreateMultipleHumanIds:
		return m.handleMsgCreateMultipleHumanIds(index, tx, cosmosMsg)
	case *identitytypes.MsgSetPrimaryNetworkWallet:
		return m.handleMsgSetPrimaryNetworkWallet(index, tx, cosmosMsg)
	case *identitytypes.MsgAcceptLinkWalletToHumanProposal:
		return m.handleMsgAcceptLinkWalletToHumanProposal(index, tx, cosmosMsg)
	case *identitytypes.MsgProposeLinkAccountToHuman:
		return m.handleMsgProposeLinkAccountToHuman(index, tx, cosmosMsg)
	case *identitytypes.MsgCreateNetwork:
		return m.handleMsgCreateNetwork(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) handleMsgVerifyAccount(index int, tx *juno.Tx, msg *identitytypes.MsgVerifyAccount) error {
	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtVerifiedAccount{}, "humanId")
	if err != nil {
		return fmt.Errorf("error finding humanId from event: %s", err)
	}

	st, err := m.src.GetKycStatus(tx.Height, identitytypes.QueryGetKycStatusRequest{
		IdProvider: msg.IdProvider,
		HumanId:    humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting kyc status: %s", err)
	}
	status := st.KycStatus

	if err := m.db.SaveOrUpdateKycStatus(msg.IdProvider, &status); err != nil {
		return fmt.Errorf("error saving kyc status: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRevokeAccount(index int, tx *juno.Tx, msg *identitytypes.MsgRevokeAccount) error {
	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtRevokedAccount{}, "humanId")
	if err != nil {
		return fmt.Errorf("error finding humanId from event: %s", err)
	}

	st, err := m.src.GetKycStatus(tx.Height, identitytypes.QueryGetKycStatusRequest{
		IdProvider: msg.IdProvider,
		HumanId:    humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting kyc status: %s", err)
	}
	status := st.KycStatus

	if err := m.db.SaveOrUpdateKycStatus(msg.IdProvider, &status); err != nil {
		return fmt.Errorf("error saving kyc status: %s", err)
	}

	return nil
}

func (m *Module) handleMsgCreateInvite(index int, tx *juno.Tx, msg *identitytypes.MsgCreateInvite) error {
	return m.db.SaveIdentityInvite(msg.Network, &identitytypes.Invite{
		Challenge:        msg.Challenge,
		Registered:       false,
		ConfirmedAccount: "",
		InviteCreator:    msg.Creator,
		HumanId:          msg.HumanId,
		GivenRoles:       msg.GivenRoles,
	})
}

func (m *Module) handleMsgClaimInvite(index int, tx *juno.Tx, msg *identitytypes.MsgClaimInvite) error {
	if err := m.db.UpdateIdentityInvite(msg.Network, msg.Challenge, msg.Creator, true); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}

	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtClaimedInvite{}, "humanId")
	if err != nil {
		return fmt.Errorf("error finding humanId from event: %s", err)
	}

	acc := &identitytypes.Account{
		Index:      msg.Creator,
		HumanId:    humanId,
		PrivateAcc: false,
	}
	if err := m.db.SaveOrUpdateIdentityAccount(acc); err != nil {
		return fmt.Errorf("error updating user networks: %s", err)
	}

	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}

	if err := m.db.SaveOrUpdateIdentityHuman(&h.Human); err != nil {
		return fmt.Errorf("error updating human: %s", err)
	}

	if err := m.db.SaveOrAppendIdentityAccountNetworks(msg.Creator, msg.Network); err != nil {
		return fmt.Errorf("error updating user networks: %s", err)
	}

	alwnc, err := m.feeGrantSrc.GetAllowances(tx.Height, feegranttypes.QueryAllowancesRequest{
		Grantee: msg.Creator,
		Pagination: &query.PageRequest{
			Limit: 1,
		},
	})
	if err != nil {
		return fmt.Errorf("error getting allowances: %s", err)
	}
	allowances := alwnc.Allowances

	if err := m.db.SaveFeeGrantAllowance(types.NewFeeGrant(*allowances[0], tx.Height)); err != nil {
		return fmt.Errorf("error saving fee grant allowance: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRescindInvite(index int, tx *juno.Tx, msg *identitytypes.MsgRescindInvite) error {
	if err := m.db.DeleteIdentityInvite(msg.Network, msg.Challenge); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}
	return nil
}

func (m *Module) handleMsgCreateMultipleInvites(index int, tx *juno.Tx, msg *identitytypes.MsgCreateMultipleInvites) error {
	inv := make([]*identitytypes.Invite, len(msg.Challenges))
	for i, c := range msg.Challenges {
		inv[i] = &identitytypes.Invite{
			Challenge:        c,
			Registered:       false,
			ConfirmedAccount: "",
			InviteCreator:    msg.Creator,
			HumanId:          msg.HumanIds[i],
			GivenRoles:       msg.GivenRoles,
		}
	}

	err := m.db.SaveMultipleIdentityInvites(msg.Network, inv)
	if err != nil {
		return fmt.Errorf("error inserting multiple invites: %s", err)
	}

	return nil
}

func (m *Module) handleMsgVerifyNetwork(index int, tx *juno.Tx, msg *identitytypes.MsgVerifyNetwork) error {
	acc, err := m.src.GetNetworkKyb(tx.Height, identitytypes.QueryGetNetworkKybRequest{
		Index: msg.Network,
	})
	if err != nil {
		return fmt.Errorf("error getting network kyb: %s", err)
	}
	kyb := acc.NetworkKyb

	if err := m.db.SaveOrUpdateIdentityNetworkKyb(&kyb); err != nil {
		return fmt.Errorf("error saving network kyb: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRevokeNetwork(index int, tx *juno.Tx, msg *identitytypes.MsgRevokeNetwork) error {
	acc, err := m.src.GetNetworkKyb(tx.Height, identitytypes.QueryGetNetworkKybRequest{
		Index: msg.Network,
	})
	if err != nil {
		return fmt.Errorf("error getting network kyb: %s", err)
	}
	kyb := acc.NetworkKyb

	if err := m.db.SaveOrUpdateIdentityNetworkKyb(&kyb); err != nil {
		return fmt.Errorf("error saving network kyb: %s", err)
	}

	return nil
}

func (m *Module) handleMsgCreateHumanId(index int, tx *juno.Tx, msg *identitytypes.MsgCreateHumanId) error {
	idx, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtCreatedHumanId{}, "humanId")
	if err != nil {
		return fmt.Errorf("error getting human id from created event: %s", err)
	}
	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: idx,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateIdentityHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRegisterIdentityProvider(index int, tx *juno.Tx, msg *identitytypes.MsgRegisterIdentityProvider) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, identitytypes.QueryGetIdentityProviderRequest{
		Index: msg.Name,
	})
	if err != nil {
		return fmt.Errorf("error getting identity provider: %s", err)
	}
	provider := ip.IdentityProvider

	if err := m.db.SaveOrUpdateIdentityProvider(provider); err != nil {
		return fmt.Errorf("error saving identity provider: %s", err)
	}

	return nil
}

func (m *Module) handleMsgSetIdentityProviderAdminAccounts(index int, tx *juno.Tx, msg *identitytypes.MsgSetIdentityProviderAdminAccounts) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, identitytypes.QueryGetIdentityProviderRequest{
		Index: msg.Name,
	})
	if err != nil {
		return fmt.Errorf("error getting identity provider: %s", err)
	}
	provider := ip.IdentityProvider

	if err := m.db.SaveOrUpdateIdentityProvider(provider); err != nil {
		return fmt.Errorf("error updating identity provider: %s", err)
	}

	return nil
}

func (m *Module) handleMsgSetIdentityProviderProviderAccounts(index int, tx *juno.Tx, msg *identitytypes.MsgSetIdentityProviderProviderAccounts) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, identitytypes.QueryGetIdentityProviderRequest{
		Index: msg.Name,
	})
	if err != nil {
		return fmt.Errorf("error getting identity provider: %s", err)
	}
	provider := ip.IdentityProvider

	if err := m.db.SaveOrUpdateIdentityProvider(provider); err != nil {
		return fmt.Errorf("error updating identity provider: %s", err)
	}

	return nil
}

func (m *Module) handleMsgJoinNetwork(index int, tx *juno.Tx, msg *identitytypes.MsgJoinNetwork) error {
	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtJoinNetwork{}, "humanId")
	if err != nil {
		return fmt.Errorf("error getting human id from join network event: %s", err)
	}

	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateIdentityHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	ip, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtJoinNetwork{}, "identityProvider")
	if err != nil {
		return fmt.Errorf("error getting identity provider from join network event: %s", err)
	}

	st, err := m.src.GetKycStatus(tx.Height, identitytypes.QueryGetKycStatusRequest{
		IdProvider: ip,
		HumanId:    humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting kyc status: %s", err)
	}
	status := st.KycStatus

	if err := m.db.SaveOrUpdateKycStatus(ip, &status); err != nil {
		return fmt.Errorf("error saving kyc status: %s", err)
	}

	return nil
}

func (m *Module) handleMsgCreateMultipleHumanIds(index int, tx *juno.Tx, msg *identitytypes.MsgCreateMultipleHumanIds) error {
	// TODO: How do we get multiple events being sent out?
	// Or we need a different event being sent out?
	return nil
}

func (m *Module) handleMsgSetPrimaryNetworkWallet(index int, tx *juno.Tx, msg *identitytypes.MsgSetPrimaryNetworkWallet) error {
	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtSetPrimaryNetworkWallet{}, "humanId")
	if err != nil {
		return fmt.Errorf("error getting human id from set primary network wallet event: %s", err)
	}

	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateIdentityHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	return nil
}

func (m *Module) handleMsgCreateNetwork(index int, tx *juno.Tx, msg *identitytypes.MsgCreateNetwork) error {
	if err := m.db.SaveIdentityNetwork(&identitytypes.Network{
		Index:            msg.ShortName,
		Active:           true,
		FullName:         msg.FullName,
		IdentityProvider: msg.IdentityProvider,
		InviteOnly:       msg.InviteOnly,
	}); err != nil {
		return fmt.Errorf("error saving network: %s", err)
	}

	ownerAuthorization := msg.ShortName + identitytypes.NamespaceSeparator + identitytypes.InitAdminGroupName
	auth, err := m.rbacSrc.GetAuthorizations(tx.Height, rbactypes.QueryGetAuthorizationsRequest{
		Index: ownerAuthorization,
	})
	if err != nil {
		return fmt.Errorf("error getting authorizations: %s", err)
	}
	authorizations := auth.Authorizations

	if err := m.db.SaveOrUpdateRbacAuthorization(&authorizations); err != nil {
		return fmt.Errorf("error saving authorizations: %s", err)
	}

	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtJoinNetwork{}, "humanId")
	if err != nil {
		return fmt.Errorf("error getting human id from join network event: %s", err)
	}

	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateIdentityHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	acc := identitytypes.Account{
		Index:      msg.Creator,
		HumanId:    humanId,
		PrivateAcc: false,
	}
	if err := m.db.SaveOrUpdateIdentityAccount(&acc); err != nil {
		return fmt.Errorf("error saving account: %s", err)
	}

	st, err := m.src.GetKycStatus(tx.Height, identitytypes.QueryGetKycStatusRequest{
		IdProvider: msg.IdentityProvider,
		HumanId:    humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting kyc status: %s", err)
	}
	status := st.KycStatus

	if err := m.db.SaveOrUpdateKycStatus(msg.IdentityProvider, &status); err != nil {
		return fmt.Errorf("error saving kyc status: %s", err)
	}

	if err := m.db.SaveOrAppendIdentityAccountNetworks(msg.Creator, msg.ShortName); err != nil {
		return fmt.Errorf("error saving account networks: %s", err)
	}

	return nil
}

func (m *Module) handleMsgAcceptLinkWalletToHumanProposal(index int, tx *juno.Tx, msg *identitytypes.MsgAcceptLinkWalletToHumanProposal) error {
	if err := m.db.DeleteIdentityAccountLinkProposal(msg.ProposalId); err != nil {
		return fmt.Errorf("error deleting account link proposal: %s", err)
	}

	proposerAcc, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtLinkAccountToHumanApproved{}, "proposerAccount")
	if err != nil {
		return fmt.Errorf("error getting proposer account from link account to human approved event: %s", err)
	}

	humanId, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtLinkAccountToHumanApproved{}, "humanId")
	if err != nil {
		return fmt.Errorf("error getting human id from link account to human approved event: %s", err)
	}

	acc := identitytypes.Account{
		Index:      proposerAcc,
		HumanId:    humanId,
		PrivateAcc: false,
	}
	if err := m.db.SaveOrUpdateIdentityAccount(&acc); err != nil {
		return fmt.Errorf("error saving account: %s", err)
	}

	h, err := m.src.GetHuman(tx.Height, identitytypes.QueryGetHumanRequest{
		Index: humanId,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateIdentityHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	return nil
}

func (m *Module) handleMsgProposeLinkAccountToHuman(index int, tx *juno.Tx, msg *identitytypes.MsgProposeLinkAccountToHuman) error {
	propIdx, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtProposeLinkAccountToHuman{}, "proposalIdx")
	if err != nil {
		return fmt.Errorf("error getting proposal index from propose link account to human event: %s", err)
	}

	p, err := m.src.GetAccountLinkProposal(tx.Height, identitytypes.QueryGetAccountLinkProposalRequest{
		Index: propIdx,
	})
	if err != nil {
		return fmt.Errorf("error getting account link proposal: %s", err)
	}
	proposal := p.AccountLinkProposal

	if err := m.db.SaveOrUpdateIdentityAccountLinkProposal(&proposal); err != nil {
		return fmt.Errorf("error saving account link proposal: %s", err)
	}

	return nil
}
