package remote

import (
	v1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	v1beta1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/villagelabsco/juno/v4/node/remote"

	govsource "github.com/villagelabsco/bdjuno/v3/modules/gov/source"
)

var (
	_ govsource.Source = &Source{}
)

// Source implements govsource.Source using a remote node
type Source struct {
	*remote.Source
	govClient        v1govtypes.QueryClient
	govClientv1beta1 v1beta1govtypes.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, govClient v1govtypes.QueryClient, govClientv1beta1 v1beta1govtypes.QueryClient) *Source {
	return &Source{
		Source:           source,
		govClient:        govClient,
		govClientv1beta1: govClientv1beta1,
	}
}

// Proposal implements govsource.Source
func (s Source) Proposal(height int64, id uint64) (v1beta1govtypes.Proposal, error) {
	res, err := s.govClientv1beta1.Proposal(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1beta1govtypes.QueryProposalRequest{ProposalId: id},
	)
	if err != nil {
		return v1beta1govtypes.Proposal{}, err
	}

	return res.Proposal, err
}

// ProposalDeposit implements govsource.Source
func (s Source) ProposalDeposit(height int64, id uint64, depositor string) (*v1govtypes.Deposit, error) {
	res, err := s.govClient.Deposit(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1govtypes.QueryDepositRequest{ProposalId: id, Depositor: depositor},
	)
	if err != nil {
		return nil, err
	}

	return res.Deposit, nil
}

// TallyResult implements govsource.Source
func (s Source) TallyResult(height int64, proposalID uint64) (*v1govtypes.TallyResult, error) {
	res, err := s.govClient.TallyResult(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1govtypes.QueryTallyResultRequest{ProposalId: proposalID},
	)
	if err != nil {
		return nil, err
	}

	return res.Tally, nil
}

// DepositParams implements govsource.Source
func (s Source) DepositParams(height int64) (*v1govtypes.DepositParams, error) {
	res, err := s.govClient.Params(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1govtypes.QueryParamsRequest{ParamsType: v1govtypes.ParamDeposit},
	)
	if err != nil {
		return nil, err
	}

	return res.DepositParams, nil
}

// VotingParams implements govsource.Source
func (s Source) VotingParams(height int64) (*v1govtypes.VotingParams, error) {
	res, err := s.govClient.Params(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1govtypes.QueryParamsRequest{ParamsType: v1beta1govtypes.ParamVoting},
	)
	if err != nil {
		return nil, err
	}

	return res.VotingParams, nil
}

// TallyParams implements govsource.Source
func (s Source) TallyParams(height int64) (*v1govtypes.TallyParams, error) {
	res, err := s.govClient.Params(
		remote.GetHeightRequestContext(s.Ctx, height),
		&v1govtypes.QueryParamsRequest{ParamsType: v1beta1govtypes.ParamTallying},
	)
	if err != nil {
		return nil, err
	}

	return res.TallyParams, nil
}
