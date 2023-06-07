package buyer_aggregate

import (
	"fmt"
	"time"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type PaymentMethod struct {
	seed_work.Entity
	_alias          string
	_cardNumber     string
	_securityNumber string
	_cardHolderName string
	_expiration     time.Time
	_cardType       CardType
}

func NewPaymentMethod(
	cardType CardType,
	alias string,
	cardNumber string,
	securityNumber string,
	cardHolderName string,
	expiration time.Time,
) (*PaymentMethod, error) {
	if cardNumber == "" {
		return nil, fmt.Errorf("cardNumber is required")
	}

	if securityNumber == "" {
		return nil, fmt.Errorf("securityNumber is required")
	}

	if cardHolderName == "" {
		return nil, fmt.Errorf("cardHolderName is required")
	}

	if expiration.Before(time.Now().UTC()) {
		return nil, fmt.Errorf("expiration date must be in the future")
	}

	paymentMethod := &PaymentMethod{
		_alias:          alias,
		_cardNumber:     cardNumber,
		_securityNumber: securityNumber,
		_cardHolderName: cardHolderName,
		_expiration:     expiration,
		_cardType:       cardType,
	}

	return paymentMethod, nil
}

func (pm *PaymentMethod) IsEqualTo(cardType CardType, cardNumber string, expiration time.Time) bool {
	return pm._cardType == cardType &&
		pm._cardNumber == cardNumber &&
		pm._expiration.Equal(expiration)
}

func (pm PaymentMethod) GetCardType() CardType {
	return pm._cardType
}
