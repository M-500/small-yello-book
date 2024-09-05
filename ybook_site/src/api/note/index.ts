import request from '@/utils/request'
import { type NoteFeedQuery } from './types'
import { da } from 'element-plus/es/locales.mjs'

enum NoteApi {
  // 文章列表
  NoteListAPI = '/api/v1/feed/notes'
}

export const getNoteListRequest = (data: NoteFeedQuery) => {
  // Create an instance of NoteFeedQuery
  return request.get(NoteApi.NoteListAPI, { params: data }) // Pass the instance as an argument
}
