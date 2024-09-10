import request from '@/utils/request'

enum CommentApi {
  // 文章列表
  CommentListAPI = '/api/v1/comments/:uuid',
  NoteDetialApi = '/api/v1/notes/detail/:uuid'
}


export const getCommentListRequest = (uuid: string) => {
	return request.get(CommentApi.CommentListAPI.replace(':uuid', uuid))
}