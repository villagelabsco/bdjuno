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

package types

import (
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	econtypes "github.com/villagelabsco/village/x/economics/types"
)

type DbEconomicsTransactionHook struct {
	Network string                      `db:"network"`
	Index   uint64                      `db:"index"`
	Trigger econtypes.LegacyTriggerType `db:"trigger"`
	// Trigger     econtypes.TriggerType         `db:"trigger"`
	Type        econtypes.TransactionHookType `db:"type"`
	NameId      string                        `db:"name_id"`
	Description string                        `db:"description"`
	Params      sqlxtypes.JSONText            `db:"params"`
}

func (DbEconomicsTransactionHook) FromProto(h *econtypes.TransactionHook) (DbEconomicsTransactionHook, error) {
	var p proto.Unmarshaler
	if len(h.HookParams.Value) > 0 {
		switch h.HookType {
		case econtypes.TransactionHookType_TRANSHOOKTYPE_PREMINT_TASK:
			p = new(econtypes.HookParamsPreMintTask)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_PREMINT_PRODUCT:
			p = new(econtypes.HookParamsPreMintProduct)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_PREMINT_ACCOUNTING_TOKEN_FOR_TASKS:
			p = new(econtypes.HookParamsPremintAccountingTokenForTasks)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_PREMINT_GSV_TRACKING_TOKEN:
			p = new(econtypes.HookParamsPremintGsvTrackingToken)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_MINT_ACCOUNTING_TOKEN_ON_TRANSACTIONS:
			p = new(econtypes.HookParamsMintAccountingTokenOnTransactions)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_MINT_SHARE_TOKEN:
			p = new(econtypes.HookParamsMintShareToken)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_AUTO_SWAP_PRODUCT_FOR_DENOM:
			p = new(econtypes.HookParamsAutoSwapProductForDenom)
		case econtypes.TransactionHookType_TRANSHOOKTYPE_APPLY_MARKETPLACE_FEES:
			p = new(econtypes.HookParamsApplyMarketplaceFees)
		}

		if err := p.Unmarshal(h.HookParams.Value); err != nil {
			return DbEconomicsTransactionHook{}, fmt.Errorf("error unmarshalling hook params: %v", err)
		}
	}

	bParams, err := json.Marshal(p)
	if err != nil {
		return DbEconomicsTransactionHook{}, fmt.Errorf("error marshalling hook params: %v", err)
	}

	return DbEconomicsTransactionHook{
		Network:     h.Network,
		Index:       h.HookIdx,
		Trigger:     h.Trigger,
		Type:        h.HookType,
		NameId:      h.NameId,
		Description: h.HookDescription,
		Params:      bParams,
	}, nil
}
