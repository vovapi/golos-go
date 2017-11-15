package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

// FC_REFLECT( steemit::chain::convert_operation,
//             (owner)
//             (requestid)
//             (amount) )
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::feed_publish_operation,
//             (publisher)
//             (exchange_rate) )
func (op *FeedPublishOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFeedPublish.Code()))
	enc.Encode(op.Publisher)
	enc.EncodeMoney(op.ExchangeRate.Base)
	enc.EncodeMoney(op.ExchangeRate.Quote)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::transfer_operation,
//             (from)
//             (to)
//             (amount)
//             (memo) )
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::transfer_to_vesting_operation,
//             (from)
//             (to)
//             (amount) )
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::withdraw_vesting_operation,
//             (account)
//             (vesting_shares) )
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.EncodeMoney(op.VestingShares)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::account_witness_vote_operation,
//             (account)
//             (witness)(approve) )
func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::account_witness_proxy_operation,
//             (account)
//             (proxy) )
func (op *AccountWitnessProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::comment_operation,
//             (parent_author)
//             (parent_permlink)
//             (author)
//             (permlink)
//             (title)
//             (body)
//             (json_metadata) )

// CommentOperation represents either a new post or a comment.
//
// In case Title is filled in and ParentAuthor is empty, it is a new post.
// The post category can be read from ParentPermlink.
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

// FC_REFLECT( steemit::chain::vote_operation,
//             (voter)
//             (author)
//             (permlink)
//             (weight) )
func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::limit_order_create_operation,
//             (owner)
//             (orderid)
//             (amount_to_sell)
//             (min_to_receive)
//             (fill_or_kill)
//             (expiration) )
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

// FC_REFLECT( steemit::chain::limit_order_cancel_operation,
//             (owner)
//             (orderid) )
func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::delete_comment_operation,
//             (author)
//             (permlink) )
func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::comment_options_operation,
//             (author)
//             (permlink)
//             (max_accepted_payout)
//             (percent_steem_dollars)
//             (allow_votes)
//             (allow_curation_rewards)
//             (extensions) )
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

func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}

func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	enc.Encode(byte(0))
	return enc.Err()
}

func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

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

func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	return enc.Err()
}

func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}
