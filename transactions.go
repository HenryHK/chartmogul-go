package chartmogul

import (
	"fmt"
	"strings"
)

const transactionsEndpoint = "import/invoices/:invoiceUUID/transactions"

// Transaction is either payment/refund on an invoice, for its full value.
type Transaction struct {
	UUID       string `json:"uuid,omitempty"` // only on return
	Date       string `json:"date"`
	ExternalID string `json:"external_id,omitempty"`
	Result     string `json:"result"`
	Type       string `json:"type"`
	Errors     Errors `json:"errors,omitempty"`
}

func (t Transaction) String() string {
	return fmt.Sprintf("Transaction(%v) %v %v (%v)", t.ExternalID, t.Result, t.Type, t.Date)
}

// CreateTransaction loads an transaction to a customer in Chartmogul.
// Customer must have a valid UUID! (use return value of API)
//
// See https://dev.chartmogul.com/reference#import-customers-transactions
func (api API) CreateTransaction(transaction *Transaction, invoiceUUID string) (*Transaction, error) {
	result := &Transaction{}
	path := strings.Replace(transactionsEndpoint, ":invoiceUUID", invoiceUUID, 1)
	return result, api.create(path, transaction, result)

}
