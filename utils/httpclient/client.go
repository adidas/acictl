package httpclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jorgechato/acictl/utils"
)

// Run simple sugar code http Client
func (c Client) Run(body bytes.Buffer, callback func(res *http.Response, b string), v, d bool) error {
	client := http.Client{}

	req, err := http.NewRequest(c.Method, c.Url, &body)

	if v || d {
		fmt.Printf("+ curl -X %s %v <<< %v\n", c.Method, c.Url, body.String())
		if err != nil && d {
			log.Fatal(err)
		}
	}
	req.Header.Add("Content-Type", c.ContentType)
	req.SetBasicAuth(c.User, c.Password)
	resp, err := client.Do(req)

	if err != nil && d {
		log.Fatal(err)
	}

	if v || d {
		log.Println(resp.Status)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	bodyString := string(bodyBytes)

	if d {
		if _, err := io.Copy(os.Stderr, resp.Body); err != nil {
			log.Fatal(err)
		}
	}

	if resp.StatusCode < 399 && resp.StatusCode > 100 {
		callback(resp, bodyString)
	} else {
		return errors.New(
			fmt.Sprintf(
				utils.ERROR["client"],
				resp.Status,
			),
		)
	}

	return nil
}
