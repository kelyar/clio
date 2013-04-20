package clio

import (
    "net/http"
    "fmt"
)

var (
    routes = make(map[string] map[string] func () string)
    splat = []string{}
    params = make(map[string]string)
    ctx = context {}
)

type context struct {
    Request *http.Request
    ResponseWriter http.ResponseWriter
}


/**
 *  Creating namespace to place routes by specific http method
 */
func init () {
    methods := []string{"GET", "POST", "PUT", "DELETE"}
    for index := range methods {
        routes[methods[index]] = make (map[string] func () string)
    }
}


/**
 *  Finding correct handler to certain method:path
 */
func router (w http.ResponseWriter, req *http.Request) {

    // finding correct handler
    for rawPattern, _ := range routes[req.Method] {
        pattern := prepearePattern(rawPattern)

        // splitting whole path into parts
        path, paramsString := splitPath(req.URL.String())

        if pattern.MatchString(path) {

            // homage to Sinatra's splat
            splat = pattern.FindAllStringSubmatch(path, 100)[0][1:]

            // filling params
            params = parseParams(paramsString)

            // calling matched handler
            fmt.Fprintln(w, routes[req.Method][rawPattern]())

            // terminal debugging
            println(req.Method, req.URL.String())
            break
        }
    }
}

// vim: noai:ts=4:sw=4