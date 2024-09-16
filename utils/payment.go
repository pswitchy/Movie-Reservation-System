package utils

import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/charge"
)

func ProcessPayment(amount int64, currency string, source string) (*stripe.Charge, error) {
    stripe.Key = "your_stripe_secret_key"
    params := &stripe.ChargeParams{
        Amount:   stripe.Int64(amount),
        Currency: stripe.String(currency),
        Source:   &stripe.SourceParams{Token: stripe.String(source)},
    }
    return charge.New(params)
}
