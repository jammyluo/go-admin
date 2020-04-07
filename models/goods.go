package models

import (
	"time"
)

type Goods struct {
	Id            		int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	DataokeId     		int       `json:"dataokeId" xorm:"not null default 0 bigint(20)"`
	GoodsId       		int       `json:"goodsId" xorm:"not null default 0 comment('淘宝商品id')  bigint(20)"`
	Title         		string    `json:"title" xorm:"not null default '' comment('淘宝标题') VARCHAR(200)"`
	Dtitle        		string    `json:"dtitle" xorm:"not null default '' comment('大淘客短标题') unique VARCHAR(100)"`
	Desc          		string    `json:"desc" xorm:"not null comment('推广文案') index VARCHAR(512)"`
	MainPic       		string    `json:"mainPic" xorm:"not null comment('商品主图链接') VARCHAR(512)"`
	MarketingMainPic 	string    `json:"marketingMainPic" xorm:"not null comment('营销主图链接') VARCHAR(512)"`
	Cid           		int       `json:"cid" xorm:"not null default 1 comment('商品在大淘客的分类id') INT(11)"`
	Tbcid         		int       `json:"tbcid" xorm:"not null default 1 comment('商品在淘宝的分类id') INT(11)"`
	OriginalPrice       float32   `json:"originalPrice" xorm:"not null default 1 comment('商品原价') float(10,1)"`
	ActualPrice         float32   `json:"actualPrice" xorm:"not null default 1 comment('券后价') float(10,1)"`
	CouponPrice         int       `json:"couponPrice" xorm:"not null default 1 comment('优惠券金额') INT(11)"`
	CouponConditions 	string    `json:"couponConditions" xorm:"not null comment('优惠券使用条件') VARCHAR(20)"`
	CouponStartTime 	time.Time `json:"couponStartTime" xorm:"not null default '0000-00-00 00:00:00' comment('上次登录时间') DATETIME"`
	CouponEndTime 		time.Time `json:"couponEndTime" xorm:"not null default '0000-00-00 00:00:00' comment('上次登录时间') DATETIME"`
	Utime         		time.Time `json:"utime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

type SearchGoods struct  {
	Name string `json:"name" xorm:"not null comment('姓名') VARCHAR(50)"`
}
var tGoods = "goods"

func(u *Goods) GetRow() bool {
	has, err := mEngine.Get(u)
	if err==nil &&  has  {
		return true
	}
	return false
}
func (u *Goods) GetAll()([]Goods,error) {
	var goodsList []Goods
	err:=mEngine.Find(&GoodsList)
	return goodsList,err
}
func (u *Goods) GetAllByTitle(title string)([]Goods,error) {
	var goodsList []Goods

	err:=mEngine.Table(tGoods).Where("title like ?",title+"%").Find(&goodsList)
	return goodsList,err
}

func (u *Goods) GetAllPage(paging *common.Paging)([]Goods,error) {
	var goodsList []Goods
	var err error
	paging.Total,err=mEngine.Count(u)
	paging.GetPages()
	if paging.Total<1 {
		return goodsList,err
	}
	err=mEngine.Limit(int(paging.PageSize),int(paging.StartNums)).Find(&goodsList)
	return goodsList,err
}