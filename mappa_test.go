package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
)

type MappaLoginRequest struct {
  Type string `json:"type"`
  UserName string `json:"username"`
  Password string `json:"password"`
}
type MappaLoginResponse struct {
  Id string `json:"id"`
  Ttl int `json:"ttl"`
  Created string `json:"created"`
  UserId int `json:"user_id"`
  Error string `json:"error"`
}

func TestMappa(t *testing.T) {
  ts := httptest.NewServer(setupServer())

  defer ts.Close()

  resp, err := http.Get(fmt.Sprintf("%s/hc", ts.URL))
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }
  if resp.StatusCode != 200 {
    t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
  }
  loginRequest:=&MappaLoginRequest{Type: "LOGIN_REQUEST",
    UserName: "guionardo",
  Password: "A1GU"}
  jsonStr,_:=json.Marshal(loginRequest)

  resp,err=http.Post(fmt.Sprintf("%s/mappa/api/escotistas/login",ts.URL),
    "application/json",
    bytes.NewBuffer(jsonStr))
  if resp.StatusCode != 200 {
    t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
  }

  var loginResponse MappaLoginResponse

  err=json.NewDecoder(resp.Body).Decode(&loginResponse)
  if err!=nil{
    t.Fatalf("Error when decoding login response: %v",err)
  }
  t.Logf("Login Response: %+v",loginResponse)
}
