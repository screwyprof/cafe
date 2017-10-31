package handler

//
//import (
//	"arch/tasks/usecase/query"
//	"fmt"
//)
//
//type QueryHandler interface {
//	Handle(c query.UserByIDQuery, result interface{}) error
//}
//
//type UserByIDHandler struct {
//	uh QueryHandler
//}
//
//func NewUserByIDHandler(uh QueryHandler) UserByIDHandler {
//	return UserByIDHandler{uh: uh}
//}
//
//func (h UserByIDHandler) GetUser(ID int) error {
//
//	u := query.UserByIDQueryResult{}
//	err := h.uh.Handle(query.UserByIDQuery{ID: ID}, &u)
//	if err != nil {
//		return err
//	}
//
//	fmt.Print(u)
//	return nil
//}
