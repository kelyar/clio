package routes

import (
    . "../../../../clio"
    "../../app/controllers"
)

func BooksRoutes () {
    // root
    Get("/",            controllers.Index)
    Post ("/books",     controllers.BooksCreate)
    Get ("/books",      controllers.Books)
    Get ("/books/*",    controllers.Book)
    Put ("/books/*",    controllers.BookUpdate)
    Delete ("/books",   controllers.BooksRemove)
    Delete ("/books/*", controllers.BookRemove)
}
