package main

import "fmt"

// PaymentProcessor interface 정의
type PaymentProcessor interface {
    Pay(amount float64) string
}

// CreditCard struct
type CreditCard struct {
    CardNumber string
}

// Pay 메서드 구현 (CreditCard 타입)
func (cc CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// PayPal struct
type PayPal struct {
    Email string
}

// Pay 메서드 구현 (PayPal 타입)
func (pp PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid $%.2f using PayPal account: %s", amount, pp.Email)
}

// BankTransfer 구조체
type BankTransfer struct {
    AccountNumber string
}

// Pay 메서드 구현 (BankTransfer 타입)
func (bt BankTransfer) Pay(amount float64) string {
    return fmt.Sprintf("Paid $%.2f using Bank Transfer to account: %s", amount, bt.AccountNumber)
}

// 결제 함수 (PaymentProcessor 인터페이스 활용)
func processPayment(p PaymentProcessor, amount float64) {
    fmt.Println(p.Pay(amount))
}

func main() {
    // 다양한 결제 방식 인스턴스 생성
    creditCard := CreditCard{CardNumber: "1234-5678-9876-5432"}
    payPal := PayPal{Email: "user@example.com"}
    bankTransfer := BankTransfer{AccountNumber: "110-123-456789"}

    // 결제 처리
    processPayment(creditCard, 100.50)   // 신용카드로 결제
    processPayment(payPal, 50.75)       // 페이팔로 결제
    processPayment(bankTransfer, 200.00) // 가상 계좌로 결제
}
