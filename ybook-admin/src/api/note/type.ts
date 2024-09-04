export interface publishNoteForm {
  noteTitle: string
  noteContent: string
  imgList: image[]
  statement: string
  publishTime: number
  private: boolean
}

export interface image {
  name: string
  url: string
}

export interface queryNoteListForm {
  page: number
  size: number
  state: number
}
