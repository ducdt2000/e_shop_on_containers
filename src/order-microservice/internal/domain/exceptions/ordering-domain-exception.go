package exceptions

type OrderingDomainException struct {
	message string
	cause   error
}

func NewOrderingDomainException(message string) *OrderingDomainException {
	return &OrderingDomainException{
		message: message,
	}

}

func NewOrderingDomainExceptionWithCause(message string, cause error) *OrderingDomainException {
	return &OrderingDomainException{
		message: message,
		cause:   cause,
	}
}
