package models

type (
	Bank struct {
		Gorm
		Balance        uint64 `gorm:"not null" json:"balance" form:"balance" valid:"required~Your balance is required"`
		BalanceAchieve uint64 `json:"balance_achieve" form:"balance_achieve"`
		Code           string `gorm:"not null;uniqueIndex" json:"code" form:"code" valid:"required~Your Code is required"`
		Enable         bool   `json:"enable"`
		UserId         uint   `json:"user_id"`
	}

	CreateAccountBankRequest struct {
		Balance        uint64 `json:"balance" form:"balance" valid:"required~Your balance is required"`
		BalanceAchieve uint64 `json:"balance_achieve" form:"balance_achieve"`
		Code           string `json:"code" form:"code" valid:"required~Your Code is required"`
	}

	AddDepositRequest struct {
		Balance        uint64 `json:"balance" form:"balance" valid:"required~Your balance is required"`
		BalanceAchieve uint64 `json:"balance_achieve" form:"balance_achieve"`
		Code           string `json:"code" form:"code" valid:"required~Your Code is required"`
	}

	BankBalanceHistory struct {
		Gorm
		BankBalanceId uint   `json:"bank_balance_id"`
		BalanceBefore uint64 `json:"balance_before"`
		BalanceAfter  uint64 `json:"balance_after"`
		Activity      string `json:"activity"`
		Type          string `json:"type"`
		Ip            string `json:"ip"`
		Location      string `json:"location"`
		UserAgent     string `json:"user_agent"`
		Author        string `json:"author"`
	}
)
