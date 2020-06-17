package resolvers

//import (
//    "context"
//    "fmt"
//    "gqlgen-todos/db"
//    "gqlgen-todos/graph/model"
//    "log"
//)
//func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    fmt.Println("input", input.QuestionText, input.PubDate)
//    question := model.Question{}
//    question.QuestionText = input.QuestionText
//    question.PubDate = input.PubDate
//    db.Create(&question)
//    return &question, nil
//}
//
//
//
//func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    db.Find(&r.questions)
//    for _, question := range r.questions {
//        var choices []*model.Choice
//        db.Where(&model.Choice{QuestionID: question.ID}).Find(&choices)
//        question.Choices = choices
//    }
//    return r.questions, nil
//}
//
//func (r *queryResolver) Question(ctx context.Context, id *string) (*model.Question, error) {
//    db, err := database.GetDatabase()
//    if err != nil {
//        log.Println("Unable to connect to database", err)
//        return nil, err
//    }
//    defer db.Close()
//    r.question = &model.Question{}
//    db.Find(&r.question.ID)
//    //for _, question := range r.questions {
//        var choices []*model.Choice
//        db.Where(&model.Choice{QuestionID: r.question.ID}).Find(&choices)
//        r.question.Choices = choices
//    //}
//    return r.question, nil
//}
