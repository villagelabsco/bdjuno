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
		return m.HandleMsgVerifyAccount(cosmosMsg)
	case *kyctypes.MsgRevokeAccount:
		return m.HandleMsgRevokeAccount(cosmosMsg)
	case *kyctypes.MsgCreateInvite:
		return m.HandleMsgCreateInvite(cosmosMsg)
	case *kyctypes.MsgClaimInvite:
		return m.HandleMsgClaimInvite(cosmosMsg)
	case *kyctypes.MsgRescindInvite:
		return m.HandleMsgRescindInvite(cosmosMsg)
	case *kyctypes.MsgCreateMultipleInvites:
		return m.HandleMsgCreateMultipleInvites(cosmosMsg)
	case *kyctypes.MsgVerifyNetwork:
		return m.HandleMsgVerifyNetwork(cosmosMsg)
	case *kyctypes.MsgRevokeNetwork:
		return m.HandleMsgRevokeNetwork(cosmosMsg)
	case *kyctypes.MsgCreateHumanId:
		return m.HandleMsgCreateHumanId(cosmosMsg)
	case *kyctypes.MsgRegisterIdentityProvider:
		return m.HandleMsgRegisterIdentityProvider(cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderAdminAccounts:
		return m.HandleMsgSetIdentityProviderAdminAccounts(cosmosMsg)
	case *kyctypes.MsgSetIdentityProviderProviderAccounts:
		return m.HandleMsgSetIdentityProviderProviderAccounts(cosmosMsg)
	case *kyctypes.MsgJoinNetwork:
		return m.HandleMsgJoinNetwork(cosmosMsg)
	case *kyctypes.MsgCreateMultipleHumanIds:
		return m.HandleMsgCreateMultipleHumanIds(cosmosMsg)
	case *kyctypes.MsgSetPrimaryNetworkWallet:
		return m.HandleMsgSetPrimaryNetworkWallet(cosmosMsg)
	case *kyctypes.MsgAcceptLinkWalletToHumanProposal:
		return m.HandleMsgAcceptLinkWalletToHumanProposal(cosmosMsg)
	case *kyctypes.MsgProposeLinkAccountToHuman:
		return m.HandleMsgProposeLinkAccountToHuman(cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgVerifyAccount(msg *kyctypes.MsgVerifyAccount) error {
	return nil
}

func (m *Module) HandleMsgRevokeAccount(msg *kyctypes.MsgRevokeAccount) error {
	return nil
}

func (m *Module) HandleMsgCreateInvite(msg *kyctypes.MsgCreateInvite) error {
	return m.db.InsertInvite(msg.Network, &kyctypes.Invite{
		Challenge:        msg.Challenge,
		Registered:       false,
		ConfirmedAccount: "",
		InviteCreator:    msg.Creator,
		HumanId:          msg.HumanId,
		GivenRoles:       msg.GivenRoles,
	})
}

func (m *Module) HandleMsgClaimInvite(msg *kyctypes.MsgClaimInvite) error {
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
		if err := m.db.InsertUserNetworks(un); err != nil {
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

func (m *Module) HandleMsgRescindInvite(msg *kyctypes.MsgRescindInvite) error {
	return nil
}

func (m *Module) HandleMsgCreateMultipleInvites(msg *kyctypes.MsgCreateMultipleInvites) error {
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

	err := m.db.InsertMultipleInvites(msg.Network, inv)
	if err != nil {
		return fmt.Errorf("error inserting multiple invites: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgVerifyNetwork(msg *kyctypes.MsgVerifyNetwork) error {
	return nil
}

func (m *Module) HandleMsgRevokeNetwork(msg *kyctypes.MsgRevokeNetwork) error {
	return nil
}

func (m *Module) HandleMsgCreateHumanId(msg *kyctypes.MsgCreateHumanId) error {
	return nil
}

func (m *Module) HandleMsgRegisterIdentityProvider(msg *kyctypes.MsgRegisterIdentityProvider) error {
	return nil
}

func (m *Module) HandleMsgSetIdentityProviderAdminAccounts(msg *kyctypes.MsgSetIdentityProviderAdminAccounts) error {
	return nil
}

func (m *Module) HandleMsgSetIdentityProviderProviderAccounts(msg *kyctypes.MsgSetIdentityProviderProviderAccounts) error {
	return nil
}

func (m *Module) HandleMsgJoinNetwork(msg *kyctypes.MsgJoinNetwork) error {
	return nil
}

func (m *Module) HandleMsgCreateMultipleHumanIds(msg *kyctypes.MsgCreateMultipleHumanIds) error {
	return nil
}

func (m *Module) HandleMsgSetPrimaryNetworkWallet(msg *kyctypes.MsgSetPrimaryNetworkWallet) error {
	return nil
}

func (m *Module) HandleMsgAcceptLinkWalletToHumanProposal(msg *kyctypes.MsgAcceptLinkWalletToHumanProposal) error {
	return nil
}

func (m *Module) HandleMsgProposeLinkAccountToHuman(msg *kyctypes.MsgProposeLinkAccountToHuman) error {
	return nil
}
