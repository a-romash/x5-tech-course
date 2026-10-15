package analytics

func AnalyzeSales(sales []Sale) Report {
	report := NewEmptyReport()

	soldProducts := make(map[string]int)

	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if !isValid(sale) {
			report.InvalidSales++
			continue
		}

		report.AddSale(sale)

		soldProducts[sale.Product] += sale.Quantity
	}

	topQuantity := 0
	for product, quantity := range soldProducts {
		if quantity > topQuantity || (quantity == topQuantity && product < report.TopProduct) {
			report.TopProduct = product
			topQuantity = quantity
		}
	}

	return report
}
