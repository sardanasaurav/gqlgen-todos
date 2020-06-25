package resolvers

//
//import (
//    "context"
//    "fmt"
//    "gqlgen-todos/db"
//    "gqlgen-todos/graph/model"
//    "log"
//)
//
//func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    choice := model.Choice{}
//    question := model.Question{}
//    choice.QuestionID = input.QuestionID
//    choice.ChoiceText = input.ChoiceText
//    db.First(&question, choice.QuestionID)
//    choice.Question = &question
//    db.Create(&choice)
//    return &choice, nil
//}
//
//
////func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
////    db, err := database.GetDatabase()
////    if err != nil {
////        log.Println("Unable to connect to database", err)
////        return nil, err
////    }
////    defer db.Close()
////    db.Find(&r.choices)
////    return r.choices, nil
////}
//
//func (r *queryResolver) Choice(ctx context.Context, id *string) (*model.Choice, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    r.choice = &model.Choice{}
//    db.Find(&r.choice)
//    fmt.Println("##################################")
//    fmt.Println(r.choice)
//    fmt.Println("##################################")
//    return r.choice, nil
//}
//
//func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    fmt.Println("###############")
//    fmt.Println(r.choices)
//    fmt.Println("###############")
//    db.Find(&r.choices)
//    return r.choices, nil
//}
