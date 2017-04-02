package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "rubusidaeus"
	app.Usage = "Serve image form raspberry pi camera, but quickly"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  8080,
			Usage:  "port for serving",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "username,u",
			Value:  "admin",
			Usage:  "username required to view camera, empty for no account required",
			EnvVar: "USERNAME",
		},
		cli.StringFlag{
			Name:   "password,p",
			Value:  "",
			Usage:  "password for account, random by default, see log output",
			EnvVar: "PASSWORD",
		},
	}

	app.Action = start
	app.Run(os.Args)
}

func start(ctx *cli.Context) {
	r := gin.Default()

	// a health check is nice to have
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"alive": true,
		})
	})

	var authMiddlewares []gin.HandlerFunc

	if u := ctx.GlobalString("username"); u != "" {
		log.Println("username: " + u)

		password := randPassword(12)
		if p := ctx.GlobalString("password"); p != "" {
			password = p
		} else {
			log.Println("password: " + password)
		}

		authMiddlewares = append(authMiddlewares, gin.BasicAuth(gin.Accounts{
			u: password,
		}))
	}

	authorized := r.Group("/camera", authMiddlewares...)
	authorized.GET("/image.jpg", getImage)

	addr := fmt.Sprintf(":%d", ctx.GlobalInt("port"))
	err := r.Run(addr)
	if err != nil {
		log.Panic(err)
	}
}

func randPassword(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var chars = []rune("ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz123456789")
	passwd := make([]rune, length)
	for i := range passwd {
		passwd[i] = chars[rand.Intn(len(chars))]
	}
	return string(passwd)
}
