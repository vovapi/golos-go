package types

import (
	"encoding/json"
)

// Add-on struct

type ExchRate struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type POW struct {
	Worker    string `json:"worker"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Work      string `json:"work"`
}

type ChainProperties struct {
	AccountCreationFee string `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

type SignedBlockHeader struct {
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
	WitnessSignature      string        `json:"witness_signature"`
}

type Authority struct {
	AccountAuths    StringInt64Map `json:"account_auths"`
	KeyAuths        StringInt64Map `json:"key_auths"`
	WeightThreshold uint32         `json:"weight_threshold"`
}

type POW2Input struct {
	WorkerAccount string `json:"worker_account"`
	PrevBlock     []byte `json:"prev_block"`
	Nonce         uint64 `json:"nonce"`
}

type AssetOptions struct {
	MaxSupply            uint64        `json:"max_supply"`
	MarketFeePercent     uint16        `json:"market_fee_percent"`
	MaxMarketFee         uint64        `json:"max_market_fee"`
	IssuerPermissions    uint16        `json:"issuer_permissions"`
	Flags                uint16        `json:"flags"`
	CoreExchangeRate     ExchRate      `json:"core_exchange_rate"`
	WhitelistAuthorities []string      `json:"whitelist_authorities"`
	BlacklistAuthorities []string      `json:"blacklist_authorities"`
	WhitelistMarkets     []string      `json:"whitelist_markets"`
	BlacklistMarkets     []string      `json:"blacklist_markets"`
	Description          string        `json:"description"`
	Extensions           []interface{} `json:"extensions"`
}

type BitassetOptions struct {
	FeedLifetimeSec              uint32        `json:"feed_lifetime_sec"`
	MinimumFeeds                 uint8         `json:"minimum_feeds"`
	ForceSettlementDelaySec      uint32        `json:"force_settlement_delay_sec"`
	ForceSettlementOffsetPercent uint16        `json:"force_settlement_offset_percent"`
	MaximumForceSettlementVolume uint16        `json:"maximum_force_settlement_volume"`
	ShortBackingAsset            string        `json:"short_backing_asset"`
	Extensions                   []interface{} `json:"extensions"`
}

type PriceFeed struct {
	SettlementPrice            ExchRate `json:"settlement_price"`
	MaintenanceCollateralRatio uint16   `json:"maintenance_collateral_ratio"`
	MaximumShortSqueezeRatio   uint16   `json:"maximum_short_squeeze_ratio"`
	CoreExchangeRate           ExchRate `json:"core_exchange_rate"`
}

type Beneficiarie struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiarie `json:"beneficiaries"`
}

// struct Vote16Operation{}
type Vote16Operation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

func (op *Vote16Operation) Type() OpType {
	return TypeVote16
}

func (op *Vote16Operation) Data() interface{} {
	return op
}

// struct Comment16Operation{}
type Comment16Operation struct {
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Author         string `json:"author"`
	Permlink       string `json:"permlink"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	JsonMetadata   string `json:"json_metadata"`
}

func (op *Comment16Operation) Type() OpType {
	return TypeComment16
}

func (op *Comment16Operation) Data() interface{} {
	return op
}

// struct Transfer16Operation{}
type Transfer16Operation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *Transfer16Operation) Type() OpType {
	return TypeTransfer16
}

func (op *Transfer16Operation) Data() interface{} {
	return op
}

// struct TransferToVesting16Operation{}
type TransferToVesting16Operation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (op *TransferToVesting16Operation) Type() OpType {
	return TypeTransferToVesting16
}

func (op *TransferToVesting16Operation) Data() interface{} {
	return op
}

// struct WithdrawVesting16Operation{}
type WithdrawVesting16Operation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
}

func (op *WithdrawVesting16Operation) Type() OpType {
	return TypeWithdrawVesting16
}

func (op *WithdrawVesting16Operation) Data() interface{} {
	return op
}

// struct LimitOrderCreate16Operation{}
type LimitOrderCreate16Operation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"orderid"`
	AmountToSell string `json:"amount_to_sell"`
	MinToReceive string `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

func (op *LimitOrderCreate16Operation) Type() OpType {
	return TypeLimitOrderCreate16
}

func (op *LimitOrderCreate16Operation) Data() interface{} {
	return op
}

// struct LimitOrderCancel16Operation{}
type LimitOrderCancel16Operation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

func (op *LimitOrderCancel16Operation) Type() OpType {
	return TypeLimitOrderCancel16
}

func (op *LimitOrderCancel16Operation) Data() interface{} {
	return op
}

// struct FeedPublish16Operation{}
type FeedPublish16Operation struct {
	Publisher    string   `json:"publisher"`
	ExchangeRate ExchRate `json:"exchange_rate"`
}

func (op *FeedPublish16Operation) Type() OpType {
	return TypeFeedPublish16
}

func (op *FeedPublish16Operation) Data() interface{} {
	return op
}

// struct Convert16Operation{}
type Convert16Operation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"request_id"`
	Amount    string `json:"amount"`
}

func (op *Convert16Operation) Type() OpType {
	return TypeConvert16
}

func (op *Convert16Operation) Data() interface{} {
	return op
}

// struct AccountCreate16Operation{}
type AccountCreate16Operation struct {
	Fee            string     `json:"fee"`
	Creator        string     `json:"creator"`
	NewAccountName string     `json:"new_account_name"`
	Owner          *Authority `json:"owner"`
	Active         *Authority `json:"active"`
	Posting        *Authority `json:"posting"`
	MemoKey        string     `json:"memo_key"`
	JsonMetadata   string     `json:"json_metadata"`
}

func (op *AccountCreate16Operation) Type() OpType {
	return TypeAccountCreate16
}

func (op *AccountCreate16Operation) Data() interface{} {
	return op
}

// struct AccountUpdate16Operation{}
type AccountUpdate16Operation struct {
	Account      string     `json:"account"`
	Owner        *Authority `json:"owner"`
	Active       *Authority `json:"active"`
	Posting      *Authority `json:"posting"`
	MemoKey      string     `json:"memo_key"`
	JsonMetadata string     `json:"json_metadata"`
}

