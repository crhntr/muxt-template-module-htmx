package hypertext

//go:generate go run github.com/crhntr/muxt/cmd/muxt generate --receiver-static-type=Server --routes-func=TemplateRoutes
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o=../fake/server.go --fake-name=Server . RoutesReceiver
