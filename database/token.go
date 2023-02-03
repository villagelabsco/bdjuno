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

package database

import (
	"fmt"
	"github.com/villagelabsco/bdjuno/v3/database/types"
	tokentypes "github.com/villagelabsco/villaged/x/token/types"
)

func (db *Db) TokenDenom(denom string) (*tokentypes.Token, error) {
	stmt := `
		SELECT (network, denom, ticker, description, nb_decimals, transferable, backing_asset, admin, name_id, incentive_type, icon_path, referenced_denom, offramp_enabled, clawback_enabled, clawback_period_sec)
		FROM token_tokens
		WHERE denom = $1
		LIMIT 1;
	`

	var t types.DbToken
	if err := db.Sqlx.Select(&t, stmt, denom); err != nil {
		return nil, fmt.Errorf("error while getting token: %s", err)
	}

	return t.ToProto(), nil
}

func (db *Db) SaveOrUpdateTokenDenom(token *tokentypes.Token) error {
	stmt := `
		INSERT INTO token_tokens (network, denom, ticker, description, nb_decimals, transferable, backing_asset, admin, name_id, incentive_type, icon_path, referenced_denom, offramp_enabled, clawback_enabled, clawback_period_sec) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		ON CONFLICT (denom) 
		    DO UPDATE SET
		    	ticker = $3,
		    	description = $4,
		    	nb_decimals = $5,
		    	transferable = $6,
		    	backing_asset = $7,
		    	admin = $8,
		    	name_id = $9,
		    	incentive_type = $10,
		    	icon_path = $11,
		    	referenced_denom = $12,
		    	offramp_enabled = $13,
		    	clawback_enabled = $14,
		    	clawback_period_sec = $15;
	`

	t := types.DbToken{}.FromProto(token)
	_, err := db.SQL.Exec(stmt,
		t.Network,
		t.Denom,
		t.Ticker,
		t.Description,
		t.NbDecimals,
		t.Transferable,
		t.BackingAsset,
		t.Admin,
		t.NameId,
		t.IncentiveType,
		t.IconPath,
		t.ReferencedDenom,
		t.OfframpEnabled,
		t.ClawbackEnabled,
		t.ClawbackPeriodSec,
	)
	if err != nil {
		return fmt.Errorf("error while storing or updating token: %s", err)
	}

	return nil
}
