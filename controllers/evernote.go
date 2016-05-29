package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mrjones/oauth"
	"github.com/corps/evernote"
	"github.com/corps/evernote/edamtypes"
	_ "net/http"
	_ "fmt"
	_ "errors"
)


type EvernoteController struct {
	beego.Controller
}

type EvernoteClient struct {
	accessToken  string
	sandbox bool
}

/*
func (o *EvernoteClient) getTmpToken() (string, error) {
	var err error
	tmpToken := "aaa"
	destUrl := "https://sandbox.evernote.com/oauth"
	if o.sandbox == false {
		destUrl = "https://www.evernote.com/oauth"
	}

	destUrl = fmt.Sprintf("%s?oauth_callback=%s&oauth_consumer_key=%s" , destUrl, o.callback, o.appkey)
	resp, err := http.Get(destUrl)
	if err != nil {
		msg := fmt.Sprintf("getTmpToken: %s", err)
		beego.Error(msg)
		return "", errors.New(msg)
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	beego.Debug(string(buf))
	beego.Debug("Got tmpToken: ", tmpToken)
	return tmpToken, nil
}
*/
/*
func newEvernoteClient() *EvernoteClient {
	var err error;
	return &EvernoteClient{
		appkey:appkey, 
		callback:callback, 
		sandbox:sandbox,
	}		
}
*/

/* 使用init的话，需要全局变量，但是这个control不能由自己实例化
还考虑到Evernote和印象笔记用同一个control，有请求时再构造client
func init() {
	
}
*/

func newConsumer() *oauth.Consumer {
	consumerKey := beego.AppConfig.String("evernote_consumer_key")	
	consumerSecret := beego.AppConfig.String("evernote_consumer_secret")
	//sandbox, err := beego.AppConfig.Bool("evernote_sandbox")
	//if err != nil {
	//		sandbox = true
	//}

	return oauth.NewConsumer(
	consumerKey,
	consumerSecret,
	oauth.ServiceProvider{
		RequestTokenUrl:   "https://sandbox.evernote.com/oauth",
		AuthorizeTokenUrl: "https://sandbox.evernote.com/OAuth.action",
		AccessTokenUrl:    "https://sandbox.evernote.com/oauth",
	})
}

func (o *EvernoteController) Auth() {
	callback := beego.AppConfig.String("evernote_callback")	
	var err error
	c := newConsumer()
	requestToken, url, err := c.GetRequestTokenAndUrl(callback)
	if err != nil {
		beego.Error(err)
	}

	beego.Debug(requestToken)
	o.SetSession("requestToken", requestToken)

	o.Redirect(url, 302)

}

func (o *EvernoteController) Callback() {
	var err error
	requestToken := o.GetSession("requestToken")
	verifier := o.GetString("oauth_verifier")

	c := newConsumer()
	accessToken, err := c.AuthorizeToken(requestToken.(*oauth.RequestToken), verifier)
	if err != nil {
		beego.Error(err)
	}	
	beego.Debug(accessToken)
	
	token := accessToken.Token
	client := evernote.NewEvernoteClient(evernote.SANDBOX, token)
    url, err := client.GetUserStore().GetNoteStoreUrl(&token)
    if err != nil {
        panic(err)
    }   
    beego.Debug("Got note store url: ", *url)

    noteStore, err := client.FetchNoteStore()
    if err != nil {
        panic(err)
    }   

    notebook, err := noteStore.GetDefaultNotebook(&token)
    if err != nil {
        panic(err)
    }   
    beego.Debug("Got default notebook's name: ", *notebook.Name)

    content := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
    <!DOCTYPE en-note SYSTEM "http://xml.evernote.com/pub/enml2.dtd">
    <en-note><h1>This note is soooo gooood</h1></en-note>`
    title := "Yo momma"
    note, err := noteStore.CreateNote(&token, 
		&edamtypes.Note{Title: &title, Content: &content,  NotebookGuid: notebook.Guid})
    if err != nil {
        panic(err)
    }   
    beego.Debug("Created a note: ", *note.Guid)

}



