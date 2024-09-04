import request from '@/utils/requests'
import type{  publishNoteForm,queryNoteListForm } from './type'

enum NoteApi {
  PUBLISH_NOTE = '/api/v1/notes',
  NOTE_LIST = '/api/v1/notes'
}

export const publishNote = (data: publishNoteForm) => {
  return request.post(NoteApi.PUBLISH_NOTE, data)
}


export const getNoteList = (data: queryNoteListForm) => {
  return request.get(NoteApi.NOTE_LIST, { params: data })
}
