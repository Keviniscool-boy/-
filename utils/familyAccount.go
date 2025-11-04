package utils

import (
	"fmt"
	"time"
)

type Transaction struct {
	Type      string    // "收入" or "支出"
	Amount    float64   // 金额
	Balance   float64   // 余额
	Note      string    // 说明
	Category  string    // 类别
	Timestamp time.Time // 时间戳
}

type FamilyAccount struct {
	loop         bool
	balance      float64
	transactions []Transaction // 存储所有交易记录
}
// 显示收支明细
func (this *FamilyAccount) ShowDetails() {
	fmt.Println("\n================ 当前收支明细 ================")
	if len(this.transactions) == 0 {
		fmt.Println("当前你还没有一笔收支，请来一笔吧！")
		return
	}
	
	fmt.Printf("%-20s %-10s %-12s %-12s %-15s %s\n", 
		"时间", "类型", "金额", "余额", "类别", "说明")
	fmt.Println("-----------------------------------------------------------------------------------")
	
	for _, t := range this.transactions {
		fmt.Printf("%-20s %-10s %-12.2f %-12.2f %-15s %s\n",
			t.Timestamp.Format("2006-01-02 15:04:05"),
			t.Type,
			t.Amount,
			t.Balance,
			t.Category,
			t.Note)
	}
	fmt.Println("===========================================")
}
// 登记收入
func (this *FamilyAccount) ADDIncome() {
	var money float64
	var note string
	var category string
	
	fmt.Println("\n========== 登记收入 ==========")
	fmt.Print("请输入收入金额: ")
	_, err := fmt.Scanln(&money)
	if err != nil || money <= 0 {
		fmt.Println("输入金额无效，请输入正数！")
		return
	}
	
	fmt.Print("请选择收入类别 (1.工资 2.奖金 3.投资 4.其他): ")
	var catChoice int
	fmt.Scanln(&catChoice)
	switch catChoice {
	case 1:
		category = "工资"
	case 2:
		category = "奖金"
	case 3:
		category = "投资"
	case 4:
		category = "其他"
	default:
		category = "未分类"
	}
	
	fmt.Print("请输入收入说明: ")
	fmt.Scanln(&note)
	
	this.balance += money
	
	transaction := Transaction{
		Type:      "收入",
		Amount:    money,
		Balance:   this.balance,
		Note:      note,
		Category:  category,
		Timestamp: time.Now(),
	}
	this.transactions = append(this.transactions, transaction)
	
	fmt.Printf("\n✓ 收入登记成功！当前余额: %.2f\n", this.balance)
}
// 登记支出
func (this *FamilyAccount) ADDExpense() {
	var money float64
	var note string
	var category string
	
	fmt.Println("\n========== 登记支出 ==========")
	fmt.Print("请输入支出金额: ")
	_, err := fmt.Scanln(&money)
	if err != nil || money <= 0 {
		fmt.Println("输入金额无效，请输入正数！")
		return
	}
	
	if this.balance < money {
		fmt.Printf("余额不足！当前余额: %.2f，需要支出: %.2f\n", this.balance, money)
		return
	}
	
	fmt.Print("请选择支出类别 (1.餐饮 2.交通 3.购物 4.娱乐 5.医疗 6.其他): ")
	var catChoice int
	fmt.Scanln(&catChoice)
	switch catChoice {
	case 1:
		category = "餐饮"
	case 2:
		category = "交通"
	case 3:
		category = "购物"
	case 4:
		category = "娱乐"
	case 5:
		category = "医疗"
	case 6:
		category = "其他"
	default:
		category = "未分类"
	}
	
	fmt.Print("请输入支出说明: ")
	fmt.Scanln(&note)
	
	this.balance -= money
	
	transaction := Transaction{
		Type:      "支出",
		Amount:    money,
		Balance:   this.balance,
		Note:      note,
		Category:  category,
		Timestamp: time.Now(),
	}
	this.transactions = append(this.transactions, transaction)
	
	fmt.Printf("\n✓ 支出登记成功！当前余额: %.2f\n", this.balance)
}
// 显示余额汇总
func (this *FamilyAccount) ShowBalance() {
	fmt.Println("\n================ 账户余额汇总 ================")
	fmt.Printf("当前余额: %.2f 元\n", this.balance)
	
	var totalIncome, totalExpense float64
	for _, t := range this.transactions {
		if t.Type == "收入" {
			totalIncome += t.Amount
		} else if t.Type == "支出" {
			totalExpense += t.Amount
		}
	}
	
	fmt.Printf("总收入: %.2f 元\n", totalIncome)
	fmt.Printf("总支出: %.2f 元\n", totalExpense)
	fmt.Printf("交易笔数: %d 笔\n", len(this.transactions))
	fmt.Println("===========================================")
}

// 显示分类统计
func (this *FamilyAccount) ShowStatistics() {
	fmt.Println("\n================ 分类统计 ================")
	
	if len(this.transactions) == 0 {
		fmt.Println("暂无交易记录！")
		return
	}
	
	incomeByCategory := make(map[string]float64)
	expenseByCategory := make(map[string]float64)
	
	for _, t := range this.transactions {
		if t.Type == "收入" {
			incomeByCategory[t.Category] += t.Amount
		} else if t.Type == "支出" {
			expenseByCategory[t.Category] += t.Amount
		}
	}
	
	fmt.Println("\n【收入分类统计】")
	if len(incomeByCategory) == 0 {
		fmt.Println("  暂无收入记录")
	} else {
		for category, amount := range incomeByCategory {
			fmt.Printf("  %s: %.2f 元\n", category, amount)
		}
	}
	
	fmt.Println("\n【支出分类统计】")
	if len(expenseByCategory) == 0 {
		fmt.Println("  暂无支出记录")
	} else {
		for category, amount := range expenseByCategory {
			fmt.Printf("  %s: %.2f 元\n", category, amount)
		}
	}
	fmt.Println("=========================================")
}

// 退出软件
func (this *FamilyAccount) Exit() {
	fmt.Print("\n你真的要退出软件吗？(y/n): ")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("请重新输入 (y/n): ")
	}
	
	if choice == "y" {
		this.loop = false
		fmt.Println("\n感谢使用家庭收支记账软件，再见！")
	}
}
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		loop:         true,
		balance:      10000.0,
		transactions: make([]Transaction, 0),
	}
}

// 显示主菜单
func (this *FamilyAccount) ShowMenu() {
	for this.loop {
		fmt.Println("\n╔════════════════════════════════════════╗")
		fmt.Println("║      家庭收支记账软件 v2.0           ║")
		fmt.Println("╠════════════════════════════════════════╣")
		fmt.Println("║  1. 收支明细                          ║")
		fmt.Println("║  2. 登记收入                          ║")
		fmt.Println("║  3. 登记支出                          ║")
		fmt.Println("║  4. 余额汇总                          ║")
		fmt.Println("║  5. 分类统计                          ║")
		fmt.Println("║  6. 退出系统                          ║")
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Print("请选择 (1-6): ")
		
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Scanln() // 清除输入缓冲
			fmt.Println("输入无效，请输入数字！")
			continue
		}
		
		switch choice {
		case 1:
			this.ShowDetails()
		case 2:
			this.ADDIncome()
		case 3:
			this.ADDExpense()
		case 4:
			this.ShowBalance()
		case 5:
			this.ShowStatistics()
		case 6:
			this.Exit()
		default:
			fmt.Println("输入有误，请选择 1-6！")
		}
		
		if !this.loop {
			break
		}
	}
}