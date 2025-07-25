package models

import (
	"errors"
	"fmt"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)


type Payment struct {
	StripeSecretKey string
	SuccessUrl string
	CancelUrl string
}

func (p Payment) CreatePayment(orderPayload OrderPayload)(*stripe.CheckoutSession, error){
	fmt.Println("payment : ", p)
	stripe.Key = p.StripeSecretKey //----> Get the stripe key.
	orderDetails := orderPayload.OrderDetails//----> Cart line-items.
	
	//----> Fill in the cart line items.
	lineItems := getLineItems(orderDetails)
	
	//----> Create a new checkout session with the generated line items
	params := &stripe.CheckoutSessionParams{
		LineItems: lineItems,
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:      stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(p.SuccessUrl),
		CancelURL:  stripe.String(p.CancelUrl),
	}

	//----> Create stripe session.
	s, err := session.New(params)

	//----> Check for payment error.
	 if err != nil {
		return nil, errors.New("payment is not successful")
	}

	//----> Send back the response.
	return s, nil
}

func getLineItems(orderDetails []OrderDetail)[]*stripe.CheckoutSessionLineItemParams{
	var lineItems []*stripe.CheckoutSessionLineItemParams
	 
	for _, item := range orderDetails {
		lineItem := &stripe.CheckoutSessionLineItemParams{
			PriceData:  &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(string(item.Name)),
					Images: stripe.StringSlice([]string{item.Image}),
				},
				UnitAmount: stripe.Int64(int64(item.Price * 100)),
			},
			Quantity: stripe.Int64(int64(item.Quantity)),
			
		} 
		
		lineItems = append(lineItems, lineItem)
	} 

	return lineItems
} 

func MakeSuccessAndCancelUrls(origin string) (string, string){
	origin = "http://localhost:5173"
	successUrl := fmt.Sprintf("%v/orders/payment-success", origin)
	cancelUrl := fmt.Sprintf("%v/orders/payment-failure", origin)

	//----> Send back the success and cancel urls.
	return cancelUrl, successUrl
}