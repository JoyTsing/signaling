#!/bin/bash
# 本地创建一个https证书
openssl genrsa -out ca.key 2048  
openssl req -x509 -new -key ca.key -out ca.crt  
openssl genrsa -out server.key 2048  
openssl req -new -key server.key -out server.csr  
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -out server.crt

