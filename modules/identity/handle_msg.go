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
	"github.com/forbole/bdjuno/v3/utils"
	juno "github.com/forbole/juno/v3/types"
	identitytypes "github.com/villagelabs/villaged/x/identity/types"
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
	default:
		return fmt.Errorf("unrecognized kyc message type: %T", msg)
	}
}

func (m *Module) handleMsgVerifyAccount(index int, tx *juno.Tx, msg *identitytypes.MsgVerifyAccount) error {
	return nil
}

func (m *Module) handleMsgRevokeAccount(index int, tx *juno.Tx, msg *identitytypes.MsgRevokeAccount) error {
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
	if err := m.db.UpdateIdentityInvite(msg.Network, msg.Challenge, msg.Creator); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}

	// Update the user networks mapping
	un, err := m.db.IdentityAccountNetworks(msg.Creator)
	if err != nil {
		return fmt.Errorf("error getting user networks: %s", err)
	}

	if un.Index == "" {
		un.Index = msg.Creator
		un.Networks = make([]string, 1)
		un.Networks[0] = msg.Network
		if err := m.db.SaveIdentityAccountNetworks(un); err != nil {
			return fmt.Errorf("error inserting user networks: %s", err)
		}
	} else {
		un.Networks = append(un.Networks, msg.Network)
		if err := m.db.UpdateIdentityAccountNetworks(un); err != nil {
			return fmt.Errorf("error updating user networks: %s", err)
		}
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
	idx, err := utils.FindEventAndAttr(index, tx, &identitytypes.EvtCreatedHumanId{}, "HumanId")
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
	return nil
}

func (m *Module) handleMsgCreateMultipleHumanIds(index int, tx *juno.Tx, msg *identitytypes.MsgCreateMultipleHumanIds) error {
	return nil
}

func (m *Module) handleMsgSetPrimaryNetworkWallet(index int, tx *juno.Tx, msg *identitytypes.MsgSetPrimaryNetworkWallet) error {
	return nil
}

func (m *Module) handleMsgAcceptLinkWalletToHumanProposal(index int, tx *juno.Tx, msg *identitytypes.MsgAcceptLinkWalletToHumanProposal) error {
	return nil
}

func (m *Module) handleMsgProposeLinkAccountToHuman(index int, tx *juno.Tx, msg *identitytypes.MsgProposeLinkAccountToHuman) error {
	return nil
}

func (m *Module) handleMsgCreateNetwork(index int, tx *juno.Tx, msg *identitytypes.MsgCreateNetwork) error {
	return m.db.SaveIdentityNetwork(&identitytypes.Network{
		Index:            msg.ShortName,
		Active:           true,
		FullName:         msg.FullName,
		IdentityProvider: msg.IdentityProvider,
		InviteOnly:       msg.InviteOnly,
	})
}
