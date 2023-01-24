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

package kyc

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v3/utils"
	juno "github.com/forbole/juno/v3/types"
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *kyctypes.MsgVerifyAccount:
		return m.handleMsgVerifyAccount(index, tx, cosmosMsg)
	case *kyctypes.MsgRevokeAccount:
		return m.handleMsgRevokeAccount(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateInvite:
		return m.handleMsgCreateInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgClaimInvite:
		return m.handleMsgClaimInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgRescindInvite:
		return m.handleMsgRescindInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateMultipleInvites:
		return m.handleMsgCreateMultipleInvites(index, tx, cosmosMsg)
	case *kyctypes.MsgVerifyNetwork:
		return m.handleMsgVerifyNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgRevokeNetwork:
		return m.handleMsgRevokeNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateHumanId:
		return m.handleMsgCreateHumanId(index, tx, cosmosMsg)
	case *kyctypes.MsgRegisterIdentityProvider:
		return m.handleMsgRegisterIdentityProvider(index, tx, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderAdminAccounts:
		return m.handleMsgSetIdentityProviderAdminAccounts(index, tx, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderProviderAccounts:
		return m.handleMsgSetIdentityProviderProviderAccounts(index, tx, cosmosMsg)
	case *kyctypes.MsgJoinNetwork:
		return m.handleMsgJoinNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateMultipleHumanIds:
		return m.handleMsgCreateMultipleHumanIds(index, tx, cosmosMsg)
	case *kyctypes.MsgSetPrimaryNetworkWallet:
		return m.handleMsgSetPrimaryNetworkWallet(index, tx, cosmosMsg)
	case *kyctypes.MsgAcceptLinkWalletToHumanProposal:
		return m.handleMsgAcceptLinkWalletToHumanProposal(index, tx, cosmosMsg)
	case *kyctypes.MsgProposeLinkAccountToHuman:
		return m.handleMsgProposeLinkAccountToHuman(index, tx, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized kyc message type: %T", msg)
	}
}

func (m *Module) handleMsgVerifyAccount(index int, tx *juno.Tx, msg *kyctypes.MsgVerifyAccount) error {
	return nil
}

func (m *Module) handleMsgRevokeAccount(index int, tx *juno.Tx, msg *kyctypes.MsgRevokeAccount) error {
	return nil
}

func (m *Module) handleMsgCreateInvite(index int, tx *juno.Tx, msg *kyctypes.MsgCreateInvite) error {
	return m.db.SaveInvite(msg.Network, &kyctypes.Invite{
		Challenge:        msg.Challenge,
		Registered:       false,
		ConfirmedAccount: "",
		InviteCreator:    msg.Creator,
		HumanId:          msg.HumanId,
		GivenRoles:       msg.GivenRoles,
	})
}

func (m *Module) handleMsgClaimInvite(index int, tx *juno.Tx, msg *kyctypes.MsgClaimInvite) error {
	if err := m.db.UpdateInvite(msg.Network, msg.Challenge, msg.Creator); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}

	// Update the user networks mapping
	un, err := m.db.UserNetworks(msg.Creator)
	if err != nil {
		return fmt.Errorf("error getting user networks: %s", err)
	}

	if un.Index == "" {
		un.Index = msg.Creator
		un.Networks = make([]string, 1)
		un.Networks[0] = msg.Network
		if err := m.db.SaveUserNetworks(un); err != nil {
			return fmt.Errorf("error inserting user networks: %s", err)
		}
	} else {
		un.Networks = append(un.Networks, msg.Network)
		if err := m.db.UpdateUserNetworks(un); err != nil {
			return fmt.Errorf("error updating user networks: %s", err)
		}
	}

	return nil
}

func (m *Module) handleMsgRescindInvite(index int, tx *juno.Tx, msg *kyctypes.MsgRescindInvite) error {
	if err := m.db.DeleteInvite(msg.Network, msg.Challenge); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}
	return nil
}

