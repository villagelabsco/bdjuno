package source

import (
	v1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	v1betagovtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type Source interface {
	Proposal(height int64, id uint64) (v1betagovtypes.Proposal, error)
	ProposalDeposit(height int64, id uint64, depositor string) (*v1govtypes.Deposit, error)
	TallyResult(height int64, proposalID uint64) (*v1govtypes.TallyResult, error)
	DepositParams(height int64) (*v1govtypes.DepositParams, error)
	VotingParams(height int64) (*v1govtypes.VotingParams, error)
	TallyParams(height int64) (*v1govtypes.TallyParams, error)
}
