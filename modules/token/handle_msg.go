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

package token

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/villagelabsco/bdjuno/v3/utils"
	juno "github.com/villagelabsco/juno/v4/types"
	tokentypes "github.com/villagelabsco/village/x/token/types"
	"strconv"
)

func (m Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *tokentypes.MsgCreateToken:
		return m.handleMsgCreateToken(index, tx, cosmosMsg)
	case *tokentypes.MsgUpdateToken:
		return m.handleMsgUpdateToken(index, tx, cosmosMsg)
	case *tokentypes.MsgCreateAccountingToken:
		return m.handleMsgCreateAccountingToken(index, tx, cosmosMsg)
	case *tokentypes.MsgCreateRootCurrencyToken:
		return m.handleMsgCreateRootCurrencyToken(index, tx, cosmosMsg)
	case *tokentypes.MsgTransferTokenOwnership:
		return m.handleMsgTransferTokenOwnership(index, tx, cosmosMsg)
	case *tokentypes.MsgOracleExecuteOnrampMintForTreasury:
		return m.handleMsgOracleExecuteOnrampMintForTreasury(index, tx, cosmosMsg)
	case *tokentypes.MsgOracleExecuteOnrampMintForAccount:
		return m.handleMsgOracleExecuteOnrampMintForAccount(index, tx, cosmosMsg)
	case *tokentypes.MsgOracleExecuteOfframpBurn:
		return m.handleMsgOracleExecuteOfframpBurn(index, tx, cosmosMsg)
	case *tokentypes.MsgRequestOfframpBurn:
		return m.handleMsgRequestOfframpBurn(index, tx, cosmosMsg)
	case *tokentypes.MsgCancelOfframpRequest:
		return m.handleMsgCancelOfframpRequest(index, tx, cosmosMsg)
		//case *tokentypes.MsgMintTokens:
		//	return m.handleMsgMintTokens(index, tx, cosmosMsg)
		//case *tokentypes.MsgBurnTokens:
		//	return m.handleMsgBurnTokens(index, tx, cosmosMsg)
		//case *tokentypes.MsgSwapAccountingToken:
		//	return m.handleMsgSwapAccountingToken(index, tx, cosmosMsg)
		//case *tokentypes.MsgClaimPendingBalance:
		//	return m.handleMsgClaimPendingBalance(index, tx, cosmosMsg)
		//case *tokentypes.MsgClawbackTokens:
		//	return m.handleMsgClawbackTokens(index, tx, cosmosMsg)
	}

	return nil
}

func (m Module) handleMsgCreateToken(index int, tx *juno.Tx, msg *tokentypes.MsgCreateToken) error {
	denom, err := utils.FindEventAndAttr(index, tx, &tokentypes.EvtCreatedToken{}, "denom")
	if err != nil {
		return fmt.Errorf("error while getting token denom from tx events: %s", err)
	}

	return m.handleCreateToken(tx, denom)
}

func (m Module) handleCreateToken(tx *juno.Tx, denom string) error {
	t, err := m.src.GetToken(tx.Height, tokentypes.QueryGetTokenRequest{
		Denom: denom,
	})
	if err != nil {
		return fmt.Errorf("error while getting token from source: %s", err)
	}

	if err := m.db.SaveTokenDenom(&t.Token); err != nil {
		return fmt.Errorf("error while saving token: %s", err)
	}

	return nil
}

func (m Module) handleMsgUpdateToken(index int, tx *juno.Tx, msg *tokentypes.MsgUpdateToken) error {
	t := &tokentypes.Token{
		Denom:       msg.Denom,
		Ticker:      msg.Ticker,
		Description: msg.Description,
		IconPath:    msg.IconPath,
	}
	return m.db.UpdateTokenDenom(t)
}

func (m Module) handleMsgTransferTokenOwnership(index int, tx *juno.Tx, msg *tokentypes.MsgTransferTokenOwnership) error {
	return m.db.UpdateTokenAdminAccount(msg.Denom, msg.NewAdminAccount)
}

func (m Module) handleOnramp(tx *juno.Tx, paymentRef string) error {
	op, err := m.src.GetOnRampOperations(tx.Height, tokentypes.QueryGetOnrampOperationsRequest{
		PaymentRef: paymentRef,
	})
	if err != nil {
		return fmt.Errorf("error while getting onramp operation from source: %s", err)
	}
	operation := op.OnrampOperations

	if err := m.db.SaveTokenOnrampOperation(operation); err != nil {
		return fmt.Errorf("error while saving onramp operation: %s", err)
	}

	return nil
}

