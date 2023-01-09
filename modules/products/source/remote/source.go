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

package remote

import (
	"fmt"
	"github.com/forbole/juno/v3/node/remote"
	productstypes "github.com/villagelabs/villaged/x/products/types"
)

type Source struct {
	*remote.Source
	q productstypes.QueryClient
}

func NewSource(source *remote.Source, q productstypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (productstypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &productstypes.QueryParamsRequest{})
	if err != nil {
		return productstypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetProduct(height int64, req productstypes.QueryGetProductRequest) (productstypes.QueryGetProductResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Product(ctx, &req)
	if err != nil {
		return productstypes.QueryGetProductResponse{}, fmt.Errorf("error while getting product: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllProduct(height int64, req productstypes.QueryAllProductRequest) (productstypes.QueryAllProductResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ProductAll(ctx, &req)
	if err != nil {
		return productstypes.QueryAllProductResponse{}, fmt.Errorf("error while getting all products: %s", err)
	}
	return *res, nil
}
