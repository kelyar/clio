package routes

import (
    "../../../../clio"
    "../../app/controllers"
)

func BooksRoutes () {
    // root
    clio.Get("/",            controllers.Index)
    clio.Post ("/books",     controllers.BooksCreate)
    clio.Get ("/books",      controllers.Books)
    clio.Get ("/books/*",    controllers.Book)
    clio.Put ("/books/*",    controllers.BookUpdate)
    clio.Delete ("/books",   controllers.BooksRemove)
    clio.Delete ("/books/*", controllers.BookRemove)
}
