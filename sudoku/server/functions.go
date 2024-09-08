package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func RespondGame(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

func RespondSuccess(w http.ResponseWriter, success bool) {
	w.Header().Set("Content-Type", "application/json")
	successJson, successErr := json.Marshal(map[string]bool{"success": success})
	if successErr != nil {
		http.Error(w, successErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(successJson))
}

func RespondToken(w http.ResponseWriter, token string) {
	tokenStructure := map[string]string{"token": token}
	tokenJson, tokenErr := json.Marshal(tokenStructure)
	if tokenErr != nil {
		http.Error(w, tokenErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokenJson))
}

func GetToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	return token
}

func GetUsername(r *http.Request) string {
	return r.Header.Get("X-Username")
}

func GetFromBody(r *http.Request, key []string) map[string]string {
	var data map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		return nil
	}

	// only get items in key
	outputData := make(map[string]string)
	for _, k := range key {
		outputData[k] = data[k]
	}

	return outputData
}

func GetOneFromBody(r *http.Request, key string) string {
	data := GetFromBody(r, []string{key})
	if data == nil {
		return ""
	}
	return data[key]
}

func CheckAuthorization(db *sql.DBData, token string, w http.ResponseWriter, r *http.Request) bool {
	// check for authorization by getting token and username and see if they match in the database
	username := GetUsername(r)

	if token == "" || username == "" {
		fmt.Println("Unauthorized: No token or username")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	valid, err := db.ConfirmToken(token, username)
	if err != nil {
		fmt.Println("Unauthorized: Error checking token")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	if !valid {
		fmt.Println("Unauthorized: Token not valid")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	return true
}

func GetRowColNum(params map[string]string) (int, int, int, error) {
	row, err := strconv.Atoi(params["row"])
	if err != nil {
		return 0, 0, 0, err
	}
	col, err := strconv.Atoi(params["col"])
	if err != nil {
		return 0, 0, 0, err
	}
	num, err := strconv.Atoi(params["num"])
	if err != nil {
		return 0, 0, 0, err
	}
	return row, col, num, nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// add a new function to the server package for:
// handling GET and POST functions and allow {param} to be passed in, then run a function, and have an option to check for authorization

func HandleMethod(
	db *sql.DBData,
	method string,
	path string,
	f func(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string),
	checkAuth bool,
) {
	var params []string
	// go through path and extract all names with {} around it to add to params
	for i := 0; i < len(path); i++ {
		if path[i] == '{' {
			// get the name of the param
			param := ""
			for j := i + 1; j < len(path); j++ {
				if path[j] == '}' {
					params = append(params, param)
					break
				}
				param += string(path[j])
			}
		}
	}
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		fmt.Println("Handling", r.Method, "request to", r.URL.Path)

		// check method if valid
		if method != "ALL" && r.Method != method {
			fmt.Println("Method not allowed", r.Method, "when it should have been", method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// check authorization
		token := GetToken(r)
		if checkAuth && !CheckAuthorization(db, token, w, r) {
			fmt.Println("Unauthorized attempt", path)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// get params from url
		urlParams := make(map[string]string)
		for _, param := range params {
			urlParams[param] = r.PathValue(param)
		}
		// if checkAuth and token then add token to urlParams
		if checkAuth && token != "" {
			urlParams["token"] = token
		}

		// run function
		f(db, w, r, urlParams)
	});
}

// GETHandler is a function that handles GET requests
func HandleGET(
	db *sql.DBData,
	path string,
	f func(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string),
	checkAuth bool,
) {
	HandleMethod(db, "GET", path, f, checkAuth)
}

// POSTHandler is a function that handles POST requests
func HandlePOST(
	db *sql.DBData,
	path string,
	f func(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string),
	checkAuth bool,
) {
	HandleMethod(db, "POST", path, f, checkAuth)
}

// ALLHandler is a function that handles all requests
func HandleALL(
	db *sql.DBData,
	path string,
	f func(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string),
	checkAuth bool,
) {
	HandleMethod(db, "ALL", path, f, checkAuth)
}
