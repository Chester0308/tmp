### ec2 インスタンス内で instance id を取得
```
instance_id=`ec2-metadata --instance-id | sed -e "s@.*: \(.*\)@\1@g"`
```

### ec2 の情報を取得
```
aws ec2 describe-instances --instance-ids ${instance_id} --region ${REGION}

// vpc id 
vpc_id=`aws ec2 describe-instances --instance-ids ${instance_id} --region ${REGION} | jq -r '.Reservations[].Instances[].NetworkInterfaces[].VpcId'`
```

### vpc の cidr を取得
```
cider=`aws ec2 describe-vpcs --vpc-ids ${vpc_id} --region ${REGION} | jq -r '.Vpcs[].CidrBlockAssociationSet[].CidrBlock'`
```
