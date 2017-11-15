package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

// struct CustomOperation{}
// struct CustomJSONOperation{}
// struct CustomBinaryOperation{}
// struct VoteOperation{}
func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}

// struct CommentOperation{}
func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}

func (op *CommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeComment.Code()))
	if !op.IsStoryOperation() {
		enc.Encode(op.ParentAuthor)
	} else {
		enc.Encode(byte(0))
	}
	enc.Encode(op.ParentPermlink)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Title)
	enc.Encode(op.Body)
	enc.Encode(op.JsonMetadata)
	return enc.Err()
}

// struct TransferOperation{}
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// struct TransferToVestingOperation{}
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// struct WithdrawVestingOperation{}
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.EncodeMoney(op.VestingShares)
	return enc.Err()
}

// struct LimitOrderCreateOperation{}
func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.EncodeMoney(op.AmountToSell)
	enc.EncodeMoney(op.MinToReceive)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
}

// struct LimitOrderCancelOperation{}
func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
}

// struct FeedPublishOperation{}
func (op *FeedPublishOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFeedPublish.Code()))
	enc.Encode(op.Publisher)
	enc.EncodeMoney(op.ExchangeRate.Base)
	enc.EncodeMoney(op.ExchangeRate.Quote)
	return enc.Err()
}

// struct ConvertOperation{}
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// struct AccountCreateOperation{}
// struct AccountUpdateOperation{}
// struct WitnessUpdateOperation{}
// struct AccountWitnessVoteOperation{}
func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}

// struct AccountWitnessProxyOperation{}
func (op *AccountWitnessProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
}

// struct POWOperation{}
// struct ReportOverProductionOperation{}
// struct DeleteCommentOperation{}
func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}

// struct CommentOptionsOperation{}
func (op *CommentOptionsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommentOptions.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.EncodeMoney(op.MaxAcceptedPayout)
	enc.Encode(op.PercentSteemDollars)
	enc.EncodeBool(op.AllowVotes)
	enc.EncodeBool(op.AllowCurationRewards)
	enc.Encode(byte(0))
	return enc.Err()
}

// struct SetWithdrawVestingRouteOperation{}
func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}

// struct LimitOrderCreate2Operation{}
// struct ChallengeAuthorityOperation{}
// struct ProveAuthorityOperation{}
// struct RequestAccountRecoveryOperation{}
// struct RecoverAccountOperation{}
// struct ChangeRecoveryAccountOperation{}
func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	enc.Encode(byte(0))
	return enc.Err()
}

// struct EscrowTransferOperation{}
// struct EscrowDisputeOperation{}
// struct EscrowReleaseOperation{}
// struct POW2Operation{}
// struct EscrowApproveOperation{}
// struct TransferToSavingsOperation{}
func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// struct TransferFromSavingsOperation{}
func (op *TransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// struct CancelTransferFromSavingsOperation{}
func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	return enc.Err()
}

// struct DeclineVotingRightsOperation{}
func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}

// struct ResetAccountOperation{}
// struct SetResetAccountOperation{}
// struct CommentBenefactorRewardOperation{}
// struct DelegateVestingSharesOperation{}
// struct AccountCreateWithDelegationOperation{}
// struct CommentPayoutExtensionOperation{}
// struct AssetCreateOperation{}
// struct AssetUpdateOperation{}
// struct AssetUpdateBitassetOperation{}
// struct AssetUpdateFeedProducersOperation{}
// struct AssetIssueOperation{}
// struct AssetReserveOperation{}
// struct AssetFundFeePoolOperation{}
// struct AssetSettleOperation{}
// struct AssetForceSettleOperation{}
// struct AssetGlobalSettleOperation{}
// struct AssetPublishFeedOperation{}
// struct AssetClaimFeesOperation{}
// struct CallOrderUpdateOperation{}
// struct AccountWhitelistOperation{}
// struct OverrideTransferOperation{}
// struct ProposalCreateOperation{}
// struct ProposalUpdateOperation{}
// struct ProposalDeleteOperation{}
// struct BidCollateralOperation{}
