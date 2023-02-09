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

package economics

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/villagelabsco/bdjuno/v3/utils"
	juno "github.com/villagelabsco/juno/v4/types"
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
	"strconv"
)

func (m Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *econtypes.MsgRemoveHook:
		return m.handleMsgRemoveHook(index, tx, cosmosMsg)
	case *econtypes.MsgPostTransaction:
		return m.handleMsgPostTransaction(index, tx, cosmosMsg)
	case *econtypes.MsgPostTask:
		return m.handleMsgPostTask(index, tx, cosmosMsg)
	case *econtypes.MsgEnableDisableNetworkEconomics:
		return m.handleMsgEnableDisableNetworkEconomics(index, tx, cosmosMsg)
	case *econtypes.MsgTriggerScheduledHooks:
		return m.handleMsgTriggerScheduledHooks(index, tx, cosmosMsg)
	//case *econtypes.MsgExecuteOneShotMint:
	//	return m.handleMsgExecuteOneShotMint(index, tx, cosmosMsg)
	//case *econtypes.MsgExecuteOneShotTransfer:
	//	return m.handleMsgExecuteOneShotTransfer(index, tx, cosmosMsg)
	//case *econtypes.MsgExecuteOneShotBurn:
	//	return m.handleMsgExecuteOneShotBurn(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintTask:
		return m.handleMsgSetTransactionalHookPreMintTask(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintAccountingTokenForTasks:
		return m.handleMsgSetTransactionalHookPreMintAccountingTokenForTasks(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookMintShareToken:
		return m.handleMsgSetTransactionalHookMintShareToken(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookAutoSwapProductForDenom:
		return m.handleMsgSetTransactionalHookAutoSwapProductForDenom(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookApplyMarketplaceFees:
		return m.handleMsgSetTransactionalHookApplyMarketplaceFees(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintProduct:
		return m.handleMsgSetTransactionalHookPreMintProduct(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookMintAccountingTokenOnTransactions:
		return m.handleMsgSetTransactionalHookMintAccountingTokenOnTransactions(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintGsvTrackingToken:
		return m.handleMsgSetTransactionalHookPreMintGsvTrackingToken(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookTransferDenomToShareholders:
		return m.handleMsgSetScheduledHookTransferDenomToShareholders(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookAutoSwapDenom:
		return m.handleMsgSetScheduledHookAutoSwapDenom(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookDecayDenomForInactiveAccounts:
		return m.handleMsgSetScheduledHookDecayDenomForInactiveAccounts(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookRecurringMintToken:
		return m.handleMsgSetScheduledHookRecurringMintToken(index, tx, cosmosMsg)
	}

	return nil
}

func (m Module) handleMsgRemoveHook(index int, tx *juno.Tx, msg *econtypes.MsgRemoveHook) error {
	switch msg.HookType {
	case econtypes.HookType_HOOKTYPE_TRANSACTIONAL:
		return m.db.RemoveEconomicsTransactionHook(msg.Network, msg.HookIdx)
	case econtypes.HookType_HOOKTYPE_SCHEDULED:
		return m.db.RemoveEconomicsScheduledHook(msg.Network, msg.HookIdx)
	default:
		return fmt.Errorf("unrecognized hook type: %s", msg.HookType)
	}
}

// TODO: Need an event for this
func (m Module) handleMsgPostTransaction(index int, tx *juno.Tx, msg *econtypes.MsgPostTransaction) error {
	return nil
}

func (m Module) handleMsgPostTask(index int, tx *juno.Tx, msg *econtypes.MsgPostTask) error {
	return nil
}

func (m Module) handleMsgTriggerScheduledHooks(index int, tx *juno.Tx, msg *econtypes.MsgTriggerScheduledHooks) error {
	return m.db.SaveEconomicsScheduledHookManualTrigger(msg)
}

func (m Module) handleMsgEnableDisableNetworkEconomics(index int, tx *juno.Tx, msg *econtypes.MsgEnableDisableNetworkEconomics) error {
	active, err := utils.FindEventAndAttr(index, tx, &econtypes.NetworkEnabledDisabled{}, "status")
	if err != nil {
		return fmt.Errorf("error while getting economics active from event: %s", err)
	}
	a, err := strconv.ParseBool(active)
	if err != nil {
		return fmt.Errorf("error while parsing economics active from event: %s", err)
	}

	return m.db.SaveOrUpdateEconomicsNetworkEnabled(msg.Network, a)
}

func (m Module) handleMsgSetTransactionalHookPreMintTask(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintTask) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookPreMintAccountingTokenForTasks(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintAccountingTokenForTasks) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookMintShareToken(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookMintShareToken) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookAutoSwapProductForDenom(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookAutoSwapProductForDenom) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookApplyMarketplaceFees(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookApplyMarketplaceFees) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetScheduledHookTransferDenomToShareholders(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookTransferDenomToShareholders) error {
	return m.handleCreateScheduledHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetScheduledHookAutoSwapDenom(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookAutoSwapDenom) error {
	return m.handleCreateScheduledHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetScheduledHookDecayDenomForInactiveAccounts(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookDecayDenomForInactiveAccounts) error {
	return m.handleCreateScheduledHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookPreMintProduct(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintProduct) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookMintAccountingTokenOnTransactions(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookMintAccountingTokenOnTransactions) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetTransactionalHookPreMintGsvTrackingToken(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintGsvTrackingToken) error {
	return m.handleCreateTransactionHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleMsgSetScheduledHookRecurringMintToken(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookRecurringMintToken) error {
	return m.handleCreateScheduledHook(tx.Height, msg.Network, msg.HookIdx)
}

func (m Module) handleCreateTransactionHook(height int64, network string, index uint64) error {
	hk, err := m.src.GetTransactionHook(height, econtypes.QueryGetTransactionHookRequest{
		Network: network,
		Idx:     index,
	})
	if err != nil {
		return fmt.Errorf("error while getting transaction hook from source: %s", err)
	}
	hook := hk.TransactionHook

	return m.db.SaveEconomicsTransactionHook(&hook)
}

func (m Module) handleCreateScheduledHook(height int64, network string, index uint64) error {
	hk, err := m.src.GetScheduledHook(height, econtypes.QueryGetScheduledHookRequest{
		Network: network,
		HookIdx: index,
	})
	if err != nil {
		return fmt.Errorf("error while getting scheduled hook from source: %s", err)
	}
	hook := hk.ScheduledHook

	return m.db.SaveEconomicsScheduledHook(&hook)
}
