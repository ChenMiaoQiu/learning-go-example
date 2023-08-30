package main

import (
	"fmt"

	"github.com/ChenMiaoQiu/learning-go-example/ch4/4.11/github"
)

func main() {
	fmt.Println("please enter your oper:")
	fmt.Println("1.create issue")
	fmt.Println("2.read issue")
	fmt.Println("3.update issue")
	fmt.Println("4.close issue")

	var oper int
	fmt.Scan(&oper)

	switch oper {
	case 1:
		createIssue()
	case 2:
		getIssues()
	case 3:
		editIssue()
	case 4:
		closeIssue()
	default:
		fmt.Println("please enter right oper")
	}
}

func createIssue() {
	var p github.Params
	fmt.Print("please enter the owner")
	fmt.Scan(&p.Owner)
	fmt.Print("please enter the repo")
	fmt.Scan(&p.Repo)
	fmt.Print("please enter the token")
	fmt.Scan(&p.Token)
	fmt.Printf("please enter the title")
	fmt.Scan(&p.Title)
	fmt.Printf("please enter the body")
	fmt.Scan(&p.Body)
	isSuccess := p.CreateIssue()

	if isSuccess {
		fmt.Println("create success")
	} else {
		fmt.Println("create false")
	}
}

func getIssues() {
	var p github.Params
	fmt.Print("please enter the owner")
	fmt.Scan(&p.Owner)
	fmt.Print("please enter the repo")
	fmt.Scan(&p.Repo)

	resp, err := p.GetIssues()

	if err != nil {
		fmt.Println("get issues false")
	} else {
		for _, item := range resp {
			fmt.Printf("%.55s\n", item.Title)
		}
	}
}

func editIssue() {
	var p github.Params
	fmt.Print("please enter the owner")
	fmt.Scan(&p.Owner)
	fmt.Print("please enter the repo")
	fmt.Scan(&p.Repo)
	fmt.Print("please enter the number")
	fmt.Scan(&p.Number)
	fmt.Print("please enter the token")
	fmt.Scan(&p.Token)
	fmt.Printf("please enter the title")
	fmt.Scan(&p.Title)
	fmt.Printf("please enter the body")
	fmt.Scan(&p.Body)
	isSuccess := p.EditIssue()

	if isSuccess {
		fmt.Println("create success")
	} else {
		fmt.Println("create false")
	}
}

func closeIssue() {
	var p github.Params
	fmt.Print("please enter the owner")
	fmt.Scan(&p.Owner)
	fmt.Print("please enter the repo")
	fmt.Scan(&p.Repo)
	fmt.Print("please enter the number")
	fmt.Scan(&p.Number)
	fmt.Print("please enter the token")
	fmt.Scan(&p.Token)
	isSuccess := p.CloseIssue()

	if isSuccess {
		fmt.Println("create success")
	} else {
		fmt.Println("create false")
	}
}
