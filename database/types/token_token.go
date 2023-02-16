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

import tokentypes "github.com/villagelabsco/village/x/token/types"

type DbToken struct {
	Network           string `db:"network"`
	Denom             string `db:"denom"`
	Ticker            string `db:"ticker"`
	Description       string `db:"description"`
	NbDecimals        uint32 `db:"nb_decimals"`
	Transferable      bool   `db:"transferable"`
	BackingAsset      string `db:"backing_asset"`
	Admin             string `db:"admin"`
	NameId            string `db:"name_id"`
	IncentiveType     string `db:"incentive_type"`
	IconPath          string `db:"icon_path"`
	ReferencedDenom   string `db:"referenced_denom"`
	OfframpEnabled    bool   `db:"offramp_enabled"`
	ClawbackEnabled   bool   `db:"clawback_enabled"`
	ClawbackPeriodSec uint64 `db:"clawback_period_sec"`
}

func (DbToken) FromProto(t *tokentypes.Token) DbToken {
	return DbToken{
		Network:           t.Network,
		Denom:             t.Denom,
		Ticker:            t.Ticker,
		Description:       t.Description,
		NbDecimals:        t.NbDecimals,
		Transferable:      t.Transferable,
		BackingAsset:      t.BackingAsset,
		Admin:             t.Admin,
		NameId:            t.NameId,
		IncentiveType:     t.IncentiveType,
		IconPath:          t.IconPath,
		ReferencedDenom:   t.ReferencedDenom,
		OfframpEnabled:    t.OfframpEnabled,
		ClawbackEnabled:   t.ClawbackEnabled,
		ClawbackPeriodSec: t.ClawbackPeriodSec,
	}
}

func (t DbToken) ToProto() *tokentypes.Token {
	return &tokentypes.Token{
		Network:           t.Network,
		Denom:             t.Denom,
		Ticker:            t.Ticker,
		Description:       t.Description,
		NbDecimals:        t.NbDecimals,
		Transferable:      t.Transferable,
		BackingAsset:      t.BackingAsset,
		Admin:             t.Admin,
		NameId:            t.NameId,
		IncentiveType:     t.IncentiveType,
		IconPath:          t.IconPath,
		ReferencedDenom:   t.ReferencedDenom,
		OfframpEnabled:    t.OfframpEnabled,
		ClawbackEnabled:   t.ClawbackEnabled,
		ClawbackPeriodSec: t.ClawbackPeriodSec,
	}
}