func (m *Module) handleMsgCreateMultipleInvites(index int, tx *juno.Tx, msg *kyctypes.MsgCreateMultipleInvites) error {
	inv := make([]*kyctypes.Invite, len(msg.Challenges))
	for i, c := range msg.Challenges {
		inv[i] = &kyctypes.Invite{
			Challenge:        c,
			Registered:       false,
			ConfirmedAccount: "",
			InviteCreator:    msg.Creator,
			HumanId:          msg.HumanIds[i],
			GivenRoles:       msg.GivenRoles,
		}
	}

	err := m.db.SaveMultipleInvites(msg.Network, inv)
	if err != nil {
		return fmt.Errorf("error inserting multiple invites: %s", err)
	}

	return nil
}

func (m *Module) handleMsgVerifyNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgVerifyNetwork) error {
	acc, err := m.src.GetNetworkKyb(tx.Height, kyctypes.QueryGetNetworkKybRequest{
		Index: msg.Network,
	})
	if err != nil {
		return fmt.Errorf("error getting network kyb: %s", err)
	}
	kyb := acc.NetworkKyb

	if err := m.db.SaveOrUpdateNetworkKyb(&kyb); err != nil {
		return fmt.Errorf("error saving network kyb: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRevokeNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgRevokeNetwork) error {
	acc, err := m.src.GetNetworkKyb(tx.Height, kyctypes.QueryGetNetworkKybRequest{
		Index: msg.Network,
	})
	if err != nil {
		return fmt.Errorf("error getting network kyb: %s", err)
	}
	kyb := acc.NetworkKyb

	if err := m.db.SaveOrUpdateNetworkKyb(&kyb); err != nil {
		return fmt.Errorf("error saving network kyb: %s", err)
	}

	return nil
}

func (m *Module) handleMsgCreateHumanId(index int, tx *juno.Tx, msg *kyctypes.MsgCreateHumanId) error {
	idx, err := utils.FindEventAndAttr(index, tx, &kyctypes.EvtCreatedHumanId{}, "HumanId")
	if err != nil {
		return fmt.Errorf("error getting human id from created event: %s", err)
	}
	h, err := m.src.GetHuman(tx.Height, kyctypes.QueryGetHumanRequest{
		Index: idx,
	})
	if err != nil {
		return fmt.Errorf("error getting human: %s", err)
	}
	human := h.Human

	if err := m.db.SaveOrUpdateHuman(&human); err != nil {
		return fmt.Errorf("error saving human: %s", err)
	}

	return nil
}

func (m *Module) handleMsgRegisterIdentityProvider(index int, tx *juno.Tx, msg *kyctypes.MsgRegisterIdentityProvider) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) handleMsgSetIdentityProviderAdminAccounts(index int, tx *juno.Tx, msg *kyctypes.MsgSetIdentityProviderAdminAccounts) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) handleMsgSetIdentityProviderProviderAccounts(index int, tx *juno.Tx, msg *kyctypes.MsgSetIdentityProviderProviderAccounts) error {
	ip, err := m.src.GetIdentityProvider(tx.Height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) handleMsgJoinNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgJoinNetwork) error {
	return nil
}

func (m *Module) handleMsgCreateMultipleHumanIds(index int, tx *juno.Tx, msg *kyctypes.MsgCreateMultipleHumanIds) error {
	return nil
}

func (m *Module) handleMsgSetPrimaryNetworkWallet(index int, tx *juno.Tx, msg *kyctypes.MsgSetPrimaryNetworkWallet) error {
	return nil
}

func (m *Module) handleMsgAcceptLinkWalletToHumanProposal(index int, tx *juno.Tx, msg *kyctypes.MsgAcceptLinkWalletToHumanProposal) error {
	return nil
}

func (m *Module) handleMsgProposeLinkAccountToHuman(index int, tx *juno.Tx, msg *kyctypes.MsgProposeLinkAccountToHuman) error {
	return nil
}
