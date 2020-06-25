package tool

import (
	"net/http"
	"net/url"
)

type Form struct {
	FormData url.Values
}

func NewForm(r *http.Request) (a *Form, err error) {
    err = r.ParseForm()
    if err != nil {
        return nil, err
    }
    return &Form{
    	FormData: r.Form,
    }, nil
}

func (a *Form)HaveField(b ...string) (ok bool) {
	ok = true
	for _, val := range b {
		if _,ok := a.FormData[val]; !ok {
			return false
		}
	}
	return
}
