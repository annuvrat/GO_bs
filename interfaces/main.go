package main
import "fmt"

type payment struct{
gateway stripe
}

func (p payment) makePayment(amount float32)  {
//  razorpayPaymentGw:=razorpay{}
// stripePaymentGw:=stripe{}

//  razorpayPaymentGw.pay(amount)
p.gateway.pay(amount)
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
func main(){
	stripePaymentGw:=stripe{}

newPayment:= payment{gateway: stripePaymentGw}
	newPayment.makePayment(100)
}
