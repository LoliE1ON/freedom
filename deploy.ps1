#!/bin/bash

export GOOS=linux
go build server.go

ssh e1on@5.180.138.37 "sudo systemctl stop freedom-server"
scp ./server e1on@5.180.138.37:/usr/sbin/freedom/server
scp ./.env e1on@5.180.138.37:/usr/sbin/freedom/.env
ssh e1on@5.180.138.37 "chmod +x /usr/sbin/freedom/server"
ssh e1on@5.180.138.37 "sudo systemctl start freedom-server"