func (op *AccountUpdate16Operation) Type() OpType {
	return TypeAccountUpdate16
}

func (op *AccountUpdate16Operation) Data() interface{} {
	return op
}

// struct WitnessUpdate16Operation{}
type WitnessUpdate16Operation struct {
	Owner           string           `json:"owner"`
	Url             string           `json:"url"`
	BlockSigningKey string           `json:"block_signing_key"`
	Props           *ChainProperties `json:"props"`
	Fee             string           `json:"fee"`
}

func (op *WitnessUpdate16Operation) Type() OpType {
	return TypeWitnessUpdate16
}

func (op *WitnessUpdate16Operation) Data() interface{} {
	return op
}

// struct AccountWitnessVote16Operation{}
type AccountWitnessVote16Operation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

func (op *AccountWitnessVote16Operation) Type() OpType {
	return TypeAccountWitnessVote16
}

func (op *AccountWitnessVote16Operation) Data() interface{} {
	return op
}

// struct AccountWitnessProxy16Operation{}
type AccountWitnessProxy16Operation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

func (op *AccountWitnessProxy16Operation) Type() OpType {
	return TypeAccountWitnessProxy16
}

func (op *AccountWitnessProxy16Operation) Data() interface{} {
	return op
}

// struct Pow16Operation{}
type Pow16Operation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         uint64           `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

func (op *Pow16Operation) Type() OpType {
	return TypePow16
}

func (op *Pow16Operation) Data() interface{} {
	return op
}

// struct CustomOperation{}
type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	Id            uint16   `json:"id"`
	Datas         []byte   `json:"data"`
}

func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

func (op *CustomOperation) Data() interface{} {
	return op
}

// struct ReportOverProduction16Operation{}
type ReportOverProduction16Operation struct {
	Reporter    string             `json:"reporter"`
	FirstBlock  *SignedBlockHeader `json:"first_block"`
	SecondBlock *SignedBlockHeader `json:"second_block"`
}

func (op *ReportOverProduction16Operation) Type() OpType {
	return TypeReportOverProduction16
}

func (op *ReportOverProduction16Operation) Data() interface{} {
	return op
}

// struct DeleteComment16Operation{}
type DeleteComment16Operation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *DeleteComment16Operation) Type() OpType {
	return TypeDeleteComment16
}

func (op *DeleteComment16Operation) Data() interface{} {
	return op
}

// struct CustomJSONOperation{} in to file operation_custom_json.go

// struct CommentOptions16Operation{}
type CommentOptions16Operation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    string        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *CommentOptions16Operation) Type() OpType {
	return TypeCommentOptions16
}

func (op *CommentOptions16Operation) Data() interface{} {
	return op
}

// struct SetWithdrawVestingRoute16Operation{}
type SetWithdrawVestingRoute16Operation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

func (op *SetWithdrawVestingRoute16Operation) Type() OpType {
	return TypeSetWithdrawVestingRoute16
}

func (op *SetWithdrawVestingRoute16Operation) Data() interface{} {
	return op
}

// struct LimitOrderCreate216Operation{}
type LimitOrderCreate216Operation struct {
	Qwner        string   `json:"owner"`
	Orderid      uint32   `json:"orderid"`
	AmountToSell string   `json:"amount_to_sell"`
	ExchangeRate ExchRate `json:"exchange_rate"`
	FillOrKill   bool     `json:"fill_or_kill"`
	Expiration   uint32   `json:"expiration"`
}

func (op *LimitOrderCreate216Operation) Type() OpType {
	return TypeLimitOrderCreate216
}

func (op *LimitOrderCreate216Operation) Data() interface{} {
	return op
}

// struct ChallengeAuthority16Operation{}
type ChallengeAuthority16Operation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ChallengeAuthority16Operation) Type() OpType {
	return TypeChallengeAuthority16
}

func (op *ChallengeAuthority16Operation) Data() interface{} {
	return op
}

// struct ProveAuthority16Operation{}
type ProveAuthority16Operation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ProveAuthority16Operation) Type() OpType {
	return TypeProveAuthority16
}

func (op *ProveAuthority16Operation) Data() interface{} {
	return op
}

// struct RequestAccountRecovery16Operation{}
type RequestAccountRecovery16Operation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority []interface{} `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

func (op *RequestAccountRecovery16Operation) Type() OpType {
	return TypeRequestAccountRecovery16
}

func (op *RequestAccountRecovery16Operation) Data() interface{} {
	return op
}

