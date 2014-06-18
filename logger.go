package dr

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// AccessLogger wraps a DockerRegistry instance and logs all the API calls it handles
// to the io.Writer out.
type AccessLogger struct {
	reg DockerRegistry
	out io.Writer
}

// logrecord contains the information necessary to create Common Log Format log lines.
type logrecord struct {
	ip                    string
	time                  time.Time
	method, uri, protocol string
	status                int
	responseBytes         int64
	elapsedTime           time.Duration
}

// responseWriterWrapper captures the response status for use in log lines.
type responseWriterWrapper struct {
	status int
	http.ResponseWriter
}

// ServeHTTP handles all of the logging boilerplate and bookkeeping for API calls to
// its DockerRegistry instance.
func (a *AccessLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
		clientIP = clientIP[:colon]
	}

	record := &logrecord{
		ip:          clientIP,
		time:        time.Time{},
		method:      r.Method,
		uri:         r.RequestURI,
		protocol:    r.Proto,
		status:      http.StatusOK,
		elapsedTime: time.Duration(0),
	}

	rww := &responseWriterWrapper{http.StatusOK, w}

	start := time.Now()
	a.reg.ServeHTTP(rww, r)
	end := time.Now()

	record.status = rww.status
	record.time = end.UTC()
	record.elapsedTime = end.Sub(start)

	record.log(a.out)
}

func (lr logrecord) log(out io.Writer) {
	tf := lr.time.Format("02/Jan/2006 03:04:05")
	rl := fmt.Sprintf("%s %s %s", lr.method, lr.uri, lr.protocol)
	fmt.Fprintf(out, "%s - - [%s] \"%s\" %d %d %.4f\n", lr.ip, tf, rl, lr.status, lr.responseBytes, lr.elapsedTime.Seconds())
}

func (rww *responseWriterWrapper) WriteHeader(code int) {
	rww.status = code
	rww.ResponseWriter.WriteHeader(code)
}
