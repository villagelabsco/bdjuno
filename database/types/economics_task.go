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
	sqlxtypes "github.com/jmoiron/sqlx/types"
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
	"time"
)

type DbEconomicsTask struct {
	Id                     uint64             `db:"id"`
	Network                string             `db:"network"`
	Creator                string             `db:"creator"`
	Tasker                 string             `db:"tasker"`
	Buyer                  string             `db:"buyer"`
	TaskCount              uint64             `db:"task_count"`
	TaskClassId            string             `db:"task_class_id"`
	Force                  bool               `db:"force"`
	Ref                    string             `db:"ref"`
	Timestamp              time.Time          `db:"timestamp"`
	Memo                   string             `db:"memo"`
	HooksCumulativeResult  sqlxtypes.JSONText `db:"hooks_cumulative_result"`
	HooksIndividualResults sqlxtypes.JSONText `db:"hooks_individual_results"`
}

func (DbEconomicsTask) FromProto(msg *econtypes.MsgPostTask, retValB64 string) (DbEconomicsTask, error) {
	retValB, err := base64.StdEncoding.DecodeString(retValB64)
	if err != nil {
		return DbEconomicsTask{}, fmt.Errorf("error while decoding hooks ret val b64: %s", err)
	}

	var retVal econtypes.HookSequenceRetVal
	if err := json.Unmarshal(retValB, &retVal); err != nil {
		return DbEconomicsTask{}, fmt.Errorf("error while unmarshaling hooks ret val: %s", err)
	}

	cumulativeResultB, err := json.Marshal(retVal.CumulativeResult)
	if err != nil {
		return DbEconomicsTask{}, fmt.Errorf("error while marshaling cumulative result: %s", err)
	}

	individualResultsB, err := json.Marshal(retVal.IndivHookResults)
	if err != nil {
		return DbEconomicsTask{}, fmt.Errorf("error while marshaling individual results: %s", err)
	}

	return DbEconomicsTask{
		Network:                msg.Network,
		Creator:                msg.Creator,
		Tasker:                 msg.Tasker,
		Buyer:                  msg.Buyer,
		TaskCount:              msg.TaskCount,
		TaskClassId:            msg.TaskClassId,
		Force:                  msg.Force,
		Ref:                    msg.Ref,
		Timestamp:              time.Unix(int64(msg.Ts), 0),
		Memo:                   msg.Memo,
		HooksCumulativeResult:  cumulativeResultB,
		HooksIndividualResults: individualResultsB,
	}, nil
}
