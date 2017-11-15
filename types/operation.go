package types

import (
	// Stdlib
	"bytes"
	"encoding/json"
	"reflect"

	// Vendor
	"github.com/pkg/errors"
)

// dataObjects keeps mapping operation type -> operation data object.
// This is used later on to unmarshal operation data based on the operation type.
var dataObjects = map[OpType]Operation{
	TypeVote16:                      &Vote16Operation{},
	TypeComment16:                   &Comment16Operation{},
	TypeTransfer16:                  &Transfer16Operation{},
	TypeTransferToVesting16:         &TransferToVesting16Operation{},
	TypeWithdrawVesting16:           &WithdrawVesting16Operation{},
	TypeLimitOrderCreate16:          &LimitOrderCreate16Operation{},
	TypeLimitOrderCancel16:          &LimitOrderCancel16Operation{},
	TypeFeedPublish16:               &FeedPublish16Operation{},
	TypeConvert16:                   &Convert16Operation{},
	TypeAccountCreate16:             &AccountCreate16Operation{},
	TypeAccountUpdate16:             &AccountUpdate16Operation{},
	TypeWitnessUpdate16:             &WitnessUpdate16Operation{},
	TypeAccountWitnessVote16:        &AccountWitnessVote16Operation{},
	TypeAccountWitnessProxy16:       &AccountWitnessProxy16Operation{},
	TypePow16:                       &Pow16Operation{},
	TypeCustom:                      &CustomOperation{},
	TypeReportOverProduction16:      &ReportOverProduction16Operation{},
	TypeDeleteComment16:             &DeleteComment16Operation{},
	TypeCustomJSON:                  &CustomJSONOperation{},
	TypeCommentOptions16:            &CommentOptions16Operation{},
	TypeSetWithdrawVestingRoute16:   &SetWithdrawVestingRoute16Operation{},
	TypeLimitOrderCreate216:         &LimitOrderCreate216Operation{},
	TypeChallengeAuthority16:        &ChallengeAuthority16Operation{},
	TypeProveAuthority16:            &ProveAuthority16Operation{},
	TypeRequestAccountRecovery16:    &RequestAccountRecovery16Operation{},
	TypeRecoverAccount16:            &RecoverAccount16Operation{},
	TypeChangeRecoveryAccount16:     &ChangeRecoveryAccount16Operation{},
	TypeEscrowTransfer16:            &EscrowTransfer16Operation{},
	TypeEscrowDispute16:             &EscrowDispute16Operation{},
	TypeEscrowRelease16:             &EscrowRelease16Operation{},
	TypePow216:                      &Pow216Operation{},
	TypeEscrowApprove16:             &EscrowApprove16Operation{},
	TypeTransferToSavings16:         &TransferToSavings16Operation{},
	TypeTransferFromSavings16:       &TransferFromSavings16Operation{},
	TypeCancelTransferFromSavings16: &CancelTransferFromSavings16Operation{},
	TypeCustomBinary:                &CustomBinaryOperation{},
	TypeDeclineVotingRights16:       &DeclineVotingRights16Operation{},
	TypeResetAccount16:              &ResetAccount16Operation{},
	TypeSetResetAccount16:           &SetResetAccount16Operation{},
	TypeCommentBenefactorReward16:   &CommentBenefactorReward16Operation{},
	TypeVote:                        &VoteOperation{},
	TypeComment:                     &CommentOperation{},
	TypeTransfer:                    &TransferOperation{},
	TypeTransferToVesting:           &TransferToVestingOperation{},
	TypeWithdrawVesting:             &WithdrawVestingOperation{},
	TypeLimitOrderCreate:            &LimitOrderCreateOperation{},
	TypeLimitOrderCancel:            &LimitOrderCancelOperation{},
	TypeFeedPublish:                 &FeedPublishOperation{},
	TypeConvert:                     &ConvertOperation{},
	TypeAccountCreate:               &AccountCreateOperation{},
	TypeAccountUpdate:               &AccountUpdateOperation{},
	TypeWitnessUpdate:               &WitnessUpdateOperation{},
	TypeAccountWitnessVote:          &AccountWitnessVoteOperation{},
	TypeAccountWitnessProxy:         &AccountWitnessProxyOperation{},
	TypePOW:                         &POWOperation{},
	TypeReportOverProduction:        &ReportOverProductionOperation{},
	TypeDeleteComment:               &DeleteCommentOperation{},
	TypeCommentOptions:              &CommentOptionsOperation{},
	TypeSetWithdrawVestingRoute:     &SetWithdrawVestingRouteOperation{},
	TypeLimitOrderCreate2:           &LimitOrderCreate2Operation{},
	TypeChallengeAuthority:          &ChallengeAuthorityOperation{},
	TypeProveAuthority:              &ProveAuthorityOperation{},
	TypeRequestAccountRecovery:      &RequestAccountRecoveryOperation{},
	TypeRecoverAccount:              &RecoverAccountOperation{},
	TypeChangeRecoveryAccount:       &ChangeRecoveryAccountOperation{},
	TypeEscrowTransfer:              &EscrowTransferOperation{},
	TypeEscrowDispute:               &EscrowDisputeOperation{},
	TypeEscrowRelease:               &EscrowReleaseOperation{},
	TypePOW2:                        &POW2Operation{},
	TypeEscrowApprove:               &EscrowApproveOperation{},
	TypeTransferToSavings:           &TransferToSavingsOperation{},
	TypeTransferFromSavings:         &TransferFromSavingsOperation{},
	TypeCancelTransferFromSavings:   &CancelTransferFromSavingsOperation{},
	TypeDeclineVotingRights:         &DeclineVotingRightsOperation{},
	TypeResetAccount:                &ResetAccountOperation{},
	TypeSetResetAccount:             &SetResetAccountOperation{},
	TypeCommentBenefactorReward:     &CommentBenefactorRewardOperation{},
	TypeDelegateVestingShares:       &DelegateVestingSharesOperation{},
	TypeAccountCreateWithDelegation: &AccountCreateWithDelegationOperation{},
	TypeCommentPayoutExtension:      &CommentPayoutExtensionOperation{},
	TypeAssetCreate:                 &AssetCreateOperation{},
	TypeAssetUpdate:                 &AssetUpdateOperation{},
	TypeAssetUpdateBitasset:         &AssetUpdateBitassetOperation{},
	TypeAssetUpdateFeedProducers:    &AssetUpdateFeedProducersOperation{},
	TypeAssetIssue:                  &AssetIssueOperation{},
	TypeAssetReserve:                &AssetReserveOperation{},
	TypeAssetFundFeePool:            &AssetFundFeePoolOperation{},
	TypeAssetSettle:                 &AssetSettleOperation{},
	TypeAssetForceSettle:            &AssetForceSettleOperation{},
	TypeAssetGlobalSettle:           &AssetGlobalSettleOperation{},
	TypeAssetPublishFeed:            &AssetPublishFeedOperation{},
	TypeAssetClaimFees:              &AssetClaimFeesOperation{},
	TypeCallOrderUpdate:             &CallOrderUpdateOperation{},
	TypeAccountWhitelist:            &AccountWhitelistOperation{},
	TypeOverrideTransfer:            &OverrideTransferOperation{},
	TypeProposalCreate:              &ProposalCreateOperation{},
	TypeProposalUpdate:              &ProposalUpdateOperation{},
	TypeProposalDelete:              &ProposalDeleteOperation{},
	TypeBidCollateral:               &BidCollateralOperation{},
	TypeFillConvertRequest16:        &FillConvertRequest16Operation{},      //Virtual
	TypeAuthorReward16:              &AuthorReward16Operation{},            //Virtual
	TypeCurationReward16:            &CurationReward16Operation{},          //Virtual
	TypeCommentReward16:             &CommentReward16Operation{},           //Virtual
	TypeLiquidityReward16:           &LiquidityReward16Operation{},         //Virtual
	TypeInterest16:                  &Interest16Operation{},                //Virtual
	TypeFillVestingWithdraw16:       &FillVestingWithdraw16Operation{},     //Virtual
	TypeFillOrder16:                 &FillOrder16Operation{},               //Virtual
	TypeShutdownWitness16:           &ShutdownWitness16Operation{},         //Virtual
	TypeFillTransferFromSavings16:   &FillTransferFromSavings16Operation{}, //Virtual
	TypeHardfork16:                  &Hardfork16Operation{},                //Virtual
	TypeCommentPayoutUpdate16:       &CommentPayoutUpdate16Operation{},     //Virtual
	TypeFillConvertRequest:          &FillConvertRequestOperation{},        //Virtual
	TypeAuthorReward:                &AuthorRewardOperation{},              //Virtual
	TypeCurationReward:              &CurationRewardOperation{},            //Virtual
	TypeCommentReward:               &CommentRewardOperation{},             //Virtual
	TypeLiquidityReward:             &LiquidityRewardOperation{},           //Virtual
	TypeInterest:                    &InterestOperation{},                  //Virtual
	TypeFillVestingWithdraw:         &FillVestingWithdrawOperation{},       //Virtual
	TypeFillOrder:                   &FillOrderOperation{},                 //Virtual
	TypeShutdownWitness:             &ShutdownWitnessOperation{},           //Virtual
	TypeFillTransferFromSavings:     &FillTransferFromSavingsOperation{},   //Virtual
	TypeHardfork:                    &HardforkOperation{},                  //Virtual
	TypeCommentPayoutUpdate:         &CommentPayoutUpdateOperation{},       //Virtual
	TypeReturnVestingDelegation:     &ReturnVestingDelegationOperation{},   //Virtual
	TypeAssetSettleCancel:           &AssetSettleCancelOperation{},         //Virtual
	TypeFillCallOrder:               &FillCallOrderOperation{},             //Virtual
	TypeFillSettlementOrder:         &FillSettlementOrderOperation{},       //Virtual
	TypeExecuteBid:                  &ExecuteBidOperation{},                //Virtual
}

