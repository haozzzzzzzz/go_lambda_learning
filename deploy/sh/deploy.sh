#!/usr/bin/env bash

command=$1

# 编译指令
export GOOS=linux
export GOARCH=amd64

# deploy params
LambdaFuncName=${ProgramName}
LambdaMemory=128
LambdaRegion="us-east-1"
LambdaAccountId="842913648961"

# 构建
FuncBuild() {
    echo "deleting old files..."
    if [ -e ${ProgramName} ]
    then
        rm ${ProgramName}
    fi

    if [ -e ${ProgramName}.zip ]
    then
        rm ${ProgramName}.zip
    fi

    echo "building program..."
    go build -o ${ProgramName} ./main.go
    zip ${ProgramName}.zip ${ProgramName}
}


# 发布
FuncDeploy() {
    echo "deploying..."
    source /Users/hao/Documents/Executable/Python3/Env/AWS/bin/activate
    aws lambda create-function \
        --region ${LambdaRegion} \
        --function-name ${ProgramName} \
        --memory ${LambdaMemory} \
        --role "arn:aws:iam::${LambdaAccountId}:role/${LambdaExecutionRole}" \
        --runtime go1.x \
        --zip-file "fileb://./${ProgramName}.zip" \
        --handler ${ProgramName}
    DeployResult=$?
    deactivate

    if [ ${DeployResult} == 0 ]
    then
        echo "deploy successfully."
    else
        echo "deploy failed." >&2;
    fi
}

case ${command} in
    "build")
        FuncBuild
    ;;
    "deploy")
        FuncDeploy
    ;;
    "redeploy")
    ;;
esac
