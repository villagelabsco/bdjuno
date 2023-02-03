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

import (
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
)

type Source interface {
	GetParams(height int64, req econtypes.QueryParamsRequest) (econtypes.QueryParamsResponse, error)
	GetTransactionHook(height int64, req econtypes.QueryGetTransactionHookRequest) (econtypes.QueryGetTransactionHookResponse, error)
	GetAllTransactionHook(height int64, req econtypes.QueryAllTransactionHookRequest) (econtypes.QueryAllTransactionHookResponse, error)
	GetScheduledHook(height int64, req econtypes.QueryGetScheduledHookRequest) (econtypes.QueryGetScheduledHookResponse, error)
	GetAllScheduledHook(height int64, req econtypes.QueryAllScheduledHookRequest) (econtypes.QueryAllScheduledHookResponse, error)
	GetNetworkEnabled(height int64, req econtypes.QueryGetNetworkEnabledRequest) (econtypes.QueryGetNetworkEnabledResponse, error)
	GetAllNetworkEnabled(height int64, req econtypes.QueryAllNetworkEnabledRequest) (econtypes.QueryAllNetworkEnabledResponse, error)
	GetNbTxPerDay(height int64, req econtypes.QueryGetNbTxPerDayRequest) (econtypes.QueryGetNbTxPerDayResponse, error)
	GetAllNbTxPerDay(height int64, req econtypes.QueryAllNbTxPerDayRequest) (econtypes.QueryAllNbTxPerDayResponse, error)
	GetScheduledHookExecutionState(height int64, req econtypes.QueryGetScheduledHookExecutionStateRequest) (econtypes.QueryGetScheduledHookExecutionStateResponse, error)
	GetAllScheduledHookExecutionState(height int64, req econtypes.QueryAllScheduledHookExecutionStateRequest) (econtypes.QueryAllScheduledHookExecutionStateResponse, error)
}
