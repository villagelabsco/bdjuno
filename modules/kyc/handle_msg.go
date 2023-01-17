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
	juno "github.com/forbole/juno/v3/types"
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *kyctypes.MsgVerifyAccount:
		return m.HandleMsgVerifyAccount(tx.Height, cosmosMsg)
	case *kyctypes.MsgRevokeAccount:
		return m.HandleMsgRevokeAccount(tx.Height, cosmosMsg)
	case *kyctypes.MsgCreateInvite:
		return m.HandleMsgCreateInvite(tx.Height, cosmosMsg)
	case *kyctypes.MsgClaimInvite:
		return m.HandleMsgClaimInvite(tx.Height, cosmosMsg)
	case *kyctypes.MsgRescindInvite:
		return m.HandleMsgRescindInvite(tx.Height, cosmosMsg)
	case *kyctypes.MsgCreateMultipleInvites:
		return m.HandleMsgCreateMultipleInvites(tx.Height, cosmosMsg)
	case *kyctypes.MsgVerifyNetwork:
		return m.HandleMsgVerifyNetwork(tx.Height, cosmosMsg)
	case *kyctypes.MsgRevokeNetwork:
		return m.HandleMsgRevokeNetwork(tx.Height, cosmosMsg)
	case *kyctypes.MsgCreateHumanId:
		return m.HandleMsgCreateHumanId(tx.Height, cosmosMsg)
	case *kyctypes.MsgRegisterIdentityProvider:
		return m.HandleMsgRegisterIdentityProvider(tx.Height, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderAdminAccounts:
		return m.HandleMsgSetIdentityProviderAdminAccounts(tx.Height, cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderProviderAccounts:
		return m.HandleMsgSetIdentityProviderProviderAccounts(tx.Height, cosmosMsg)
	case *kyctypes.MsgJoinNetwork:
		return m.HandleMsgJoinNetwork(tx.Height, cosmosMsg)
	case *kyctypes.MsgCreateMultipleHumanIds:
		return m.HandleMsgCreateMultipleHumanIds(tx.Height, cosmosMsg)
	case *kyctypes.MsgSetPrimaryNetworkWallet:
		return m.HandleMsgSetPrimaryNetworkWallet(tx.Height, cosmosMsg)
	case *kyctypes.MsgAcceptLinkWalletToHumanProposal:
		return m.HandleMsgAcceptLinkWalletToHumanProposal(tx.Height, cosmosMsg)
	case *kyctypes.MsgProposeLinkAccountToHuman:
		return m.HandleMsgProposeLinkAccountToHuman(tx.Height, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized kyc message type: %T", msg)
	}
}

func (m *Module) HandleMsgVerifyAccount(height int64, msg *kyctypes.MsgVerifyAccount) error {
	return nil
}

func (m *Module) HandleMsgRevokeAccount(height int64, msg *kyctypes.MsgRevokeAccount) error {
	return nil
}

func (m *Module) HandleMsgCreateInvite(height int64, msg *kyctypes.MsgCreateInvite) error {
	return m.db.SaveInvite(msg.Network, &kyctypes.Invite{
		Challenge:        msg.Challenge,
		Registered:       false,
		ConfirmedAccount: "",
		InviteCreator:    msg.Creator,
		HumanId:          msg.HumanId,
		GivenRoles:       msg.GivenRoles,
	})
}

func (m *Module) HandleMsgClaimInvite(height int64, msg *kyctypes.MsgClaimInvite) error {
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

func (m *Module) HandleMsgRescindInvite(height int64, msg *kyctypes.MsgRescindInvite) error {
	if err := m.db.DeleteInvite(msg.Network, msg.Challenge); err != nil {
		return fmt.Errorf("error updating invite: %s", err)
	}
	return nil
}

func (m *Module) HandleMsgCreateMultipleInvites(height int64, msg *kyctypes.MsgCreateMultipleInvites) error {
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

func (m *Module) HandleMsgVerifyNetwork(height int64, msg *kyctypes.MsgVerifyNetwork) error {
	acc, err := m.src.GetNetworkKyb(height, kyctypes.QueryGetNetworkKybRequest{
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

func (m *Module) HandleMsgRevokeNetwork(height int64, msg *kyctypes.MsgRevokeNetwork) error {
	acc, err := m.src.GetNetworkKyb(height, kyctypes.QueryGetNetworkKybRequest{
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

func (m *Module) HandleMsgCreateHumanId(height int64, msg *kyctypes.MsgCreateHumanId) error {
	return nil
}

func (m *Module) HandleMsgRegisterIdentityProvider(height int64, msg *kyctypes.MsgRegisterIdentityProvider) error {
	ip, err := m.src.GetIdentityProvider(height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) HandleMsgSetIdentityProviderAdminAccounts(height int64, msg *kyctypes.MsgSetIdentityProviderAdminAccounts) error {
	ip, err := m.src.GetIdentityProvider(height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) HandleMsgSetIdentityProviderProviderAccounts(height int64, msg *kyctypes.MsgSetIdentityProviderProviderAccounts) error {
	ip, err := m.src.GetIdentityProvider(height, kyctypes.QueryGetIdentityProviderRequest{
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

func (m *Module) HandleMsgJoinNetwork(height int64, msg *kyctypes.MsgJoinNetwork) error {
	return nil
}

func (m *Module) HandleMsgCreateMultipleHumanIds(height int64, msg *kyctypes.MsgCreateMultipleHumanIds) error {
	return nil
}

func (m *Module) HandleMsgSetPrimaryNetworkWallet(height int64, msg *kyctypes.MsgSetPrimaryNetworkWallet) error {
	return nil
}

func (m *Module) HandleMsgAcceptLinkWalletToHumanProposal(height int64, msg *kyctypes.MsgAcceptLinkWalletToHumanProposal) error {
	return nil
}

func (m *Module) HandleMsgProposeLinkAccountToHuman(height int64, msg *kyctypes.MsgProposeLinkAccountToHuman) error {
	return nil
}
