package simplegae

import (
	"encoding/json"
	"google.golang.org/appengine"
	"net/http"
)

type ServerInfo struct {
	AppID                  string
	BackendInstanceIndex   int
	Datacenter             string
	DefaultVersionHostname string
	InstanceID             string
	IsDevAppServer         bool
	IsOverQuota            bool
	IsTimeoutError         bool
	ModuleHostname         string
	ModuleName             string
	RequestID              string
	ServerSoftware         string
	ServiceAccount         string
	VersionID              string
}

func init() {
	http.HandleFunc("/serverinfo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var err error
	s := ServerInfo{}
	s.AppID = appengine.AppID(c)
	s.Datacenter = appengine.Datacenter(c)
	s.DefaultVersionHostname = appengine.DefaultVersionHostname(c)
	s.InstanceID = appengine.InstanceID()
	s.IsDevAppServer = appengine.IsDevAppServer()
	s.IsOverQuota = appengine.IsOverQuota(err)
	s.IsTimeoutError = appengine.IsTimeoutError(err)
	s.ModuleHostname, err = appengine.ModuleHostname(c, "default", "1", "instance")
	if err != nil {
		c.Errorf("%s", err)
	}
	s.ModuleName = appengine.ModuleName(c)
	s.RequestID = appengine.RequestID(c)
	s.ServerSoftware = appengine.ServerSoftware()
	s.ServiceAccount, err = appengine.ServiceAccount(c)
	s.VersionID = appengine.VersionID(c)

	w.Header().Add("Content-type", "application/json")
	w.Header().Add("charset", "utf-8")
	json.NewEncoder(w).Encode(s)

}