// struct RecoverAccount16Operation{}
type RecoverAccount16Operation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    string        `json:"new_owner_authority"`
	RecentOwnerAuthority string        `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *RecoverAccount16Operation) Type() OpType {
	return TypeRecoverAccount16
}

func (op *RecoverAccount16Operation) Data() interface{} {
	return op
}

// struct ChangeRecoveryAccount16Operation{}
type ChangeRecoveryAccount16Operation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

func (op *ChangeRecoveryAccount16Operation) Type() OpType {
	return TypeChangeRecoveryAccount16
}

func (op *ChangeRecoveryAccount16Operation) Data() interface{} {
	return op
}

// struct EscrowTransfer16Operation{}
type EscrowTransfer16Operation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	SbdAmount            string `json:"sbd_amount"`
	SteemAmount          string `json:"steem_amount"`
	EscrowId             uint32 `json:"escrow_id"`
	Agent                string `json:"agent"`
	Fee                  string `json:"fee"`
	JsonMeta             string `json:"json_meta"`
	RatificationDeadline uint32 `json:"ratification_deadline"`
	EscrowExpiration     uint32 `json:"escrow_expiration"`
}

func (op *EscrowTransfer16Operation) Type() OpType {
	return TypeEscrowTransfer16
}

func (op *EscrowTransfer16Operation) Data() interface{} {
	return op
}

// struct EscrowDispute16Operation{}
type EscrowDispute16Operation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
}

func (op *EscrowDispute16Operation) Type() OpType {
	return TypeEscrowDispute16
}

func (op *EscrowDispute16Operation) Data() interface{} {
	return op
}

// struct EscrowRelease16Operation{}
type EscrowRelease16Operation struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Agent       string `json:"agent"`
	Who         string `json:"who"`
	Receiver    string `json:"receiver"`
	EscrowId    uint32 `json:"escrow_id"`
	SbdAmount   string `json:"sbd_amount"`
	SteemAmount string `json:"steem_amount"`
}

func (op *EscrowRelease16Operation) Type() OpType {
	return TypeEscrowRelease16
}

func (op *EscrowRelease16Operation) Data() interface{} {
	return op
}

// struct Pow216Operation{}
type Pow216Operation struct {
	Work        []interface{}    `json:"work"`
	NewOwnerKey string           `json:"new_owner_key"`
	Props       *ChainProperties `json:"props"`
}

func (op *Pow216Operation) Type() OpType {
	return TypePow216
}

func (op *Pow216Operation) Data() interface{} {
	return op
}

// struct EscrowApprove16Operation{}
type EscrowApprove16Operation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

func (op *EscrowApprove16Operation) Type() OpType {
	return TypeEscrowApprove16
}

func (op *EscrowApprove16Operation) Data() interface{} {
	return op
}

// struct TransferToSavings16Operation{}
type TransferToSavings16Operation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferToSavings16Operation) Type() OpType {
	return TypeTransferToSavings16
}

func (op *TransferToSavings16Operation) Data() interface{} {
	return op
}

// struct TransferFromSavings16Operation{}
type TransferFromSavings16Operation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Memo      string `json:"memo"`
}

func (op *TransferFromSavings16Operation) Type() OpType {
	return TypeTransferFromSavings16
}

func (op *TransferFromSavings16Operation) Data() interface{} {
	return op
}

// struct CancelTransferFromSavings16Operation{},
type CancelTransferFromSavings16Operation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
}

func (op *CancelTransferFromSavings16Operation) Type() OpType {
	return TypeCancelTransferFromSavings16
}

func (op *CancelTransferFromSavings16Operation) Data() interface{} {
	return op
}

// struct CustomBinaryOperation{}
type CustomBinaryOperation struct {
	RequiredOwnerAuths   []string `json:"required_owner_auths"`
	RequiredActiveAuths  []string `json:"required_active_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	RequiredAuths        []string `json:"required_auths"`
	Id                   string   `json:"id"`
	Datas                []byte   `json:"data"`
}

func (op *CustomBinaryOperation) Type() OpType {
	return TypeCustomBinary
}

func (op *CustomBinaryOperation) Data() interface{} {
	return op
}

// struct DeclineVotingRights16Operation{}
type DeclineVotingRights16Operation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

func (op *DeclineVotingRights16Operation) Type() OpType {
	return TypeDeclineVotingRights16
}

func (op *DeclineVotingRights16Operation) Data() interface{} {
	return op
}

// struct ResetAccount16Operation{}
type ResetAccount16Operation struct {
	ResetAccount      string `json:"reset_account"`
	AccountToReset    string `json:"Account_to_reset"`
	NewOwnerAuthority string `json:"new_owner_authority"`
}

func (op *ResetAccount16Operation) Type() OpType {
	return TypeResetAccount16
}

func (op *ResetAccount16Operation) Data() interface{} {
	return op
}

// struct SetResetAccount16Operation{}
type SetResetAccount16Operation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

func (op *SetResetAccount16Operation) Type() OpType {
	return TypeSetResetAccount16
}

func (op *SetResetAccount16Operation) Data() interface{} {
	return op
}

// struct CommentBenefactorReward16Operation{}
type CommentBenefactorReward16Operation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     string `json:"reward"`
}

func (op *CommentBenefactorReward16Operation) Type() OpType {
	return TypeCommentBenefactorReward16
}

func (op *CommentBenefactorReward16Operation) Data() interface{} {
	return op
}

// struct VoteOperation{}
type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

func (op *VoteOperation) Type() OpType {
	return TypeVote
}

func (op *VoteOperation) Data() interface{} {
	return op
}

// struct CommentOperation{}
type CommentOperation struct {
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Author         string `json:"author"`
	Permlink       string `json:"permlink"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	JsonMetadata   string `json:"json_metadata"`
}

func (op *CommentOperation) Type() OpType {
	return TypeComment
}

func (op *CommentOperation) Data() interface{} {
	return op
}

// struct TransferOperation{}
type TransferOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferOperation) Type() OpType {
	return TypeTransfer
}

func (op *TransferOperation) Data() interface{} {
	return op
}

// struct TransferToVestingOperation{}
type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (op *TransferToVestingOperation) Type() OpType {
	return TypeTransferToVesting
}

func (op *TransferToVestingOperation) Data() interface{} {
	return op
}

// struct WithdrawVestingOperation{}
type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
}

func (op *WithdrawVestingOperation) Type() OpType {
	return TypeWithdrawVesting
}

func (op *WithdrawVestingOperation) Data() interface{} {
	return op
}

// struct LimitOrderCreateOperation{}
type LimitOrderCreateOperation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"order_id"`
	AmountToSell string `json:"amount_to_sell"`
	MinToReceive string `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

func (op *LimitOrderCreateOperation) Type() OpType {
	return TypeLimitOrderCreate
}

func (op *LimitOrderCreateOperation) Data() interface{} {
	return op
}

// struct LimitOrderCancelOperation{}
type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"order_id"`
}

func (op *LimitOrderCancelOperation) Type() OpType {
	return TypeLimitOrderCancel
}

func (op *LimitOrderCancelOperation) Data() interface{} {
	return op
}

// struct FeedPublishOperation{}
type FeedPublishOperation struct {
	Publisher    string   `json:"publisher"`
	ExchangeRate ExchRate `json:"exchange_rate"`
}

func (op *FeedPublishOperation) Type() OpType {
	return TypeFeedPublish
}

