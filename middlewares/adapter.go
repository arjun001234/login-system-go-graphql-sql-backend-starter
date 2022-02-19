package middlewares

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(index http.Handler, as ...Adapter) http.Handler {
	for _, v := range as {
		index = v(index)
	}
	return index
}
