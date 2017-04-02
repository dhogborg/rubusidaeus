package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"

	"gopkg.in/gin-gonic/gin.v1"
)

func getImage(c *gin.Context) {
	cmd := exec.Command("raspistill", "-o", "-")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := cmd.Start(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	imgdata, err := ioutil.ReadAll(stdout)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := cmd.Wait(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Data(200, "image/jpeg", imgdata)
}
