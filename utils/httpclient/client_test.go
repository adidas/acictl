package httpclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jorgechato/acictl/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	Server          *httptest.Server
	LastRequest     *http.Request
	LastRequestBody string
	ResponseFunc    func(http.ResponseWriter, *http.Request)
}

func (s *ServerSuite) SetupSuite() {
	s.Server = httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := ioutil.ReadAll(r.Body)
			s.LastRequestBody = string(body)
			s.LastRequest = r

			if s.ResponseFunc != nil {
				s.ResponseFunc(w, r)
			}
		}),
	)
}

func (s *ServerSuite) TearDownSuite() {
	s.Server.Close()
}

func (s *ServerSuite) SetupTest() {
	s.ResponseFunc = nil
	s.LastRequest = nil
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) TestRunOK() {
	s.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	c := NewClient(
		s.Server.URL,
		"GET",
		"application/json",
		"",
		"",
	)

	err := c.Run(
		bytes.Buffer{},
		func(res *http.Response, b string) {
			assert.Equal(
				s.T(),
				http.StatusOK,
				res.StatusCode,
				"StatusCodes should be equals",
			)
		},
		false,
		false,
	)

	assert.Nil(s.T(), err, "Error should be nil")
}

func (s *ServerSuite) TestRunFail() {
	s.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}

	c := NewClient(
		s.Server.URL,
		"GET",
		"application/json",
		"",
		"",
	)

	err := c.Run(
		bytes.Buffer{},
		func(res *http.Response, b string) {
			assert.Equal(
				s.T(),
				http.StatusNotFound,
				res.StatusCode,
				"StatusCodes should be equals",
			)
		},
		true,
		true,
	)

	assert.EqualError(
		s.T(),
		err,
		fmt.Sprintf(
			utils.Error().Client,
			fmt.Sprintf(
				"%v %v",
				http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
			),
		),
		"Error should not be nil",
	)
}
