package logger_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/michimani/evidentlylocal/logger"

	"github.com/stretchr/testify/assert"
)

func Test_NewEvidentlyLocalLogger(t *testing.T) {
	cases := []struct {
		name      string
		out       io.Writer
		wantError bool
	}{
		{
			name:      "success",
			out:       os.Stdout,
			wantError: false,
		},
		{
			name:      "fail",
			out:       nil,
			wantError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			l, err := logger.NewEvidentlyLocalLogger(c.out)

			if c.wantError {
				asst.Nil(l)
				asst.Error(err)
				return
			}

			asst.NoError(err)
			asst.NotNil(l)
		})
	}
}

func Test_ELLogger_Info(t *testing.T) {
	cases := []struct {
		name   string
		msg    string
		expect string
	}{
		{
			name:   "success",
			msg:    "test",
			expect: `"msg":"test"`,
		},
		{
			name:   "success: empty message",
			msg:    "",
			expect: "",
		},
	}

	out := bytes.Buffer{}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Cleanup(func() {
				out.Reset()
			})

			asst := assert.New(tt)
			l, err := logger.NewEvidentlyLocalLogger(&out)
			asst.NoError(err)
			asst.NotNil(l)

			l.Info(c.msg)

			asst.Contains(out.String(), `"level":"INFO"`, out.String())
			asst.Contains(out.String(), c.expect, out.String())
		})
	}
}

func Test_ELLogger_Warn(t *testing.T) {
	cases := []struct {
		name   string
		msg    string
		expect string
	}{
		{
			name:   "success",
			msg:    "test",
			expect: `"msg":"test"`,
		},
		{
			name:   "success: empty message",
			msg:    "",
			expect: "",
		},
	}

	out := bytes.Buffer{}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Cleanup(func() {
				out.Reset()
			})

			asst := assert.New(tt)
			l, err := logger.NewEvidentlyLocalLogger(&out)
			asst.NoError(err)
			asst.NotNil(l)

			l.Warn(c.msg)

			asst.Contains(out.String(), `"level":"WARN"`, out.String())
			asst.Contains(out.String(), c.expect, out.String())
		})
	}
}

func Test_ELLogger_Error(t *testing.T) {
	cases := []struct {
		name    string
		msg     string
		err     error
		expects []string
	}{
		{
			name:    "success",
			msg:     "test",
			expects: []string{`"msg":"test"`},
		},
		{
			name:    "success: empty message",
			msg:     "",
			expects: []string{},
		},
		{
			name:    "success with error",
			msg:     "test",
			err:     errors.New("test error"),
			expects: []string{`"msg":"test"`, `"error":"test error"`},
		},
	}

	out := bytes.Buffer{}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Cleanup(func() {
				out.Reset()
			})

			asst := assert.New(tt)
			l, err := logger.NewEvidentlyLocalLogger(&out)
			asst.NoError(err)
			asst.NotNil(l)

			l.Error(c.msg, c.err)

			asst.Contains(out.String(), `"level":"ERROR"`, out.String())
			for _, expect := range c.expects {
				asst.Contains(out.String(), expect, out.String())
			}
		})
	}
}
