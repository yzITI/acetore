package controller

import (
	"acetore/utils"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// TODO: file info
// TODO: large file 断点续传 & 多线程下载
type fileInfo struct {
	name   string
	length int
	hash   string
}

func CopyFile(src string, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

// Receive ...
func Receive(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(400, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	tokens := c.Request.Header.Get("token")
	fmt.Println(tokens)
	v := utils.Verify(tokens, 864000)
	if v == nil {
		c.String(403, fmt.Sprintf("Auth failed"))
		return
	}
	filename := header.Filename
	filepath := "tmp/" + filename
	// if the file doesn't support resume
	if c.Request.Header.Get("Accept-Ranges") != "bytes" {
		// store in tmp
		out, err := os.Create(filepath)
		if err != nil {
			c.String(500, fmt.Sprintf("Internal Server Error : %s", err.Error()))
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			c.String(500, fmt.Sprintf("Internal Server Error : %s", err.Error()))
			return
		}
		hash, err := utils.HashFileMd5(filepath)
		// copy to public
		err = CopyFile(filepath, "public/"+hash)
		if err != nil {
			c.String(500, fmt.Sprintf("Internal Server Error : %s", err.Error()))
			return
		}
		utils.Log("uploaded file " + filename + " at " + v[0] + " through " + v[1])
		c.JSON(200, gin.H{"hash": hash})
	} else {

	}
}
