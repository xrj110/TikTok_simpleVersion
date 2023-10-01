# 使用官方 Python 镜像作为基础镜像
FROM python:3.9

# 设置工作目录
WORKDIR /app
RUN apt-get update && apt-get install -y libgl1-mesa-glx
# 复制应用程序代码到镜像中
COPY . /app

# 安装 Flask 和其他依赖项
RUN pip install -r requirements.txt

# 暴露 Flask 应用程序运行的端口
EXPOSE 5000

# 启动 Flask 应用程序
CMD ["python", "generateCover.py"]
