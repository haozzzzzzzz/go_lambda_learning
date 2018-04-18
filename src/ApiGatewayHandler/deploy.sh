#!/usr/bin/env bash

command=$1

source ./params.sh

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
    if [ -z $1 ]
    then
        echo "not body json specified."
        return
    fi

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
        --principal s3.amazonaws.com \
        --source-arn arn:aws:s3:::${bucket} \
        --source-account ${LambdaAccountId}
}

# 允许ApiGateway调用Lambda测试
FuncPermitApiGatewayTest() {
    echo "add api gateway permission..."
    aws lambda add-permission \
        --function-name ${ProgramName} \
        --statement-id apigateway-test \
        --action lambda:InvokeFunction \
        --principal apigateway.amazonaws.com \
        --source-arn "arn:aws:execute-api:${LambdaRegion}:${LambdaAccountId}:${1}/*/POST/DynamoDBManager"
}

# 允许ApiGateway调用Lambda 正式
FuncPermitApiGatewayProd() {
    echo "add api gateway permission..."
    aws lambda add-permission \
        --function-name ${ProgramName} \
        --statement-id apigateway-prod \
        --action lambda:InvokeFunction \
        --principal apigateway.amazonaws.com \
        --source-arn "arn:aws:execute-api:${LambdaRegion}:${LambdaAccountId}:${1}/prod/POST/DynamoDBManager"
}

# 访问策略
FuncPolicy() {
    echo "getting policy..."
    aws lambda get-policy \
        --function-name ${ProgramName}
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

    "permit-apigateway-test")
        FuncPermitApiGatewayTest $2
    ;;

    "permit-apigateway-prod")
        FuncPermitApiGatewayProd $2
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
        permit-apigateway-test api-id: 授权ApiGateway调用
        permit-apigateway-prod api-id: 授权ApiGateway调用
    "
    ;;
esac
