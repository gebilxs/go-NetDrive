package meta

import "sort"

//FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta: 新增/更新文件元信息
func UpdateFileMetas(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
	return
}

//GetFileMeta:通过Sha1 值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}
