// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) ClientSide(
		hash string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "hash", hash)
	return revel.MainRouter.Reverse("App.ClientSide", args).URL
}

func (_ tApp) Timing(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Timing", args).URL
}

func (_ tApp) Timing_Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Timing_Login", args).URL
}

func (_ tApp) Auth1_Login(
		jwt string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "jwt", jwt)
	return revel.MainRouter.Reverse("App.Auth1_Login", args).URL
}

func (_ tApp) Auth1(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Auth1", args).URL
}

func (_ tApp) Bypass(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Bypass", args).URL
}

func (_ tApp) Expired_JWT_Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Expired_JWT_Login", args).URL
}

func (_ tApp) Expired_JWT(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Expired_JWT", args).URL
}

func (_ tApp) Leaky_JWT_Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Leaky_JWT_Login", args).URL
}

func (_ tApp) Leaky_JWT(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Leaky_JWT", args).URL
}

func (_ tApp) UserAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.UserAgent", args).URL
}

func (_ tApp) UserAgent_Ping(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.UserAgent_Ping", args).URL
}

func (_ tApp) JWT_None_Check(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.JWT_None_Check", args).URL
}

func (_ tApp) JWT_None(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.JWT_None", args).URL
}

func (_ tApp) GoogleOAuth(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GoogleOAuth", args).URL
}

func (_ tApp) GoogleOAuthCallback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GoogleOAuthCallback", args).URL
}

func (_ tApp) GoogleOAuth_Login(
		jwt string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "jwt", jwt)
	return revel.MainRouter.Reverse("App.GoogleOAuth_Login", args).URL
}


