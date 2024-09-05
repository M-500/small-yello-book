export interface publishNoteForm {
  noteTitle: string
  noteContent: string
  imgList: image[]
  statement: string
  publishTime: number
  private: boolean
}

export interface image {
  name: string | undefined
  url: string | undefined
}

export interface queryNoteListForm {
  page: number
  size: number
  state: number
}
