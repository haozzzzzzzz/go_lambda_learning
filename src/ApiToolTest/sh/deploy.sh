#!/usr/bin/env bash
source params.sh

# 复制文件到跳板机
cd ../stage/${stage}
scp -i ${jumpServerKey} ./deploy_${stage}_${serviceName}.zip ${jumpServer}:~/luohao/

cd ../../sh/
scp -i ${jumpServerKey} target_server.sh ${jumpServer}:~/luohao/

# 在跳板机执行操作
ssh -i ${jumpServerKey} ${jumpServer} 'bash -s' < "jump_server.sh" ${stage} ${serviceName} ${targetServerKey} ${targetServer}
