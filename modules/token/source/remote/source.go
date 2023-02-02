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
	"github.com/forbole/juno/v4/node/remote"
	tokentypes "github.com/villagelabsco/villaged/x/token/types"
)

type Source struct {
	*remote.Source
	q tokentypes.QueryClient
}

func NewSource(source *remote.Source, q tokentypes.QueryClient) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

func (s Source) GetParams(height int64) (tokentypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &tokentypes.QueryParamsRequest{})
	if err != nil {
		return tokentypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetToken(height int64, req tokentypes.QueryGetTokenRequest) (tokentypes.QueryGetTokenResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Token(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetTokenResponse{}, fmt.Errorf("error while getting token: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllToken(height int64, req tokentypes.QueryAllTokenRequest) (tokentypes.QueryAllTokenResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.TokenAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllTokenResponse{}, fmt.Errorf("error while getting all tokens: %s", err)
	}
	return *res, nil
}

func (s Source) GetOffRampOperations(height int64, req tokentypes.QueryGetOfframpOperationsRequest) (tokentypes.QueryGetOfframpOperationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.OfframpOperations(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetOfframpOperationsResponse{}, fmt.Errorf("error while getting offramp operations: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllOffRampOperations(height int64, req tokentypes.QueryAllOfframpOperationsRequest) (tokentypes.QueryAllOfframpOperationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.OfframpOperationsAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllOfframpOperationsResponse{}, fmt.Errorf("error while getting all offramp operations: %s", err)
	}
	return *res, nil
}

func (s Source) GetOnRampOperations(height int64, req tokentypes.QueryGetOnrampOperationsRequest) (tokentypes.QueryGetOnrampOperationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.OnrampOperations(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetOnrampOperationsResponse{}, fmt.Errorf("error while getting onramp operations: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllOnRampOperations(height int64, req tokentypes.QueryAllOnrampOperationsRequest) (tokentypes.QueryAllOnrampOperationsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.OnrampOperationsAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllOnrampOperationsResponse{}, fmt.Errorf("error while getting all onramp operations: %s", err)
	}
	return *res, nil
}

func (s Source) GetImmobilizedFunds(height int64, req tokentypes.QueryGetImmobilizedFundsRequest) (tokentypes.QueryGetImmobilizedFundsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ImmobilizedFunds(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetImmobilizedFundsResponse{}, fmt.Errorf("error while getting immobilized funds: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllImmobilizedFunds(height int64, req tokentypes.QueryAllImmobilizedFundsRequest) (tokentypes.QueryAllImmobilizedFundsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ImmobilizedFundsAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllImmobilizedFundsResponse{}, fmt.Errorf("error while getting all immobilized funds: %s", err)
	}
	return *res, nil
}

func (s Source) GetTokenDetails(height int64, req tokentypes.QueryGetTokenDetailsRequest) (tokentypes.QueryGetTokenDetailsResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.TokenDetails(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetTokenDetailsResponse{}, fmt.Errorf("error while getting token details: %s", err)
	}
	return *res, nil
}

func (s Source) GetPendingBalance(height int64, req tokentypes.QueryGetPendingBalanceRequest) (tokentypes.QueryGetPendingBalanceResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingBalance(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetPendingBalanceResponse{}, fmt.Errorf("error while getting pending balance: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllPendingBalance(height int64, req tokentypes.QueryAllPendingBalanceRequest) (tokentypes.QueryAllPendingBalanceResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingBalanceAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllPendingBalanceResponse{}, fmt.Errorf("error while getting all pending balance: %s", err)
	}
	return *res, nil
}

func (s Source) GetNbTokenCreationPerDay(height int64, req tokentypes.QueryGetNbTokenCreationPerDayRequest) (tokentypes.QueryGetNbTokenCreationPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbTokenCreationPerDay(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetNbTokenCreationPerDayResponse{}, fmt.Errorf("error while getting nb token creation per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllNbTokenCreationPerDay(height int64, req tokentypes.QueryAllNbTokenCreationPerDayRequest) (tokentypes.QueryAllNbTokenCreationPerDayResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.NbTokenCreationPerDayAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllNbTokenCreationPerDayResponse{}, fmt.Errorf("error while getting all nb token creation per day: %s", err)
	}
	return *res, nil
}

func (s Source) GetPendingClawbackableOperation(height int64, req tokentypes.QueryGetPendingClawbackableOperationRequest) (tokentypes.QueryGetPendingClawbackableOperationResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingClawbackableOperation(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetPendingClawbackableOperationResponse{}, fmt.Errorf("error while getting pending clawbackable operation: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllPendingClawbackableOperation(height int64, req tokentypes.QueryAllPendingClawbackableOperationRequest) (tokentypes.QueryAllPendingClawbackableOperationResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingClawbackableOperationAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllPendingClawbackableOperationResponse{}, fmt.Errorf("error while getting all pending clawbackable operation: %s", err)
	}
	return *res, nil
}

func (s Source) GetPendingClawbackableMultiOperation(height int64, req tokentypes.QueryGetPendingClawbackableMultiOperationRequest) (tokentypes.QueryGetPendingClawbackableMultiOperationResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingClawbackableMultiOperation(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetPendingClawbackableMultiOperationResponse{}, fmt.Errorf("error while getting pending clawbackable multi operation: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllPendingClawbackableMultiOperation(height int64, req tokentypes.QueryAllPendingClawbackableMultiOperationRequest) (tokentypes.QueryAllPendingClawbackableMultiOperationResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.PendingClawbackableMultiOperationAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllPendingClawbackableMultiOperationResponse{}, fmt.Errorf("error while getting all pending clawbackable multi operation: %s", err)
	}
	return *res, nil
}

func (s Source) GetLastInputActivity(height int64, req tokentypes.QueryGetLastInputActivityRequest) (tokentypes.QueryGetLastInputActivityResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.LastInputActivity(ctx, &req)
	if err != nil {
		return tokentypes.QueryGetLastInputActivityResponse{}, fmt.Errorf("error while getting last input activity: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllLastInputActivity(height int64, req tokentypes.QueryAllLastInputActivityRequest) (tokentypes.QueryAllLastInputActivityResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.LastInputActivityAll(ctx, &req)
	if err != nil {
		return tokentypes.QueryAllLastInputActivityResponse{}, fmt.Errorf("error while getting all last input activity: %s", err)
	}
	return *res, nil
}

func (s Source) GetShowClawbackableAmount(height int64, req tokentypes.QueryShowClawbackableAmountRequest) (tokentypes.QueryShowClawbackableAmountResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.ShowClawbackableAmount(ctx, &req)
	if err != nil {
		return tokentypes.QueryShowClawbackableAmountResponse{}, fmt.Errorf("error while getting show clawbackable amount: %s", err)
	}
	return *res, nil
}
