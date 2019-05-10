package main

import(
  "log"
  "net/http"
  "os/exec"
)

// Auth : return value of Authorization header or a new User
func Auth(r *http.Request) []byte {
  token := []byte(r.Header.Get("Authorization"))
  if len(token) > 0 {
    return token
  }
  return user()
}

func user() []byte{
  out, err := exec.Command("uuidgen").Output()
     if err != nil {
         log.Fatal(err)
     }
     return out
}
