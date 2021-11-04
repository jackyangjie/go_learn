package main

import "fmt"

type notifier interface {
	notify()
}

type play interface {
	weichar()
	zfb()
}

type user struct {
	name  string
	email string
}

// user 实现 notify 接口
func (u user) notify() {
	fmt.Println("发送邮件给 %s,%s", u.name, u.email)
}

// user 的指针 实现 play接口
func (u *user) weichar() {
	fmt.Println("微信支付...")
}
func (u *user) zfb() {
	fmt.Println("支付宝。。。")
}

func main() {
	u := user{"张三", "www.trs.com"}
	sendEmail(u)
	goPlay(&u)
}

func sendEmail(u notifier) {
	u.notify()

}

func goPlay(p play) {
	p.weichar()
}
