#!/bin/sh -x

while getopts u OPT
do
  case $OPT in
    "u" ) IS_UPDATE="TRUE" ;;
  esac
done


S3_TEMPLATES_BUCKET=cf-templates-c3hpvs53r8ht-ap-northeast-1
APP_ENV=stg
APP_NAME=ohtsuki
STACK=${APP_ENV}-${APP_NAME}
CONFIG_PROFILE=default
EC2_KeyName=adastria-swagger
RDS_SUBNETGROUP=none
GITHUB_BRANCH=master
AWS_REGION=ap-northeast-1

# s3 に template をアップロード & response.yml 取得
aws cloudformation package \
    --template-file ./main.yaml \
    --s3-bucket ${S3_TEMPLATES_BUCKET} \
    --output-template-file ./response.yml \
    --profile ${CONFIG_PROFILE}


# -u の場合アップデート
if [ "$IS_UPDATE" = "TRUE" ]; then
    COMMAND='update-stack'
    DISABLE_ROLLBACK=''
else
    COMMAND='create-stack'
    DISABLE_ROLLBACK='--disable-rollback'
fi

aws cloudformation ${COMMAND} \
    --stack-name ${STACK} \
    --region ${AWS_REGION} \
    --capabilities CAPABILITY_NAMED_IAM \
    --template-body file://./response.yml \
    --parameters \
        ParameterKey=GitHubUser,ParameterValue=yumemi \
        ParameterKey=GitHubToken,ParameterValue=cc06aa8e37f0da3917a338a0dc1f08f2912de079 \
        ParameterKey=GitHubRepo,ParameterValue=adastria_campaign \
        ParameterKey=GitHubBranch,ParameterValue=${GITHUB_BRANCH} \
        ParameterKey=VpcId,ParameterValue=vpc-03256a6632e8e1241 \
        ParameterKey=PublicSubnet1,ParameterValue=subnet-0943a4cc80b816d2e \
        ParameterKey=PublicSubnet2,ParameterValue=subnet-03b3eddb30e61febf \
        ParameterKey=PublicSubnet1CIDR,ParameterValue=10.192.10.0/24 \
        ParameterKey=PublicSubnet2CIDR,ParameterValue=10.192.11.0/24 \
        ParameterKey=RDSMultiAZSubnetGroup,ParameterValue=${RDS_SUBNETGROUP} \
        ParameterKey=EC2KeyName,ParameterValue=${EC2_KeyName} \
        ParameterKey=RDSMasterUserPassword,ParameterValue=SDrriW7g \
        ParameterKey=RDSAppUserPassword,ParameterValue=beZoKnET \
        ParameterKey=RDSAppDatabaseName,ParameterValue=campaign \
        ParameterKey=GameTop,ParameterValue=https://stg-styletap-web.dot-st.com/campaign/index.html \
        ParameterKey=GameEnd,ParameterValue=https://stg-styletap-web.dot-st.com/campaign/thankyou.html \
        ParameterKey=FtpEndpoint,ParameterValue=ftp-stg.internal.dot-st.com \
        ParameterKey=FtpUserPassword,ParameterValue=sK3vVXdT \
        ParameterKey=FtpDir,ParameterValue=/coredb/SEND/GRANTREDUCTION \
    --profile ${CONFIG_PROFILE} \
    ${DISABLE_ROLLBACK}
