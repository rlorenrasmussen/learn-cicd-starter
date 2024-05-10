package auth

import (
  "errors"
  "net/http"
  "reflect"
  "testing"
)

func TestGetAPIKey(t *testing.T) {
  type test struct {
    header http.Header
    result string
    err error
  }

  tests := []test{
    {header:make(http.Header), result:"", err:errors.New("no authorization header included")},
    {header:make(http.Header), result:"hunter2", err:nil},
    {header:make(http.Header), result:"", err:errors.New("malformed authorization header")},
    {header:make(http.Header), result:"", err:errors.New("malformed authorization header")},
  }

  tests[1].header.Add("Authorization","ApiKey hunter2")
  tests[2].header.Add("Authorization","ApiKey")
  tests[3].header.Add("Authorization","ApiToken hunter2")

  for _, tc := range tests {
    got,got_err := GetAPIKey(tc.header)
    if !reflect.DeepEqual(tc.err,got_err) {
      t.Fatalf("expected %v error, got: %v",tc.err,got_err)
    }
    if !reflect.DeepEqual(tc.result, got) {
      t.Fatalf("expected %v, got: %v", tc.result, got)
    }
  }
}
