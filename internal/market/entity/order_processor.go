package entity

type OrderProcessor struct {
	Transaction *Transaction
}

func NewOrderProcessor(transation *Transaction) *OrderProcessor {
	return &OrderProcessor{
		Transaction: transation,
	}
}

func (op *OrderProcessor) Process() {
	shares := op.calculateShares()
	op.updatePositions(shares)
	op.updateOrders(shares)
	op.Transaction.Total = float64(shares) * op.Transaction.Price
}

func (op *OrderProcessor) calculateShares() int {
	availableShares := op.Transaction.Shares

	if op.Transaction.BuyingOrder.PendingShares < availableShares {
		availableShares = op.Transaction.BuyingOrder.PendingShares
	}

	if op.Transaction.SellingOrder.PendingShares < availableShares {
		availableShares = op.Transaction.SellingOrder.PendingShares
	}

	return availableShares
}

func (op *OrderProcessor) updatePositions(qtdShares int) {
	op.Transaction.SellingOrder.Investor.AdjustAssetPosition(op.Transaction.SellingOrder.Asset.ID, -qtdShares)
	op.Transaction.BuyingOrder.Investor.AdjustAssetPosition(op.Transaction.BuyingOrder.Asset.ID, qtdShares)
}

func (op *OrderProcessor) updateOrders(qtdShares int) {
	op.Transaction.BuyingOrder.ApplyTrade(qtdShares)
	op.Transaction.SellingOrder.ApplyTrade(qtdShares)
}
