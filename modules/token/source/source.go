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

import tokentypes "github.com/villagelabs/villaged/x/token/types"

type Source interface {
	GetParams(height int64) (tokentypes.Params, error)
	GetToken(height int64, req tokentypes.QueryGetTokenRequest) (tokentypes.QueryGetTokenResponse, error)
	GetAllToken(height int64, req tokentypes.QueryAllTokenRequest) (tokentypes.QueryAllTokenResponse, error)
	GetOffRampOperations(height int64, req tokentypes.QueryGetOfframpOperationsRequest) (tokentypes.QueryGetOfframpOperationsResponse, error)
	GetAllOffRampOperations(height int64, req tokentypes.QueryAllOfframpOperationsRequest) (tokentypes.QueryAllOfframpOperationsResponse, error)
	GetOnRampOperations(height int64, req tokentypes.QueryGetOnrampOperationsRequest) (tokentypes.QueryGetOnrampOperationsResponse, error)
	GetAllOnRampOperations(height int64, req tokentypes.QueryAllOnrampOperationsRequest) (tokentypes.QueryAllOnrampOperationsResponse, error)
	GetImmobilizedFunds(height int64, req tokentypes.QueryGetImmobilizedFundsRequest) (tokentypes.QueryGetImmobilizedFundsResponse, error)
	GetAllImmobilizedFunds(height int64, req tokentypes.QueryAllImmobilizedFundsRequest) (tokentypes.QueryAllImmobilizedFundsResponse, error)
	GetTokenDetails(height int64, req tokentypes.QueryGetTokenDetailsRequest) (tokentypes.QueryGetTokenDetailsResponse, error)
	GetPendingBalance(height int64, req tokentypes.QueryGetPendingBalanceRequest) (tokentypes.QueryGetPendingBalanceResponse, error)
	GetAllPendingBalance(height int64, req tokentypes.QueryAllPendingBalanceRequest) (tokentypes.QueryAllPendingBalanceResponse, error)
	GetNbTokenCreationPerDay(height int64, req tokentypes.QueryGetNbTokenCreationPerDayRequest) (tokentypes.QueryGetNbTokenCreationPerDayResponse, error)
	GetAllNbTokenCreationPerDay(height int64, req tokentypes.QueryAllNbTokenCreationPerDayRequest) (tokentypes.QueryAllNbTokenCreationPerDayResponse, error)
	GetPendingClawbackableOperation(height int64, req tokentypes.QueryGetPendingClawbackableOperationRequest) (tokentypes.QueryGetPendingClawbackableOperationResponse, error)
	GetAllPendingClawbackableOperation(height int64, req tokentypes.QueryAllPendingClawbackableOperationRequest) (tokentypes.QueryAllPendingClawbackableOperationResponse, error)
	GetPendingClawbackableMultiOperation(height int64, req tokentypes.QueryGetPendingClawbackableMultiOperationRequest) (tokentypes.QueryGetPendingClawbackableMultiOperationResponse, error)
	GetAllPendingClawbackableMultiOperation(height int64, req tokentypes.QueryAllPendingClawbackableMultiOperationRequest) (tokentypes.QueryAllPendingClawbackableMultiOperationResponse, error)
	GetLastInputActivity(height int64, req tokentypes.QueryGetLastInputActivityRequest) (tokentypes.QueryGetLastInputActivityResponse, error)
	GetAllLastInputActivity(height int64, req tokentypes.QueryAllLastInputActivityRequest) (tokentypes.QueryAllLastInputActivityResponse, error)
	GetShowClawbackableAmount(height int64, req tokentypes.QueryShowClawbackableAmountRequest) (tokentypes.QueryShowClawbackableAmountResponse, error)
}
