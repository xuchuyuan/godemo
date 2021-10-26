package cases

import "fmt"

type User struct {
	ID     int
	Name   string
	Avatar string
}

func initString() *string {
	str := new(string)
	*str = "EDDYCJY"
	return str
}

func getUserInfo() *User {
	return &User{ID: 10001, Name: "xuchuyuan", Avatar: "http://avatar.wx.com/xcy.jpg"}
}

func main() {
	_ = getUserInfo()
	_ = initString()

	str := new(string)
	*str = "EDDYCJY"
	fmt.Println(str)
}
