package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

const (
	ListDir      = 0x001
	STATIC_DIR   = "./photoweb/public"
	UPLOAD_DIR   = "./photoweb/uploads"
	TEMPLATE_DIR = "./photoweb/tpls"
)

var templates = make(map[string]*template.Template)

func init() {
	log.Println("index init start...")
	for _, tmpl := range []string{"upload", "list"} {
		//template.Must()确保了模板不能解析成功时，一定会触发错误处理流程
		templates[tmpl] = template.Must(template.ParseFiles(TEMPLATE_DIR + "/" + tmpl + ".html"))
	}
}

func main() {
	//mux := http.NewServeMux()
	//staticDirHandler(mux, "assets", STATIC_DIR, 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
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
	check(err)

	locals := make(map[string]interface{})
	images := []string{} //@todo
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images

	renderHtml(w, "list", locals)
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	return templates[tmpl].Execute(w, locals)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	//@todo 未完全理解
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", err)
				log.Println("WARN: panic in %v. -%v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
