import request from '@/utils/request'
import { type NoteFeedQuery, type queryNoteListForm } from './types'

enum NoteApi {
  // Feed流文章列表
  NoteListAPI = '/api/v1/feed/notes',
  // 获取文章列表
  NOTE_LIST = '/api/v1/notes',
  NoteDetialApi = '/api/v1/notes/detail/:uuid',
  NoteByUserApi = '/api/v1/notes/:uuid/published',
  NoteByUserCollectedApi = '/api/v1/notes/:uuid/collected',
  NoteByUserLikedApi = '/api/v1/notes/:uuid/liked'
}

export const getNoteListRequest = (data: NoteFeedQuery) => {
  // Create an instance of NoteFeedQuery
  return request.get(NoteApi.NoteListAPI, { params: data }) // Pass the instance as an argument
}

export const getNoteDetailRequest = (uuid: string) => {
  return request.get(NoteApi.NoteDetialApi.replace(':uuid', uuid))
}

export const getUserNoteListRequest = (data: queryNoteListForm) => {
  return request.get(NoteApi.NOTE_LIST, { params: data })
}

export const getUserNoteByUserRequest = (uuid: string, data: queryNoteListForm) => {
  return request.get(NoteApi.NoteByUserApi.replace(':uuid', uuid), { params: data })
}

export const getUserNoteByUserCollectedRequest = (uuid: string, data: queryNoteListForm) => {
  return request.get(NoteApi.NoteByUserCollectedApi.replace(':uuid', uuid), { params: data })
}

export const getUserNoteByUserLikedRequest = (uuid: string, data: queryNoteListForm) => {
  return request.get(NoteApi.NoteByUserLikedApi.replace(':uuid', uuid), { params: data })
}
