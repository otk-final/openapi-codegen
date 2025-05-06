package golang

import "fmt"

// ApiResult ApiResult[T any]
type ApiResult[T any] struct {
	Code string `json:"code,omitempty"` //业务响应码
	Data T      `json:"data,omitempty"` //数据
	Err  string `json:"err,omitempty"`  //业务错误信息
	Msg  string `json:"msg,omitempty"`  //业务响应信息

}

// PageData PageData[T any]
type PageData[T any] struct {
	Entities []T   `json:"entities,omitempty"` //数据
	Page     int32 `json:"page,omitempty"`     //页码
	Size     int32 `json:"size,omitempty"`     //条数
	Total    int64 `json:"total,omitempty"`    //总数

}

// MapResult MapResult[T any]
type MapResult[T any] struct {
	Code string       `json:"code,omitempty"` //响应码
	Data map[string]T `json:"data,omitempty"` //数据

}

// Cat 猫
type Cat struct {
	Foods    []string `json:"foods,omitempty"`    //食物
	Nickname string   `json:"nickname,omitempty"` //昵称

}

// Dog 狗
type Dog struct {
	Color    string `json:"color,omitempty"`    //颜色
	Id       int64  `json:"id,omitempty"`       //ID
	Nickname string `json:"nickname,omitempty"` //昵称
	Price    int    `json:"price,omitempty"`    //售价

}

// User 用户
type User struct {
	Age       int32             `json:"age,omitempty"`       //年龄
	Birthday  string            `json:"birthday,omitempty"`  //生日
	Cats      []*Cat            `json:"cats,omitempty"`      //集合
	Enable    bool              `json:"enable,omitempty"`    //是否启用
	Gods      map[string]*Dog   `json:"gods,omitempty"`      //引用map
	Id        int64             `json:"id,omitempty"`        //ID
	Kv        map[string]string `json:"kv,omitempty"`        //基础map
	Name      string            `json:"name,omitempty"`      //名称
	Profile   any               `json:"profile,omitempty"`   //
	Tags      []string          `json:"tags,omitempty"`      //标签
	Timestamp int64             `json:"timestamp,omitempty"` //时间戳

}

// AnimalControllerAddCat "/animal/cat/add"
type AnimalControllerAddCat struct {
	Client ApiClient[int]
}

// Call 响应long
func (receiver *AnimalControllerAddCat) Call(body *Cat) (int, error) {

	return receiver.Client.Post("/animal/cat/add", body)
}

// AnimalControllerAddDog fmt.Sprintf("/animal/add/%v/dog",kind)
type AnimalControllerAddDog struct {
	Client ApiClient[*Dog]
}

// Call 路径+Body
func (receiver *AnimalControllerAddDog) Call(kind string, body *Dog) (*Dog, error) {

	return receiver.Client.Post(fmt.Sprintf("/animal/add/%v/dog", kind), body)
}

// AnimalControllerCats "/animal/cats"
type AnimalControllerCats struct {
	Client ApiClient[[]*Cat]
}

// Call 响应list
func (receiver *AnimalControllerCats) Call() ([]*Cat, error) {
	params := map[string]any{}
	return receiver.Client.Get("/animal/cats", params)
}

// AnimalControllerDog fmt.Sprintf("/animal/dog/%v",id)
type AnimalControllerDog struct {
	Client ApiClient[*Dog]
}

// Call 响应对象
func (receiver *AnimalControllerDog) Call(id int32) (*Dog, error) {
	params := map[string]any{}
	return receiver.Client.Get(fmt.Sprintf("/animal/dog/%v", id), params)
}

// AnimalControllerMapCats "/animal/map/cats"
type AnimalControllerMapCats struct {
	Client ApiClient[map[string]*Cat]
}

// Call 响应map
func (receiver *AnimalControllerMapCats) Call() (map[string]*Cat, error) {
	params := map[string]any{}
	return receiver.Client.Get("/animal/map/cats", params)
}

// AnimalControllerMapDog "/animal/map/dogs"
type AnimalControllerMapDog struct {
	Client ApiClient[MapResult[*Dog]]
}

// Call 响应Map泛型
func (receiver *AnimalControllerMapDog) Call() (MapResult[*Dog], error) {
	params := map[string]any{}
	return receiver.Client.Get("/animal/map/dogs", params)
}

