from dtk_open_platform import DtkOpenPlatform
import json

appKey = '5e804206da2ae'
appSecret = 'ee621cf7c8b5d2d072cff7f6f37f8bc2'
version = 'v1.2.1'
url = 'https://openapi.dataoke.com/api/goods/get-goods-list'
method = 'Get'

send = DtkOpenPlatform()
data = {'appKey': appKey,
        'version': version,
        'tchaoshi':1,
        'pageId':1
}
response = send.open_platform_send(method=method,url=url,args=data,key=appSecret)
print(response)
f1 = open('test1.txt', 'w')
f1.write(json.dumps(response))


url = 'https://openapi.dataoke.com/api/goods/get-goods-details'
send = DtkOpenPlatform()
data = {'appKey': appKey,
        'version': version,
        'id':25424137
}
response = send.open_platform_send(method=method,url=url,args=data,key=appSecret)
print(response)




#转链接
url = 'https://openapi.dataoke.com/api/tb-service/get-privilege-link'
goodsId = ''
couponId = ''


#28439328
#c73c99e2969c298e06d0074e3142634b