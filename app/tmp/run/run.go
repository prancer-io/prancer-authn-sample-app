// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	"reflect"
	"github.com/revel/revel"
	controllers "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/revel"
	_ "github.com/revel/revel/cache"
	_ "prancer-io/prancer-authn-sample-app/app"
	controllers0 "prancer-io/prancer-authn-sample-app/app/controllers"
	tests "prancer-io/prancer-authn-sample-app/tests"
	"github.com/revel/revel/testing"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					27: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "ClientSide",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "hash", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					44: []string{ 
						"logged_in",
					},
				},
			},
			&revel.MethodType{
				Name: "Timing",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Timing_Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Auth1_Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "jwt", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Auth1",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					159: []string{ 
						"tokenString",
						"username",
					},
				},
			},
			&revel.MethodType{
				Name: "Bypass",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					169: []string{ 
						"forwarded",
						"logged_in",
					},
				},
			},
			&revel.MethodType{
				Name: "Expired_JWT_Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Expired_JWT",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					206: []string{ 
						"tokenString",
						"username",
					},
				},
			},
			&revel.MethodType{
				Name: "Leaky_JWT_Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Leaky_JWT",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					243: []string{ 
						"tokenString",
						"username",
					},
				},
			},
			&revel.MethodType{
				Name: "UserAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					261: []string{ 
						"app",
						"comment",
					},
				},
			},
			&revel.MethodType{
				Name: "UserAgent_Ping",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					272: []string{ 
						"app",
					},
				},
			},
			&revel.MethodType{
				Name: "JWT_None_Check",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "JWT_None",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					475: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "GoogleOAuth",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					481: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "GoogleOAuthCallback",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					502: []string{ 
						"status",
					},
					506: []string{ 
						"result",
						"status",
					},
				},
			},
			&revel.MethodType{
				Name: "GoogleOAuth_Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "jwt", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}
}
