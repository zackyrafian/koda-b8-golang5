package main

import (
	"fmt"

)
type User struct { 
  FirstName string
  LastName string
  Email string
  Password string
}

var token = User {}
var accounts = []User { 
  {
    FirstName: "John",
    LastName: "Doe",
    Email: "test@example.com",
    Password: "test",
  },
  {
    FirstName: "Lawak",
    LastName: "Test",
    Email: "lawak@example.com",
    Password: "test",
  },
}
func clearTerminal() { 
  fmt.Print("\x1B[2J\x1B[H")
}

func dashboard() { 
  clearTerminal()
  fmt.Print("Welcome to system")
  fmt.Printf("\nHello %s %s", token.FirstName, token.LastName)
  fmt.Print("\n1. List All Users\n2. Logout\n\n0. Exit")
}

func register(){ 
  clearTerminal()
  fmt.Println("Register")
  fmt.Print("")
  reg := User{}

  fmt.Print("What is you first name: ")
  fmt.Scan(&reg.FirstName)
  fmt.Print("What is you last name: ")
  fmt.Scan(&reg.LastName)
  fmt.Print("What is you email: ")
  fmt.Scan(&reg.Email)
  fmt.Print("Enter a strong password : ")
  fmt.Scan(&reg.Password)
  fmt.Print("Enter a strong password : ")
  fmt.Scan(&reg.Password)

  fmt.Print("\n")

  var confirm string
  fmt.Print("Is it true?\n")
  fmt.Printf("First Name: %s\n", reg.FirstName)
  fmt.Printf("Last Name: %s\n", reg.LastName)
  fmt.Printf("Email: %s\n", reg.Email)
  fmt.Print("Continue (y/n): ")
  fmt.Scan(&confirm)

  if confirm != "y" { 
    register()
  }
  accounts = append(accounts, reg)
  fmt.Print("Register Success, press enter to back..")
  defer home()
}

func login() {  
    clearTerminal()
    loginForm := User{}
    fmt.Println("Login")
    fmt.Print("Enter your email: ")
    fmt.Scan(&loginForm.Email)
    idx := -1
    for index, x := range accounts{ 
      if x.Email == loginForm.Email {
        idx = index
      }
    }

    if idx != -1 { 
      fmt.Print("Enter your password: ")
      fmt.Scan(&loginForm.Password)
      if loginForm.Password != accounts[idx].Password { 
        login()
      } else {
        token = accounts[idx]
        dashboard()
      }
    } else { 
      login()
    }
}

func forgetPassword(){ 
  fmt.Println("Forget Password")
}

func home() { 
  clearTerminal()
  var q string
  fmt.Println("Welcome to system")
  fmt.Print("1. Register\n2. Login\n3. Forget Password\n\n0. Exit")
  fmt.Print("\nChoose a menu: ")
  fmt.Scan(&q)
  
  switch q { 
    case "1": 
      register()
    case "2": 
      login()
    case "3": 
      forgetPassword()
    default: 
      home()
  }
}


func main() { 
  home()
  // register()
  // login()
}