import request from '@/utils/requests'
import { type publishNoteForm } from './type'

enum NoteApi {
  PUBLISH_NOTE = '/api/v1/notes'
}

export const publishNote = (data: publishNoteForm) => {
  return request.post(NoteApi.PUBLISH_NOTE, data)
}
