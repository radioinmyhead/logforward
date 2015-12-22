package logforward

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"

	"clog/parse"
	//"clog/send"

	"github.com/ant0ine/go-json-rest/rest"
)

// parse
type parser interface {
	Decode() (map[string]string, error)
	Set([]byte)
}

// send
type sender interface {
	Send()
	Set(interface{}) error
}

var (
	decode = map[string]parser{
		"json":      &parse.StrJson{},
		"normal":    &parse.StrNormal{},
		"ng_access": &parse.StrNgAccess{},
	}
	sendto = map[string]sender{}
)

func getDecode(t string) (p parser) {
	if p, ok := decode[t]; ok {
		return p
	}
	return nil
}
func getSend(t string) (s sender) {
	if s, ok := sendto[t]; ok {
		return s
	}
	return nil
}

func GetPost(r *rest.Request) (ret []byte, err error) {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return ret, err
	}
	if len(content) == 0 {
		return ret, fmt.Errorf("content is empty")
	}
	return content, nil
}
func GetIp(r *rest.Request) (ret string, err error) {
	ips := strings.Split(r.RemoteAddr, ":")
	return ips[0], nil
}
func LogCol(w rest.ResponseWriter, r *rest.Request) {
	prefix := r.PathParam("prefix")
	host := r.PathParam("host")
	if prefix == "" {
		rest.Error(w, "prefix required", 400)
		return
	}
	if host == "" {
		rest.Error(w, "host required", 400)
		return
	}
	postdata, err := GetPost(r)
	if err != nil {
		rest.Error(w, "postdata required", 400)
		return
	}
	typ := config.getForm(prefix)
	if typ == "" {
		log.Println("get type fail")
		rest.Error(w, "get type fail", 500)
		return
	}
	p := getDecode(typ)
	if p == nil {
		log.Println("no found parser")
		rest.Error(w, "no found parser", 500)
		return
	}

	p.Set(postdata)
	ret, err := p.Decode()
	if err != nil {
		log.Println("data parse fail")
		rest.Error(w, "data parse fail", 500)
		return
	}
	// add required field
	ret["prefix"] = prefix
	ret["hostname"] = host

	// debug
	log.Println(string(postdata), ret)

	ts := config.getTo(prefix)
	for _, t := range ts {
		s := getSend(t)
		if s == nil {
			rest.Error(w, "no found sender", 500)
			return
		}
		if err := s.Set(ret); err != nil {
			rest.Error(w, "set sender fail", 500)
			return
		}
		if err := s.Send(); err != nil {
			log.Println("send data fail")
			rest.Error(w, "send data fail", 500)
			return
		}
	}
	rest.Error(w, "success", 200)
}
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
	getOpt()
}
func main() {
	// http
	api := rest.NewApi()
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Post("/log/:prefix", LogCol),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":"+config.Port, api.MakeHandler()))

}
