package main

import (
	"clothingShop/api"
	"clothingShop/config"
	"clothingShop/db"
	"clothingShop/engine"
	"clothingShop/entity"
	"log"
	"net/http"
	"reflect"
	"slices"
	"strings"
	"time"
)

type Info struct {
	Name   string
	Access string
	Link   reflect.Value
}

var types map[string]bool
var hdl *api.Handler
var apiMap map[string]map[string]Info
var accessExceptions []string

func init() {
	cfg := config.Get()
	apiMap = make(map[string]map[string]Info)
	apiMap["POST"] = make(map[string]Info)
	apiMap["PUT"] = make(map[string]Info)
	apiMap["DELETE"] = make(map[string]Info)
	apiMap["GET"] = make(map[string]Info)
	maps := cfg.Api

	types = make(map[string]bool)
	types[".png"] = true
	types[".html"] = true
	types[".js"] = true
	types[".svg"] = true
	types[".css"] = true

	hdl = &api.Handler{}
	services := reflect.ValueOf(hdl)
	_struct := reflect.TypeOf(hdl)

	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}
		if _, ok := apiMap[val.Method]; !ok {

		}
		apiMap[val.Method][val.Url] = Info{
			Name:   method.Name,
			Access: "",
			Link:   services.Method(methodNum),
		}
	}

	accessExceptions = cfg.List
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	ctx := engine.Context{
		Response: w,
		Request:  r,
	}

	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")
	pathName := pathArr[0]

	if pathArr[0] == "" {
		sendFile("./static/index.html", ctx)
		return
	}

	if staticUrl, ok := static(path); ok {
		sendFile("./static/"+staticUrl, ctx)
		return
	}

	maps, ok := apiMap[r.Method]
	if !ok {
		w.Write([]byte("No such method"))
	}

	if len(pathArr) > 1 {
		pathName += "/{id}"
	}

	if fun, ok := maps[pathName]; ok {
		if slices.Contains(accessExceptions, fun.Name) || checkAccess(r.Header) {
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(&ctx)
			fun.Link.Call(in)
		} else {
			log.Println("Forbidden")
		}
	}
}

func checkAccess(h http.Header) bool {
	authStr := h.Get("Authorization")

	log.Println(authStr)

	if len(authStr) < 1 {
		return false
	}

	var token entity.Token
	db.DB().Where("token = ? and expired > ?",
		authStr, time.Now()).First(&token)

	return token.Tid != 0
}

func static(path string) (string, bool) {
	splitPath := strings.Split(path, "/")
	fileName := splitPath[len(splitPath)-1]
	splitName := strings.Split(fileName, ".")
	fileExt := "." + splitName[len(splitName)-1]
	if types[fileExt] {
		return path, true
	}
	return "", false
}

func sendFile(path string, ctx engine.Context) {
	//ctx.Response.Write([]byte{})
	http.ServeFile(ctx.Response, ctx.Request, path)
}
