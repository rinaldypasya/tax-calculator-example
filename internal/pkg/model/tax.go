package model

import "time"

type Tax struct {
	ID        int64
	UserID    int64
	Name      string
	TaxCode   TaxCode
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tax) GetTaxCodeString() string {
	switch t.TaxCode {
	case TaxCodeFood:
		return "Food & Beverage"
	case TaxCodeTobacco:
		return "Tobacco"
	case TaxCodeEntertainment:
		return "Entertainment"
	default:
		return "unknown"
	}
}

func (t *Tax) IsRefundable() bool {
	switch t.TaxCode {
	case TaxCodeFood:
		return true
	case TaxCodeTobacco:
		return false
	case TaxCodeEntertainment:
		return false
	default:
		return false
	}
}

func (t *Tax) GetTaxValue() float64 {
	switch t.TaxCode {
	case TaxCodeFood:
		// 10% of Price
		return (float64(10) / float64(100)) * float64(t.Price)
	case TaxCodeTobacco:
		// 10 + (2% of Price )
		return float64(10) + float64((float64(2)/float64(100))*float64(t.Price))
	case TaxCodeEntertainment:
		// Price >= 100: 1% of ( Price - 100)
		if t.Price >= 100 {
			return float64(1) / float64(100) * (float64(t.Price) - float64(100))
		}

		// 0 < Price < 100: tax-free
		return 0
	default:
		return 0
	}
}

func (t *Tax) GetAmount() float64 {
	return float64(t.Price) + t.GetTaxValue()
}
