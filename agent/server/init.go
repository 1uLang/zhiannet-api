package server

import (
	"fmt"
	"path/filepath"

	param "github.com/1uLang/zhiannet-api/agent/const"
	"github.com/1uLang/zhiannet-api/agent/model"
	cm "github.com/1uLang/zhiannet-api/common/model"
	"github.com/rs/xid"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(param.CONFIG_PATH)
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	if size := viper.GetInt("Agent.DescribeSize"); size != 0 {
		param.Describe_SIZE = size
	}
	if maxSize := viper.GetInt("Agent.FileMaxSize"); maxSize != 0 {
		param.MAX_FILE_SIZE = maxSize
	}
	if fp := viper.GetString("Agent.StorePath"); fp != "" {
		param.FILE_STORE_PATH = fp
	}
}

// generateSN 生成文件唯一名称
func generateSN(fn string) string {
	guid := xid.New()

	sn := guid.String()
	ext := filepath.Ext(fn)

	return sn + ext
}

// getFPWithID 通过id获取文件的存储路径
func getFPWithID(id int64) (string, error) {
	af := model.AgentFile{}
	cm.MysqlConn.First(&af, id)
	if af.ID == 0 {
		return "", fmt.Errorf("请输入正确的文件ID")
	}

	return af.Path, nil
}

// formatBytes 格式化字节大小
func formatBytes(bytes int) string {
	fb := float64(bytes)
	switch {
	case fb > param.GB:
		return fmt.Sprintf("%.1fGB", fb/param.GB)
	case fb > param.MB:
		return fmt.Sprintf("%.1fMB", fb/param.MB)
	case fb > param.KB:
		return fmt.Sprintf("%.1fKB", fb/param.KB)
	default:
		return fmt.Sprintf("%.1fB", fb)
	}
}
