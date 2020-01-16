package main

import (
	"gody/models"
	"gody/pkgs/douyin"
	"net/http"
)

func main()  {



	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		param := r.URL.Query()
		url, ok := param["url"]

		if(ok){
			newUrl := url[0]

			//url :="https://v.douyin.com/qae6LL/"
			data,ok := douyin.Dy(newUrl)
			if(ok){
				models.InsertOrUpdate(data)
			}else{
				_, _ = w.Write([]byte("no data"))
			}
			_, _ = w.Write([]byte("succ"))
		}else{
			_, _ = w.Write([]byte("no url"))
		}

	})
	_ = http.ListenAndServe("127.0.0.1:8080", nil)

}
