package host

import (
	"fmt"
	"net/http"
)

type Host struct {
	HostPort     string
	HostHandlers []Handler
}

type Handler struct {
	Pattern        string
	HandleFunction func(w *http.ResponseWriter, r *http.Request)
}

func (host *Host) New(port string, handlers []Handler) *Host {
	return &Host{
		HostPort:     port,
		HostHandlers: handlers,
	}
}

func (host *Host) Run() error {
	err := http.ListenAndServe(host.HostPort, nil)
	if err != nil {
		return err
	}
	return nil
}

func (host *Host) RunAsync() {
	go func() {
		err := http.ListenAndServe(host.HostPort, nil)
		if err != nil {
			fmt.Println(err)
		}
	}()
}
