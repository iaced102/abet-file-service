package accesscontrol

// import (
// 	"fmt"
// 	"aBet/adapters/auth"
// 	"aBet/adapters/request"
// 	"aBet/library"
// 	"os"
// )

// func CheckActionWithCurrentRole(authObject auth.AuthObject, objectIdentifier, action string) bool {
// 	if authObject.IsBa() {
// 		return true
// 	}
// 	roleIdentifier := authObject.GetUserRole()
// 	roleIdentifier = "orgchart:118:196c70e7-024e-486e-8f62-2181baa9c3a6"
// 	return CheckRoleActionRemote(authObject, roleIdentifier, objectIdentifier, action)
// }

// func CheckRoleActionRemote(authObject auth.AuthObject, roleIdentifier, objectIdentifier, action string) bool {
// 	listAction := GetRoleActionRemote(authObject, roleIdentifier, objectIdentifier)
// 	fmt.Println("listAction")
// 	fmt.Println(listAction)
// 	return library.StringInSlice(action, listAction)
// }

// func GetRoleActionRemote(authObject auth.AuthObject, roleIdentifier, objectIdentifier string) []string {
// 	listAction := make([]string, 0)
// 	dataRes, e := request.Make(os.Getenv("ACCESS_CONTROL_SERVICE") + "/roles/" + roleIdentifier + "/accesscontrol/" + objectIdentifier).SetHeaders(map[string]string{"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEzMTEiLCJmaXJzdE5hbWUiOiJUXHUwMGY5bmcxIiwibGFzdE5hbWUiOiJUclx1MWVhN24gTFx1MDBlYSIsInVzZXJOYW1lIjoidHVuZ3RsQGNlbnRvZ3JvdXAudm4iLCJkaXNwbGF5TmFtZSI6IlRyXHUxZWE3biBMXHUwMGVhIFRcdTAwZjluZzEiLCJlbWFpbCI6InR1bmd0bEBjZW50b2dyb3VwLnZuIiwidGVuYW50SWQiOiIxIiwicmVzZXRQc3dkSW5mbyI6bnVsbCwidHlwZSI6InVzZXIiLCJpcCI6IjU4LjE4Ny4yNTAuMTM4IiwidXNlckFnZW50IjoiTW96aWxsYVwvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0XC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWVcLzEwOS4wLjAuMCBTYWZhcmlcLzUzNy4zNiIsImRlbGVnYXRlZEJ5Ijp7ImlkIjoiMTkiLCJlbWFpbCI6ImhvYW5nbmRAY2VudG9ncm91cC52biIsIm5hbWUiOiJOZ3V5XHUxZWM1biBcdTAxMTBcdTAwZWNuaCBIb1x1MDBlMG5nIiwiZXhwIjoiMCJ9LCJ0ZW5hbnQiOnsiaWQiOiIxIn0sImlzX2Nsb3VkIjp0cnVlLCJ0ZW5hbnRfZG9tYWluIjoic3ltcGVyLnZuIiwiZXhwIjoiMCIsImlhdCI6MTY3NTE0MDE2Niwicm9sZSI6Im9yZ2NoYXJ0OjExODoxOTZjNzBlNy0wMjRlLTQ4NmUtOGY2Mi0yMTgxYmFhOWMzYTYifQ==.NTk5NTY5YWRjZWRmN2ZiOWI4ZjljYzYxNDhmZWQ2YmRmMGI1OTBhYTU0M2QzYTI1ZWQzOTYwODAyYzU1YmI0ZjViNjMwYTM2OTE3NzI4YmFkNDllYjU1NWI2NWQ1NWNmZWQ4NzBmMDQxYjYzYjE1ZjY1ODU1ZTAwZWIyOWUxZGZiNmZkYTRkNTU1YTljZGYyYTI0NzAxOTMyMmJjOGJmMjRiZWEzNzY3MDY3ZjRkM2YxMmZjOGU0Yzk3YmNjYjZhMjQ1YzIwNzMzMDgxZTQ1NDRlNTAxMTI2NTIwNWU5MjJhYWRhYWM0MmExNjQ5NDM2NGQxZGRkZjkxZjg4NjhmNTE4ZDkyNjYyZDcyMWUxMDcyYzc3OWU0MDc4NzM1M2IwMWYzNTk4NGE4YWMzODVlYmIwMzZkNmQ1OGExNWYwZTBmNDBmYTBmZWQ1OWQ3YWRiOGM5ZTgxNzRkOWJhN2IzNDJmNGQ1Y2MwNDQzN2ZkZjFkMzNiYjc4MGE1OGIzZTRhYjY4MjAxZTE2YjMxNDdmODY5ZmM2MjQ2Y2Y4NDgwNTc4OWNkMjFhYTQyNmVkMDMxZTNiMWEwODk2OGQ1YTJhNmU3ODBmY2MyNWE0Y2Q5MDNlYTdkZjE5Y2QxNjg1NDNmNWE1NjdkMTgxYjE1N2NjYzQ5MGM5ZGU1YjQxNTVkYTQ="}).Get()
// 	if e != nil || dataRes.Status != 200 {
// 		return listAction
// 	}
// 	dataJsonStr := dataRes.Data.(map[string]interface{})
// 	dataApi := dataJsonStr["data"].([]interface{})
// 	for _, v := range dataApi {
// 		vJson := v.(map[string]interface{})
// 		listAction = append(listAction, vJson["action"].(string))
// 	}
// 	return listAction
// }