func (op *FeedPublishOperation) Data() interface{} {
	return op
}

// struct ConvertOperation{}
type ConvertOperation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"request_id"`
	Amount    string `json:"amount"`
}

func (op *ConvertOperation) Type() OpType {
	return TypeConvert
}

func (op *ConvertOperation) Data() interface{} {
	return op
}

// struct AccountCreateOperation{}
type AccountCreateOperation struct {
	Fee            string     `json:"fee"`
	Creator        string     `json:"creator"`
	NewAccountName string     `json:"new_account_name"`
	Owner          *Authority `json:"owner"`
	Active         *Authority `json:"active"`
	Posting        *Authority `json:"posting"`
	MemoKey        string     `json:"memo_key"`
	JsonMetadata   string     `json:"json_metadata"`
}

func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

func (op *AccountCreateOperation) Data() interface{} {
	return op
}

// struct AccountUpdateOperation{}
type AccountUpdateOperation struct {
	Account      string     `json:"account"`
	Owner        *Authority `json:"owner"`
	Active       *Authority `json:"active"`
	Posting      *Authority `json:"posting"`
	MemoKey      string     `json:"memo_key"`
	JsonMetadata string     `json:"json_metadata"`
}

func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

// struct WitnessUpdateOperation{}
type WitnessUpdateOperation struct {
	Owner           string           `json:"owner"`
	Url             string           `json:"url"`
	BlockSigningKey string           `json:"block_signing_key"`
	Props           *ChainProperties `json:"props"`
	Fee             string           `json:"fee"`
}

func (op *WitnessUpdateOperation) Type() OpType {
	return TypeWitnessUpdate
}

func (op *WitnessUpdateOperation) Data() interface{} {
	return op
}

// struct AccountWitnessVoteOperation{}
type AccountWitnessVoteOperation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

func (op *AccountWitnessVoteOperation) Type() OpType {
	return TypeAccountWitnessVote
}

func (op *AccountWitnessVoteOperation) Data() interface{} {
	return op
}

// struct AccountWitnessProxyOperation{}
type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

func (op *AccountWitnessProxyOperation) Type() OpType {
	return TypeAccountWitnessProxy
}

func (op *AccountWitnessProxyOperation) Data() interface{} {
	return op
}

// struct POWOperation{}
type POWOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         uint64           `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

func (op *POWOperation) Type() OpType {
	return TypePOW
}

func (op *POWOperation) Data() interface{} {
	return op
}

// struct ReportOverProductionOperation{}
type ReportOverProductionOperation struct {
	Reporter    string             `json:"reporter"`
	FirstBlock  *SignedBlockHeader `json:"first_block"`
	SecondBlock *SignedBlockHeader `json:"second_block"`
}

func (op *ReportOverProductionOperation) Type() OpType {
	return TypeReportOverProduction
}

func (op *ReportOverProductionOperation) Data() interface{} {
	return op
}

// struct DeleteCommentOperation{}
type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *DeleteCommentOperation) Type() OpType {
	return TypeDeleteComment
}

func (op *DeleteCommentOperation) Data() interface{} {
	return op
}

// struct CommentOptionsOperation{}
type CommentOptionsOperation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    string        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *CommentOptionsOperation) Type() OpType {
	return TypeCommentOptions
}

func (op *CommentOptionsOperation) Data() interface{} {
	return op
}

// struct SetWithdrawVestingRouteOperation{}
type SetWithdrawVestingRouteOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

func (op *SetWithdrawVestingRouteOperation) Type() OpType {
	return TypeSetWithdrawVestingRoute
}

func (op *SetWithdrawVestingRouteOperation) Data() interface{} {
	return op
}

// struct LimitOrderCreate2Operation{}
type LimitOrderCreate2Operation struct {
	Qwner        string   `json:"owner"`
	Orderid      uint32   `json:"orderid"`
	AmountToSell string   `json:"amount_to_sell"`
	Expiration   uint32   `json:"expiration"`
	ExchangeRate ExchRate `json:"exchange_rate"`
	FillOrKill   bool     `json:"fill_or_kill"`
}

func (op *LimitOrderCreate2Operation) Type() OpType {
	return TypeLimitOrderCreate2
}

func (op *LimitOrderCreate2Operation) Data() interface{} {
	return op
}

// struct ChallengeAuthorityOperation{}
type ChallengeAuthorityOperation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ChallengeAuthorityOperation) Type() OpType {
	return TypeChallengeAuthority
}

func (op *ChallengeAuthorityOperation) Data() interface{} {
	return op
}

// struct ProveAuthorityOperation{}
type ProveAuthorityOperation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ProveAuthorityOperation) Type() OpType {
	return TypeProveAuthority
}

func (op *ProveAuthorityOperation) Data() interface{} {
	return op
}

// struct RequestAccountRecoveryOperation{}
type RequestAccountRecoveryOperation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority []interface{} `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

// struct RecoverAccountOperation{}
type RecoverAccountOperation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    string        `json:"new_owner_authority"`
	RecentOwnerAuthority string        `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

// struct ChangeRecoveryAccountOperation{}
type ChangeRecoveryAccountOperation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

func (op *ChangeRecoveryAccountOperation) Type() OpType {
	return TypeChangeRecoveryAccount
}

func (op *ChangeRecoveryAccountOperation) Data() interface{} {
	return op
}

// struct EscrowTransferOperation{}
type EscrowTransferOperation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	SbdAmount            string `json:"sbd_amount"`
	SteemAmount          string `json:"steem_amount"`
	EscrowId             uint32 `json:"escrow_id"`
	Agent                string `json:"agent"`
	Fee                  string `json:"fee"`
	JsonMeta             string `json:"json_meta"`
	RatificationDeadline uint32 `json:"ratification_deadline"`
	EscrowExpiration     uint32 `json:"escrow_expiration"`
}

func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

// struct EscrowDisputeOperation{}
type EscrowDisputeOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
}

func (op *EscrowDisputeOperation) Type() OpType {
	return TypeEscrowDispute
}

