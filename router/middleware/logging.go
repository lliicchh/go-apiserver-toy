package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lliicchh/apiserver/handler"
	"github.com/lliicchh/apiserver/pkg/errno"
	"github.com/willf/pad"
	"io/ioutil"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" {
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method

		ip := c.ClientIP()

		bodylogwriter := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = bodylogwriter

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		var response handler.Response
		if err := json.Unmarshal(bodylogwriter.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response: %s", bodylogwriter.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}",
			latency, ip, pad.Right(method, 5, ""), path, code, message)

	}

}
