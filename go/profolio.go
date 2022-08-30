package main

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	profolio package
// Date:	2022/08/30

type Profolio []Money

func (p Profolio) Add(money Money) Profolio {
	p = append(p, money)
	return p
}

func (p Profolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}
	return Money{amount: total, currency: currency}
}