func (op *EscrowDisputeOperation) Data() interface{} {
	return op
}

// struct EscrowReleaseOperation{}
type EscrowReleaseOperation struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Agent       string `json:"agent"`
	Who         string `json:"who"`
	Receiver    string `json:"receiver"`
	EscrowId    uint32 `json:"escrow_id"`
	SbdAmount   string `json:"sbd_amount"`
	SteemAmount string `json:"steem_amount"`
}

func (op *EscrowReleaseOperation) Type() OpType {
	return TypeEscrowRelease
}

func (op *EscrowReleaseOperation) Data() interface{} {
	return op
}

// struct POW2Operation{}
type POW2Operation struct {
	Input      *POW2Input `json:"input"`
	PowSummary uint32     `json:"pow_summary"`
}

func (op *POW2Operation) Type() OpType {
	return TypePOW2
}

func (op *POW2Operation) Data() interface{} {
	return op
}

// struct EscrowApproveOperation{}
type EscrowApproveOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

func (op *EscrowApproveOperation) Type() OpType {
	return TypeEscrowApprove
}

func (op *EscrowApproveOperation) Data() interface{} {
	return op
}

// struct TransferToSavingsOperation{}
type TransferToSavingsOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferToSavingsOperation) Type() OpType {
	return TypeTransferToSavings
}

func (op *TransferToSavingsOperation) Data() interface{} {
	return op
}

// struct TransferFromSavingsOperation{}
type TransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Memo      string `json:"memo"`
}

func (op *TransferFromSavingsOperation) Type() OpType {
	return TypeTransferFromSavings
}

func (op *TransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct CancelTransferFromSavingsOperation{}
type CancelTransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
}

func (op *CancelTransferFromSavingsOperation) Type() OpType {
	return TypeCancelTransferFromSavings
}

func (op *CancelTransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct DeclineVotingRightsOperation{}
type DeclineVotingRightsOperation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

func (op *DeclineVotingRightsOperation) Type() OpType {
	return TypeDeclineVotingRights
}

func (op *DeclineVotingRightsOperation) Data() interface{} {
	return op
}

// struct ResetAccountOperation{}
type ResetAccountOperation struct {
	ResetAccount      string `json:"reset_account"`
	AccountToReset    string `json:"Account_to_reset"`
	NewOwnerAuthority string `json:"new_owner_authority"`
}

func (op *ResetAccountOperation) Type() OpType {
	return TypeResetAccount
}

func (op *ResetAccountOperation) Data() interface{} {
	return op
}

// struct SetResetAccountOperation{}
type SetResetAccountOperation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

func (op *SetResetAccountOperation) Type() OpType {
	return TypeSetResetAccount
}

func (op *SetResetAccountOperation) Data() interface{} {
	return op
}

// struct CommentBenefactorRewardOperation{}
type CommentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     string `json:"reward"`
}

func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}

// struct DelegateVestingSharesOperation{}
type DelegateVestingSharesOperation struct {
	Delegator     string `json:"delegator"`
	Delegatee     string `json:"delegatee"`
	VestingShares string `json:"vesting_shares"`
}

func (op *DelegateVestingSharesOperation) Type() OpType {
	return TypeDelegateVestingShares
}

func (op *DelegateVestingSharesOperation) Data() interface{} {
	return op
}

// struct AccountCreateWithDelegationOperation{}
type AccountCreateWithDelegationOperation struct {
	Fee            string        `json:"fee"`
	Delegation     string        `json:"delegation"`
	Creator        string        `json:"creator"`
	NewAccountName string        `json:"new_account_name"`
	Owner          string        `json:"owner"`
	Active         string        `json:"active"`
	Posting        string        `json:"posting"`
	MemoKey        string        `json:"memo_key"`
	JsonMetadata   string        `json:"json_metadata"`
	Extensions     []interface{} `json:"extensions"`
}

func (op *AccountCreateWithDelegationOperation) Type() OpType {
	return TypeAccountCreateWithDelegation
}

func (op *AccountCreateWithDelegationOperation) Data() interface{} {
	return op
}

// struct CommentPayoutExtensionOperation{}
type CommentPayoutExtensionOperation struct {
	Payer         string `json:"payer"`
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	ExtensionTime uint32 `json:"extension_time"`
	Amount        string `json:"amount"`
}

func (op *CommentPayoutExtensionOperation) Type() OpType {
	return TypeCommentPayoutExtension
}

func (op *CommentPayoutExtensionOperation) Data() interface{} {
	return op
}

// struct AssetCreateOperation{}
type AssetCreateOperation struct {
	Issuer             string          `json:"issuer"`
	AssetName          string          `json:"asset_name"`
	Precision          uint8           `json:"precision"`
	CommonOptions      AssetOptions    `json:"common_options"`
	BitassetOpts       BitassetOptions `json:"bitasset_opts"`
	IsPredictionMarket bool            `json:"is_prediction_market"`
	Extensions         []interface{}   `json:"extensions"`
}

func (op *AssetCreateOperation) Type() OpType {
	return TypeAssetCreate
}

func (op *AssetCreateOperation) Data() interface{} {
	return op
}

// struct AssetUpdateOperation{}
type AssetUpdateOperation struct {
	Issuer        string        `json:"issuer"`
	AssetToUpdate string        `json:"asset_to_update"`
	NewIssuer     string        `json:"new_issuer"`
	NewOptions    AssetOptions  `json:"new_options"`
	Extensions    []interface{} `json:"extensions"`
}

func (op *AssetUpdateOperation) Type() OpType {
	return TypeAssetUpdate
}

func (op *AssetUpdateOperation) Data() interface{} {
	return op
}

// struct AssetUpdateBitassetOperation{}
type AssetUpdateBitassetOperation struct {
	Issuer        string          `json:"issuer"`
	AssetToUpdate string          `json:"asset_to_update"`
	NewOptions    BitassetOptions `json:"new_options"`
	Extensions    []interface{}   `json:"extensions"`
}

