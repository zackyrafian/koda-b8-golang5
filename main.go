package main

import (
	"crypto/md5"
	"encoding/hex"
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
  {
    FirstName: "a",
    LastName: "a",
    Email: "a",
    Password: "a",
  },
}
func clearTerminal() { 
  fmt.Print("\x1B[2J\x1B[H")
}

func MD5(text string) string {
  hash := md5.Sum([]byte(text))
  return hex.EncodeToString(hash[:])
}

func dashboard() { 
  clearTerminal()
  var question string
  fmt.Print("Welcome to system\n")
  fmt.Printf("\nHello %s %s", token.FirstName, token.LastName)
  fmt.Print("\n1. List All Users\n2. Logout\n\n0. Exit\n\n")

  fmt.Print("Choose a menu: ")
  fmt.Scan(&question)
  switch question {
    case "1": listAllUsers()
    case "2":
    case "0": 
    default: dashboard()
  }
}

func listAllUsers(){ 
  fmt.Print("List all users\n")
  for idx, user := range accounts{ 
    fmt.Printf("%d.\n", idx + 1)
    fmt.Printf("Fullname: %s %s\n", user.FirstName, user.LastName)
    fmt.Printf("Email: %s\n", user.Email)
    fmt.Printf("Password: %s\n\n", user.Password)
  }
  fmt.Print("Press enter to back ")
  fmt.Scanln()
  defer dashboard()
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
  var password string
  fmt.Print("Enter a strong password : ")
  fmt.Scan(&password)
  reg.Password = MD5(password)

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
      if MD5(loginForm.Password) != accounts[idx].Password { 
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
  // 
  //  
  // 
  
}