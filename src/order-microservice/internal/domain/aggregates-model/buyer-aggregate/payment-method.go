package buyer_aggregate

import (
	"fmt"
	"time"

	seed_work "github.com/ducdt2000/e_shop_on_containers/order-microservice/internal/domain/seed-work"
)

type PaymentMethod struct {
	seed_work.Entity
	alias          string
	cardNumber     string
	securityNumber string
	cardHolderName string
	expiration     time.Time
	cardType       CardType
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
		alias:          alias,
		cardNumber:     cardNumber,
		securityNumber: securityNumber,
		cardHolderName: cardHolderName,
		expiration:     expiration,
		cardType:       cardType,
	}

	return paymentMethod, nil
}

func (pm *PaymentMethod) IsEqualTo(cardType CardType, cardNumber string, expiration time.Time) bool {
	return pm.cardType == cardType &&
		pm.cardNumber == cardNumber &&
		pm.expiration.Equal(expiration)
}

func (pm PaymentMethod) GetCardType() CardType {
	return pm.cardType
}
