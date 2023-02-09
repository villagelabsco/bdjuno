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

func (db *Db) SaveTokenDenom(token *tokentypes.Token) error {
	stmt := `
		INSERT INTO token_tokens (network, denom, ticker, description, nb_decimals, transferable, backing_asset, admin, name_id, incentive_type, icon_path, referenced_denom, offramp_enabled, clawback_enabled, clawback_period_sec) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);
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

func (db *Db) UpdateTokenDenom(token *tokentypes.Token) error {
	stmt := `
		UPDATE token_tokens
		SET ticker = $2,
			description = $3,
			icon_path = $4
		WHERE denom = $1;
	`

	_, err := db.SQL.Exec(stmt, token.Denom, token.Ticker, token.Description, token.IconPath)
	if err != nil {
		return fmt.Errorf("error while updating token: %s", err)
	}

	return nil
}

func (db *Db) UpdateTokenAdminAccount(denom, adminAccount string) error {
	stmt := `
		UPDATE token_tokens
		SET admin = $2
		WHERE denom = $1;
	`

	_, err := db.SQL.Exec(stmt, denom, adminAccount)
	if err != nil {
		return fmt.Errorf("error while updating token admin account: %s", err)
	}

	return nil
}

func (db *Db) SaveTokenOnrampOperation(op tokentypes.OnrampOperations) error {
	stmt := `
		INSERT INTO token_onramp_operations (payment_ref, account, amount)
		VALUES ($1, $2, $3);
	`

	dbOp := types.DbTokenOnrampOperation{}.FromProto(op)
	_, err := db.SQL.Exec(stmt,
		dbOp.PaymentRef,
		dbOp.Account,
		dbOp.Amount,
	)
	if err != nil {
		return fmt.Errorf("error while storing token onramp operation: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateTokenOfframpOperation(op tokentypes.OfframpOperations) error {
	stmt := `
		INSERT INTO token_offramp_operations (id, account, human_id, executed, amount, creation_block, execution_block, funds_transfer_method_pseudo_id, id_provider) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			account = $2,
			human_id = $3,
			executed = $4,
			amount = $5,
			creation_block = $6,
			execution_block = $7,
			funds_transfer_method_pseudo_id = $8,
			id_provider = $9;
	`

	dbOp := types.DbTokenOfframpOperation{}.FromProto(op)
	_, err := db.SQL.Exec(stmt,
		dbOp.Id,
		dbOp.Account,
		dbOp.HumanId,
		dbOp.Executed,
		dbOp.Amount,
		dbOp.CreationBlock,
		dbOp.ExecutionBlock,
		dbOp.FundsTransferMethodPseudoId,
		dbOp.IdProvider,
	)
	if err != nil {
		return fmt.Errorf("error while storing or updating token offramp operation: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateTokenImmobilizedFunds(i tokentypes.ImmobilizedFunds) error {
	stmt := `
		INSERT INTO token_immobilized_funds (account, amount) 
		VALUES ($1, $2) ON CONFLICT (account) 
		DO UPDATE SET amount = $2;
	`

	dbImm := types.DbTokenImmobilizedFunds{}.FromProto(i)
	_, err := db.SQL.Exec(stmt,
		dbImm.Account,
		dbImm.Amount,
	)
	if err != nil {
		return fmt.Errorf("error while storing or updating token immobilized funds: %s", err)
	}

	return nil
}

func (db *Db) DeleteTokenOfframpOperation(id uint64) error {
	stmt := `
		DELETE FROM token_offramp_operations WHERE id = $1;
	`

	_, err := db.SQL.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("error while deleting token offramp operation: %s", err)
	}

	return nil
}
