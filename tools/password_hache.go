package tools

import(
   "golang.org/x/crypto/bcrypt"
)

func HachePassword(password string)(string, error){
   hachedpassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
   if err !=nil{
     return "",err
   }
   return string(hachedpassword),nil
}
 
func Verifierpassword(hachedpassword string,password string)bool{
  err := bcrypt.CompareHashAndPassword([]byte(hachedpassword),[]byte(password))
  return err == nil
}


