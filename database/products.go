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
	"encoding/json"
	"fmt"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/forbole/bdjuno/v3/database/db_types"
	productstypes "github.com/villagelabs/villaged/x/products/types"
)

func (db *Db) SaveOrUpdateProductClass(
	pc productstypes.ProductClassInfo,
	nftClass *nfttypes.Class,
	metadata productstypes.StdClassData,
	specificMetadata productstypes.ProductClassData) error {
	specificMetadataB, err := json.Marshal(specificMetadata)
	if err != nil {
		return fmt.Errorf("error while marshalling specific metadata: %s", err)
	}

	metadata.SpecificMetadata = nil
	metadataB, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("error while marshalling metadata: %s", err)
	}

	pci := db_types.DbProductClassInfo{}.FromProto(pc, nftClass, metadataB, specificMetadataB)
	return db.saveOrUpdateProductClass(pci)
}

func (db *Db) SaveOrUpdateTaskClass(
	tc productstypes.ProductClassInfo,
	nftClass *nfttypes.Class,
	metadata productstypes.StdClassData,
	specificMetadata productstypes.TaskClassData) error {
	specificMetadataB, err := json.Marshal(specificMetadata)
	if err != nil {
		return fmt.Errorf("error while marshalling specific metadata: %s", err)
	}

	metadata.SpecificMetadata = nil
	metadataB, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("error while marshalling metadata: %s", err)
	}

	tci := db_types.DbProductClassInfo{}.FromProto(tc, nftClass, metadataB, specificMetadataB)
	return db.saveOrUpdateProductClass(tci)
}

func (db *Db) UpdateProductClassDisabled(fullClassId string, val bool) error {
	stmt := `
		UPDATE products_product_class_infos
		SET metadata = jsonb_set(metadata, '{disabled}', $1::jsonb, true)
		WHERE full_class_id = $2
	`

	_, err := db.Sql.Exec(stmt, val, fullClassId)
	if err != nil {
		return fmt.Errorf("error while updating product class metadata: %s", err)
	}
	return nil
}

func (db *Db) saveOrUpdateProductClass(pci db_types.DbProductClassInfo) error {
	stmt := `
		INSERT INTO products_product_class_infos (network, class_id, full_class_id, class_type, name, description, metadata, specific_metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		ON CONFLICT (full_class_id) DO 
		UPDATE
		    SET 
		        name = $5,
		        description = $6,
		        metadata = $7,
		        specific_metadata = $8
	`

	_, err := db.Sql.Exec(stmt,
		pci.Network,
		pci.ClassId,
		pci.FullClassId,
		pci.ClassType,
		pci.Name,
		pci.Description,
		pci.Metadata,
		pci.SpecificMetadata,
	)
	if err != nil {
		return fmt.Errorf("error while saving product class info: %s", err)
	}

	return nil
}
