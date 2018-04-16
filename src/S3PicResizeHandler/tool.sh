#!/usr/bin/env bash
ProgramName="S3PicResizeHandler" # Lambda处理器名称
LambdaExecutionRole="lambda-s3-execution-role" #执行所用的角色名
LambdaRegion="us-east-1" # 可用去
LambdaAccountId="842913648961" # 账号
LambdaMemory=128 # 占用内容

source ../../deploy/sh/deploy.sh