// Operation represents an operation stored in a transaction.
type Operation interface {
	// Type returns the operation type as present in the operation object, element [0].
	Type() OpType

	// Data returns the operation data as present in the operation object, element [1].
	//
	// When the operation type is known to this package, this field contains
	// the operation data object associated with the given operation type,
	// e.g. Type is TypeVote -> Data contains *VoteOperation.
	// Otherwise this field contains raw JSON (type *json.RawMessage).
	Data() interface{}
}

type Operations []Operation

func (ops *Operations) UnmarshalJSON(data []byte) error {
	var tuples []*operationTuple
	if err := json.Unmarshal(data, &tuples); err != nil {
		return err
	}

	items := make([]Operation, 0, len(tuples))
	for _, tuple := range tuples {
		items = append(items, tuple.Data)
	}

	*ops = items
	return nil
}

/*func (ops Operations) MarshalJSON() ([]byte, error) {
	tuples := make([]*operationTuple, 0, len(ops))
	for _, op := range ops {
		tuples = append(tuples, &operationTuple{
			Type: op.Type(),
			Data: op.Data().(Operation),
		})
	}
	return json.Marshal(tuples)
}*/

func (ops Operations) MarshalJSON() ([]byte, error) {
	tuples := make([]*operationTuple, 0, len(ops))
	for _, op := range ops {
		tuples = append(tuples, &operationTuple{
			Type: op.Type(),
			Data: op.Data().(Operation),
		})
	}
	return JSONMarshal(tuples)
}

