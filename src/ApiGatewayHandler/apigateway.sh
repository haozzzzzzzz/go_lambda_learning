#!/usr/bin/env bash

command=$1

source ./params.sh

FuncCreateAPI() {
    aws apigateway create-rest-api \
        --name DynamoDBOperations \
        --region ${LambdaRegion}
}

FuncGetAPI() {
    aws apigateway get-resources \
        --rest-api-id $1
}

FuncCreateResource() {
    aws apigateway create-resource \
        --rest-api-id $1 \
        --parent-id $2 \
        --path-part DynamoDBManager
}

FuncPutMethod() {
    aws apigateway put-method \
        --rest-api-id $1 \
        --resource-id $2 \
        --http-method POST \
        --authorization-type NONE
}

FuncPutIntegration() {
    aws apigateway put-integration \
        --rest-api-id $1 \
        --resource-id $2 \
        --http-method POST \
        --type AWS_PROXY \
        --integration-http-method POST \
        --uri "arn:aws:apigateway:${LambdaRegion}:lambda:path/2015-03-31/functions/arn:aws:lambda:${LambdaRegion}:${LambdaAccountId}:function:${ProgramName}/invocations"
}

FuncPutMethodResponse() {
    aws apigateway put-method-response \
        --rest-api-id $1 \
        --resource-id $2 \
        --http-method POST \
        --status-code 200 \
        --response-models "{\"application/json\": \"Empty\"}"
}

FuncPutIntegrationResponse() {
    aws apigateway put-integration-response \
        --rest-api-id $1 \
        --resource-id $2 \
        --http-method POST \
        --status-code 200 \
        --response-templates "{\"application/json\": \"\"}"
}

FuncCreateDeployment() {
    aws apigateway create-deployment \
        --rest-api-id $1 \
        --stage-name prod
}

FuncTestInvokeMethod() {
    aws apigateway test-invoke-method \
        --rest-api-id $1 \
        --resource-id $2 \
        --http-method POST \
        --path-with-query-string ""
}

case ${command} in
    "create-api")
        FuncCreateAPI
    ;;

    "get-api")
        FuncGetAPI $2
    ;;

    "create-resource")
        FuncCreateResource $2 $3
    ;;

    "put-method")
        FuncPutMethod $2 $3
    ;;

    "put-integration")
        FuncPutIntegration $2 $3
    ;;

    "put-method-response")
        FuncPutMethodResponse $2 $3
    ;;

    "put-integration-response")
        FuncPutIntegrationResponse $2 $3
    ;;

    "create-deployment")
        FuncCreateDeployment $2
    ;;

    "test-invoke-method")
        FuncTestInvokeMethod $2 $3 $4
    ;;

    *)
        echo "
            create-api: 创建Restful API
        "
    ;;

esac