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
		return m.HandleMsgVerifyAccount(index, tx, cosmosMsg)
	case *kyctypes.MsgRevokeAccount:
		return m.HandleMsgRevokeAccount(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateInvite:
		return m.HandleMsgCreateInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgClaimInvite:
		return m.HandleMsgClaimInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgRescindInvite:
		return m.HandleMsgRescindInvite(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateMultipleInvites:
		return m.HandleMsgCreateMultipleInvites(index, tx, cosmosMsg)
	case *kyctypes.MsgVerifyNetwork:
		return m.HandleMsgVerifyNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgRevokeNetwork:
		return m.HandleMsgRevokeNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateHumanId:
		return m.HandleMsgCreateHumanId(index, tx, cosmosMsg)
	case *kyctypes.MsgRegisterIdentityProvider:
		return m.HandleMsgRegisterIdentityProvider(index, tx, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderAdminAccounts:
		return m.HandleMsgSetIdentityProviderAdminAccounts(index, tx, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderProviderAccounts:
		return m.HandleMsgSetIdentityProviderProviderAccounts(index, tx, cosmosMsg)
	case *kyctypes.MsgJoinNetwork:
		return m.HandleMsgJoinNetwork(index, tx, cosmosMsg)
	case *kyctypes.MsgCreateMultipleHumanIds:
		return m.HandleMsgCreateMultipleHumanIds(index, tx, cosmosMsg)
	case *kyctypes.MsgSetPrimaryNetworkWallet:
		return m.HandleMsgSetPrimaryNetworkWallet(index, tx, cosmosMsg)
	case *kyctypes.MsgAcceptLinkWalletToHumanProposal:
		return m.HandleMsgAcceptLinkWalletToHumanProposal(index, tx, cosmosMsg)
	case *kyctypes.MsgProposeLinkAccountToHuman:
		return m.HandleMsgProposeLinkAccountToHuman(index, tx, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized kyc message type: %T", msg)
	}
}

func (m *Module) HandleMsgVerifyAccount(index int, tx *juno.Tx, msg *kyctypes.MsgVerifyAccount) error {
	return nil
}

func (m *Module) HandleMsgRevokeAccount(index int, tx *juno.Tx, msg *kyctypes.MsgRevokeAccount) error {
	return nil
}

func (m *Module) HandleMsgCreateInvite(index int, tx *juno.Tx, msg *kyctypes.MsgCreateInvite) error {
	return m.db.SaveInvite(msg.Network, &kyctypes.Invite{
		Challenge:        msg.Challenge,
		Registered:       false,
		ConfirmedAccount: "",
		InviteCreator:    msg.Creator,
		HumanId:          msg.HumanId,
		GivenRoles:       msg.GivenRoles,
	})
}

func (m *Module) HandleMsgClaimInvite(index int, tx *juno.Tx, msg *kyctypes.MsgClaimInvite) error {
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

func (m *Module) HandleMsgRescindInvite(index int, tx *juno.Tx, msg *kyctypes.MsgRescindInvite) error {
	if err := m.db.DeleteInvite(msg.Network, msg.Challenge); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}
	return nil
}

func (m *Module) HandleMsgCreateMultipleInvites(index int, tx *juno.Tx, msg *kyctypes.MsgCreateMultipleInvites) error {
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

func (m *Module) HandleMsgVerifyNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgVerifyNetwork) error {
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

func (m *Module) HandleMsgRevokeNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgRevokeNetwork) error {
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

func (m *Module) HandleMsgCreateHumanId(index int, tx *juno.Tx, msg *kyctypes.MsgCreateHumanId) error {
	tx.FindEventByType(tx.Height, utils.ProtoMsgName(msg))
}

func (m *Module) HandleMsgRegisterIdentityProvider(index int, tx *juno.Tx, msg *kyctypes.MsgRegisterIdentityProvider) error {
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

func (m *Module) HandleMsgSetIdentityProviderAdminAccounts(index int, tx *juno.Tx, msg *kyctypes.MsgSetIdentityProviderAdminAccounts) error {
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

func (m *Module) HandleMsgSetIdentityProviderProviderAccounts(index int, tx *juno.Tx, msg *kyctypes.MsgSetIdentityProviderProviderAccounts) error {
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

func (m *Module) HandleMsgJoinNetwork(index int, tx *juno.Tx, msg *kyctypes.MsgJoinNetwork) error {
	return nil
}

func (m *Module) HandleMsgCreateMultipleHumanIds(index int, tx *juno.Tx, msg *kyctypes.MsgCreateMultipleHumanIds) error {
	return nil
}

func (m *Module) HandleMsgSetPrimaryNetworkWallet(index int, tx *juno.Tx, msg *kyctypes.MsgSetPrimaryNetworkWallet) error {
	return nil
}

func (m *Module) HandleMsgAcceptLinkWalletToHumanProposal(index int, tx *juno.Tx, msg *kyctypes.MsgAcceptLinkWalletToHumanProposal) error {
	return nil
}

func (m *Module) HandleMsgProposeLinkAccountToHuman(index int, tx *juno.Tx, msg *kyctypes.MsgProposeLinkAccountToHuman) error {
	return nil
}
