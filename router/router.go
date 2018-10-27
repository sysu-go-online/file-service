package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/file-service/controller"
	"github.com/sysu-go-online/public-service/types"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	r.Handle("/users/{username}/projects/{projectname}/files", types.ErrorHandler(controller.GetFileStructureHandler)).Methods("GET")
	r.Handle("/users/{username}/projects/{projectname}/files/{filepath:.*}", types.ErrorHandler(controller.GetFileContentHandler)).Methods("GET")
	r.Handle("/users/{username}/projects/{projectname}/files/{filepath:.*}", types.ErrorHandler(controller.UpdateFileHandler)).Methods("PATCH")
	r.Handle("/users/{username}/projects/{projectname}/files/{filepath:.*}", types.ErrorHandler(controller.CreateFileHandler)).Methods("POST")
	r.Handle("/users/{username}/projects/{projectname}/files/{filepath:.*}", types.ErrorHandler(controller.DeleteFileHandler)).Methods("DELETE")

	// project collection

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
