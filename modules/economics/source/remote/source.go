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
	"github.com/villagelabsco/juno/v4/node/remote"
	econtypes "github.com/villagelabsco/village/x/economics/types"
)

type Source struct {
	*remote.Source
	q econtypes.QueryClient
}

func NewSource(source *remote.Source, q econtypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64, req econtypes.QueryParamsRequest) (econtypes.QueryParamsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &req)
	if err != nil {
		return econtypes.QueryParamsResponse{}, fmt.Errorf("error while getting economics params: %s", err)
	}
	return *res, nil
}

func (s Source) GetTransactionHook(height int64, req econtypes.QueryGetTransactionHookRequest) (econtypes.QueryGetTransactionHookResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.TransactionHook(ctx, &req)
	if err != nil {
		return econtypes.QueryGetTransactionHookResponse{}, fmt.Errorf("error while getting transaction hook: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllTransactionHook(height int64, req econtypes.QueryAllTransactionHookRequest) (econtypes.QueryAllTransactionHookResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.TransactionHookAll(ctx, &req)
	if err != nil {
		return econtypes.QueryAllTransactionHookResponse{}, fmt.Errorf("error while getting all transaction hooks: %s", err)
	}
	return *res, nil
}

func (s Source) GetScheduledHook(height int64, req econtypes.QueryGetScheduledHookRequest) (econtypes.QueryGetScheduledHookResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ScheduledHook(ctx, &req)
	if err != nil {
		return econtypes.QueryGetScheduledHookResponse{}, fmt.Errorf("error while getting scheduled hook: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllScheduledHook(height int64, req econtypes.QueryAllScheduledHookRequest) (econtypes.QueryAllScheduledHookResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ScheduledHookAll(ctx, &req)
	if err != nil {
		return econtypes.QueryAllScheduledHookResponse{}, fmt.Errorf("error while getting all scheduled hooks: %s", err)
	}
	return *res, nil
}

func (s Source) GetNetworkEnabled(height int64, req econtypes.QueryGetNetworkEnabledRequest) (econtypes.QueryGetNetworkEnabledResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkEnabled(ctx, &req)
	if err != nil {
		return econtypes.QueryGetNetworkEnabledResponse{}, fmt.Errorf("error while getting network enabled: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNetworkEnabled(height int64, req econtypes.QueryAllNetworkEnabledRequest) (econtypes.QueryAllNetworkEnabledResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NetworkEnabledAll(ctx, &req)
	if err != nil {
		return econtypes.QueryAllNetworkEnabledResponse{}, fmt.Errorf("error while getting all network enabled: %s", err)
	}
	return *res, nil
}

func (s Source) GetNbTxPerDay(height int64, req econtypes.QueryGetNbTxPerDayRequest) (econtypes.QueryGetNbTxPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbTxPerDay(ctx, &req)
	if err != nil {
		return econtypes.QueryGetNbTxPerDayResponse{}, fmt.Errorf("error while getting nb tx per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNbTxPerDay(height int64, req econtypes.QueryAllNbTxPerDayRequest) (econtypes.QueryAllNbTxPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbTxPerDayAll(ctx, &req)
	if err != nil {
		return econtypes.QueryAllNbTxPerDayResponse{}, fmt.Errorf("error while getting all nb tx per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetScheduledHookExecutionState(height int64, req econtypes.QueryGetScheduledHookExecutionStateRequest) (econtypes.QueryGetScheduledHookExecutionStateResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ScheduledHookExecutionState(ctx, &req)
	if err != nil {
		return econtypes.QueryGetScheduledHookExecutionStateResponse{}, fmt.Errorf("error while getting scheduled hook execution state: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllScheduledHookExecutionState(height int64, req econtypes.QueryAllScheduledHookExecutionStateRequest) (econtypes.QueryAllScheduledHookExecutionStateResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ScheduledHookExecutionStateAll(ctx, &req)
	if err != nil {
		return econtypes.QueryAllScheduledHookExecutionStateResponse{}, fmt.Errorf("error while getting all scheduled hook execution states: %s", err)
	}
	return *res, nil
}
