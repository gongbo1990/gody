package douyin

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)


func Dy(url string) (map[string]interface{},bool){
	a1 := strings.Split(url, "//")[1]
	uid := strings.Split(a1, "/")[1]
	htm := getHtml(url)
	htm = replaceIcon(htm)
	return getData(htm,uid)
}


func getData(htm,uid string) (map[string]interface{},bool)  {

	var data = make(map[string]interface{})

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(htm))
	if err!=nil{
		log.Fatalln(err)
		return data,false
	}

	//封面图
	avatar, _ := dom.Find(".author .avatar").Attr("src")
	avatar = trim(avatar)

	//昵称
	nickname := trim(dom.Find(".nickname").Text())

	//抖音id
	shortid := trim(dom.Find(".shortid").Text())
	shortid = strings.Split(shortid, "：")[1]

	//简介 signature
	signature := trim(dom.Find(".signature").Text())

	//关注
	focusTxt := trim(dom.Find(".focus .num").Text())
	focusnum := getNum(focusTxt)

	//粉丝
	followTxt := trim(dom.Find(".follower .num").Text())
	follownum := getNum(followTxt)

	//赞
	likedTxt := trim(dom.Find(".liked-num .num").Text())
	likednum := getNum(likedTxt)


	data["uid"] = uid
	data["nickname"] = nickname
	data["shortid"] = shortid
	data["avatar"] = avatar
	data["signature"] = signature
	data["focusnum"] = focusnum
	data["followernum"] = follownum
	data["likednum"] = likednum

	return data,true
}


func replaceIcon(htm string) string{


	arr := map[int][3]string{
		0: {"&#xe603;", "&#xe60d;", "&#xe616;"},
		1: {"&#xe602;", "&#xe60e;", "&#xe618;"},
		2: {"&#xe605;", "&#xe610;", "&#xe617;"},
		3: {"&#xe604;", "&#xe611;", "&#xe61a;"},
		4: {"&#xe606;", "&#xe60c;", "&#xe619;"},
		5: {"&#xe607;", "&#xe60f;", "&#xe61b;"},
		6: {"&#xe608;", "&#xe612;", "&#xe61f;"},
		7: {"&#xe60a;", "&#xe613;", "&#xe61c;"},
		8: {"&#xe60b;", "&#xe614;", "&#xe61d;"},
		9: {"&#xe609;", "&#xe615;", "&#xe61e;"}}

	newMap := make(map[string]int)
	for k,v := range arr  {
		for _,v2:= range v{
			newMap[v2] = k
		}
	}

	for k,v:= range newMap{

		//1、过正则来判断字符串是否匹配
		//上面的例子也可以通过MatchString实现
		if ok, _ := regexp.MatchString(k, htm); ok {
			reg, _ := regexp.Compile(k)
			htm = reg.ReplaceAllString(htm, strconv.Itoa(v))
		}
	}
	return htm
}



func getHtml(url_ string) string {
	req, _ := http.NewRequest("GET", url_, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && data == nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%s", data)
}


//去空
func trim(str string) string {
	reg, _ := regexp.Compile(" ")
	return reg.ReplaceAllString(str, "")
}

//判断是否有w
func getNum(str string) int {
	if ok, _ := regexp.MatchString("w", str); ok {
		reg, _ := regexp.Compile("w")
		str = reg.ReplaceAllString(str, "")

		nums , _ := strconv.ParseFloat(str,64)
		return int(nums*10000)
	}else{
		nums , _ := strconv.ParseFloat(str,64)
		return int(nums)
	}
}