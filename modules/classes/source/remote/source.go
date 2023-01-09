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
	classestypes "github.com/villagelabs/villaged/x/classes/types"
)

type Source struct {
	*remote.Source
	q classestypes.QueryClient
}

func NewSource(source *remote.Source, q classestypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (classestypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &classestypes.QueryParamsRequest{})
	if err != nil {
		return classestypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetClass(height int64, req classestypes.QueryGetClassRequest) (classestypes.QueryGetClassResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Class(ctx, &req)
	if err != nil {
		return classestypes.QueryGetClassResponse{}, fmt.Errorf("error while getting class: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllClass(height int64, req classestypes.QueryAllClassesRequest) (classestypes.QueryAllClassesResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ClassAll(ctx, &req)
	if err != nil {
		return classestypes.QueryAllClassesResponse{}, fmt.Errorf("error while getting all classes: %s", err)
	}
	return *res, nil
}
