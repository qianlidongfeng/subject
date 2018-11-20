package server
import(
	"flag"
	"net/http"
	"github.com/qianlidongfeng/log"
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"sync"
)
var(
	port *string
	dir string
	mSubjects map[string]interface{}
	rmSubjects map[string]int16
	lock sync.Mutex
)


func init(){
	var err error
	port =flag.String("p","2580","subject server port")
	flag.Parse()
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(dir+"/subject_conf/subject.json")
	if err != nil{
		log.Fatal(err)
	}
	mSubjects=make(map[string]interface{})
	rmSubjects=make(map[string]int16)
	err = json.Unmarshal(data, &mSubjects)
	if err != nil{
		log.Fatal(err)
	}
	for k,v := range mSubjects{
		rmSubject,err:=strconv.Atoi(k)
		if err != nil{
			log.Fatal(err)
		}
		rmSubjects[v.(string)]=int16(rmSubject)
	}
}

type Server struct{
	port string
}


func New() *Server{
	return &Server{*port}
}

func (this *Server) Run() error{
	router := NewRouter()
	err:=http.ListenAndServe(":"+this.port, router)
	return err
}