package ps

type UploadResp struct {
	File struct {
		Name     string
		Type     string
		Tmp_name string
		Error    int
		Size     int
		Filename string
	}
}
