package routes

import (
	"context"
	"fmt"

	log "github.com/bayusamudra5502/go-backend-template/lib/log"
	"github.com/bayusamudra5502/go-backend-template/lib/output"
	"github.com/go-chi/chi/v5"
)

var colorMap = map[string][]output.Color{
	"GET" 	 : {output.BackGreen,   output.ForeWhite},
	"POST"	 : {output.BackBlue,    output.ForeWhite},
	"PUT" 	 : {output.BackCyan,    output.ForeWhite},
	"PATCH"  : {output.BackMagenta, output.ForeWhite},
	"DELETE" : {output.BackRed, 		output.ForeWhite},
}

func colorizeMethod(name string) string {
	res := ""
	val, ok := colorMap[name]

	if ok {
		res = string(val[0]) + string(val[1])
	} else {
		res = string(output.ForeBlack) + string(output.BackWhite) 
	}

	res = res + " "

	res = res + name

	for i := 0; i < 8 - (len(name) + 1); i++ {
		res = res + " "
	}

	return res + string(output.Reset)
}

func PrintChiRoutes(r *chi.Mux, l log.Log) {
	routeData := map[string][]string{}

	for _, route := range r.Routes() {
		for method := range route.Handlers {
			name := route.Pattern

			if routeData[method] == nil {
				routeData[method] = []string{name}
			} else {
				routeData[method] = append(routeData[method], name)
			}
		}
	}

	l.Info(context.Background(), "Routes Information:")
	l.Info(context.Background(), "")

	loggedMethod := []string{
		"GET","POST","PUT","PATCH","DELETE",
	}

	for _, method := range loggedMethod {
		for _, pattern := range routeData[method] {
		l.Info(
			context.Background(), 
			fmt.Sprintf("%s %s", 
					colorizeMethod(method),
					pattern),
				)
			}
		}
			
			
	l.Info(context.Background(), "")
}