package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"strconv"
	"syscall"
	"time"
)

func main() {
	url := "https://www.amazon.co.jp/ap/signin?_encoding=UTF8&ignoreAuthState=1&openid.assoc_handle=jpflex&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.ns.pape=http%3A%2F%2Fspecs.openid.net%2Fextensions%2Fpape%2F1.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.co.jp%2Fgp%2Fhelp%2Fcustomer%2Fdisplay.html%2Fref%3Dnav_signin%3FnodeId%3D201945460&switch_account="
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))

	page := open_page(url, driver)

	ap_email := page.FindByID("ap_email")
	ap_password := page.FindByID("ap_password")

	var email string

	fmt.Print("emailを入力してください: ")
	fmt.Scanf(email)

	fmt.Print("passwordを入力してください: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))

	ap_email.Fill(email)
	ap_password.Fill(string(password))



	err =

	time.Sleep(3 * time.Second)
}

func open_page(url string, driver *agouti.WebDriver) (page *agouti.Page) {
	err := driver.Start()
	if err != nil {
		log.Fatalf("ドライバの起動に失敗しました:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("ページの表示に失敗しました:%v", err)
	}

	err = page.Navigate(url)
	if err != nil {
		log.Fatalf("ページのアクセスに失敗しました%v", err)
	}
}