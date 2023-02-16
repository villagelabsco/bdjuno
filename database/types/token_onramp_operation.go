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

type DbTokenOnrampOperation struct {
	PaymentRef string             `json:"payment_ref"`
	Account    string             `json:"account"`
	Amount     sqlxtypes.JSONText `json:"amount"`
}

func (DbTokenOnrampOperation) FromProto(op tokentypes.OnrampOperations) (DbTokenOnrampOperation, error) {
	amt, err := json.Marshal(op.Amount)
	if err != nil {
		return DbTokenOnrampOperation{}, fmt.Errorf("error while marshaling amount: %s", err)
	}

	return DbTokenOnrampOperation{
		PaymentRef: op.PaymentRef,
		Account:    op.Account,
		Amount:     amt,
	}, nil
}
