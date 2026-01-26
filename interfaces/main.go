package main
import "fmt"



type paymenter interface{
	pay(amount float32)
	refund(amount float32)
}

type payment struct{
gateway paymenter // dependency inversion
}

func (p payment) makePayment(amount float32)  {
//  razorpayPaymentGw:=razorpay{}
// stripePaymentGw:=stripe{}

//  razorpayPaymentGw.pay(amount)
p.gateway.pay(amount)
}

func (p payment) refundPayment(amount float32)  {
//  razorpayPaymentGw:=razorpay{}
// stripePaymentGw:=stripe{}

//  razorpayPaymentGw.pay(amount)
p.gateway.refund(amount)
}


type razorpay struct{

}
func (r razorpay) pay(amount float32)  {
	

	fmt.Println("Payment of", amount, "made using Razorpay")
}


type stripe struct{}

func (s stripe) pay(amount float32)  {
	fmt.Println("Payment of", amount, "made using Stripe")
}

type paypal struct{}

func (p paypal)pay(amount float32)  {
	fmt.Println("Payment of", amount, "made using Paypal")
}

func (p paypal)refund(amount float32)  {
	fmt.Println("Refund of", amount, "processed using Paypal")
}
func main(){
	PaymentGw:=paypal{}

newPayment:= payment{gateway: PaymentGw}
	newPayment.refundPayment(100)
}
