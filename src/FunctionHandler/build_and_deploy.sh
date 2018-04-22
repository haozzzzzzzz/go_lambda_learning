#!/usr/bin/env bash
ProgramName="FunctionHandler" # Lambda处理器名称
LambdaExecutionRole="lambda_basic_execution" #执行所用的角色名
LambdaRegion="us-east-1" # 可用去
LambdaAccountId="842913648961" # 账号
LambdaMemory=128 # 占用内容

export HTTP_PROXY=http://127.0.0.1:49825
export HTTPS_PROXY=http://127.0.0.1:49825

source ./deploy.sh