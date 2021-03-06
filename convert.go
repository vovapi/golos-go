package client

import (
	"strconv"
)

//SbdMedianPrice returns the average cost of GBG when converting GOLOS.
func (client *Client) SbdMedianPrice() (float64, error) {
	smpreq, errsmp := client.Witness.GetFeedHistory()
	if errsmp != nil {
		return 0, errsmp
	}

	base := smpreq.CurrentMedianHistory.Base.Amount
	quote := smpreq.CurrentMedianHistory.Quote.Amount
	smptmp := base / quote

	str := strconv.FormatFloat(smptmp, 'f', 3, 64)
	smp, errsmp := strconv.ParseFloat(str, 64)
	if errsmp != nil {
		return 0, errsmp
	}

	return smp, nil
}

//SteemPerMvest returns the ratio of TotalVersingFundSteem to TotalVestingShares.
func (client *Client) SteemPerMvest() (float64, error) {
	dgp, errdgp := client.Database.GetDynamicGlobalProperties()
	if errdgp != nil {
		return 0, errdgp
	}

	tvfs := dgp.TotalVersingFundSteem.Amount
	tvs := dgp.TotalVestingShares.Amount

	spmtmp := (tvfs / tvs) * 1000000

	str := strconv.FormatFloat(spmtmp, 'f', 3, 64)
	spm, errspm := strconv.ParseFloat(str, 64)
	if errspm != nil {
		return 0, errspm
	}

	return spm, nil
}
