package routes

import (
	"github.com/gorilla/mux"
	"github.com/thekabi19/CSP3341_A2_code/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Book routes
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// magazine routes
	router.HandleFunc("/magazine/", controllers.CreateMagazine).Methods("POST")
	router.HandleFunc("/magazine/", controllers.GetAllMagazines).Methods("GET")
	router.HandleFunc("/magazine/{magazineId}", controllers.GetMagazineByID).Methods("GET")

	// Author routes
	router.HandleFunc("/author/", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/", controllers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.GetAuthorByID).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.DeleteAuthor).Methods("DELETE")
	router.HandleFunc("/author/{authorId}/books", controllers.GetBooksByAuthor).Methods("GET")

	// Member routes
	router.HandleFunc("/member/", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/member/{memberId}", controllers.GetMemberByID).Methods("GET")
	router.HandleFunc("/members/{memberId}/fees", controllers.GetMemberFees).Methods("GET")
	//router.HandleFunc("/member/{memberId}", controllers.DeleteMember).Methods("DELETE")

	// BookLoanInformation routes (Loan records)
	router.HandleFunc("/member/{memberID}/loans", controllers.GetLoansForMember).Methods("GET")
	router.HandleFunc("/loan/", controllers.CreateLoanInformation).Methods("POST")

}
