package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lava-game/protocol"
)

// RediCon 为redis实例
var RediCon redis.Conn

// ReadHandler 读取庄家信息
func ReadHandler(c *gin.Context) {
	// 获取最新高度
	var chainInfo protocol.ChainInfo
	blockChainInfo, chainErr := redis.String(RediCon.Do("get", "blockchaininfo"))
	if chainErr != nil {
		fmt.Println("sorry,blockchaininfo has some error:", chainErr)
		return
	}
	json.Unmarshal([]byte(blockChainInfo), &chainInfo)
	IndexNumber := strconv.Itoa(int(chainInfo.SlotIndex))

	// 获取slotinfo信息
	accountInfo, slotErr := redis.String(RediCon.Do("get", "order_"+IndexNumber))
	if slotErr != nil {
		fmt.Println("sorry,get slotinfo has some error:", slotErr)
		return
	}
	var slotInfo protocol.SlotInfo
	json.Unmarshal([]byte(accountInfo), &slotInfo)

	c.HTML(200, "read.tmpl", gin.H{
		"longAddr":    slotInfo.LongInfo.Addr,
		"shortAddr":   slotInfo.ShortInfo.Addr,
		"slotIndex":   chainInfo.SlotIndex,
		"total":       slotInfo.Total / 100000000,
		"rate":        float64(slotInfo.LongInfo.Amount) / float64(slotInfo.ShortInfo.Amount),
		"longAmount":  float64(slotInfo.LongInfo.Amount) / 100000000,
		"shortAmount": float64(slotInfo.ShortInfo.Amount) / 100000000,
		"longRight":   float64(slotInfo.Total) / float64(slotInfo.LongInfo.Amount),
		"shortRight":  float64(slotInfo.Total) / float64(slotInfo.ShortInfo.Amount)})
}

// ImageHandler 返回地址二维码
// addr.png
func ImageHandler(c *gin.Context) {
	addrPath := c.Param("addr")
	ext := path.Ext(addrPath)
	addr := addrPath[:len(addrPath)-len(ext)]
	if ext == "" || addr == "" {
		http.NotFound(c.Writer, c.Request)
		return
	}

	// Create the barcode
	qrCode, _ := qr.Encode(addr, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// encode the barcode as png
	w := c.Writer
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	png.Encode(&content, qrCode)
	http.ServeContent(w, c.Request, "image", time.Time{}, bytes.NewReader(content.Bytes()))
}

// OrderHandler 列出所有涨跌地址相关信息
func OrderHandler(c *gin.Context) {
	// 获取最新高度
	var chainInfo protocol.ChainInfo
	blockChainInfo, chainErr := redis.String(RediCon.Do("get", "blockchaininfo"))
	if chainErr != nil {
		fmt.Println("sorry,blockchaininfo has some error:", chainErr)
		return
	}
	json.Unmarshal([]byte(blockChainInfo), &chainInfo)
	IndexNumber := strconv.Itoa(int(chainInfo.SlotIndex))

	// 获取slotinfo信息
	accountInfo, slotErr := redis.String(RediCon.Do("get", "order_"+IndexNumber))
	if slotErr != nil {
		fmt.Println("sorry,get slotinfo has some error:", slotErr)
		return
	}
	var slotInfo protocol.SlotInfo
	json.Unmarshal([]byte(accountInfo), &slotInfo)

	c.JSON(200, gin.H{
		"longAddr":    slotInfo.LongInfo.Addr,
		"shortAddr":   slotInfo.ShortInfo.Addr,
		"slotIndex":   chainInfo.SlotIndex,
		"total":       slotInfo.Total / 100000000,
		"longAmount":  float64(slotInfo.LongInfo.Amount) / 100000000,
		"shortAmount": float64(slotInfo.ShortInfo.Amount) / 100000000})
}