func (m Module) handleOfframp(tx *juno.Tx, payload struct {
	id      uint64
	account string
	denom   string
}) error {
	op, err := m.src.GetOffRampOperations(tx.Height, tokentypes.QueryGetOfframpOperationsRequest{
		Account: payload.account,
		Id:      payload.id,
	})
	if err != nil {
		return fmt.Errorf("error while getting offramp operation from source: %s", err)
	}
	operation := op.OfframpOperations

	if err := m.db.SaveOrUpdateTokenOfframpOperation(operation); err != nil {
		return fmt.Errorf("error while saving offramp operation: %s", err)
	}

	imm, err := m.src.GetImmobilizedFunds(tx.Height, tokentypes.QueryGetImmobilizedFundsRequest{
		Denom:   payload.denom,
		Account: payload.account,
	})
	if err != nil {
		return fmt.Errorf("error while getting immobilized funds from source: %s", err)
	}
	immobilizedFunds := imm.ImmobilizedFunds

	if err := m.db.SaveOrUpdateTokenImmobilizedFunds(immobilizedFunds); err != nil {
		return fmt.Errorf("error while saving immobilized funds: %s", err)
	}

	return nil
}

func (m Module) handleMsgOracleExecuteOnrampMintForTreasury(index int, tx *juno.Tx, msg *tokentypes.MsgOracleExecuteOnrampMintForTreasury) error {
	return m.handleOnramp(tx, msg.PaymentRef)
}

func (m Module) handleMsgOracleExecuteOnrampMintForAccount(index int, tx *juno.Tx, msg *tokentypes.MsgOracleExecuteOnrampMintForAccount) error {
	return m.handleOnramp(tx, msg.PaymentRef)
}

func (m Module) handleMsgOracleExecuteOfframpBurn(index int, tx *juno.Tx, msg *tokentypes.MsgOracleExecuteOfframpBurn) error {
	denom, err := utils.FindEventAndAttr(index, tx, &tokentypes.EvtOracleBurnedVillageUsd{}, "amount.denom")
	if err != nil {
		return fmt.Errorf("error while getting operation amount.denom from events: %s", err)
	}

	return m.handleOfframp(tx, struct {
		id      uint64
		account string
		denom   string
	}{id: uint64(msg.OfframpOperationIdx), account: msg.Creator, denom: denom})
}

func (m Module) handleMsgRequestOfframpBurn(index int, tx *juno.Tx, msg *tokentypes.MsgRequestOfframpBurn) error {
	idx, err := utils.FindEventAndAttr(index, tx, &tokentypes.EvtRequestedBurnVillageUsd{}, "itemIdx")
	if err != nil {
		return fmt.Errorf("error while getting operation itemIdx from events: %s", err)
	}
	id, err := strconv.Atoi(idx)
	if err != nil {
		return fmt.Errorf("error while converting operation itemIdx to int: %s", err)
	}

	return m.handleOfframp(tx, struct {
		id      uint64
		account string
		denom   string
	}{id: uint64(id), account: msg.Creator, denom: msg.Amount.Denom})
}

func (m Module) handleMsgCancelOfframpRequest(index int, tx *juno.Tx, msg *tokentypes.MsgCancelOfframpRequest) error {
	op, err := m.src.GetOffRampOperations(tx.Height, tokentypes.QueryGetOfframpOperationsRequest{
		Account: msg.Creator,
		Id:      msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while getting offramp operation from source: %s", err)
	}
	operation := op.OfframpOperations

	imm, err := m.src.GetImmobilizedFunds(tx.Height, tokentypes.QueryGetImmobilizedFundsRequest{
		Denom:   op.OfframpOperations.Amount.Denom,
		Account: msg.Creator,
	})
	if err != nil {
		return fmt.Errorf("error while getting immobilized funds from source: %s", err)
	}
	immobilizedFunds := imm.ImmobilizedFunds

	if err := m.db.DeleteTokenOfframpOperation(operation.Id); err != nil {
		return fmt.Errorf("error while saving offramp operation: %s", err)
	}

	if err := m.db.SaveOrUpdateTokenImmobilizedFunds(immobilizedFunds); err != nil {
		return fmt.Errorf("error while saving immobilized funds: %s", err)
	}

	return nil
}

func (m Module) handleMsgCreateAccountingToken(index int, tx *juno.Tx, msg *tokentypes.MsgCreateAccountingToken) error {
	denom, err := utils.FindEventAndAttr(index, tx, &tokentypes.EvtCreatedAccountingToken{}, "denom")
	if err != nil {
		return fmt.Errorf("error while getting token denom from events: %s", err)
	}

	return m.handleCreateToken(tx, denom)
}

func (m Module) handleMsgCreateRootCurrencyToken(index int, tx *juno.Tx, msg *tokentypes.MsgCreateRootCurrencyToken) error {
	denom, err := utils.FindEventAndAttr(index, tx, &tokentypes.EvtCreatedRootCurrencyToken{}, "denom")
	if err != nil {
		return fmt.Errorf("error while getting token denom from events: %s", err)
	}

	return m.handleCreateToken(tx, denom)
}
