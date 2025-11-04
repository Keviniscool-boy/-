package utils
import "fmt"
type FamilyAccount struct {
	//和面向过程一样
	//我们需要余额balance float64， 收支money float64，说明details string，note sring
	//声明一个变量，用于接收用户的输入
	Num string
	loop bool
	balance float64
	money float64
	note string
	details string
	flag bool
}
//将收支明细写一个方法
func (this *FamilyAccount) ShowDetails(){
	fmt.Println("\n-------当前收支明细-------")
			if this.flag {
				fmt.Println(this.details)
			}else{
				fmt.Println("当前你还没有一笔收支，请来一笔吧！")
			}

}
//将登记收入写一个方法
func(this *FamilyAccount) ADDIncome(){
	fmt.Println("本次输入金额:\n")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:\n")
	fmt.Scanln(&this.note)
	//现在我要将收支情况拼接到details里面去
	this.details +=fmt.Sprintf("\n收入\t%v\t          %v\t         %v",this.balance,this.money,this.note)
	this.flag = true

}
//将登记支出写一个方法
func(this *FamilyAccount) ADDExpense(){fmt.Println("本次支出金额:\n")
	fmt.Scanln(&this.money)
	if this.balance < this.money{
		fmt.Println("余额不足")
		return 
	}else{
		this.balance -= this.money
	}
	fmt.Println("支出说明:")
	fmt.Scanln(&this.note)
	//拼接details
			this.details +=fmt.Sprintf("\n支出\t%v\t          %v\t         %v",this.balance,this.money,this.note)
	this.flag = true
	}
//将退出软件写一个方法
func (this *FamilyAccount)Exit(){
	fmt.Println("你真的要退出软件吗？(y/n)")
		//要退出软件，这里不适用break，因为switch分支在for循环里面
			choice :=""
			for{	fmt.Scanln(&choice)

				if choice =="y" || choice =="n"{
					break
				}
				fmt.Println("请重新输入:")
			}
			if choice == "y"{
		//可以在for循环外定义一个loop用于退出软件

			this.loop = false
			fmt.Println("退出程序成功！")
			}else{
				this.loop = true
			}
}
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		Num :"",
		loop :true,
		balance :10000.0,
		money :0.0,
		note :"",
		details :"收支\t账户余额\t收支金额\t说明",
		flag :false,
	}
}

//给这个结构体绑定方法
func (this *FamilyAccount) ShowMenu(){
	for this.loop{
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println(" ------------1.收支明细----------")
		fmt.Println("  -----------2.登记收入----------")
		fmt.Println("  -----------3.登记支出----------")
		fmt.Println("-------------4.退出--------------")
		fmt.Println("请选择（1-4）:")
		//定义一个变量接收用户输入的数字
		var Num int
		fmt.Scanln(&Num)
		//使用switch分支，分别处理1-4的功能
		switch Num{
		case 1:
			this.ShowDetails()
		case 2:
			this.ADDIncome()
		case 3:
			this.ADDExpense()
		case 4:
			this.Exit()
			
		default :
			fmt.Println("您输入的数字有误,请重新输入:")

		}
		//完成switch 后，如果loop还是true那就继续显示这个主菜单，如果loop为false那就要退出程序
		//所以要用if判断现在loop值
		if !this.loop {
			//此时loop为false，说明了用户要退出程序
			break
		}

	}
}