// func GetRoleActionRemoteAllObject(authObject auth.AuthObject) map[string]interface{} {
// 	roleIdentifier := authObject.GetUserRole()
// 	roleIdentifier = "orgchart:118:196c70e7-024e-486e-8f62-2181baa9c3a6"
// 	dataRes, e := request.Make(os.Getenv("ACCESS_CONTROL_SERVICE") + "/roles/" + roleIdentifier + "/accesscontrol").SetHeaders(map[string]string{"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEzMTEiLCJmaXJzdE5hbWUiOiJUXHUwMGY5bmcxIiwibGFzdE5hbWUiOiJUclx1MWVhN24gTFx1MDBlYSIsInVzZXJOYW1lIjoidHVuZ3RsQGNlbnRvZ3JvdXAudm4iLCJkaXNwbGF5TmFtZSI6IlRyXHUxZWE3biBMXHUwMGVhIFRcdTAwZjluZzEiLCJlbWFpbCI6InR1bmd0bEBjZW50b2dyb3VwLnZuIiwidGVuYW50SWQiOiIxIiwicmVzZXRQc3dkSW5mbyI6bnVsbCwidHlwZSI6InVzZXIiLCJpcCI6IjU4LjE4Ny4yNTAuMTM4IiwidXNlckFnZW50IjoiTW96aWxsYVwvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0XC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWVcLzEwOS4wLjAuMCBTYWZhcmlcLzUzNy4zNiIsImRlbGVnYXRlZEJ5Ijp7ImlkIjoiMTkiLCJlbWFpbCI6ImhvYW5nbmRAY2VudG9ncm91cC52biIsIm5hbWUiOiJOZ3V5XHUxZWM1biBcdTAxMTBcdTAwZWNuaCBIb1x1MDBlMG5nIiwiZXhwIjoiMCJ9LCJ0ZW5hbnQiOnsiaWQiOiIxIn0sImlzX2Nsb3VkIjp0cnVlLCJ0ZW5hbnRfZG9tYWluIjoic3ltcGVyLnZuIiwiZXhwIjoiMCIsImlhdCI6MTY3NTE0MDE2Niwicm9sZSI6Im9yZ2NoYXJ0OjExODoxOTZjNzBlNy0wMjRlLTQ4NmUtOGY2Mi0yMTgxYmFhOWMzYTYifQ==.NTk5NTY5YWRjZWRmN2ZiOWI4ZjljYzYxNDhmZWQ2YmRmMGI1OTBhYTU0M2QzYTI1ZWQzOTYwODAyYzU1YmI0ZjViNjMwYTM2OTE3NzI4YmFkNDllYjU1NWI2NWQ1NWNmZWQ4NzBmMDQxYjYzYjE1ZjY1ODU1ZTAwZWIyOWUxZGZiNmZkYTRkNTU1YTljZGYyYTI0NzAxOTMyMmJjOGJmMjRiZWEzNzY3MDY3ZjRkM2YxMmZjOGU0Yzk3YmNjYjZhMjQ1YzIwNzMzMDgxZTQ1NDRlNTAxMTI2NTIwNWU5MjJhYWRhYWM0MmExNjQ5NDM2NGQxZGRkZjkxZjg4NjhmNTE4ZDkyNjYyZDcyMWUxMDcyYzc3OWU0MDc4NzM1M2IwMWYzNTk4NGE4YWMzODVlYmIwMzZkNmQ1OGExNWYwZTBmNDBmYTBmZWQ1OWQ3YWRiOGM5ZTgxNzRkOWJhN2IzNDJmNGQ1Y2MwNDQzN2ZkZjFkMzNiYjc4MGE1OGIzZTRhYjY4MjAxZTE2YjMxNDdmODY5ZmM2MjQ2Y2Y4NDgwNTc4OWNkMjFhYTQyNmVkMDMxZTNiMWEwODk2OGQ1YTJhNmU3ODBmY2MyNWE0Y2Q5MDNlYTdkZjE5Y2QxNjg1NDNmNWE1NjdkMTgxYjE1N2NjYzQ5MGM5ZGU1YjQxNTVkYTQ="}).Get()
// 	if e != nil || dataRes.Status != 200 {
// 		return map[string]interface{}{}
// 	}
// 	dataJsonStr := dataRes.Data.(map[string]interface{})
// 	return dataJsonStr
// }

// func GetFilterString(authObject auth.AuthObject, objectIdentifier, action string, andAsPrefix bool) string {
// 	if authObject.IsBa() {
// 		return ""
// 	}

// 	return ""
// }

