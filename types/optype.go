package types

// OpType represents a Steem operation type, i.e. vote, comment, pow and so on.
type OpType string

// Code returns the operation code associated with the given operation type.
func (kind OpType) Code() uint16 {
	return opCodes[kind]
}

const (
	TypeVote16                      OpType = "vote_16"
	TypeComment16                   OpType = "comment_16"
	TypeTransfer16                  OpType = "transfer_16"
	TypeTransferToVesting16         OpType = "transfer_to_vesting_16"
	TypeWithdrawVesting16           OpType = "withdraw_vesting_16"
	TypeLimitOrderCreate16          OpType = "limit_order_create_16"
	TypeLimitOrderCancel16          OpType = "limit_order_cancel_16"
	TypeFeedPublish16               OpType = "feed_publish_16"
	TypeConvert16                   OpType = "convert_16"
	TypeAccountCreate16             OpType = "account_create_16"
	TypeAccountUpdate16             OpType = "account_update_16"
	TypeWitnessUpdate16             OpType = "witness_update_16"
	TypeAccountWitnessVote16        OpType = "account_witness_vote_16"
	TypeAccountWitnessProxy16       OpType = "account_witness_proxy_16"
	TypePow16                       OpType = "pow_16"
	TypeCustom                      OpType = "custom"
	TypeReportOverProduction16      OpType = "report_over_production_16"
	TypeDeleteComment16             OpType = "delete_comment_16"
	TypeCustomJSON                  OpType = "custom_json"
	TypeCommentOptions16            OpType = "comment_options_16"
	TypeSetWithdrawVestingRoute16   OpType = "set_withdraw_vesting_route_16"
	TypeLimitOrderCreate216         OpType = "limit_order_create2_16"
	TypeChallengeAuthority16        OpType = "challenge_authority_16"
	TypeProveAuthority16            OpType = "prove_authority_16"
	TypeRequestAccountRecovery16    OpType = "request_account_recovery_16"
	TypeRecoverAccount16            OpType = "recover_account_16"
	TypeChangeRecoveryAccount16     OpType = "change_recovery_account_16"
	TypeEscrowTransfer16            OpType = "escrow_transfer_16"
	TypeEscrowDispute16             OpType = "escrow_dispute_16"
	TypeEscrowRelease16             OpType = "escrow_release_16"
	TypePow216                      OpType = "pow2_16"
	TypeEscrowApprove16             OpType = "escrow_approve_16"
	TypeTransferToSavings16         OpType = "transfer_to_savings_16"
	TypeTransferFromSavings16       OpType = "transfer_from_savings_16"
	TypeCancelTransferFromSavings16 OpType = "cancel_transfer_from_savings_16"
	TypeCustomBinary                OpType = "custom_binary"
	TypeDeclineVotingRights16       OpType = "decline_voting_rights_16"
	TypeResetAccount16              OpType = "reset_account_16"
	TypeSetResetAccount16           OpType = "set_reset_account_16"
	TypeCommentBenefactorReward16   OpType = "comment_benefactor_reward_16"
	TypeVote                        OpType = "vote"
	TypeComment                     OpType = "comment"
	TypeTransfer                    OpType = "transfer"
	TypeTransferToVesting           OpType = "transfer_to_vesting"
	TypeWithdrawVesting             OpType = "withdraw_vesting"
	TypeLimitOrderCreate            OpType = "limit_order_create"
	TypeLimitOrderCancel            OpType = "limit_order_cancel"
	TypeFeedPublish                 OpType = "feed_publish"
	TypeConvert                     OpType = "convert"
	TypeAccountCreate               OpType = "account_create"
	TypeAccountUpdate               OpType = "account_update"
	TypeWitnessUpdate               OpType = "witness_update"
	TypeAccountWitnessVote          OpType = "account_witness_vote"
	TypeAccountWitnessProxy         OpType = "account_witness_proxy"
	TypePOW                         OpType = "pow"
	TypeReportOverProduction        OpType = "report_over_production"
	TypeDeleteComment               OpType = "delete_comment"
	TypeCommentOptions              OpType = "comment_options"
	TypeSetWithdrawVestingRoute     OpType = "set_withdraw_vesting_route"
	TypeLimitOrderCreate2           OpType = "limit_order_create2"
	TypeChallengeAuthority          OpType = "challenge_authority"
	TypeProveAuthority              OpType = "prove_authority"
	TypeRequestAccountRecovery      OpType = "request_account_recovery"
	TypeRecoverAccount              OpType = "recover_account"
	TypeChangeRecoveryAccount       OpType = "change_recovery_account"
	TypeEscrowTransfer              OpType = "escrow_transfer"
	TypeEscrowDispute               OpType = "escrow_dispute"
	TypeEscrowRelease               OpType = "escrow_release"
	TypePOW2                        OpType = "pow2"
	TypeEscrowApprove               OpType = "escrow_approve"
	TypeTransferToSavings           OpType = "transfer_to_savings"
	TypeTransferFromSavings         OpType = "transfer_from_savings"
	TypeCancelTransferFromSavings   OpType = "cancel_transfer_from_savings"
	TypeDeclineVotingRights         OpType = "decline_voting_rights"
	TypeResetAccount                OpType = "reset_account"
	TypeSetResetAccount             OpType = "set_reset_account"
	TypeCommentBenefactorReward     OpType = "comment_benefactor_reward"
	TypeDelegateVestingShares       OpType = "delegate_vesting_shares"
	TypeAccountCreateWithDelegation OpType = "account_create_with_delegation"
	TypeCommentPayoutExtension      OpType = "comment_payout_extension"
	TypeAssetCreate                 OpType = "asset_create"
	TypeAssetUpdate                 OpType = "asset_update"
	TypeAssetUpdateBitasset         OpType = "asset_update_bitasset"
	TypeAssetUpdateFeedProducers    OpType = "asset_update_feed_producers"
	TypeAssetIssue                  OpType = "asset_issue"
	TypeAssetReserve                OpType = "asset_reserve"
	TypeAssetFundFeePool            OpType = "asset_fund_fee_pool"
	TypeAssetSettle                 OpType = "asset_settle"
	TypeAssetForceSettle            OpType = "asset_force_settle"
	TypeAssetGlobalSettle           OpType = "asset_global_settle"
	TypeAssetPublishFeed            OpType = "asset_publish_feed"
	TypeAssetClaimFees              OpType = "asset_claim_fees"
	TypeCallOrderUpdate             OpType = "call_order_update"
	TypeAccountWhitelist            OpType = "account_whitelist"
	TypeOverrideTransfer            OpType = "override_transfer"
	TypeProposalCreate              OpType = "proposal_create"
	TypeProposalUpdate              OpType = "proposal_update"
	TypeProposalDelete              OpType = "proposal_delete"
	TypeBidCollateral               OpType = "bid_collateral"
	TypeFillConvertRequest16        OpType = "fill_convert_request_16"       //Virtual
	TypeAuthorReward16              OpType = "author_reward_16"              //Virtual
	TypeCurationReward16            OpType = "curation_reward_16"            //Virtual
	TypeCommentReward16             OpType = "comment_reward_16"             //Virtual
	TypeLiquidityReward16           OpType = "liquidity_reward_16"           //Virtual
	TypeInterest16                  OpType = "interest_16"                   //Virtual
	TypeFillVestingWithdraw16       OpType = "fill_vesting_withdraw_16"      //Virtual
	TypeFillOrder16                 OpType = "fill_order_16"                 //Virtual
	TypeShutdownWitness16           OpType = "shutdown_witness_16"           //Virtual
	TypeFillTransferFromSavings16   OpType = "fill_transfer_from_savings_16" //Virtual
	TypeHardfork16                  OpType = "hardfork_16"                   //Virtual
	TypeCommentPayoutUpdate16       OpType = "comment_payout_update_16"      //Virtual
	TypeFillConvertRequest          OpType = "fill_convert_request"          //Virtual
	TypeAuthorReward                OpType = "author_reward"                 //Virtual
	TypeCurationReward              OpType = "curation_reward"               //Virtual
	TypeCommentReward               OpType = "comment_reward"                //Virtual
	TypeLiquidityReward             OpType = "liquidity_reward"              //Virtual
	TypeInterest                    OpType = "interest"                      //Virtual
	TypeFillVestingWithdraw         OpType = "fill_vesting_withdraw"         //Virtual
	TypeFillOrder                   OpType = "fill_order"                    //Virtual
	TypeShutdownWitness             OpType = "shutdown_witness"              //Virtual
	TypeFillTransferFromSavings     OpType = "fill_transfer_from_savings"    //Virtual
	TypeHardfork                    OpType = "hardfork"                      //Virtual
	TypeCommentPayoutUpdate         OpType = "comment_payout_update"         //Virtual
	TypeReturnVestingDelegation     OpType = "return_vesting_delegation"     //Virtual
	TypeAssetSettleCancel           OpType = "asset_settle_cancel"           //Virtual
	TypeFillCallOrder               OpType = "fill_call_order"               //Virtual
	TypeFillSettlementOrder         OpType = "fill_settlement_order"         //Virtual
	TypeExecuteBid                  OpType = "execute_bid"                   //Virtual
)

