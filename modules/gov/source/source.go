package source

import govbeta1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

type Source interface {
	Proposal(height int64, id uint64) (govbeta1types.Proposal, error)
	ProposalDeposit(height int64, id uint64, depositor string) (govbeta1types.Deposit, error)
	TallyResult(height int64, proposalID uint64) (govbeta1types.TallyResult, error)
	DepositParams(height int64) (govbeta1types.DepositParams, error)
	VotingParams(height int64) (govbeta1types.VotingParams, error)
	TallyParams(height int64) (govbeta1types.TallyParams, error)
}
