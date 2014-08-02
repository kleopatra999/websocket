package simplegae

import (
	"appengine"
	"encoding/json"
	"net/http"
)

type ServerInfo struct {
	AppID                  string
	BackendHostname        string
	BackendInstanceName    string
	BackendInstanceIndex   int
	Datacenter             string
	DefaultVersionHostname string
	InstanceID             string
	IsCapabilityDisabled   bool
	IsDevAppServer         bool
	IsOverQuota            bool
	IsTimeoutError         bool
	ModuleHostname         string
	ModuleName             string
	// PublicCertificates
	RequestID      string
	ServerSoftware string
	ServiceAccount string
	VersionID      string
}

func init() {
	http.HandleFunc("/serverinfo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var err error
	s := ServerInfo{}
	s.AppID = appengine.AppID(c)
	s.BackendHostname = appengine.BackendHostname(c, "name", 0)
	s.BackendInstanceName, s.BackendInstanceIndex = appengine.BackendInstance(c)
	s.Datacenter = appengine.Datacenter()
	s.DefaultVersionHostname = appengine.DefaultVersionHostname(c)
	s.InstanceID = appengine.InstanceID()
	s.IsCapabilityDisabled = appengine.IsCapabilityDisabled(err)
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