// AnimalControllerPageDogs "/animal/page/dogs"
type AnimalControllerPageDogs struct {
	Client ApiClient[PageData[*Dog]]
}

// Call 响应Array泛型
func (receiver *AnimalControllerPageDogs) Call() (PageData[*Dog], error) {
	params := map[string]any{}
	return receiver.Client.Get("/animal/page/dogs", params)
}

// AnimalControllerSearchDog fmt.Sprintf("/animal/query/%v/dog",kind)
type AnimalControllerSearchDog struct {
	Client ApiClient[*Dog]
}

// Call 路径+Query
func (receiver *AnimalControllerSearchDog) Call(color string, kind string, name string) (*Dog, error) {

	params := map[string]any{}
	params["color"] = color
	params["name"] = name

	return receiver.Client.Get(fmt.Sprintf("/animal/query/%v/dog", kind), params)
}

// OtherApiGet "/other/long"
type OtherApiGet struct {
	Client ApiClient[string]
}

// Call 响应基础类型
func (receiver *OtherApiGet) Call() (string, error) {
	params := map[string]any{}
	return receiver.Client.Get("/other/long", params)
}

// OtherApiGetContent "/other/string"
type OtherApiGetContent struct {
	Client ApiClient[int]
}

// Call 响应基础类型
func (receiver *OtherApiGetContent) Call() (int, error) {
	params := map[string]any{}
	return receiver.Client.Get("/other/string", params)
}

// OtherApiMultiplePathVariable fmt.Sprintf("/other/boolean/%v/%v",type2,id)
type OtherApiMultiplePathVariable struct {
	Client ApiClient[[]*Dog]
}

// Call 多路径参数
func (receiver *OtherApiMultiplePathVariable) Call(id string, type2 string) ([]*Dog, error) {
	params := map[string]any{}
	return receiver.Client.Get(fmt.Sprintf("/other/boolean/%v/%v", type2, id), params)
}

// UserApiAdd "/user/add"
type UserApiAdd struct {
	Client ApiClient[ApiResult[int64]]
}

// Call 添加用户
func (receiver *UserApiAdd) Call(body *User) (ApiResult[int64], error) {

	return receiver.Client.Post("/user/add", body)
}

// UserApiDelete fmt.Sprintf("/user/delete/%v",id)
type UserApiDelete struct {
	Client ApiClient[ApiResult[bool]]
}

// Call 删除用户
func (receiver *UserApiDelete) Call(id int64) (ApiResult[bool], error) {
	params := map[string]any{}
	return receiver.Client.Delete(fmt.Sprintf("/user/delete/%v", id), params)
}

// UserApiEdit fmt.Sprintf("/user/edit/%v",id)
type UserApiEdit struct {
	Client ApiClient[ApiResult[int64]]
}

// Call 修改用户
func (receiver *UserApiEdit) Call(id int64, body *User) (ApiResult[int64], error) {

	return receiver.Client.Put(fmt.Sprintf("/user/edit/%v", id), body)
}

// UserApiEdit_1 fmt.Sprintf("/user/detail/%v",id)
type UserApiEdit_1 struct {
	Client ApiClient[ApiResult[*User]]
}

// Call 查询用户
func (receiver *UserApiEdit_1) Call(id int64, keyword string) (ApiResult[*User], error) {

	params := map[string]any{}
	params["keyword"] = keyword

	return receiver.Client.Get(fmt.Sprintf("/user/detail/%v", id), params)
}

// UserApiList "/user/list"
type UserApiList struct {
	Client ApiClient[ApiResult[[]*User]]
}

// Call 用户列表
func (receiver *UserApiList) Call() (ApiResult[[]*User], error) {
	params := map[string]any{}
	return receiver.Client.Get("/user/list", params)
}

// UserApiPage "/user/page"
type UserApiPage struct {
	Client ApiClient[ApiResult[PageData[*User]]]
}

// Call 用户分页
func (receiver *UserApiPage) Call(page int32, size int32) (ApiResult[PageData[*User]], error) {

	params := map[string]any{}
	params["page"] = page
	params["size"] = size

	return receiver.Client.Get("/user/page", params)
}
