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

package identity

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/villagelabsco/bdjuno/v4/types"
	identitytypes "github.com/villagelabsco/village/x/identity/types"
)

func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", identitytypes.ModuleName).Int64("height", height).
		Msg("updating params")

	p, err := m.src.GetParams(height)
	if err != nil {
		return fmt.Errorf("error while getting identity params: %s", err)
	}

	return m.db.SaveIdentityParams(&types.IdentityParams{
		Params: p,
		Height: height,
	})
}
