package local

import (
	"fmt"
	v1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v1beta1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/villagelabsco/juno/v4/node/local"

	govsource "github.com/villagelabsco/bdjuno/v3/modules/gov/source"
)

var (
	_ govsource.Source = &Source{}
)

// Source implements govsource.Source by using a local node
type Source struct {
	*local.Source
	q       v1govtypes.QueryServer
	qv1beta v1beta1govtypes.QueryServer
}

// NewSource returns a new Source instance
func NewSource(source *local.Source, govKeeper v1govtypes.QueryServer, govKeeperv1beta1 v1beta1govtypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		q:       govKeeper,
		qv1beta: govKeeperv1beta1,
	}
}

// Proposal implements govsource.Source
func (s Source) Proposal(height int64, id uint64) (v1beta1govtypes.Proposal, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return v1beta1govtypes.Proposal{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.qv1beta.Proposal(sdk.WrapSDKContext(ctx), &v1beta1govtypes.QueryProposalRequest{ProposalId: id})
	if err != nil {
		return v1beta1govtypes.Proposal{}, err
	}

	return res.Proposal, nil
}

// ProposalDeposit implements govsource.Source
func (s Source) ProposalDeposit(height int64, id uint64, depositor string) (*v1govtypes.Deposit, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Deposit(sdk.WrapSDKContext(ctx), &v1govtypes.QueryDepositRequest{ProposalId: id, Depositor: depositor})
	if err != nil {
		return nil, err
	}

	return res.Deposit, nil
}

// TallyResult implements govsource.Source
func (s Source) TallyResult(height int64, proposalID uint64) (*v1govtypes.TallyResult, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.TallyResult(sdk.WrapSDKContext(ctx), &v1govtypes.QueryTallyResultRequest{ProposalId: proposalID})
	if err != nil {
		return nil, err
	}

	return res.Tally, nil
}

// DepositParams implements govsource.Source
func (s Source) DepositParams(height int64) (*v1govtypes.DepositParams, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(sdk.WrapSDKContext(ctx), &v1govtypes.QueryParamsRequest{ParamsType: v1beta1govtypes.ParamDeposit})
	if err != nil {
		return nil, err
	}

	return res.DepositParams, nil
}

// VotingParams implements govsource.Source
func (s Source) VotingParams(height int64) (*v1govtypes.VotingParams, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(sdk.WrapSDKContext(ctx), &v1govtypes.QueryParamsRequest{ParamsType: v1beta1govtypes.ParamVoting})
	if err != nil {
		return nil, err
	}

	return res.VotingParams, nil
}

// TallyParams implements govsource.Source
func (s Source) TallyParams(height int64) (*v1govtypes.TallyParams, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(sdk.WrapSDKContext(ctx), &v1govtypes.QueryParamsRequest{ParamsType: v1beta1govtypes.ParamTallying})
	if err != nil {
		return nil, err
	}

	return res.TallyParams, nil
}
