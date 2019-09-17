package oauth

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type exchangeTask struct {
	client  *http.Client
	values  url.Values
	promise chan exchangeTaskResult
}

type exchangeTaskResult struct {
	token Token
	err   error
}

var (
	exchangeTasksPool = make([]*exchangeTask, 0, 10)
)

func pullExchangeTask() *exchangeTask {
	if len(exchangeTasksPool) == 0 {
		return &exchangeTask{
			client:  &http.Client{},
			values:  url.Values{},
			promise: make(chan exchangeTaskResult, 0),
		}
	}

	task := exchangeTasksPool[0]
	exchangeTasksPool = exchangeTasksPool[1:]

	return task
}

func (t *exchangeTask) dispose() {
	exchangeTasksPool = append(exchangeTasksPool, t)
}

func (t *exchangeTask) exec(s Service, auth AuthResponse) {
	defer t.dispose()

	t.values.Set("code", auth.Code)
	t.values.Set("grant_type", "authorization_code")
	t.values.Set("redirect_uri", auth.RedirectURI)

	req, err := http.NewRequest("POST", BaseURL+"/token", strings.NewReader(t.values.Encode()))
	if err != nil {
		t.promise <- exchangeTaskResult{err: err}
		return
	}

	req.Header.Set("User-Agent", auth.UserAgent)
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(s.ClientID.Get()+":"+s.ClientSecret.Get())),
	)

	resp, err := t.client.Do(req)
	if err != nil {
		t.promise <- exchangeTaskResult{err: err}
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.promise <- exchangeTaskResult{err: err}
		return
	}

	if err = ParseErrorResponse(body); err != nil {
		t.promise <- exchangeTaskResult{err: err}
		return
	}

	token := Token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		t.promise <- exchangeTaskResult{err: err}
		return
	}

	t.promise <- exchangeTaskResult{token: token}
}