// func GetOperations(authObject auth.AuthObject, objectIdentifier, action string) []map[string]interface{} {
// 	roleIdentifier := authObject.GetUserRole()
// 	roleIdentifier = "orgchart:118:196c70e7-024e-486e-8f62-2181baa9c3a6"
// 	mapOperation := make([]map[string]interface{}, 0)
// 	dataRes, e := request.Make(os.Getenv("ACCESS_CONTROL_SERVICE") + "/roles/" + roleIdentifier + "/accesscontrol/" + objectIdentifier).SetHeaders(map[string]string{"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEzMTEiLCJmaXJzdE5hbWUiOiJUXHUwMGY5bmcxIiwibGFzdE5hbWUiOiJUclx1MWVhN24gTFx1MDBlYSIsInVzZXJOYW1lIjoidHVuZ3RsQGNlbnRvZ3JvdXAudm4iLCJkaXNwbGF5TmFtZSI6IlRyXHUxZWE3biBMXHUwMGVhIFRcdTAwZjluZzEiLCJlbWFpbCI6InR1bmd0bEBjZW50b2dyb3VwLnZuIiwidGVuYW50SWQiOiIxIiwicmVzZXRQc3dkSW5mbyI6bnVsbCwidHlwZSI6InVzZXIiLCJpcCI6IjU4LjE4Ny4yNTAuMTM4IiwidXNlckFnZW50IjoiTW96aWxsYVwvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0XC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWVcLzEwOS4wLjAuMCBTYWZhcmlcLzUzNy4zNiIsImRlbGVnYXRlZEJ5Ijp7ImlkIjoiMTkiLCJlbWFpbCI6ImhvYW5nbmRAY2VudG9ncm91cC52biIsIm5hbWUiOiJOZ3V5XHUxZWM1biBcdTAxMTBcdTAwZWNuaCBIb1x1MDBlMG5nIiwiZXhwIjoiMCJ9LCJ0ZW5hbnQiOnsiaWQiOiIxIn0sImlzX2Nsb3VkIjp0cnVlLCJ0ZW5hbnRfZG9tYWluIjoic3ltcGVyLnZuIiwiZXhwIjoiMCIsImlhdCI6MTY3NTE0MDE2Niwicm9sZSI6Im9yZ2NoYXJ0OjExODoxOTZjNzBlNy0wMjRlLTQ4NmUtOGY2Mi0yMTgxYmFhOWMzYTYifQ==.NTk5NTY5YWRjZWRmN2ZiOWI4ZjljYzYxNDhmZWQ2YmRmMGI1OTBhYTU0M2QzYTI1ZWQzOTYwODAyYzU1YmI0ZjViNjMwYTM2OTE3NzI4YmFkNDllYjU1NWI2NWQ1NWNmZWQ4NzBmMDQxYjYzYjE1ZjY1ODU1ZTAwZWIyOWUxZGZiNmZkYTRkNTU1YTljZGYyYTI0NzAxOTMyMmJjOGJmMjRiZWEzNzY3MDY3ZjRkM2YxMmZjOGU0Yzk3YmNjYjZhMjQ1YzIwNzMzMDgxZTQ1NDRlNTAxMTI2NTIwNWU5MjJhYWRhYWM0MmExNjQ5NDM2NGQxZGRkZjkxZjg4NjhmNTE4ZDkyNjYyZDcyMWUxMDcyYzc3OWU0MDc4NzM1M2IwMWYzNTk4NGE4YWMzODVlYmIwMzZkNmQ1OGExNWYwZTBmNDBmYTBmZWQ1OWQ3YWRiOGM5ZTgxNzRkOWJhN2IzNDJmNGQ1Y2MwNDQzN2ZkZjFkMzNiYjc4MGE1OGIzZTRhYjY4MjAxZTE2YjMxNDdmODY5ZmM2MjQ2Y2Y4NDgwNTc4OWNkMjFhYTQyNmVkMDMxZTNiMWEwODk2OGQ1YTJhNmU3ODBmY2MyNWE0Y2Q5MDNlYTdkZjE5Y2QxNjg1NDNmNWE1NjdkMTgxYjE1N2NjYzQ5MGM5ZGU1YjQxNTVkYTQ="}).Get()
// 	if e != nil || dataRes.Status != 200 {
// 		return mapOperation
// 	}
// 	dataJsonStr := dataRes.Data.(map[string]interface{})
// 	dataApi := dataJsonStr["data"].([]interface{})
// 	for _, v := range dataApi {
// 		vJson := v.(map[string]interface{})
// 		actionPackId := vJson["actionPackId"]
// 		filter := vJson["filter"]
// 		item := map[string]interface{}{"filter": filter, "actionPackId": actionPackId}
// 		mapOperation = append(mapOperation, item)
// 	}
// 	return mapOperation
// }

// func ReplaceRemoteQueryByPlaceholder(str string, mapPlaceHolder *[]string) string {

// 	// str = preg_replace("/\r\n|\r|\n/", ' ', $str);
// 	// var re = regexp.MustCompile(`\r\n|\r|\n`)
// 	// s := re.ReplaceAllString(str, " ")

// 	// tmpStr := strings.Replace(str, "'ref", "__SYMPER_PLACEHOLDER_STRING_REF ref", -1)
// 	return ""
// }
