package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	handler := new(studentsHandler)
	//这个两个地址是不同的，一个是单个页面
	http.Handle("/students", handler)
	//另一个是必须传入参数的页面
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

//所以一个链接需要处理很多情况
// /students 分割后长度是2
// /students/{id} 分割后长度是3
// /students/{id}/grades  分割后长度是4
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//简单对字符进行处理：
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2: //获取全部信息
		sh.getAll(w, r)
	case 3: //查询
		//提取id
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, r, id)
	case 4: //新增
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//获取全部学生信息
func (sh studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	//将学生信息全部转为JSON，返回给data，最后写入
	data, err := sh.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

//根据ID，搜索学生信息
func (sh studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	//逻辑与获取全部信息几乎一致
	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("学生信息序列化失败: %q", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

//增加学生成绩信息
func (sh studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	//上面的逻辑一样的，没什么好说

	//接着要从URL中获取到要传达的学生成绩信息
	var g Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	//这部分逻辑是获取成绩信息
	//如果学生成绩获取正确，就附加信息
	student.Grades = append(student.Grades, g)
	w.WriteHeader(http.StatusCreated) //201

	data, err := sh.toJSON(g)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "applicaiton/json")
	w.Write(data)
}

//转化为JSON
func (sh studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("学生信息序列化失败: %q", err)
	}
	return b.Bytes(), nil
}
