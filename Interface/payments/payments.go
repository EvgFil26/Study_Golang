package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentInfo struct {
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		description,
		usd,
		false,
	}

	p.paymentsInfo[id] = info

	return id
}

func (p PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)

	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}
	info.Cancelled = true

	p.paymentsInfo[id] = info
}

func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	return tempMap
}