func (op *AssetUpdateBitassetOperation) Type() OpType {
	return TypeAssetUpdateBitasset
}

func (op *AssetUpdateBitassetOperation) Data() interface{} {
	return op
}

// struct AssetUpdateFeedProducersOperation{}
type AssetUpdateFeedProducersOperation struct {
	Issuer           string        `json:"issuer"`
	AssetToUpdate    string        `json:"asset_to_update"`
	NewFeedProducers []string      `json:"new_feed_producers"`
	Extensions       []interface{} `json:"extensions"`
}

func (op *AssetUpdateFeedProducersOperation) Type() OpType {
	return TypeAssetUpdateFeedProducers
}

func (op *AssetUpdateFeedProducersOperation) Data() interface{} {
	return op
}

// struct AssetIssueOperation{}
type AssetIssueOperation struct {
	Issuer         string        `json:"issuer"`
	AssetToIssue   string        `json:"asset_to_issue"`
	IssueToAccount string        `json:"issue_to_account"`
	Memo           string        `json:"memo"`
	Extensions     []interface{} `json:"extensions"`
}

func (op *AssetIssueOperation) Type() OpType {
	return TypeAssetIssue
}

func (op *AssetIssueOperation) Data() interface{} {
	return op
}

// struct AssetReserveOperation{}
type AssetReserveOperation struct {
	Payer           string        `json:"payer"`
	AmountToReserve string        `json:"amount_to_reserve"`
	Extensions      []interface{} `json:"extensions"`
}

func (op *AssetReserveOperation) Type() OpType {
	return TypeAssetReserve
}

func (op *AssetReserveOperation) Data() interface{} {
	return op
}

// struct AssetFundFeePoolOperation{}
type AssetFundFeePoolOperation struct {
	FromAccount string        `json:"from_account"`
	AssetName   string        `json:"asset_name"`
	Amount      uint64        `json:"amount"`
	Extensions  []interface{} `json:"extensions"`
}

func (op *AssetFundFeePoolOperation) Type() OpType {
	return TypeAssetFundFeePool
}

func (op *AssetFundFeePoolOperation) Data() interface{} {
	return op
}

// struct AssetSettleOperation{}
type AssetSettleOperation struct {
	Account    string        `json:"account"`
	Amount     string        `json:"amount"`
	Extensions []interface{} `json:"extensions"`
}

func (op *AssetSettleOperation) Type() OpType {
	return TypeAssetSettle
}

func (op *AssetSettleOperation) Data() interface{} {
	return op
}

// struct AssetForceSettleOperation{}
type AssetForceSettleOperation struct {
	Account      string        `json:"account"`
	Amount       string        `json:"amount"`
	SettlementId uint32        `json:"settlement_id"`
	Extensions   []interface{} `json:"extensions"`
}

func (op *AssetForceSettleOperation) Type() OpType {
	return TypeAssetForceSettle
}

func (op *AssetForceSettleOperation) Data() interface{} {
	return op
}

// struct AssetGlobalSettleOperation{}
type AssetGlobalSettleOperation struct {
	Issuer        string        `json:"issuer"`
	AssetToSettle string        `json:"asset_to_settle"`
	SettlePrice   ExchRate      `json:"settle_price"`
	Extensions    []interface{} `json:"extensions"`
}

func (op *AssetGlobalSettleOperation) Type() OpType {
	return TypeAssetGlobalSettle
}

func (op *AssetGlobalSettleOperation) Data() interface{} {
	return op
}

// struct AssetPublishFeedOperation{}
type AssetPublishFeedOperation struct {
	Publisher  string        `json:"publisher"`
	AssetName  string        `json:"asset_name"`
	Feed       PriceFeed     `json:"feed"`
	Extensions []interface{} `json:"extensions"`
}

func (op *AssetPublishFeedOperation) Type() OpType {
	return TypeAssetPublishFeed
}

func (op *AssetPublishFeedOperation) Data() interface{} {
	return op
}

// struct AssetClaimFeesOperation{}
type AssetClaimFeesOperation struct {
	Issuer        string        `json:"issuer"`
	AmountToClaim string        `json:"amount_to_claim"`
	Extensions    []interface{} `json:"extensions"`
}

func (op *AssetClaimFeesOperation) Type() OpType {
	return TypeAssetClaimFees
}

func (op *AssetClaimFeesOperation) Data() interface{} {
	return op
}

// struct CallOrderUpdateOperation{}
type CallOrderUpdateOperation struct {
	FundingAccount  string `json:"funding_account"`
	DeltaCollateral string `json:"delta_collateral"`
	DeltaDebt       string `json:"delta_debt"`
}

func (op *CallOrderUpdateOperation) Type() OpType {
	return TypeCallOrderUpdate
}

func (op *CallOrderUpdateOperation) Data() interface{} {
	return op
}

// struct AccountWhitelistOperation{}
type AccountWhitelistOperation struct {
	Fee                string        `json:"fee"`
	AuthorizingAccount string        `json:"authorizing_account"`
	AccountToList      string        `json:"account_to_list"`
	NewListing         uint8         `json:"new_listing"`
	Extensions         []interface{} `json:"extensions"`
}

func (op *AccountWhitelistOperation) Type() OpType {
	return TypeAccountWhitelist
}

func (op *AccountWhitelistOperation) Data() interface{} {
	return op
}

// struct OverrideTransferOperation{}
type OverrideTransferOperation struct {
	Issuer     string        `json:"issuer"`
	From       string        `json:"from"`
	To         string        `json:"to"`
	Amount     string        `json:"amount"`
	Memo       string        `json:"memo"`
	Extensions []interface{} `json:"extensions"`
}

func (op *OverrideTransferOperation) Type() OpType {
	return TypeOverrideTransfer
}

func (op *OverrideTransferOperation) Data() interface{} {
	return op
}

