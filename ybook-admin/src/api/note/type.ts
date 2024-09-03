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
