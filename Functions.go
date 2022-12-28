package main


import (
	"bytes"
	"io"
	"os"
	"strings"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

/*

This file will hold all the function of this project since its no that big it wont be so much functions thats why i put it in a single file 

*/




func GenerateToken(songLink string) string{ // will return a string of Token that we will use to get the link to download the song 

	// obviously the headers that we need to send a request 
  headers := map[string]string{
    "Host": "s64.notube.io",
    "Content-Length": "66",
    "Sec-Ch-Ua": "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\"",
    "Accept": "text/html, */*; q=0.01",
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
    "Sec-Ch-Ua-Mobile": "?0",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36",
    "Sec-Ch-Ua-Platform": "\"Linux\"",
    "Origin": "https://notube.io",
    "Sec-Fetch-Site": "same-site",
    "Sec-Fetch-Mode": "cors",
    "Sec-Fetch-Dest": "empty",
    "Referer": "https://notube.io/",
    "Accept-Language": "en-US,en;q=0.9",
    "Connection": "close",
}
  // this is the body of the request 
  var data = []byte("url="+songLink+"&format=mp3&lang=fr")
  // using this function we generate a response from which we have access to status code , body etc 
  response := httpRequest("https://s64.notube.io:443/recover_weight.php", "POST", data, headers)
  // since response is not string thats why i used ioutil.ReadAll to get []byte 
  body, _ := ioutil.ReadAll(response.Body)
  // body is []byte but i want a string thats why i used string function 
  bodyStringed := string(body)
  // this is just some string manipulation inorder to get the token there are so many ways but i chose the worst i guess 
  bodySplited := strings.Split(bodyStringed, "\"")
  Token := bodySplited[19]
  return Token

}

// this function will generate a pointer to httpRequest that we used to get the body of the response
func httpRequest(targetUrl string, method string, data []byte, headers map[string]string) *http.Response {

	// setting up a new request with the inputs we have 
	request, error := http.NewRequest(method, targetUrl, bytes.NewBuffer(data))
	for k, v := range headers {
		request.Header.Set(k, v)

	}

	// sending the request and saving the output in a variable named response
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}
	response, error := client.Do(request)
	// error checking (obviously)
	if error != nil {
		panic(error)
	}

	return response
}


func DownloadSong(Token string, OutputName string){
	// Creating an mp3 file 
	out,_:= os.Create(OutputName+".mp3")
	defer out.Close()
	// sending a request so we can take its body which is the song 
	resp, _ := http.Get("https://s64.notube.io/download.php?token="+Token)
	/*
	Deepdive -> if we send a request to for example a blog site and take its resp.Body 
	it will be automatically encoded to somthing like ISO-8859-1 / utf-8 .. 
	but we send a request to site containing a song automatically it will be encoded in binary 
	so if copy that binary to a file we get a working mp3 song 
	*/
	io.Copy(out, resp.Body)
	defer resp.Body.Close()

}
















