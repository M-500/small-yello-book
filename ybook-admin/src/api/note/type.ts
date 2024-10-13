export interface publishNoteForm {
  noteTitle: string
  noteContent: string
  imgList: imageForm[]
  statement: string
  publishTime: number
  private: boolean
  contentType: number
}

export interface imageForm {
  name: string | undefined
  url: string | undefined
}

export interface queryNoteListForm {
  page: number
  size: number
  state: number
}
