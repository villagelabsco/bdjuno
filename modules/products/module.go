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

package products

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/villagelabsco/bdjuno/v3/database"
	nftsource "github.com/villagelabsco/bdjuno/v3/modules/nft/source"
	productssource "github.com/villagelabsco/bdjuno/v3/modules/products/source"
	"github.com/villagelabsco/juno/v4/modules"
	productstypes "github.com/villagelabsco/village/x/products/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	cdc codec.Codec
	db  *database.Db

	src    productssource.Source
	nftSrc nftsource.Source
}

func NewModule(
	cdc codec.Codec,
	db *database.Db,
	src productssource.Source,
	nftSrc nftsource.Source) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		src:    src,
		nftSrc: nftSrc,
	}
}

func (m *Module) Name() string {
	return productstypes.ModuleName
}
