import request from '@/utils/request'
import { type NoteFeedQuery } from './types'
import { da } from 'element-plus/es/locales.mjs'

enum NoteApi {
  // 文章列表
  NoteListAPI = '/api/v1/feed/notes',
  NoteDetialApi = '/api/v1/notes/detail/:uuid'
}

export const getNoteListRequest = (data: NoteFeedQuery) => {
  // Create an instance of NoteFeedQuery
  return request.get(NoteApi.NoteListAPI, { params: data }) // Pass the instance as an argument
}

export const getNoteDetailRequest = (uuid: string) => {
  return request.get(NoteApi.NoteDetialApi.replace(':uuid', uuid))
}
