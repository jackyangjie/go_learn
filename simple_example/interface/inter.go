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

type admin struct {
	user
	level string
}

// user 实现 notify 接口
func (u user) notify() {
	fmt.Printf("发送邮件给 %s,%s \n", u.name, u.email)
}

// user 的指针 实现 play接口
func (u *user) weichar() {
	fmt.Println("微信支付...用户名：", u.name, "邮箱：", u.email)
}
func (u *user) zfb() {
	fmt.Println("支付宝。。。用户名：", u.name, "邮箱：", u.email)
}

func (a *admin) weichar() {
	fmt.Println("微信支付...用户名：", a.name, "邮箱：", a.email, "级别：", a.level)
}
func (a *admin) zfb() {
	fmt.Println("支付宝...用户名：", a.name, "邮箱：", a.email, "级别：", a.level)
}

func main() {
	u := user{"张三", "www.trs.com"}
	sendEmail(u)
	goPlay(&u)
	a := admin{
		user{"李四", "www.baidu.com"},
		"23",
	}

	//a.notify()

	sendEmail(a)
	goPlay(&a)
	a.zfb()
}

func sendEmail(u notifier) {
	u.notify()

}

func goPlay(p play) {
	p.weichar()
}
