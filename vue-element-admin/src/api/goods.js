import request from '@/utils/request'
import { httphost } from '@/utils/global'

export function fetchList(query) {
  return request({
    baseURL: httphost,
    url: '/goods/list',
    method: 'get',
    params: query
  })
}


export function genQQLink(id) {
  return request({
    baseURL: httphost,
    url: '/goods/genqqlink',
    method: 'get',
    params: { id }
  })
}

