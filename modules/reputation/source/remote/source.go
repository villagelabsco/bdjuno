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
	reputationtypes "github.com/villagelabs/villaged/x/reputation/types"
)

type Source struct {
	*remote.Source
	q reputationtypes.QueryClient
}

func (s Source) GetParams(height int64) (reputationtypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Params(ctx, &reputationtypes.QueryParamsRequest{})
	if err != nil {
		return reputationtypes.Params{}, fmt.Errorf("error while getting params: %s", err)
	}
	return res.Params, nil
}

func (s Source) GetFeedback(height int64, req reputationtypes.QueryGetFeedbackRequest) (reputationtypes.QueryGetFeedbackResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.Feedback(ctx, &req)
	if err != nil {
		return reputationtypes.QueryGetFeedbackResponse{}, fmt.Errorf("error while getting feedback: %s", err)
	}
	return *res, nil
}

func (s Source) GetAllFeedback(height int64, req reputationtypes.QueryAllFeedbackRequest) (reputationtypes.QueryAllFeedbackResponse, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)
	res, err := s.q.FeedbackAll(ctx, &req)
	if err != nil {
		return reputationtypes.QueryAllFeedbackResponse{}, fmt.Errorf("error while getting all feedback: %s", err)
	}
	return *res, nil
}
