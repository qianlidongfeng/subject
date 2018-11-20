package client
import(
	"net/http"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type Subjects struct{
	subjects map[int16]string
	ids map[string]int16
	server string
}

func New(serverAddr string) Subjects{
	return Subjects{server:serverAddr}
}

func (this *Subjects)Update() error{
	client:=&http.Client{
		Timeout: time.Second * 10,
	}
	respSubjects,err:=client.Get(this.server+"/subjects")
	if err != nil{
		return err
	}
	if respSubjects.StatusCode != 200{
		return err
	}
	defer respSubjects.Body.Close()
	html,err:=ioutil.ReadAll(respSubjects.Body)
	if err != nil{
		return err
	}
	subjects := make(map[string]interface{})
	err=json.Unmarshal(html,&subjects)
	if err != nil{
		return err
	}
	tempSubjects:=make(map[int16]string)
	for k,v:=range subjects{
		id,err:=strconv.Atoi(k)
		if err != nil{
			return err
		}
		tempSubjects[int16(id)] = v.(string)
	}
	tempIds:= make(map[string]int16)
	for k,v:=range tempSubjects{
		tempIds[v]=k
	}
	this.subjects=tempSubjects
	this.ids=tempIds
	return nil
}


func (this *Subjects) GetID(subject string) (id int16,ok bool){
	id,ok=this.ids[subject]
	return
}

func (this *Subjects) GetSubject(id int16) (subject string,ok bool){
	subject,ok = this.subjects[id]
	return
}