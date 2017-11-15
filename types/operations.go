package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

// FC_REFLECT( steemit::chain::convert_operation,
//             (owner)
//             (requestid)
//             (amount) )

// FC_REFLECT( steemit::chain::feed_publish_operation,
//             (publisher)
//             (exchange_rate) )

// FC_REFLECT( steemit::chain::transfer_operation,
//             (from)
//             (to)
//             (amount)
//             (memo) )

// FC_REFLECT( steemit::chain::transfer_to_vesting_operation,
//             (from)
//             (to)
//             (amount) )

// FC_REFLECT( steemit::chain::withdraw_vesting_operation,
//             (account)
//             (vesting_shares) )

// FC_REFLECT( steemit::chain::account_witness_vote_operation,
//             (account)
//             (witness)(approve) )

// FC_REFLECT( steemit::chain::account_witness_proxy_operation,
//             (account)
//             (proxy) )

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

// FC_REFLECT( steemit::chain::vote_operation,
//             (voter)
//             (author)
//             (permlink)
//             (weight) )

// FC_REFLECT( steemit::chain::limit_order_create_operation,
//             (owner)
//             (orderid)
//             (amount_to_sell)
//             (min_to_receive)
//             (fill_or_kill)
//             (expiration) )

// FC_REFLECT( steemit::chain::limit_order_cancel_operation,
//             (owner)
//             (orderid) )

// FC_REFLECT( steemit::chain::delete_comment_operation,
//             (author)
//             (permlink) )

// FC_REFLECT( steemit::chain::comment_options_operation,
//             (author)
//             (permlink)
//             (max_accepted_payout)
//             (percent_steem_dollars)
//             (allow_votes)
//             (allow_curation_rewards)
//             (extensions) )
