package v1

import (
	go_restful "github.com/emicklei/go-restful"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

type MetricsHTTPHandler interface {
	Metrics(req *go_restful.Request, resp *go_restful.Response)
}

func RegisterMetricsHTTPServer(container *go_restful.Container, metricHandler MetricsHTTPHandler) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.Path("")
		container.Add(ws)
	}
	container.EnableContentEncoding(false)
	ws.Route(ws.GET("/metrics").
		To(metricHandler.Metrics).
		Produces(go_restful.MIME_JSON))
}
