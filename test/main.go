package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	_ "github.com/bmizerany/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//创建表格的结构
type Db struct {
	Id      int64
	State   string
	Journey string
}

//获取表格的信息
func indexdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	mods := make([]Db, 0, 10)
	mod := Db{}
	rows, err := db.Query("SELECT * FROM todolist")
	checkErr(err)
	for rows.Next() {
		var id int64
		var state string
		var journey string
		err = rows.Scan(&id, &state, &journey)
		checkErr(err)
		mod.Id = id
		mod.State = state
		mod.Journey = journey
		mods = append(mods, mod)
	}
	db.Close()
	buf, _ := json.Marshal(mods)
	w.Write(buf)

}

//删除指定数据
func Del(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	stm, err := db.Prepare("DELETE  FROM todolist WHERE id=$1")
	checkErr(err)
	_, err = stm.Exec(id)
	checkErr(err)
	db.Close()
}

//更新数据
func Updata(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	stm, err := db.Prepare("UPDATE todolist set state=$1 WHERE id=$2")
	checkErr(err)
	_, err = stm.Exec("已完成", id)
	checkErr(err)
	db.Close()
}

//插入数据
func Insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	journeyStr := r.Form.Get("journey")
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	stm, err := db.Prepare("INSERT INTO todolist(state,journey) VALUES($1,$2)")
	checkErr(err)
	_, err = stm.Exec("未完成", journeyStr)
	checkErr(err)
	db.Close()
}
func index(w http.ResponseWriter, r *http.Request) {
	//读取html文件
	f, _ := os.Open(`./views/index.html`)
	//读取数据
	buf, _ := ioutil.ReadAll(f)
	//写入到响应
	w.Write(buf)
}
func index2(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open(`./views/index2.html`)
	buf, _ := ioutil.ReadAll(f)
	w.Write(buf)
}
func main() {
	//路由
	http.HandleFunc("/", index)
	http.HandleFunc("/index", indexdata)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/del", Del)
	http.HandleFunc("/up", Updata)
	http.HandleFunc("/insert", Insert)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) //访问静态资源
	http.ListenAndServe(":8080", nil)                                                            //监听8080端口
}