type operationTuple struct {
	Type OpType
	Data Operation
}

/*func (op *operationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		op.Type,
		op.Data,
	})
}*/

func (op *operationTuple) MarshalJSON() ([]byte, error) {
	return JSONMarshal([]interface{}{
		op.Type,
		op.Data,
	})
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func (op *operationTuple) UnmarshalJSON(data []byte) error {
	// The operation object is [opType, opBody].
	raw := make([]*json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Wrapf(err, "failed to unmarshal operation object: %v", string(data))
	}
	if len(raw) != 2 {
		return errors.Errorf("invalid operation object: %v", string(data))
	}

	// Unmarshal the type.
	var opType OpType
	if err := json.Unmarshal(*raw[0], &opType); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Operation.Type: %v", string(*raw[0]))
	}

	// Unmarshal the data.
	var opData Operation
	template, ok := dataObjects[opType]
	if ok {
		opData = reflect.New(
			reflect.Indirect(reflect.ValueOf(template)).Type(),
		).Interface().(Operation)

		if err := json.Unmarshal(*raw[1], opData); err != nil {
			return errors.Wrapf(err, "failed to unmarshal Operation.Data: %v", string(*raw[1]))
		}
	} else {
		opData = &UnknownOperation{opType, raw[1]}
	}

	// Update fields.
	op.Type = opType
	op.Data = opData
	return nil
}
