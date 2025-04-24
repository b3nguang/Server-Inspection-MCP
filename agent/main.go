package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CommandRequest 定义请求体结构
type CommandRequest struct {
	Command string `json:"command" binding:"required"`
}

// CommandResponse 定义响应体结构
type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func main() {
	// 定义命令行参数
	port := flag.String("port", "", "服务器端口号 (默认: 8080)")
	flag.Parse()

	// 如果命令行参数没有指定端口，则尝试从环境变量获取
	serverPort := *port
	if serverPort == "" {
		serverPort = os.Getenv("SERVER_PORT")
	}

	// 如果环境变量也没有指定端口，则使用默认端口8080
	if serverPort == "" {
		serverPort = "8080"
	}

	// 创建Gin默认路由
	r := gin.Default()

	// 设置根路由，显示服务已启动
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":      "running",
			"message":     "命令执行服务已启动",
			"usage":       "向 /execute 发送POST请求，请求体为 {\"command\": \"你的命令\"}",
			"server_time": time.Now().Format(time.RFC3339),
		})
	})

	// 设置API路由
	r.POST("/execute", executeCommand)

	// 启动服务器
	address := fmt.Sprintf(":%s", serverPort)
	log.Printf("服务器启动在 http://localhost%s\n", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// executeCommand 处理命令执行请求
func executeCommand(c *gin.Context) {
	var req CommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求格式"})
		return
	}

	// 执行命令并获取输出
	output, err := runCommand(req.Command)

	response := CommandResponse{
		Output: output,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(200, response)
}

// runCommand 执行系统命令并返回输出
func runCommand(command string) (string, error) {
	var cmd *exec.Cmd

	// 根据操作系统选择命令解释器
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// 合并标准输出和错误输出
	output := stdout.String()
	if stderr.String() != "" {
		if output != "" {
			output += "\n"
		}
		output += stderr.String()
	}

	// 去除多余的空行
	output = strings.TrimSpace(output)

	return output, err
}
