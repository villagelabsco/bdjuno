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
	juno "github.com/forbole/juno/v3/types"
	econtypes "github.com/villagelabs/villaged/x/economics/types"
)

func (m Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *econtypes.MsgRemoveHook:
		return m.handleMsgRemoveHook(index, tx, cosmosMsg)
	case *econtypes.MsgPostTransaction:
		return m.handleMsgPostTransaction(index, tx, cosmosMsg)
	case *econtypes.MsgPostTask:
		return m.handleMsgPostTask(index, tx, cosmosMsg)
	case *econtypes.MsgTriggerScheduledHooks:
		return m.handleMsgTriggerScheduledHooks(index, tx, cosmosMsg)
	case *econtypes.MsgEnableDisableNetworkEconomics:
		return m.handleMsgEnableDisableNetworkEconomics(index, tx, cosmosMsg)
	case *econtypes.MsgExecutePendingTask:
		return m.handleMsgExecutePendingTask(index, tx, cosmosMsg)
	case *econtypes.MsgExecuteOneShotMint:
		return m.handleMsgExecuteOneShotMint(index, tx, cosmosMsg)
	case *econtypes.MsgExecuteOneShotTransfer:
		return m.handleMsgExecuteOneShotTransfer(index, tx, cosmosMsg)
	case *econtypes.MsgExecuteOneShotBurn:
		return m.handleMsgExecuteOneShotBurn(index, tx, cosmosMsg)
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
	case *econtypes.MsgSetScheduledHookTransferDenomToShareholders:
		return m.handleMsgSetScheduledHookTransferDenomToShareholders(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookAutoSwapDenom:
		return m.handleMsgSetScheduledHookAutoSwapDenom(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookDecayDenomForInactiveAccounts:
		return m.handleMsgSetScheduledHookDecayDenomForInactiveAccounts(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintProduct:
		return m.handleMsgSetTransactionalHookPreMintProduct(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookMintAccountingTokenForRevShareMarketplaces:
		return m.handleMsgSetTransactionalHookMintAccountingTokenForRevShareMarketplaces(index, tx, cosmosMsg)
	case *econtypes.MsgSetTransactionalHookPreMintGsvTrackingToken:
		return m.handleMsgSetTransactionalHookPreMintGsvTrackingToken(index, tx, cosmosMsg)
	case *econtypes.MsgSetScheduledHookRecurringMintToken:
		return m.handleMsgSetScheduledHookRecurringMintToken(index, tx, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized economics message type: %T", msg)
	}
}

func (m Module) handleMsgRemoveHook(index int, tx *juno.Tx, msg *econtypes.MsgRemoveHook) error {
	return nil
}

func (m Module) handleMsgPostTransaction(index int, tx *juno.Tx, msg *econtypes.MsgPostTransaction) error {
	return nil
}

func (m Module) handleMsgPostTask(index int, tx *juno.Tx, msg *econtypes.MsgPostTask) error {
	return nil
}

func (m Module) handleMsgTriggerScheduledHooks(index int, tx *juno.Tx, msg *econtypes.MsgTriggerScheduledHooks) error {
	return nil
}

func (m Module) handleMsgEnableDisableNetworkEconomics(index int, tx *juno.Tx, msg *econtypes.MsgEnableDisableNetworkEconomics) error {
	return nil
}

func (m Module) handleMsgExecutePendingTask(index int, tx *juno.Tx, msg *econtypes.MsgExecutePendingTask) error {
	return nil
}

func (m Module) handleMsgExecuteOneShotMint(index int, tx *juno.Tx, msg *econtypes.MsgExecuteOneShotMint) error {
	return nil
}

func (m Module) handleMsgExecuteOneShotTransfer(index int, tx *juno.Tx, msg *econtypes.MsgExecuteOneShotTransfer) error {
	return nil
}

func (m Module) handleMsgExecuteOneShotBurn(index int, tx *juno.Tx, msg *econtypes.MsgExecuteOneShotBurn) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookPreMintTask(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintTask) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookPreMintAccountingTokenForTasks(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintAccountingTokenForTasks) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookMintShareToken(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookMintShareToken) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookAutoSwapProductForDenom(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookAutoSwapProductForDenom) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookApplyMarketplaceFees(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookApplyMarketplaceFees) error {
	return nil
}

func (m Module) handleMsgSetScheduledHookTransferDenomToShareholders(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookTransferDenomToShareholders) error {
	return nil
}

func (m Module) handleMsgSetScheduledHookAutoSwapDenom(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookAutoSwapDenom) error {
	return nil
}

func (m Module) handleMsgSetScheduledHookDecayDenomForInactiveAccounts(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookDecayDenomForInactiveAccounts) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookPreMintProduct(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintProduct) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookMintAccountingTokenForRevShareMarketplaces(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookMintAccountingTokenForRevShareMarketplaces) error {
	return nil
}

func (m Module) handleMsgSetTransactionalHookPreMintGsvTrackingToken(index int, tx *juno.Tx, msg *econtypes.MsgSetTransactionalHookPreMintGsvTrackingToken) error {
	return nil
}

func (m Module) handleMsgSetScheduledHookRecurringMintToken(index int, tx *juno.Tx, msg *econtypes.MsgSetScheduledHookRecurringMintToken) error {
	return nil
}