// struct ProposalCreateOperation{}
type ProposalCreateOperation struct {
	Owner               string        `json:"owner"`
	ProposalId          uint32        `json:"proposal_id"`
	ExpirationTime      uint32        `json:"expiration_time"`
	ProposedOperations  []interface{} `json:"proposed_operations"`
	ReviewPeriodSeconds uint32        `json:"review_period_seconds"`
	Extensions          []interface{} `json:"extensions"`
}

func (op *ProposalCreateOperation) Type() OpType {
	return TypeProposalCreate
}

func (op *ProposalCreateOperation) Data() interface{} {
	return op
}

// struct ProposalUpdateOperation{}
type ProposalUpdateOperation struct {
	Owner                    string        `json:"owner"`
	ProposalId               uint32        `json:"proposal_id"`
	ActiveApprovalsToAdd     []string      `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove  []string      `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd      []string      `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove   []string      `json:"owner_approvals_to_remove"`
	PostingApprovalsToAdd    []string      `json:"posting_approvals_to_add"`
	PostingApprovalsToRemove []string      `json:"posting_approvals_to_remove"`
	KeyApprovalsToAdd        []string      `json:"key_approvals_to_add"`
	KeyApprovalsToRemove     []string      `json:"key_approvals_to_remove"`
	Extensions               []interface{} `json:"extensions"`
}

func (op *ProposalUpdateOperation) Type() OpType {
	return TypeProposalUpdate
}

func (op *ProposalUpdateOperation) Data() interface{} {
	return op
}

// struct ProposalDeleteOperation{}
type ProposalDeleteOperation struct {
	Owner               string        `json:"owner"`
	UsingOwnerAuthority bool          `json:"using_owner_authority"`
	ProposalId          uint32        `json:"proposal_id"`
	Extensions          []interface{} `json:"extensions"`
}

func (op *ProposalDeleteOperation) Type() OpType {
	return TypeProposalDelete
}

func (op *ProposalDeleteOperation) Data() interface{} {
	return op
}

// struct BidCollateralOperation{}
type BidCollateralOperation struct {
	Bidder               string        `json:"bidder"`
	AdditionalCollateral string        `json:"additional_collateral"`
	DebtCovered          string        `json:"debt_covered"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *BidCollateralOperation) Type() OpType {
	return TypeBidCollateral
}

func (op *BidCollateralOperation) Data() interface{} {
	return op
}

// struct FillConvertRequest16Operation{} //Virtual
type FillConvertRequest16Operation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  string `json:"amount_in"`
	AmountOut string `json:"amount_out"`
}

func (op *FillConvertRequest16Operation) Type() OpType {
	return TypeFillConvertRequest16
}

func (op *FillConvertRequest16Operation) Data() interface{} {
	return op
}

// struct AuthorReward16Operation{} //Virtual
type AuthorReward16Operation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     string `json:"sbd_payout"`
	SteemPayout   string `json:"steem_payout"`
	VestingPayout string `json:"vesting_payout"`
}

func (op *AuthorReward16Operation) Type() OpType {
	return TypeAuthorReward16
}

func (op *AuthorReward16Operation) Data() interface{} {
	return op
}

// struct CurationReward16Operation{} //Virtual
type CurationReward16Operation struct {
	Curator         string `json:"curator"`
	Reward          string `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

func (op *CurationReward16Operation) Type() OpType {
	return TypeCurationReward16
}

func (op *CurationReward16Operation) Data() interface{} {
	return op
}

// struct CommentReward16Operation{} //Virtual
type CommentReward16Operation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   string `json:"payout"`
}

func (op *CommentReward16Operation) Type() OpType {
	return TypeCommentReward16
}

func (op *CommentReward16Operation) Data() interface{} {
	return op
}

// struct LiquidityReward16Operation{} //Virtual
type LiquidityReward16Operation struct {
	Owner  string `json:"owner"`
	Payout string `json:"payout"`
}

func (op *LiquidityReward16Operation) Type() OpType {
	return TypeLiquidityReward16
}

func (op *LiquidityReward16Operation) Data() interface{} {
	return op
}

// struct Interest16Operation{} //Virtual
type Interest16Operation struct {
	Owner    string `json:"owner"`
	Interest string `json:"interest"`
}

func (op *Interest16Operation) Type() OpType {
	return TypeInterest16
}

func (op *Interest16Operation) Data() interface{} {
	return op
}

// struct FillVestingWithdraw16Operation{} //Virtual
type FillVestingWithdraw16Operation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Withdrawn   string `json:"withdrawn"`
	Deposited   string `json:"deposited"`
}

func (op *FillVestingWithdraw16Operation) Type() OpType {
	return TypeFillVestingWithdraw16
}

func (op *FillVestingWithdraw16Operation) Data() interface{} {
	return op
}

// struct FillOrder16Operation{} //Virtual
type FillOrder16Operation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    string `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       string `json:"open_pays"`
}

func (op *FillOrder16Operation) Type() OpType {
	return TypeFillOrder16
}

func (op *FillOrder16Operation) Data() interface{} {
	return op
}

// struct ShutdownWitness16Operation{} //Virtual
type ShutdownWitness16Operation struct {
	Owner string `json:"owner"`
}

func (op *ShutdownWitness16Operation) Type() OpType {
	return TypeShutdownWitness16
}

func (op *ShutdownWitness16Operation) Data() interface{} {
	return op
}

// struct FillTransferFromSavings16Operation{} //Virtual
type FillTransferFromSavings16Operation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	RequestId uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

func (op *FillTransferFromSavings16Operation) Type() OpType {
	return TypeFillTransferFromSavings16
}

func (op *FillTransferFromSavings16Operation) Data() interface{} {
	return op
}

// struct Hardfork16Operation{} //Virtual
type Hardfork16Operation struct {
	HardforkId uint32 `json:"hardfork_id"`
}

func (op *Hardfork16Operation) Type() OpType {
	return TypeHardfork16
}

func (op *Hardfork16Operation) Data() interface{} {
	return op
}

// struct CommentPayoutUpdate16Operation{} //Virtual
type CommentPayoutUpdate16Operation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *CommentPayoutUpdate16Operation) Type() OpType {
	return TypeCommentPayoutUpdate16
}

