package resolvers

//import (
//    "context"
//    "fmt"
//    "gqlgen-todos/graph/model"
//    "net/http"
//    "time"
//)
//
//func (r *queryResolver) State(ctx context.Context) ([]*model.State, error) {
//    req, _ := http.NewRequest("GET", "http://in.goibibo.com/api/v2/states"  , nil)
//    req.Header.Add("Authorization", "Token db681cd65fd17f542e386503b855c6136c8477d1")
//    req.Header.Add("Content-Type", "application/json")
//    q := req.URL.Query()
//    q.Add("state", "gujarat")
//    req.URL.RawQuery = q.Encode()
//    getExperimentClient := &http.Client{Timeout: time.Millisecond * time.Duration(10) ,}
//    res, _ := getExperimentClient.Do(req)
//    fmt.Println(res)
//    panic("")
//}
