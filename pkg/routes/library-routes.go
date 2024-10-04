package routes

import (
	"github.com/gorilla/mux"
	"github.com/thekabi19/CSP3341_A2_code/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// Author routes
	router.HandleFunc("/author/", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.GetAuthorByID).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.DeleteAuthor).Methods("DELETE")
	router.HandleFunc("/author/{authorId}/books", controllers.GetBooksByAuthor).Methods("GET")

	// Member routes
	router.HandleFunc("/member/", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/member/{memberId}", controllers.GetMemberByID).Methods("GET")
	//router.HandleFunc("/member/{memberId}", controllers.DeleteMember).Methods("DELETE")

	// BookLoanInformation routes (Loan records)
	router.HandleFunc("/member/{memberID}/loans", controllers.GetLoansForMember).Methods("GET")
	router.HandleFunc("/loan/", controllers.CreateBookLoanInformation).Methods("POST")

}
