package admin

import (
	"github.com/qor/exchange"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/qor/validations"
)

var ProductExchange *exchange.Resource

func init() {
	ProductExchange = exchange.NewResource(&models.Product{}, exchange.Config{PrimaryField: "Code"})
	ProductExchange.Meta(exchange.Meta{Name: "Code"})
	ProductExchange.Meta(exchange.Meta{Name: "Name"})
	ProductExchange.Meta(exchange.Meta{Name: "Price"})

	ProductExchange.AddValidator(func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
		if utils.ToInt(metaValues.Get("Price").Value) < 100 {
			return validations.NewError(record, "Price", "price don't less than 100")
		}
		return nil
	})
}
