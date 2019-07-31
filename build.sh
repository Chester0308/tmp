#!/bin/sh
 
while getopts s OPT
do
  case $OPT in
    "s" ) IS_STG="TRUE" ;;
  esac
done

# 設定項目を設定し、リポジトリのルートディレクトリで実行する

# 設定項目 ================================================
APP_ENV=stg
APP_NAME=adastria-campaign
REGION=ap-northeast-1

 
# -s の場合 stg
if [ "$IS_STG" = "TRUE" ]; then
    PROFILE=adastria
    AWS_ACCOUNT_ID=1234567890
else
    PROFILE=default
    AWS_ACCOUNT_ID=1234567890
fi

# =========================================================


ECR_URL="${AWS_ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${APP_ENV}-${APP_NAME}"
$(aws ecr get-login --no-include-email --region ap-northeast-1 --profile ${PROFILE})

# push nginx image
docker build -t ${ECR_URL}-nginx:latest docker/nginx/
docker push ${ECR_URL}-nginx:latest

## push api image
#docker build -t ${ECR_URL}-api:latest -f docker/php-fpm/Dockerfile .
#docker push ${ECR_URL}-api:latest
