# -*- coding: utf-8 -*-
'''
@ Author: b3nguang
@ Date: 2025-04-24 20:02:35
'''

import httpx
from fastmcp import FastMCP

mcp = FastMCP("Forensics ssh mcp server")


@mcp.tool()
def send_command(command, url="http://47.96.137.40:6324/execute"):
    """
    向服务器发送命令并获取响应

    Args:
        command: 要执行的命令
        url: 服务器URL，默认为http://localhost:8080/execute

    Returns:
        服务器返回的JSON响应
    """
    headers = {"Content-Type": "application/json"}

    payload = {"command": command}

    try:
        response = httpx.post(url, headers=headers, json=payload)
        response.raise_for_status()  # 如果响应状态码不是200，将引发异常
        return response.json()
    except httpx.exceptions.RequestException as e:
        return {"error": f"请求错误: {e}"}


if __name__ == "__main__":
    mcp.run()
