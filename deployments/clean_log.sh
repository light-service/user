#!/bin/bash

# 清除微服务日志
if [ -z "${DATA_PATH}" ]; then
    DATA_PATH="/app/"
fi
MAIN_LOG_PATH="${DATA_PATH}main.log"
GRPC_LOG_PATH="${DATA_PATH}grpc.log"
echo > "${MAIN_LOG_PATH}"
echo > "${GRPC_LOG_PATH}"
