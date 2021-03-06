#!/usr/bin/env bash

command=$1

# 编译指令
export GOOS=linux
export GOARCH=amd64

# deploy params
LambdaFuncName=${ProgramName}

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
    zip ${ProgramName}.zip ${ProgramName} config.yaml
}

# 运行
FuncInvoke() {
    echo "invoking..."
    aws lambda invoke \
        --function-name ${ProgramName} \
        --region ${LambdaRegion} \
        --payload "`cat ./event.json`" \
        output.txt

    cat output.txt
}

# 发布
FuncDeploy() {
    echo "deploying..."
    aws lambda create-function \
        --region ${LambdaRegion} \
        --function-name ${ProgramName} \
        --memory ${LambdaMemory} \
        --role "arn:aws:iam::${LambdaAccountId}:role/${LambdaExecutionRole}" \
        --runtime go1.x \
        --zip-file "fileb://./${ProgramName}.zip" \
        --handler ${ProgramName}

    DeployResult=$?
    if [ ${DeployResult} == 0 ]
    then
        echo "deploy successfully."
    else
        echo "deploy failed." >&2;
    fi

}

# 重新部署
FuncRedeploy() {
    echo "redeploying..."
}

# 子命令
case ${command} in
    "build")
        FuncBuild
    ;;

    "deploy")
        FuncDeploy
    ;;

    "invoke")
        FuncInvoke
    ;;

    "redeploy")
        FuncRedeploy
    ;;

    *)
    echo "
        build: 构建Lambda函数
        deploy: 部署Lambda函数到aws
        invoke: 运行Lambda函数
        redeploy: 重新部署Lambda函数到aws
    "
    ;;
esac
