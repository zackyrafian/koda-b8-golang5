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

type UserStore struct { 
  accounts []*User
  token *User
}

func clearTerminal() { 
  fmt.Print("\x1B[2J\x1B[H")
}

func MD5(text string) string {
  hash := md5.Sum([]byte(text))
  return hex.EncodeToString(hash[:])
}

func (us *UserStore) dashboard() { 
  clearTerminal()
  var question string
  fmt.Print("Welcome to system\n")
  fmt.Printf("\nHello %s %s", us.token.FirstName, us.token.LastName)
  fmt.Print("\n1. List All Users\n2. Logout\n\n0. Exit\n\n")

  fmt.Print("Choose a menu: ")
  fmt.Scan(&question)
  switch question {
    case "1": us.listAllUsers()
    case "2":
    case "0": 
    default: us.dashboard()
  }
}

func (us *UserStore) listAllUsers(){ 
  clearTerminal()
  fmt.Print("List all users\n")
  for idx, user := range us.accounts { 
    fmt.Printf("%d.\n", idx + 1)
    fmt.Printf("Fullname: %s %s\n", user.FirstName, user.LastName)
    fmt.Printf("Email: %s\n", user.Email)
    fmt.Printf("Password: %s\n\n", user.Password)
  }
  fmt.Print("Press enter to back ")
  fmt.Scanln()
  defer us.dashboard()
}

func (us *UserStore) register(){ 
  clearTerminal()
  fmt.Println("Register")
  fmt.Print("")
  reg := User{}
  var password, confirmPassword string
  
  fmt.Print("What is you first name: ")
  fmt.Scan(&reg.FirstName)
  fmt.Print("What is you last name: ")
  fmt.Scan(&reg.LastName)
  fmt.Print("What is you email: ")
  fmt.Scan(&reg.Email)
  fmt.Print("Enter a strong password : ")
  fmt.Scan(&password)
  fmt.Print("Enter a confirm password : ")
  fmt.Scan(&confirmPassword)

  if password != confirmPassword {
    us.register()
  }
  
  reg.Password = MD5(password)
  var confirm string
  fmt.Print("Is it true?\n")
  fmt.Printf("First Name: %s\n", reg.FirstName)
  fmt.Printf("Last Name: %s\n", reg.LastName)
  fmt.Printf("Email: %s\n", reg.Email)
  fmt.Print("Continue (y/n): ")
  fmt.Scan(&confirm)

  if confirm != "y" { 
    us.register()
  }
  us.accounts = append(us.accounts, &reg)
  fmt.Print("Register Success, press enter to back..")
  fmt.Scanln()
  defer us.home()
}

func (us *UserStore) login() {  
    clearTerminal()
    loginForm := &User{}
    fmt.Print("Login\n\n")
    fmt.Print("Enter your email: ")
    fmt.Scan(&loginForm.Email)
    idx := -1
    for index, x := range us.accounts{ 
      if x.Email == loginForm.Email {
        idx = index
      }
    }

    if idx != -1 { 
      fmt.Print("Enter your password: ")
      fmt.Scan(&loginForm.Password)
      if MD5(loginForm.Password) != us.accounts[idx].Password { 
        us.login()
      } else {
        us.token = us.accounts[idx]
        us.dashboard()
      }
    } else { 
      us.login()
    }
}

func (us *UserStore)forgetPassword(){ 
  fmt.Println("Forget Password")
}

func (us *UserStore)home() { 
  clearTerminal()
  var q string
  fmt.Print("Welcome to system\n\n")
  fmt.Print("1. Register\n2. Login\n3. Forget Password\n\n0. Exit")
  fmt.Print("\nChoose a menu: ")
  fmt.Scan(&q)
  
  switch q { 
    case "1": 
      us.register()
    case "2": 
      us.login()
    case "3": 
      us.forgetPassword()
    default: 
      us.home()
  }
}


func main() { 
  us := &UserStore{}
  us.home()
  // register()
  // login()
  // 
  //  
  // 
  
}