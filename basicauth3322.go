package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/base64"
)

var(
	url	= "http://members.3322.net/dyndns/update?system=dyndns&hostname=your.3322.org"
	agent = "Mozilla/4.0 (compatible; MSIE 5.00; Windows 98"
	auth = "Basic " + base64.StdEncoding.EncodeToString([]byte("username:password"))
)

func main(){
	fmt.Println("MiShell tools v0.1")
	fmt.Println("copyright(2014) mi.shell@outlook.com")
	fmt.Println("Site:your.3322.org")
	req,err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", agent)
	req.Header.Set("Authorization", auth)
	
	client := &http.Client{}
	
	resp,_ := client.Do(req)
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)	
    fmt.Println(string(body))

}

