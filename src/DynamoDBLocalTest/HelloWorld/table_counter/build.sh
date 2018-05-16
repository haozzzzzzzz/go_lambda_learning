#!/usr/bin/env bash
command=$1

endpointUrl="http://0.0.0.0:8000"
tableName="video_buddy_counter_dev"

FuncListTables() {
    aws dynamodb list-tables \
    --endpoint-url ${endpointUrl}
}

FuncCreateTable() {
    aws dynamodb create-table \
    --endpoint-url ${endpointUrl} \
    --table-name ${tableName} \
    --attribute-definitions \
        AttributeName=partition_key,AttributeType=S \
        AttributeName=sort_key,AttributeType=S \
    --key-schema \
        AttributeName=partition_key,KeyType=HASH \
        AttributeName=sort_key,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
}

FuncDescribeTable() {
    aws dynamodb describe-table \
    --endpoint-url ${endpointUrl} \
    --table-name ${tableName}
}

FuncScanTable() {
    aws dynamodb scan \
    --endpoint-url ${endpointUrl} \
    --table-name ${tableName}
}

FuncQuery() {
    aws dynamodb query \
    --endpoint-url ${endpointUrl} \
    --table-name ${tableName} \
    --key-condition-expression "home_id = :home_id" \
    --expression-attribute-values '{":home_id": {"N": "1"}}'
}

FuncDeleteTable() {
    aws dynamodb delete-table \
    --endpoint-url ${endpointUrl} \
    --table-name ${tableName}
}

case ${command} in
"list-tables")
    FuncListTables
;;
"create-table")
    FuncCreateTable
;;
"describe-table")
    FuncDescribeTable
;;
"scan-table")
    FuncScanTable
;;
"query")
    FuncQuery
;;
"delete-table")
    FuncDeleteTable
;;
*)
;;
esac