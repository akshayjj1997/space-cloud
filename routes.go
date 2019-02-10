package main

func (s *server) routes() {

	// Initialize the routes for the crud operations
	crudRouter := s.router.Methods("POST").PathPrefix("/v1/api/{project}/crud/{dbType}/{col}").Subrouter()
	crudRouter.HandleFunc("/create", s.handleCreate())
	crudRouter.HandleFunc("/read", s.handleRead())
	crudRouter.HandleFunc("/update", s.handleUpdate())
	crudRouter.HandleFunc("/delete", s.handleDelete())

	// Initialize the routes for the user management operations
	userRouter := s.router.PathPrefix("/v1/api/{project}/auth/{dbType}").Subrouter()
	userRouter.Methods("POST").Path("/email/signin").HandlerFunc(s.user.HandleEmailSignIn())
	userRouter.Methods("POST").Path("/email/signup").HandlerFunc(s.user.HandleEmailSignUp())
	userRouter.Methods("GET").Path("/profile/{id}").HandlerFunc(s.user.HandleProfile())
	userRouter.Methods("GET").Path("/profiles").HandlerFunc(s.user.HandleProfile())
}
