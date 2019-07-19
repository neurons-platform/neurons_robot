package client

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type HttpReceiver struct {
	lock sync.Mutex
	Addr string
	Mq   []string
}

func (h *HttpReceiver) Handler(w http.ResponseWriter, r *http.Request) {
	h.lock.Lock()
	defer h.lock.Unlock()
	content, err := ioutil.ReadAll(r.Body)
	if err == nil {
		h.Mq = append(h.Mq, string(content))
	}
}

func (h *HttpReceiver) Init() {
	http.HandleFunc("/receive", h.Handler)
	http.ListenAndServe(":8080", nil)
}

func (h *HttpReceiver) Recv() string {
	h.lock.Lock()
	defer h.lock.Unlock()
	if len(h.Mq) > 0 {
		x := h.Mq[0]
		h.Mq = h.Mq[1:]
		return x
	}
	return ""
}
