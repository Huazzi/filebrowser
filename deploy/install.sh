#!/bin/bash
set -e

# 创建目录结构
sudo mkdir -p /opt/filebrowser/{bin,config,data,logs}

# 复制文件
sudo cp bin/filebrowser /opt/filebrowser/bin/
sudo cp config/config.json /opt/filebrowser/config/

# 设置权限
sudo chown -R filebrowser:filebrowser /opt/filebrowser
sudo chmod 755 /opt/filebrowser/bin/filebrowser

# 初始化数据库
sudo -u filebrowser /opt/filebrowser/bin/filebrowser -d /opt/filebrowser/data/filebrowser.db config init

# 创建systemd服务
sudo tee /etc/systemd/system/filebrowser.service <<SVC_EOF
[Unit]
Description=FileBrowser
After=network.target

[Service]
User=filebrowser
Group=filebrowser
ExecStart=/opt/filebrowser/bin/filebrowser -c /opt/filebrowser/config/config.json
Restart=always

[Install]
WantedBy=multi-user.target
SVC_EOF

# 启用服务
sudo systemctl daemon-reload
sudo systemctl enable filebrowser
sudo systemctl start filebrowser

# 防火墙设置
sudo ufw allow 8080/tcp
