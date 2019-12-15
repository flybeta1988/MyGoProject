package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const TPL_DIR = "./src/photoweb/tpls"
const UPLOAD_DIR = "./src/photoweb/uploads"

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		t, err := template.ParseFiles(TPL_DIR + "/upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if !isExists(imagePath) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	//将该路径下的文件从磁盘中读取并作为服务端的返回信息输出给客户端
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//var listHtml string

	locals := make(map[string]interface{})
	images := []string{} //@todo
	for _, fileInfo := range fileInfoArr {
		//imgid := fileInfo.Name()
		//listHtml += "<li><a href=\"/view?id="+imgid+"\">"+imgid+"</a></li>"
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images

	//读取指定模板的内容并且返回一个*template.Template 值
	t, err := template.ParseFiles(TPL_DIR + "/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//根据模板语法来执行模板的渲染,并将渲染后的结果作为HTTP的返回数据输出
	t.Execute(w, locals)

	//io.WriteString(w, "<html><ol>" + listHtml + "</ol></html>")
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
