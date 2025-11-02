package main
import ("fmt")
func main(){
	//功能1：显示主菜单，并且可以退出。
	//功能2：完成可以显示明细和登记收入的功能
	//因为需要明细，所以需要定义一个变量 string来记录
	//还需要定义来记录余额（balance），每次收支的余额（money），每次收支的说明（note）
	//功能3：完成了登记支出的功能
	//定义一个flag来完善
	var flag = false
	var loop = true
	 balance := 10000.0
	 money := 0.0
	 note := ""
	details := "收支\t账户余额\t收支金额\t说明"
	//当有收支时，只需要对details进行拼接处理
	fmt.Println("欢迎使用记账小程序！")
	for loop{
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
			fmt.Println("\n-------当前收支明细-------")
			if flag {
				fmt.Println(details)
			}else{
				fmt.Println("当前你还没有一笔收支，请来一笔吧！")
			}
		case 2:
			fmt.Println("本次输入金额:\n")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:\n")
			fmt.Scanln(&note)
			//现在我要将收支情况拼接到details里面去
			details +=fmt.Sprintf("\n收入\t%v\t          %v\t         %v",balance,money,note)
			flag = true


		case 3:
			fmt.Println("本次支出金额:\n")
			fmt.Scanln(&money)
			if balance <money{
				fmt.Println("余额不足")
				break
			}else{
				balance -=money
			}
			fmt.Println("支出说明:")
			fmt.Scanln(&note)
			//拼接details
			details +=fmt.Sprintf("\n支出\t%v\t          %v\t         %v",balance,money,note)
			flag = true

		case 4:
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

			loop = false
			fmt.Println("退出程序成功！")
			}else{
				loop = true
			}
			
		default :
			fmt.Println("您输入的数字有误,请重新输入:")

		}
		//完成switch 后，如果loop还是true那就继续显示这个主菜单，如果loop为false那就要退出程序
		//所以要用if判断现在loop值
		if !loop {
			//此时loop为false，说明了用户要退出程序
			break
		}

	}
	//此时已经为退出了整个for循环，那可以prompt一下
	fmt.Println("您已经成功退出家庭记账软件")
}