package client

import (
	"log"

	"github.com/pkg/errors"

	"github.com/asuleymanov/golos-go/encoding/wif"
	"github.com/asuleymanov/golos-go/types"
)

var (
	OpTypeKey = make(map[types.OpType][]string)
)

func init() {
	OpTypeKey["vote"] = []string{"posting"}
	OpTypeKey["comment"] = []string{"posting"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["limit_order_create"] = []string{"active"}
	OpTypeKey["limit_order_cancel"] = []string{"active"}
	OpTypeKey["feed_publish"] = []string{"active"}
	OpTypeKey["convert"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"posting"}
	OpTypeKey["account_witness_proxy"] = []string{"posting"}
	OpTypeKey["pow"] = []string{"active"}
	OpTypeKey["custom"] = []string{"active"}
	OpTypeKey["report_over_production"] = []string{"active"}
	OpTypeKey["delete_comment"] = []string{"posting"}
	OpTypeKey["custom_json"] = []string{"posting"}
	OpTypeKey["comment_options"] = []string{"posting"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["limit_order_create2"] = []string{"active"}
	OpTypeKey["challenge_authority"] = []string{"posting"}
	OpTypeKey["prove_authority"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"owner"}
	OpTypeKey["change_recovery_account"] = []string{"owner"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["pow2"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["transfer_to_savings"] = []string{"active"}
	OpTypeKey["transfer_from_savings"] = []string{"active"}
	OpTypeKey["cancel_transfer_from_savings"] = []string{"active"}
	OpTypeKey["custom_binary"] = []string{"posting"}
	OpTypeKey["decline_voting_rights"] = []string{"owner"}
	OpTypeKey["reset_account"] = []string{"active"}
	OpTypeKey["set_reset_account"] = []string{"posting"}
	OpTypeKey["comment_benefactor_reward"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create_with_delegation"] = []string{"active"}
	OpTypeKey["comment_payout_extension"] = []string{"active"}
	OpTypeKey["asset_create"] = []string{"active"}
	OpTypeKey["asset_update"] = []string{"active"}
	OpTypeKey["asset_update_bitasset"] = []string{"active"}
	OpTypeKey["asset_update_feed_producers"] = []string{"active"}
	OpTypeKey["asset_issue"] = []string{"active"}
	OpTypeKey["asset_reserve"] = []string{"active"}
	OpTypeKey["asset_fund_fee_pool"] = []string{"active"}
	OpTypeKey["asset_settle"] = []string{"active"}
	OpTypeKey["asset_force_settle"] = []string{"active"}
	OpTypeKey["asset_global_settle"] = []string{"active"}
	OpTypeKey["asset_publish_feed"] = []string{"active"}
	OpTypeKey["asset_claim_fees"] = []string{"active"}
	OpTypeKey["call_order_update"] = []string{"active"}
	OpTypeKey["account_whitelist"] = []string{"posting"}
	OpTypeKey["override_transfer"] = []string{"posting"}
	OpTypeKey["proposal_create"] = []string{"posting"}
	OpTypeKey["proposal_update"] = []string{"posting"}
	OpTypeKey["proposal_delete"] = []string{"posting"}
	OpTypeKey["bid_collateral"] = []string{"posting"}
}

func (api *Client) Signing_Keys(username string, trx types.Operation) [][]byte {
	var keys [][]byte
	if _, ok := Key_List[username]; ok {
		op_keys := OpTypeKey[trx.Type()]
		for _, val := range op_keys {
			switch {
			case val == "posting":
				privKey, err := wif.Decode(Key_List[username].PKey)
				if err != nil {
					log.Println(errors.Wrapf(err, "Error decode Key: "))
				}
				keys = append(keys, privKey)
			case val == "active":
				privKey, err := wif.Decode(Key_List[username].AKey)
				if err != nil {
					log.Println(errors.Wrapf(err, "Error decode Key: "))
				}
				keys = append(keys, privKey)
			case val == "owner":
				privKey, err := wif.Decode(Key_List[username].OKey)
				if err != nil {
					log.Println(errors.Wrapf(err, "Error decode Key: "))
				}
				keys = append(keys, privKey)
			case val == "memo":
				privKey, err := wif.Decode(Key_List[username].MKey)
				if err != nil {
					log.Println(errors.Wrapf(err, "Error decode Key: "))
				}
				keys = append(keys, privKey)
			}
		}
	} else {
		log.Println(errors.New("No user keys found"))
	}
	return keys
}
