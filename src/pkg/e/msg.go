package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:            "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_GET_TAGS_FAIL:             "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:            "统计标签失败",
	ERROR_ADD_TAG_FAIL:              "新增标签失败",
	ERROR_EDIT_TAG_FAIL:             "修改标签失败",
	ERROR_DELETE_TAG_FAIL:           "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:           "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:           "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	ERROR_EXIST_VERSION:             "已存在该版本名称",
	ERROR_EXIST_VERSION_FAIL:        "获取已存在版本失败",
	ERROR_NOT_EXIST_VERSION:         "该版本不存在",
	ERROR_GET_VERSIONS_FAIL:         "获取所有版本失败",
	ERROR_COUNT_VERSION_FAIL:        "统计版本失败",
	ERROR_ADD_VERSION_FAIL:          "新增版本失败",
	ERROR_EDIT_VERSION_FAIL:         "修改版本失败",
	ERROR_DELETE_VERSION_FAIL:       "删除版本失败",
	ERROR_EXPORT_VERSION_FAIL:       "导出版本失败",
	ERROR_IMPORT_VERSION_FAIL:       "导入版本失败",
	ERROR_NOT_EXIST_RECORD:          "该脚本不存在",
	ERROR_ADD_RECORD_FAIL:           "新增脚本失败",
	ERROR_DELETE_RECORD_FAIL:        "删除脚本失败",
	ERROR_CHECK_EXIST_RECORD_FAIL:   "检查脚本是否存在失败",
	ERROR_EDIT_RECORD_FAIL:          "修改脚本失败",
	ERROR_COUNT_RECORD_FAIL:         "统计脚本失败",
	ERROR_GET_RECORDS_FAIL:          "获取多个脚本失败",
	ERROR_GET_RECORD_FAIL:           "获取单个脚本失败",
	ERROR_GEN_RECORD_POSTER_FAIL:    "生成脚本海报失败",
	ERROR_UPDATE_VERSIONS_FAIL:      "更新版本列表失败",
	ERROR_UPDATE_AFFS_FAIL:          "更新谱面文件列表失败",
	ERROR_DELETE_AFFS_FAIL:          "清除所有谱面失败",
	ERROR_GET_VERSION_FAIL:          "获取版本失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
