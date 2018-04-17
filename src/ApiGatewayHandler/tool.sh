#!/usr/bin/env bash
ProgramName="ApiGatewayHandler" # Lambda处理器名称
LambdaExecutionRole="lambda-gateway-execution-role" #执行所用的角色名
LambdaRegion="us-east-1" # 可用区
LambdaAccountId="842913648961" # 账号
LambdaMemory=128 # 占用内存

source ./deploy.sh
