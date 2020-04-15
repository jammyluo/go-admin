
CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dataoke_id` bigint(20) not null default 0  COMMENT '大淘客ID',
  `goods_id` bigint(20) not null default 0  COMMENT '淘宝商品id',
  `item_link` varchar(100) NOT NULL DEFAULT '' COMMENT '商品淘宝链接',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '淘宝标题',
  `dtitle` varchar(100) NOT NULL DEFAULT '' COMMENT '大淘客短标题',
  `desc` varchar(512) NOT NULL DEFAULT '' COMMENT '推广文案',
  `main_pic` varchar(512) NOT NULL DEFAULT '' COMMENT '商品主图链接',
  `marketing_main_pic` varchar(512) NOT NULL DEFAULT '' COMMENT '营销主图链接',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '商品在大淘客的分类id',
  `tbcid` int(11) NOT NULL DEFAULT 0 COMMENT '商品在淘宝的分类id',
  `original_price` float(10,1) NOT NULL DEFAULT 0 COMMENT '商品原价',
  `actual_price` float(10,1) NOT NULL DEFAULT 0 COMMENT '券后价',
  `commission_rate` int(11) NOT NULL DEFAULT 0 COMMENT '佣金比例',
  `coupon_price` int(11) NOT NULL DEFAULT 0 COMMENT '优惠券金额',
  `coupon_conditions` varchar(100) NOT NULL DEFAULT '' COMMENT '优惠券使用条件',
  `coupon_link` varchar(100) NOT NULL DEFAULT '' COMMENT '优惠券链接',
  `coupon_start_time` datetime NOT NULL DEFAULT 0 COMMENT '优惠券开始时间',
  `coupon_end_time` datetime NOT NULL DEFAULT 0 COMMENT '优惠券结束时间',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (dataoke_id,goods_id),
  key idx_id(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品';


sql = "INSERT INTO goods(dataoke_id,goods_id ,item_link, title,dtitle,`desc`,main_pic,marketing_main_pic,cid,tbcid,original_price,actual_price,commission_rate,coupon_price,coupon_conditions,coupon_link,coupon_start_time,coupon_end_time) VALUES (%d,'%s','%s','%s','%s','%s','%s','%s',%d,%d,%f,%f,%d,%d,'%s','%s','%s','%s' ) ON DUPLICATE KEY UPDATE item_link='%s', title='%s', dtitle='%s', main_pic='%s', marketing_main_pic='%s', original_price=%f, actual_price=%f, commission_rate=%d, coupon_price=%d, coupon_conditions='%s', coupon_link='%s', coupon_start_time='%s', coupon_end_time='%s';" %(item['id'], item['goodsId'], item['itemLink'], item['title'], item['dtitle'], item['desc'], item['mainPic'], item['marketingMainPic'], item['cid'], item['tbcid'], item['originalPrice'], item['actualPrice'],item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'], item['couponEndTime'],item['itemLink'], item['title'], item['dtitle'], item['mainPic'], item['marketingMainPic'], item['originalPrice'], item['actualPrice'], item['commissionRate'], item['couponPrice'], item['couponConditions'],item['couponLink'],item['couponStartTime'],item['couponEndTime'])
