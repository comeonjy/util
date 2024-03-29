// @Description  工具函数
// @Author  	 jiangyang  
// @Created  	 2020/11/17 3:32 下午
package tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// 解析src interface中的数据到dst pointer中
// 适用于无法断言的场景
func InterfaceToPointer(dst interface{}, src interface{}) error {
	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(marshal, dst); err != nil {
		return err
	}
	return nil
}

func MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}