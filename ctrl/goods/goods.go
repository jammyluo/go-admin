package goods

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go-admin/models"
	"go-admin/modules/response"
	"go-admin/public/common"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
)

func List(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	paging := &common.Paging{Page: page, PageSize: limit}
	goodsModel := models.Goods{}
	goodsArr, err := goodsModel.GetAllPage(paging)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data := make(map[string]interface{})
	data["items"] = goodsArr
	data["total"] = paging.Total
	response.ShowData(c, data)
	return
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func makeSign(data map[string]string, appSecret string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := ""
	for _, k := range keys {
		str = str + "&" + k + "=" + data[k]
	}
	str = str + "&key=" + appSecret
	str = strings.Trim(str, "&")
	fmt.Println("makeSign str:%s", str)
	sign := strings.ToUpper(MD5(str))
	fmt.Println("makeSign sign:%s", sign)
	return sign
}

func GenQQLink(c *gin.Context) {
	var params map[string]string
	params = make(map[string]string)
	params["appKey"] = "5e804206da2ae"
	params["version"] = "v1.1.1"
	params["goodsId"] = "609756884330"
	params["couponId"] = "2aca911c696d479bb5aabedeafc122c7"
	params["pid"] = "mm_905880039_1360650444_110096600415"
	sign := makeSign(params, "ee621cf7c8b5d2d072cff7f6f37f8bc2")
	url := "https://openapi.dataoke.com/api/tb-service/get-privilege-link"

	req := HttpRequest.NewRequest()
	res, _ := req.Get(url, map[string]interface{}{
		"appKey":   "5e804206da2ae",
		"version":  "v1.1.1",
		"goodsId":  "609756884330",
		"couponId": "2aca911c696d479bb5aabedeafc122c7",
		"pid":      "mm_905880039_1360650444_110096600415",
		"sign":     sign,
	})
	body, _ := res.Body()
	fmt.Println("RES:%v", string(body))

	return
}

func GenWXLink(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	paging := &common.Paging{Page: page, PageSize: limit}
	goodsModel := models.Goods{}
	goodsArr, err := goodsModel.GetAllPage(paging)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data := make(map[string]interface{})
	data["items"] = goodsArr
	data["total"] = paging.Total
	response.ShowData(c, data)
	return
}
