import request from '@/utils/requests'
import type { publishNoteForm, queryNoteListForm } from './type'

enum NoteApi {
  PUBLISH_NOTE = '/api/v1/notes',
  PUBLISH_NOTE_VIDEO = '/api/v1/notes/video',
  NOTE_LIST = '/api/v1/notes',
  PASS_NOTE = '/api/v1/notes/:id/pass'
}

export const publishNote = (data: publishNoteForm) => {
  return request.post(NoteApi.PUBLISH_NOTE, data)
}

export const getNoteList = (data: queryNoteListForm) => {
  return request.get(NoteApi.NOTE_LIST, { params: data })
}

export const passNote = (id: string) => {
  return request.put(NoteApi.PASS_NOTE.replace(':id', id))
}

// 发布视频笔记
export const publishVideoNote = (data: publishNoteForm) => {
  return request.post(NoteApi.PUBLISH_NOTE_VIDEO, data)
}