var opTypes = [...]OpType{
	TypeVote16,
	TypeComment16,
	TypeTransfer16,
	TypeTransferToVesting16,
	TypeWithdrawVesting16,
	TypeLimitOrderCreate16,
	TypeLimitOrderCancel16,
	TypeFeedPublish16,
	TypeConvert16,
	TypeAccountCreate16,
	TypeAccountUpdate16,
	TypeWitnessUpdate16,
	TypeAccountWitnessVote16,
	TypeAccountWitnessProxy16,
	TypePow16,
	TypeCustom,
	TypeReportOverProduction16,
	TypeDeleteComment16,
	TypeCustomJSON,
	TypeCommentOptions16,
	TypeSetWithdrawVestingRoute16,
	TypeLimitOrderCreate216,
	TypeChallengeAuthority16,
	TypeProveAuthority16,
	TypeRequestAccountRecovery16,
	TypeRecoverAccount16,
	TypeChangeRecoveryAccount16,
	TypeEscrowTransfer16,
	TypeEscrowDispute16,
	TypeEscrowRelease16,
	TypePow216,
	TypeEscrowApprove16,
	TypeTransferToSavings16,
	TypeTransferFromSavings16,
	TypeCancelTransferFromSavings16,
	TypeCustomBinary,
	TypeDeclineVotingRights16,
	TypeResetAccount16,
	TypeSetResetAccount16,
	TypeCommentBenefactorReward16,
	TypeVote,
	TypeComment,
	TypeTransfer,
	TypeTransferToVesting,
	TypeWithdrawVesting,
	TypeLimitOrderCreate,
	TypeLimitOrderCancel,
	TypeFeedPublish,
	TypeConvert,
	TypeAccountCreate,
	TypeAccountUpdate,
	TypeWitnessUpdate,
	TypeAccountWitnessVote,
	TypeAccountWitnessProxy,
	TypePOW,
	TypeReportOverProduction,
	TypeDeleteComment,
	TypeCommentOptions,
	TypeSetWithdrawVestingRoute,
	TypeLimitOrderCreate2,
	TypeChallengeAuthority,
	TypeProveAuthority,
	TypeRequestAccountRecovery,
	TypeRecoverAccount,
	TypeChangeRecoveryAccount,
	TypeEscrowTransfer,
	TypeEscrowDispute,
	TypeEscrowRelease,
	TypePOW2,
	TypeEscrowApprove,
	TypeTransferToSavings,
	TypeTransferFromSavings,
	TypeCancelTransferFromSavings,
	TypeDeclineVotingRights,
	TypeResetAccount,
	TypeSetResetAccount,
	TypeCommentBenefactorReward,
	TypeDelegateVestingShares,
	TypeAccountCreateWithDelegation,
	TypeCommentPayoutExtension,
	TypeAssetCreate,
	TypeAssetUpdate,
	TypeAssetUpdateBitasset,
	TypeAssetUpdateFeedProducers,
	TypeAssetIssue,
	TypeAssetReserve,
	TypeAssetFundFeePool,
	TypeAssetSettle,
	TypeAssetForceSettle,
	TypeAssetGlobalSettle,
	TypeAssetPublishFeed,
	TypeAssetClaimFees,
	TypeCallOrderUpdate,
	TypeAccountWhitelist,
	TypeOverrideTransfer,
	TypeProposalCreate,
	TypeProposalUpdate,
	TypeProposalDelete,
	TypeBidCollateral,
	TypeFillConvertRequest16,
	TypeAuthorReward16,
	TypeCurationReward16,
	TypeCommentReward16,
	TypeLiquidityReward16,
	TypeInterest16,
	TypeFillVestingWithdraw16,
	TypeFillOrder16,
	TypeShutdownWitness16,
	TypeFillTransferFromSavings16,
	TypeHardfork16,
	TypeCommentPayoutUpdate16,
	TypeFillConvertRequest,
	TypeAuthorReward,
	TypeCurationReward,
	TypeCommentReward,
	TypeLiquidityReward,
	TypeInterest,
	TypeFillVestingWithdraw,
	TypeFillOrder,
	TypeShutdownWitness,
	TypeFillTransferFromSavings,
	TypeHardfork,
	TypeCommentPayoutUpdate,
	TypeReturnVestingDelegation,
	TypeAssetSettleCancel,
	TypeFillCallOrder,
	TypeFillSettlementOrder,
	TypeExecuteBid,
}

// opCodes keeps mapping operation type -> operation code.
var opCodes map[OpType]uint16

func init() {
	opCodes = make(map[OpType]uint16, len(opTypes))
	for i, opType := range opTypes {
		opCodes[opType] = uint16(i)
	}
}
