#!/usr/bin/python
# -*- coding: UTF-8 -*-

import MySQLdb
import json

from dtk_open_platform import DtkOpenPlatform
import json
# 打开数据库连接
db = MySQLdb.connect("172.17.0.2", "root", "123456", "systemdb", charset='utf8' )
# 使用cursor()方法获取操作游标 
cursor = db.cursor()


appKey = '5e804206da2ae'
appSecret = 'ee621cf7c8b5d2d072cff7f6f37f8bc2'
version = 'v1.2.1'
url = 'https://openapi.dataoke.com/api/goods/get-goods-list'
method = 'Get'

MaxPage = 10
i = 0
while i < MaxPage :
	send = DtkOpenPlatform()
	data = {'appKey': appKey,
			'version': version,
			'tchaoshi':1,
			'pageId':i
	}
	i = i + 1 
	response = send.open_platform_send(method=method,url=url,args=data,key=appSecret)
	#print(response)
	items = response['data']['list']
	item_num = len(items)
	index = 0
	while index < item_num :
		item = items[index]
		index = index +1
		print(item)
		#cursor.execute("INSERT INTO goods(dataoke_id,goods_id ,item_link, title,dtitle,`desc`,main_pic,marketing_main_pic,cid,tbcid,original_price,actual_price,commission_rate,coupon_price,coupon_conditions,coupon_link,coupon_start_time,coupon_end_time) VALUES (%d,%s,'%s','%s','%s','%s','%s','%s',%d,%d,%f,%f,%d,%d,'%s','%s','%s','%s' ) ON DUPLICATE KEY UPDATE item_link='%s', title='%s', dtitle='%s', main_pic='%s', marketing_main_pic='%s', original_price=%f, actual_price=%f, commission_rate=%d, coupon_price=%d, coupon_conditions='%s', coupon_link='%s', coupon_start_time='%s', coupon_end_time='%s';",(item['id'], item['goodsId'], item['itemLink'], item['title'], item['dtitle'], item['desc'], item['mainPic'], item['marketingMainPic'], item['cid'], item['tbcid'], item['originalPrice'], item['actualPrice'],item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'], item['couponEndTime'],item['itemLink'], item['title'], item['dtitle'], item['mainPic'], item['marketingMainPic'], item['originalPrice'], item['actualPrice'], item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'],item['couponEndTime']))
		cursor.execute("INSERT INTO goods(dataoke_id,goods_id ,item_link, title,dtitle,`desc`,main_pic,marketing_main_pic,cid,tbcid,original_price,actual_price,commission_rate,coupon_price,coupon_conditions,coupon_link,coupon_start_time,coupon_end_time) VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s ) ON DUPLICATE KEY UPDATE item_link=%s, title=%s, dtitle=%s, main_pic=%s, marketing_main_pic=%s, original_price=%s, actual_price=%s, commission_rate=%s, coupon_price=%s, coupon_conditions=%s, coupon_link=%s, coupon_start_time=%s, coupon_end_time=%s;",(item['id'], item['goodsId'], item['itemLink'], item['title'], item['dtitle'], item['desc'], item['mainPic'], item['marketingMainPic'], item['cid'], item['tbcid'], item['originalPrice'], item['actualPrice'],item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'], item['couponEndTime'],item['itemLink'], item['title'], item['dtitle'], item['mainPic'], item['marketingMainPic'], item['originalPrice'], item['actualPrice'], item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'],item['couponEndTime']))
		#print(sql)
		#cursor.execute(sql)

	db.commit()

# 关闭数据库连接
db.close()
