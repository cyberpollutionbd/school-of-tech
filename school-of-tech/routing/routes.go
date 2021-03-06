package routing

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"github.com/grant/school-of-tech/school-of-tech/ws"
)

type RouteType int

const (
	File = iota
	Directory
	Ws
	NotFound
)

type Route struct {
	Type                   RouteType
	AuthenticationRequired bool
	Method                 string
	Path                   string
	PathPrefix             string
	SocketHandler          ws.SocketHandler
	HandlerFunc            http.HandlerFunc
}

type Routes []Route

const STATIC_DIR = "/static/"

func createRoute(path string, method string, handler http.HandlerFunc) Route {
	return Route{
		Type: File,
		AuthenticationRequired: false,
		Path: path,
		Method: method,
		HandlerFunc: handler,
	}
}

func createAuthenticatedRoute(path string, method string, handler http.HandlerFunc) Route {
	return Route{
		Type: File,
		AuthenticationRequired: true,
		Path: path,
		Method: method,
		HandlerFunc: handler,
	}
}

func createDirectoryRoute(pathPrefix string, method string, handler http.HandlerFunc) Route {
	return Route{
		Type: Directory,
		AuthenticationRequired: false,
		PathPrefix: pathPrefix,
		Method: method,
		HandlerFunc: handler,
	}
}

func createNotFoundRoute(handler http.HandlerFunc) Route {
	return Route{
		Type: NotFound,
		AuthenticationRequired: false,
		HandlerFunc: handler,
	}
}

func createWebsocketRoute(path string, handler ws.SocketHandler) Route {
	return Route{
		Type: Ws,
		AuthenticationRequired: false,
		Path: path,
		SocketHandler: handler,
	}
}

var handler = RouteHandler{}
var routes = Routes{
	createRoute("/", gorequest.GET, handler.Index),
	createRoute("/login", gorequest.POST, handler.Login),
	createRoute("/logout", gorequest.POST, handler.Logout),
	createRoute("/signup", gorequest.POST, handler.Signup),
	createAuthenticatedRoute("/db", gorequest.GET, handler.Db),
	createWebsocketRoute(ws.Connection, handler.WebsocketConnect),
	createDirectoryRoute(STATIC_DIR, gorequest.GET, handler.Static),
	createNotFoundRoute(handler.NotFound),
}
