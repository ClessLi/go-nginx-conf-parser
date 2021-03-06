package test

import (
	"github.com/ClessLi/bifrost/pkg/resolv/nginx"
	"testing"
)

func TestSave(t *testing.T) {
	path, caches, err := nginx.Load("./config_test/nginx.conf")

	if err != nil {
		t.Log(err)
		return
	}

	conf, err := caches.GetConfig(path)
	if err != nil {
		t.Log(err)
		return
	}

	_, serr := nginx.Save(conf)

	if serr != nil {
		t.Log(serr)
	}
}

func TestVerifyAndSave(t *testing.T) {
	path, caches, err := nginx.Load("./config_test/nginx.conf")

	if err != nil {
		t.Log(err)
		return
	}

	conf, err := caches.GetConfig(path)
	if err != nil {
		t.Log(err)
		return
	}

	tmpConf := conf

	conf = nginx.NewConf(nil, "./config_test/nginx.conf")
	//caches = nginx.NewCaches()
	//t.Log(tmpConf.string(&caches))
	t.Log(tmpConf.String())
	//caches = nginx.NewCaches()
	//t.Log(conf.string(&caches))
	t.Log(conf.String())
}
