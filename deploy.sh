GOOS=linux GOARCH=amd64 go build
scp go_uploader ubuntu@54.169.113.185:~/go_uploader
scp go_uploader ubuntu@54.169.113.185:/var/www/go_uploader

# systemd
# [Unit]
# Description= instance to serve go_uploader api
# After=network.target

# [Service]
# User=ubuntu
# Group=www-data

# ExecStart=/var/www/go_uploader/go_uploader

# [Install]
# WantedBy=multi-user.target
