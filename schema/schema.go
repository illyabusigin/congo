//************************************************************************//
// congo JSON Hyper-schema
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --url=http://localhost
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package schema

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the API JSON schema controller under "/schema.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Schema")
	service.Info("mount", "ctrl", "Schema", "action", "Show", "route", "GET /schema.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSchema)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/schema.json", h)
}

// getSchema is the httprouter handle that returns the API JSON hyper schema.
// func getSchema(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSchema(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/schema+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(schema))
}

// Generated schema
const schema = `{"$schema":"http://json-schema.org/draft-04/hyper-schema","id":"http://localhost/schema","title":"Congo - Conference Management Made Easy","type":"object","properties":{"Account":{"$ref":"#/definitions/Account"},"CreateAccountPayload":{"$ref":"#/definitions/CreateAccountPayload"},"CreateInstancePayload":{"$ref":"#/definitions/CreateInstancePayload"},"CreateSeriesPayload":{"$ref":"#/definitions/CreateSeriesPayload"},"CreateUserPayload":{"$ref":"#/definitions/CreateUserPayload"},"Instance":{"$ref":"#/definitions/Instance"},"Series":{"$ref":"#/definitions/Series"},"UpdateAccountPayload":{"$ref":"#/definitions/UpdateAccountPayload"},"UpdateInstancePayload":{"$ref":"#/definitions/UpdateInstancePayload"},"UpdateSeriesPayload":{"$ref":"#/definitions/UpdateSeriesPayload"},"UpdateUserPayload":{"$ref":"#/definitions/UpdateUserPayload"},"User":{"$ref":"#/definitions/User"},"account":{"$ref":"#/definitions/account"},"instance":{"$ref":"#/definitions/instance"},"series":{"$ref":"#/definitions/series"},"user":{"$ref":"#/definitions/user"}},"definitions":{"Account":{"title":"Mediatype identifier: application/vnd.account+json","type":"object","properties":{"href":{"type":"string","description":"API href of account"},"id":{"type":"integer","description":"ID of account"},"name":{"type":"string","description":"Name of account","minLength":2}},"description":"A tenant account","media":{"type":"application/vnd.account+json"}},"CreateAccountPayload":{"title":"CreateAccountPayload","type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"]},"CreateInstancePayload":{"title":"CreateInstancePayload","type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"]},"CreateSeriesPayload":{"title":"CreateSeriesPayload","type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"]},"CreateUserPayload":{"title":"CreateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"first_name":{"type":"string","minLength":2},"id":{"type":"integer"},"last_name":{"type":"string","minLength":2},"role":{"type":"string"}},"required":["email"]},"Instance":{"title":"Mediatype identifier: application/vnd.instance+json","type":"object","properties":{"href":{"type":"string","description":"API href of instance"},"id":{"type":"integer","description":"ID of Instance"},"name":{"type":"string","minLength":2},"series":{"description":"Series that this instance belongs to","$ref":"#/definitions/Series"}},"description":"An instance of an event or conference","media":{"type":"application/vnd.instance+json"}},"Series":{"title":"Mediatype identifier: application/vnd.series+json","type":"object","properties":{"account":{"description":"Account that owns bottle","$ref":"#/definitions/Account"},"href":{"type":"string","description":"API href of series"},"id":{"type":"integer","description":"ID of series"},"name":{"type":"string","minLength":2}},"description":"A recurring event or conference","media":{"type":"application/vnd.series+json"}},"UpdateAccountPayload":{"title":"UpdateAccountPayload","type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"]},"UpdateInstancePayload":{"title":"UpdateInstancePayload","type":"object","properties":{"name":{"type":"string","minLength":2}}},"UpdateSeriesPayload":{"title":"UpdateSeriesPayload","type":"object","properties":{"name":{"type":"string","minLength":2}}},"UpdateUserPayload":{"title":"UpdateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"first_name":{"type":"string","minLength":2},"id":{"type":"integer"},"last_name":{"type":"string","minLength":2},"role":{"type":"string"}}},"User":{"title":"Mediatype identifier: application/vnd.user+json","type":"object","properties":{"email":{"type":"string","description":"Email address of user","format":"email","minLength":2},"first_name":{"type":"string","description":"First name of user","minLength":2},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"last_name":{"type":"string","description":"Last name of user","minLength":2}},"description":"A user belonging to a tenant account","media":{"type":"application/vnd.user+json"}},"account":{"title":"account","type":"object","properties":{"href":{"type":"string","description":"API href of account"},"id":{"type":"integer","description":"ID of account"},"name":{"type":"string","description":"Name of account","minLength":2}},"description":"A tenant account","media":{"type":"application/vnd.account+json"},"links":[{"title":"create","rel":"create","href":"/api/accounts","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateAccountPayload"}},{"title":"delete","rel":"delete","href":"/api/accounts/{accountID}","method":"DELETE"},{"title":"list","rel":"list","href":"/api/accounts","method":"GET"},{"title":"show","rel":"self","href":"/api/accounts/{accountID}","method":"GET","targetSchema":{"$ref":"#/definitions/Account"},"mediaType":"application/vnd.account+json"},{"title":"update","rel":"update","href":"/api/accounts/{accountID}","method":"PUT","schema":{"description":"update payload","$ref":"#/definitions/UpdateAccountPayload"}}]},"instance":{"title":"instance","type":"object","properties":{"href":{"type":"string","description":"API href of instance"},"id":{"type":"integer","description":"ID of Instance"},"name":{"type":"string","minLength":2},"series":{"description":"Series that this instance belongs to","$ref":"#/definitions/Series"}},"description":"An instance of an event or conference","media":{"type":"application/vnd.instance+json"},"links":[{"title":"create","rel":"create","href":"/api/accounts/{accountID}/series/{seriesID}/instances","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateInstancePayload"}},{"title":"delete","rel":"delete","href":"/api/accounts/{accountID}/series/{seriesID}/instances/{instanceID}","method":"DELETE"},{"title":"list","rel":"list","href":"/api/accounts/{accountID}/series/{seriesID}/instances","method":"GET"},{"title":"show","rel":"self","href":"/api/accounts/{accountID}/series/{seriesID}/instances/{instanceID}","method":"GET","targetSchema":{"$ref":"#/definitions/Instance"},"mediaType":"application/vnd.instance+json"},{"title":"update","rel":"update","href":"/api/accounts/{accountID}/series/{seriesID}/instances/{instanceID}","method":"PATCH","schema":{"description":"update payload","$ref":"#/definitions/UpdateInstancePayload"}}]},"series":{"title":"series","type":"object","properties":{"account":{"description":"Account that owns bottle","$ref":"#/definitions/Account"},"href":{"type":"string","description":"API href of series"},"id":{"type":"integer","description":"ID of series"},"name":{"type":"string","minLength":2}},"description":"A recurring event or conference","media":{"type":"application/vnd.series+json"},"links":[{"title":"create","rel":"create","href":"/api/accounts/{accountID}/series","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateSeriesPayload"}},{"title":"delete","rel":"delete","href":"/api/accounts/{accountID}/series/{seriesID}","method":"DELETE"},{"title":"list","rel":"list","href":"/api/accounts/{accountID}/series","method":"GET"},{"title":"show","rel":"self","href":"/api/accounts/{accountID}/series/{seriesID}","method":"GET","targetSchema":{"$ref":"#/definitions/Series"},"mediaType":"application/vnd.series+json"},{"title":"update","rel":"update","href":"/api/accounts/{accountID}/series/{seriesID}","method":"PATCH","schema":{"description":"update payload","$ref":"#/definitions/UpdateSeriesPayload"}}]},"user":{"title":"user","type":"object","properties":{"email":{"type":"string","description":"Email address of user","format":"email","minLength":2},"first_name":{"type":"string","description":"First name of user","minLength":2},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"last_name":{"type":"string","description":"Last name of user","minLength":2}},"description":"A user belonging to a tenant account","media":{"type":"application/vnd.user+json"},"links":[{"title":"create","rel":"create","href":"/api/accounts/{accountID}/users","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateUserPayload"}},{"title":"delete","rel":"delete","href":"/api/accounts/{accountID}/users/{userID}","method":"DELETE"},{"title":"list","rel":"list","href":"/api/accounts/{accountID}/users","method":"GET"},{"title":"show","rel":"self","href":"/api/accounts/{accountID}/users/{userID}","method":"GET","targetSchema":{"$ref":"#/definitions/User"},"mediaType":"application/vnd.user+json"},{"title":"update","rel":"update","href":"/api/accounts/{accountID}/users/{userID}","method":"PATCH","schema":{"description":"update payload","$ref":"#/definitions/UpdateUserPayload"}}]}},"description":"Multi-tenant conference management application","links":[{"rel":"self","href":"http://localhost"},{"rel":"self","href":"/schema","method":"GET","targetSchema":{"$schema":"http://json-schema.org/draft-04/hyper-schema","additionalProperties":true}}]}`
