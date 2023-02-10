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
	"encoding/base64"
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	"github.com/villagelabsco/bdjuno/v3/utils"
	econtypes "github.com/villagelabsco/village/x/economics/types"
	"time"
)

type DbEconomicsTransaction struct {
	Id                     uint64             `db:"id"`
	Network                string             `db:"network"`
	Creator                string             `db:"creator"`
	Seller                 string             `db:"seller"`
	Buyer                  string             `db:"buyer"`
	Amount                 sqlxtypes.JSONText `db:"amount"`
	ProductClass           string             `db:"product_class"`
	Metadata               sqlxtypes.JSONText `db:"metadata"`
	Force                  bool               `db:"force"`
	Ref                    string             `db:"ref"`
	Timestamp              time.Time          `db:"timestamp"`
	Memo                   string             `db:"memo"`
	HooksCumulativeResult  sqlxtypes.JSONText `db:"hooks_cumulative_result"`
	HooksIndividualResults sqlxtypes.JSONText `db:"hooks_individual_results"`
}

func (DbEconomicsTransaction) FromProto(msg *econtypes.MsgPostTransaction, retValB64 string) (DbEconomicsTransaction, error) {
	metadata, err := utils.ParseJsonStrMap(msg.MetadataJson)
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while parsing metadata: %s", err)
	}
	metadataB, err := json.Marshal(metadata)
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while marshaling metadata: %s", err)
	}

	retValB, err := base64.StdEncoding.DecodeString(retValB64)
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while decoding hooks ret val b64: %s", err)
	}

	var retVal econtypes.HookSequenceRetVal
	if err := json.Unmarshal(retValB, &retVal); err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while unmarshaling ret val b64: %s", err)
	}

	cumulativeResultB, err := json.Marshal(retVal.CumulativeResult)
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while marshaling cumulative result: %s", err)
	}

	individualResultsB, err := json.Marshal(retVal.IndivHookResults)
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while marshaling individual results: %s", err)
	}

	amt, err := json.Marshal(sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Amount))))
	if err != nil {
		return DbEconomicsTransaction{}, fmt.Errorf("error while marshaling amount: %s", err)
	}

	return DbEconomicsTransaction{
		Network:                msg.Network,
		Creator:                msg.Creator,
		Seller:                 msg.Seller,
		Buyer:                  msg.Buyer,
		Amount:                 amt,
		ProductClass:           msg.ProductClass,
		Metadata:               metadataB,
		Force:                  msg.Force,
		Ref:                    msg.Ref,
		Timestamp:              time.Unix(int64(msg.Ts), 0),
		Memo:                   msg.Memo,
		HooksCumulativeResult:  cumulativeResultB,
		HooksIndividualResults: individualResultsB,
	}, nil
}