func (op *CommentPayoutUpdate16Operation) Data() interface{} {
	return op
}

// struct FillConvertRequestOperation{} //Virtual
type FillConvertRequestOperation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  string `json:"amount_in"`
	AmountOut string `json:"amount_out"`
}

func (op *FillConvertRequestOperation) Type() OpType {
	return TypeFillConvertRequest
}

func (op *FillConvertRequestOperation) Data() interface{} {
	return op
}

// struct AuthorRewardOperation{} //Virtual
type AuthorRewardOperation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     string `json:"sbd_payout"`
	SteemPayout   string `json:"steem_payout"`
	VestingPayout string `json:"vesting_payout"`
}

func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

func (op *AuthorRewardOperation) Data() interface{} {
	return op
}

// struct CurationRewardOperation{} //Virtual
type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          string `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

func (op *CurationRewardOperation) Data() interface{} {
	return op
}

// struct CommentRewardOperation{} //Virtual
type CommentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   string `json:"payout"`
}

func (op *CommentRewardOperation) Type() OpType {
	return TypeCommentReward
}

func (op *CommentRewardOperation) Data() interface{} {
	return op
}

// struct LiquidityRewardOperation{} //Virtual
type LiquidityRewardOperation struct {
	Owner  string `json:"owner"`
	Payout string `json:"payout"`
}

func (op *LiquidityRewardOperation) Type() OpType {
	return TypeLiquidityReward
}

func (op *LiquidityRewardOperation) Data() interface{} {
	return op
}

// struct InterestOperation{} //Virtual
type InterestOperation struct {
	Owner    string `json:"owner"`
	Interest string `json:"interest"`
}

func (op *InterestOperation) Type() OpType {
	return TypeInterest
}

func (op *InterestOperation) Data() interface{} {
	return op
}

// struct FillVestingWithdrawOperation{} //Virtual
type FillVestingWithdrawOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Withdrawn   string `json:"withdrawn"`
	Deposited   string `json:"deposited"`
}

func (op *FillVestingWithdrawOperation) Type() OpType {
	return TypeFillVestingWithdraw
}

func (op *FillVestingWithdrawOperation) Data() interface{} {
	return op
}

// struct FillOrderOperation{} //Virtual
type FillOrderOperation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    string `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       string `json:"open_pays"`
}

func (op *FillOrderOperation) Type() OpType {
	return TypeFillOrder
}

func (op *FillOrderOperation) Data() interface{} {
	return op
}

// struct ShutdownWitnessOperation{} //Virtual
type ShutdownWitnessOperation struct {
	Owner string `json:"owner"`
}

func (op *ShutdownWitnessOperation) Type() OpType {
	return TypeShutdownWitness
}

func (op *ShutdownWitnessOperation) Data() interface{} {
	return op
}

// struct FillTransferFromSavingsOperation{} //Virtual
type FillTransferFromSavingsOperation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	RequestId uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

func (op *FillTransferFromSavingsOperation) Type() OpType {
	return TypeFillTransferFromSavings
}

func (op *FillTransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct HardforkOperation{} //Virtual
type HardforkOperation struct {
	HardforkId uint32 `json:"hardfork_id"`
}

func (op *HardforkOperation) Type() OpType {
	return TypeHardfork
}

func (op *HardforkOperation) Data() interface{} {
	return op
}

// struct CommentPayoutUpdateOperation{} //Virtual
type CommentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *CommentPayoutUpdateOperation) Type() OpType {
	return TypeCommentPayoutUpdate
}

func (op *CommentPayoutUpdateOperation) Data() interface{} {
	return op
}

// struct ReturnVestingDelegationOperation{} //Virtual
type ReturnVestingDelegationOperation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
}

func (op *ReturnVestingDelegationOperation) Type() OpType {
	return TypeReturnVestingDelegation
}

func (op *ReturnVestingDelegationOperation) Data() interface{} {
	return op
}

// struct AssetSettleCancelOperation{} //Virtual
type AssetSettleCancelOperation struct {
	Settlement uint32        `json:"settlement"`
	Account    string        `json:"account"`
	Amount     string        `json:"amount"`
	Extensions []interface{} `json:"extensions"`
}

func (op *AssetSettleCancelOperation) Type() OpType {
	return TypeAssetSettleCancel
}

func (op *AssetSettleCancelOperation) Data() interface{} {
	return op
}

// struct FillCallOrderOperation{} //Virtual
type FillCallOrderOperation struct {
	OrderId  uint32 `json:"order_id"`
	Owner    string `json:"owner"`
	Pays     string `json:"pays"`
	Receives string `json:"receives"`
	Fee      string `json:"fee"`
}

func (op *FillCallOrderOperation) Type() OpType {
	return TypeFillCallOrder
}

func (op *FillCallOrderOperation) Data() interface{} {
	return op
}

// struct FillSettlementOrderOperation{} //Virtual
type FillSettlementOrderOperation struct {
	OrderId  uint32 `json:"order_id"`
	Owner    string `json:"owner"`
	Pays     string `json:"pays"`
	Receives string `json:"receives"`
	Fee      string `json:"fee"`
}

func (op *FillSettlementOrderOperation) Type() OpType {
	return TypeFillSettlementOrder
}

func (op *FillSettlementOrderOperation) Data() interface{} {
	return op
}

// struct ExecuteBidOperation{} //Virtual
type ExecuteBidOperation struct {
	Bidder     string `json:"bidder"`
	Debt       string `json:"debt"`
	Collateral string `json:"collateral"`
}

func (op *ExecuteBidOperation) Type() OpType {
	return TypeExecuteBid
}

func (op *ExecuteBidOperation) Data() interface{} {
	return op
}

type UnknownOperation struct {
	kind OpType
	data *json.RawMessage
}

func (op *UnknownOperation) Type() OpType {
	return op.kind
}

func (op *UnknownOperation) Data() interface{} {
	return op.data
}
