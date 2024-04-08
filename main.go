package main 
import "fmt"


func sendMailSimple() {
  smtp.PlainAuth(
    "",
    "clrsclrsclrsclrsclrs@gmail.com",
    "",
    "smtp.gmail.com",

    )
    
    msg := "Subject : My special subject \n this is the body"
    smtp.SendMail(
    "smtp.gmail.com:587",
    auth,
    "clrsclrsclrsclrsclrs@gmail.com",
    [] string {"clrsclrsclrsclrsclrs@gmail.com"}
    )

    if err !=nil {
      fmt.Println(err)
  }
}

func main () {
  fmt.Println("Hello World")
}
