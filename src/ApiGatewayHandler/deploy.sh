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
    requestBody="`cat $1`"
    echo ${requestBody}

    requestBody=${requestBody//"\""/"\\\""}
    request=`cat ./event.json`
    payload=${request/"##BODY##"/$requestBody}
    echo ${payload} > ./api_event.json
    aws lambda invoke \
        --function-name ${ProgramName} \
        --region ${LambdaRegion} \
        --payload  "file://./api_event.json"\
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

FuncDelete() {
    echo "deleting..."
    aws lambda delete-function \
        --function-name ${ProgramName}
}


# 允许S3调用Lambda
FuncPermitS3() {
    echo "adding s3 permission..."
    bucket=$1
    if [ -z ${bucket} ]
    then
        echo "no bucket name"
        return
    fi

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

# apigateway
FuncApiGateway() {
    echo "creating api gateway..."
}

# 子命令
case ${command} in
    "build")
        FuncBuild
    ;;

    "deploy")
        FuncBuild
        FuncDeploy
    ;;

    "delete")
        FuncDelete
    ;;

    "invoke")
        FuncInvoke $2
    ;;

    "permits3")
        FuncPermitS3 $2
    ;;

    "policy")
        FuncPolicy
    ;;

    "redeploy")
        FuncBuild
        FuncDelete
        FuncDeploy
    ;;

    "api")
    ;;

    *)
    echo "
        build: 构建Lambda函数
        deploy: 部署Lambda函数到aws
        delete: 删除Lambda函数
        invoke: 运行Lambda函数
        permits3 bucket-name: 添加访问允许
        policy: 访问策略
        redeploy: 重新部署Lambda函数到aws
    "
    ;;
esac
