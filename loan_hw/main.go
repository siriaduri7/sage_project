package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type Integer int

func (r *Integer) Set() {
	if r == nil {

	}
	fmt.Println(r)
}

type LoanCalculator interface {
	TotalLoanCalc() float64
	EMIPerMonth() (float64, float64)
	LoanToAmountRatio() float64
}

type HomeLoan struct {
	TotalCost    float64
	DownPayment  float64
	LoanTermYear float64
	IntrestRate  float64
}

type AutoLoan struct {
	CarPrice     float64
	DownPayment  float64
	LoanTermYear float64
	InterestRate float64
}

type PersonalLoan struct {
	PersonalCost        float64
	InterestRate        float64
	LoanTermMonthly     float64
	TotalIntrestPayable float64
}

func (q *PersonalLoan) TotalLoanCalc() float64 {
	return q.PersonalCost - q.LoanTermMonthly
}
func Loansta(credit int, income int) string {

	if credit >= 800 && income >= 30000 {
		return "congratulations you got loan approved"
	} else {
		return "sorry your loan is rejected"
	}
}

func (q *PersonalLoan) EMIPerMonth() (float64, float64) {
	p := q.PersonalCost
	n := q.LoanTermMonthly * 12
	r := q.InterestRate / 1200

	em := (p * r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
	princPerMonthh := em / (1 + r)

	return em, princPerMonthh
}

func (q *PersonalLoan) LoanToAmountRatio() float64 {
	return (q.PersonalCost - q.LoanTermMonthly) / q.PersonalCost * 100

}

func (a *AutoLoan) TotalLoanCalc() float64 {

	return a.CarPrice - a.DownPayment
}
func LoanStatus(credit int, income int) string {

	if credit >= 800 && income >= 30000 {
		return "congratulations you got loan approved"
	} else {
		return "sorry your loan is rejected"
	}
}

func (a *AutoLoan) LoanToAmountRatio() float64 {
	return (a.CarPrice - a.DownPayment) / a.CarPrice * 100

}

func (a *AutoLoan) EMIPerMonth() (float64, float64) {
	//p := a.TotalLoanCalc()
	p := a.CarPrice
	n := a.LoanTermYear * 12
	r := a.InterestRate / 1200

	emii := (p * r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
	principlePerMonthh := emii / (1 + r)

	return emii, principlePerMonthh

}

func (h *HomeLoan) TotalLoanCalc() float64 {
	return h.TotalCost - h.DownPayment

}
func LoanValid(credit int, income int) string {

	if credit >= 800 && income >= 30000 {
		return "congratulations you got loan approved"
	} else {
		return "sorry your loan is rejected"
	}
}

func (h *HomeLoan) LoanToAmountRatio() float64 {
	return (h.TotalCost - h.DownPayment) / h.TotalCost * 100

}

func (h *HomeLoan) EMIPerMonth() (float64, float64) {
	// formula (p*r*(1+r)pow N/ (1-r)pow N-1) // p = principal loan amount ; n = tenure in months
	//R is the prevailing interest rate.
	//(p*r*math.Pow(1+r, n) / (math.Pow(1+r, n) - 1)
	// p := h.TotalLoanCalc()
	p := h.TotalCost
	n := h.LoanTermYear * 12
	r := h.IntrestRate / 1200

	emi := (p * r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
	principlePerMonth := emi / (1 + r)

	return emi, principlePerMonth

}

func main() {
	for {
		fmt.Println("Plesae select type of loan")
		fmt.Println("1.Home loan")
		fmt.Println("2.Auto loan")
		fmt.Println("3.Personal loan")
		fmt.Println("4. Exit")
		var choice int
		fmt.Println("enter your choices(1-4): ")
		fmt.Scanln(&choice)

		calc := &HomeLoan{}

		auto := &AutoLoan{}

		per := &PersonalLoan{}

		switch choice {
		case 1:
			fmt.Println("Enter the loanamount home")
			var Totalcost float64
			fmt.Scanln(&Totalcost)

			fmt.Println("enter the downpayment home")
			var downpayment float64
			fmt.Scanln(&downpayment)

			fmt.Println("enter the loanterm home")
			var loanterm float64
			fmt.Scanln(&loanterm)

			fmt.Println("enter the interestrate home")
			var Interestrate float64
			fmt.Scanln(&Interestrate)
			calc = &HomeLoan{
				TotalCost:    Totalcost,
				DownPayment:  downpayment,
				LoanTermYear: loanterm,
				IntrestRate:  Interestrate,
			}

		case 2:
			fmt.Println("Enter the loanamount auto")
			var Totalcost float64
			fmt.Scanln(&Totalcost)

			fmt.Println("enter the downpayment auto")
			var downpayment float64
			fmt.Scanln(&downpayment)

			fmt.Println("enter the loanterm auto")
			var loanterm float64
			fmt.Scanln(&loanterm)

			fmt.Println("enter the interestrate auto")
			var Interestrate float64
			fmt.Scanln(&Interestrate)
			auto = &AutoLoan{
				CarPrice:     Totalcost,
				DownPayment:  downpayment,
				LoanTermYear: loanterm,
				InterestRate: Interestrate,
			}

		case 3:
			fmt.Println("enter the PersonalCost ")
			var pers float64
			fmt.Scanln(&pers)

			fmt.Println("enter the InterestRate")
			var intre float64
			fmt.Scanln(&intre)

			fmt.Println("enter the LoanTermMonthly")
			var term float64
			fmt.Scanln(&term)

			fmt.Println("enter the TotalIntrestPayable ")
			var pay float64
			fmt.Scanln(&pay)

			per = &PersonalLoan{
				PersonalCost:        pers,
				InterestRate:        intre,
				LoanTermMonthly:     term,
				TotalIntrestPayable: pay,
			}
		case 4:
			os.Exit(1)

		default:

			fmt.Println("invalid choice")

		}

		// a := calc.TotalLoanCalc()
		// fmt.Println(&a)
		var credit, income int
		fmt.Println("Enter your credit score between( 300 and 850):")
		fmt.Scanln(&credit)
		fmt.Println(" enter your income in dollars:")
		fmt.Scanln(&income)
		status := LoanStatus(credit, income)
		valid := LoanValid(credit, income)
		sta := Loansta(credit, income)
		fmt.Println(status)
		fmt.Println(valid)
		fmt.Println(sta)
		//approvalState := LoanStatus(credit, income)

		if choice == 1 {
			//calculate home loan

			fmt.Println("total loan amount home : ", calc.TotalLoanCalc())
			fmt.Println("Approved loan amount home:", calc.LoanToAmountRatio())
			emi, principlePerMonth := calc.EMIPerMonth()
			fmt.Println("EMI per month home", emi)
			fmt.Println("principle per month home", principlePerMonth)
		} else if choice == 2 {
			fmt.Println("total loan amount : ", auto.TotalLoanCalc())
			fmt.Println("Approved loan amount:", auto.LoanToAmountRatio())
			emii, principlePerMonthh := auto.EMIPerMonth()
			fmt.Println("EMI per month", emii)
			fmt.Println("principle per month", principlePerMonthh)
		} else if choice == 3 {
			fmt.Println("totalloan amount for personal ", per.TotalLoanCalc())
			fmt.Println("Approved loan amount:", per.LoanToAmountRatio())
			em, princPerMonthh := per.EMIPerMonth()
			fmt.Println("EMI per month", em)
			fmt.Println("principle per month", princPerMonthh)

		} else {
			fmt.Println("invalid choice: please select valid option")
			time.Sleep(500)
		}
	}
}

//}
// 	emi, principlePerMonth := calc.EMIPerMonth()
// 	fmt.Println("EMI per monthls", emi)
// 	fmt.Println("principle per month", principlePerMonth)

// }

// how can i calculate home  loan, once if i click home loan in output it should display how much is your total loan, how much can bank approve, and how much percentage can be approved, by using above deatils need code in golang

// now using interface perform above program
// include emipermonth method in interface for above program and it should display how much emi need to pay for my loan and how much its principle
// in the above code include autoloan struct and implement same fields as homeloancalculator
// type AutoLoanCalc struct {
// 	CarPrice       float64
// 	DownPayment    float64
// 	LoanTermMonths float64
// 	IntrsetRate    float64
// }
