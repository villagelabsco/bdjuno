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

package rbac

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/villagelabsco/bdjuno/v3/database"
	rbacsource "github.com/villagelabsco/bdjuno/v3/modules/rbac/source"
	"github.com/villagelabsco/juno/v4/modules"
	rbactypes "github.com/villagelabsco/village/x/rbac/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	cdc codec.Codec
	db  *database.Db

	src rbacsource.Source
}

func NewModule(cdc codec.Codec, db *database.Db, src rbacsource.Source) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
		src: src,
	}
}

func (m *Module) Name() string {
	return rbactypes.ModuleName
}
