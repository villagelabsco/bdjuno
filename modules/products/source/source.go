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

package source

import productstypes "github.com/villagelabsco/villaged/x/products/types"

type Source interface {
	GetParams(height int64) (productstypes.Params, error)
	GetProductClassInfo(height int64, req productstypes.QueryGetProductClassInfoRequest) (productstypes.QueryGetProductClassInfoResponse, error)
	GetAllProductClassInfo(height int64, req productstypes.QueryAllProductClassInfoRequest) (productstypes.QueryAllProductClassInfoResponse, error)
}
