import request from '@/utils/request'
import { type IntrLikeForm } from './types'

enum IntrApi {
  // 点赞
  LikeAPI = '/api/v1/interactive/like'
}

export const likeRequest = (data: IntrLikeForm) => {
  return request.post(IntrApi.LikeAPI, data)
}
