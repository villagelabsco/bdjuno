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
	sqlxtypes "github.com/jmoiron/sqlx/types"
	tokentypes "github.com/villagelabsco/village/x/token/types"
)

type DbTokenOfframpOperation struct {
	Id                          uint64             `json:"id"`
	Account                     string             `json:"account"`
	HumanId                     string             `json:"human_id"`
	Executed                    bool               `json:"executed"`
	Amount                      sqlxtypes.JSONText `json:"amount"`
	CreationBlock               uint64             `json:"creation_block"`
	ExecutionBlock              uint64             `json:"execution_block"`
	FundsTransferMethodPseudoId string             `json:"funds_transfer_method_pseudo_id"`
	IdProvider                  string             `json:"id_provider"`
}

func (DbTokenOfframpOperation) FromProto(op tokentypes.OfframpOperations) (DbTokenOfframpOperation, error) {
	amt, err := json.Marshal(op.Amount)
	if err != nil {
		return DbTokenOfframpOperation{}, fmt.Errorf("error while marshaling amount: %s", err)
	}

	return DbTokenOfframpOperation{
		Id:                          op.Id,
		Account:                     op.Account,
		HumanId:                     op.HumanId,
		Executed:                    op.Executed,
		Amount:                      amt,
		CreationBlock:               op.CreationBlock,
		ExecutionBlock:              op.ExecutionBlock,
		FundsTransferMethodPseudoId: op.FundsTransferMethodPseudoId,
		IdProvider:                  op.IdProvider,
	}, nil
}
