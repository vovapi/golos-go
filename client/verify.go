package client

import (
	// Stdlib
	"fmt"
	"log"
	"strings"

	// Vendor
	"github.com/pkg/errors"
)

//We check whether there is a voter on the list of those who have already voted
func (api *Client) Verify_Voter_Weight(author, permlink, voter string, weight int) bool {
	ans, err := api.Rpc.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter Weight: "))
		return false
	} else {
		for _, v := range ans {
			if v.Voter == voter && v.Percent == weight {
				return true
			}
		}
		return false
	}
}

func (api *Client) Verify_Voter(author, permlink, voter string) bool {
	ans, err := api.Rpc.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	} else {
		for _, v := range ans {
			if v.Voter == voter {
				return true
			}
		}
		return false
	}
}

//We check whether there are voted
func (api *Client) Verify_Votes(author, permlink string) bool {
	ans, err := api.Rpc.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Votes: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

func (api *Client) Verify_Comments(author, permlink string) bool {
	ans, err := api.Rpc.Database.GetContentReplies(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Comments: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

func (api *Client) Verify_Reblogs(author, permlink, rebloger string) bool {
	ans, err := api.Rpc.Follow.GetRebloggedBy(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Reblogs: "))
		return false
	} else {
		for _, v := range ans {
			if v == rebloger {
				return true
			}
		}
		return false
	}
}

func (api *Client) Verify_Follow(follower, following string) bool {
	ans, err := api.Rpc.Follow.GetFollowing(follower, following, "blog", 1)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Follow: "))
		return false
	} else {
		for _, v := range ans {
			if (v.Follower == follower) && (v.Following == following) {
				return true
			} else {
				return false
			}
		}
		return false
	}
}

func (api *Client) Verify_Post(author, permlink string) bool {
	ans, err := api.Rpc.Database.GetContent(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Post: "))
		return false
	} else {
		if (ans.Author == author) && (ans.Permlink == permlink) {
			return true
		} else {
			return false
		}
		return false
	}
}

func (api *Client) Verify_Delegate_Posting_Key_Sign(username string, arrvote []ArrVote) []ArrVote {
	var truevote []ArrVote
	var arrusr, arrtmp []string

	for _, v := range arrvote {
		arrusr = append(arrusr, v.User)
	}

	acc, err := api.Rpc.Database.GetAccounts(arrusr)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Delegate Vote Sign: "))
		return nil
	} else {
		for _, val := range acc {
			for _, v := range val.Posting.AccountAuths {
				l := strings.Split(strings.Replace(strings.Replace(fmt.Sprintf("%v", v), "[", "", -1), "]", "", -1), " ")[0]
				if l == username {
					arrtmp = append(arrtmp, val.Name)
				}
			}
		}
	}

	for _, v := range arrtmp {
		for _, val := range arrvote {
			if v == val.User {
				truevote = append(truevote, val)
			}
		}
	}

	return truevote
}
