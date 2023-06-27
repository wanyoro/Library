package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"time"

	//"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("API_SECERT_KEY"))

func Migrate() {

}

type Credentials struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type Claims struct {
	Fullname string `json:"fullname"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashPassword), nil
}

func ComparePasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var resp = map[string]interface{}{"status": "success", "message": "Registered successfully"}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := Person{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	usr, _ := user.GetUser(a.DB)
	if usr != nil {
		resp["status"] = "failed"
		resp["message"] = "User already registered, please login"
		JSON(w, http.StatusBadRequest, resp)
		return
	}
	user.Prepare()
	newPassword, err := HashPassword(user.Password)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	user.Password = newPassword

	err = user.Validate("")
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.FormatError("")
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SavePerson(a.DB)
	// if err != nil {
	// 	formattedError := FormatError(err.Error())

	// 	ERROR(w, http.StatusInternalServerError, formattedError)
	// 	return
	// }
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	JSON(w, http.StatusCreated, userCreated)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	// var credentials Credentials
	// var dbUser Person
	// var resp = map[string]interface{}{"status": "success", "message": "logged in"}

	// err := json.NewDecoder(r.Body).Decode(&credentials)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// dbName := a.DB.Debug().Find(&dbUser, "fullname", credentials.Fullname)
	// fmt.Printf("%s", dbUser.Password)

	// if dbName.RowsAffected == 0 {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(credentials.Password))
	// if err != nil {
	// 	ERROR(w, http.StatusUnauthorized, err)
	// 	fmt.Printf("%s", err)
	// 	return
	// }

	// fmt.Println("password matched")

	// expirationTime := time.Now().Add(time.Minute * 5)

	// claims := &Claims{
	// 	Fullname: credentials.Fullname,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(jwtKey)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// http.SetCookie(w,
	// 	&http.Cookie{
	// 		Name:    "token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	})
	// token, err := AuthToken(dbUser.ID)
	// fmt.Printf("%s", token)
	// if err != nil {
	// 	ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	// resp["token"] = token
	// JSON(w, http.StatusOK, resp)
	// return
	w.Header().Set("Content-Type", "Application/json")
	var resp = map[string]interface{}{"status": "success", "message": "logged in"}

	user := Person{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	usr, err := user.GetUser(a.DB)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if usr == nil {
		resp["status"] = "failed"
		resp["message"] = "Login Failed, please signup"
		JSON(w, http.StatusBadRequest, resp)
		return
	}

	err = ComparePasswordHash(user.Password, usr.Password)
	if err != nil {
		resp["status"] = "failed"
		resp["message"] = "Login failed, please try again"
		JSON(w, http.StatusForbidden, err)
		return
	}
	token, err := AuthToken(usr.ID)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	JSON(w, http.StatusOK, resp)

}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Fullname)))

}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
