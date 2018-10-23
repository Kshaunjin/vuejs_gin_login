package main

import (
	//"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type Account struct {
	Name     string
	Password string
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

const (
	SecretKey = "welcome to steins gate"
	URL       = "127.0.0.1:27017"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.Static("/", "./public")

	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	con := session.DB("accounts").C("testdata")

	//讀取html資料夾
	router.LoadHTMLGlob("public/*")

	//讀取文件
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/ans", func(c *gin.Context) {
		var form Account
		name := c.PostForm("name")
		password := c.PostForm("password")
		//c.String(http.StatusOK, fmt.Sprintf("Login success, name=%s and password=%s.", name, password))

		result := Account{}
		jin := "jin"
		err = con.Find(bson.M{"name": jin}).One(&result)

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if name != result.Name || password != result.Password {
			fmt.Println("Error logging in")
			return
		}
		//JWT PART
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		//expireCookie := time.Now().Add(time.Hour * 1)

		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintln(w, "Error extracting the key")
			fatal(err)
		}

		tokenString, err := token.SignedString([]byte(SecretKey))
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintln("Error while signing the token")
			fatal(err)
		}

		//response := Token{tokenString}

		//cookie := http.Cookie{Name: "Auth", Value: tokenString, Expires: expireCookie, HttpOnly: true}
		//http.SetCookie(c.Writer, &cookie)
		//res, err := json.Marshal(response)
		c.JSON(200, gin.H{
			"token":    tokenString,
			"password": password,
		})

		/*
			c.HTML(http.StatusOK, "ans.html", gin.H{
				"response": response,
				"res":      res,
				"err":      err,
			})
		*/

	})
	sec := router.Group("/api")
	sec.Use()
	{
		sec.GET("/info", func(c *gin.Context) {
			//c.String(http.StatusOK, "infos")
			datas := []int{1, 2, 3, 4, 5}
			test := []int{6, 7, 8, 9, 10}
			c.JSON(200, gin.H{
				"datas": datas,
				"test":  test,
			})
		})

	}
	router.Run(":8000")
}

func MyMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("Auth")
		value := cookie.Value
		token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return []byte(SecretKey), nil
		})
		/*
			token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (interface{}, error) {
					return []byte(SecretKey), nil
				})

		*/
		if err == nil {
			if token.Valid {
				c.Next()
			} else {
				c.String(http.StatusUnauthorized, "Token is not valid")
			}
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized access to this resource ")
		}
	}
}
