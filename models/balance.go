package models

type (
	Balance struct {
		Gorm
		Balance        uint64 `json:"balance"`
		BalanceAchieve uint64 `json:"balance_achieve"`
		UserId         uint   `json:"user_id"`
	}

	BalanceHistory struct {
		Gorm
		UserBalanceId uint   `json:"user_balance_id"`
		BalanceBefore uint64 `json:"balance_before"`
		BalanceAfter  uint64 `json:"balance_after"`
		Activity      string `json:"activity"`
		Type          string `json:"type"`
		Ip            string `json:"ip"`
		Location      string `json:"location"`
		UserAgent     string `json:"user_agent"`
		Author        string `json:"author"`
	}

	TopUpRequest struct {
		CodeAccountBank string `json:"code" form:"code"`
		Amount          uint64 `json:"amount" form:"amount"`
	}

	TransferBalance struct {
		Amount         uint64 `json:"amount" form:"amount"`
		EmailRecipient string `json:"email_recepient"`
	}
)
