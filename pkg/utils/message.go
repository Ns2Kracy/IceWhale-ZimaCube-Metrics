package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func AssembleMessage(rows string) string {
	message := `
{
    "header": {
        "template": "blue",
        "title": {
            "content": "Service Monitor",
            "tag": "plain_text"
        }
    },
    "elements": [
        {
            "tag": "table",
            "page_size": 12,
            "row_height": "middle",
            "header_style": {
                "bold": true,
                "background_style": "grey",
                "lines": 1,
                "text_size": "heading",
                "text_align": "center"
            },
            "columns": [
                {
                    "name": "Service",
                    "display_name": "Service",
                    "data_type": "text",
					"width": "auto",
					"align": "center"
                },
				{
					"name": "CPU",
					"display_name": "CPU",
					"data_type": "text",
					"width": "auto",
					"align": "center"
				},
				{
					"name": "MEM",
					"display_name": "MEM",
					"data_type": "text",
					"width": "auto",
					"align": "center"
				}
            ],
            "rows": [
                ` + rows + `
            ]
        }
    ]
}`
	return message
}

func SendMessage(webhookURL, message string) error {
	// 构造请求体的JSON数据
	jsonData := fmt.Sprintf(`{"msg_type":"interactive","card":%s}`, message)
	fmt.Println(jsonData)
	reqBody := bytes.NewBufferString(jsonData)

	// 创建HTTP请求
	req, err := http.NewRequest(http.MethodPost, webhookURL, reqBody)
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取并打印响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response:", string(responseBody))
	return nil
}
