import request from '@/utils/requests'


enum FileApi {
	UPLOAD_FILE = '/api/v1/file/upload'
}

export const uploadFile = (data: any) => {
	return request.post(FileApi.UPLOAD_FILE, data)
}