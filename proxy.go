package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 设置代理处理函数
	http.HandleFunc("/proxy", proxyHandler)
	
	// 启动服务器
	log.Println("Proxy server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中获取目标URL
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}
	
	// 解析目标URL
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	// 创建新的请求
	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	// 复制请求头
	copyHeaders(proxyReq.Header, r.Header)
	
	// 设置主机头为目标URL的主机
	proxyReq.Host = parsedURL.Host
	
	// 发送请求到目标服务器
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error sending request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	// 设置CORS头，允许所有来源访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	
	// 处理预检请求
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// 复制响应头
	copyHeaders(w.Header(), resp.Header)
	
	// 设置响应状态码
	w.WriteHeader(resp.StatusCode)
	
	// 复制响应体
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Error copying response body: %v", err)
	}
}

// 复制请求头
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		// 跳过一些不需要转发的头
		if strings.EqualFold(k, "Host") || 
		   strings.EqualFold(k, "Origin") || 
		   strings.EqualFold(k, "Referer") {
			continue
		}
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
