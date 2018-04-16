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
    zip ${ProgramName}.zip ${ProgramName}
}

# 运行
FuncInvoke() {
    echo "invoking..."
    aws lambda invoke \
        --invokcation-type Event \
        --function-name ${ProgramName} \
        --region ${LambdaRegion} \
        --payload "file://./event.json" \
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


# 允许S3调用Lambda
FuncPermitS3() {
    echo "adding s3 permission..."
    bucket = $2
    aws lambda add-permission \
        --function-name ${ProgramName} \
        --region ${LambdaRegion} \
        --statement-id "`date +%s%m`" \
        --action "lambda:InvokeFunction" \
        --principa s3.amazonaws.com \
        --source-arn arn:aws:s3:::${bucket} \
        --source-account ${LambdaAccountId}
}

# 访问策略
FuncPolicy() {
    echo "getting policy..."
    aws lambda get-policy \
        --function-name ${ProgramName}
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

    "permits3")
        FuncPermitS3
    ;;

    "policy")
        FuncPolicy
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
