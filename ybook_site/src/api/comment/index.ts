import request from '@/utils/request'
import type { CommentForm } from './types'

enum CommentApi {
  // 文章列表
  CommentListAPI = '/api/v1/comments/:uuid',
  SubCommentAPI = '/api/v1/comments',
  NoteDetialApi = '/api/v1/notes/detail/:uuid'
}

// 获取评论列表
export const getCommentListRequest = (uuid: string) => {
  return request.get(CommentApi.CommentListAPI.replace(':uuid', uuid))
}

// 新增评论
export const addCommentRequest = (data: CommentForm) => {
  return request.post(CommentApi.SubCommentAPI, data)
}
