# 从centos 7镜像开始构建
FROM centos:7
# 创建存放可执行文件的目录
RUN mkdir -p /transferorder
# 创建缓存图片
RUN mkdir -p /data
# 将可执行文件放入
COPY transferorder /transferorder/
# 打开GIN的release模式
ENV GIN_MODE release
# 暴露端口到物理机
EXPOSE 8000 8000
# 切换工作目录
WORKDIR /transferorder/
CMD ["/transferorder/transferorder"]