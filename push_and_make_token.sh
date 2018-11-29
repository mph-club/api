#
# RUN me where kubectl is available,& make sure to replace account,region etc
#
AWS_REGION=us-east-1
SECRET_NAME=${AWS_REGION}-ecr-registry
EMAIL=dev@mphclub.com

#
# Fetch account number
#
AWS_ACCOUNT_NUMBER=`aws sts get-caller-identity --output text --query 'Account'`

#
# Fetch token (which will expire in 12 hours)
#

TOKEN=`aws ecr get-authorization-token --output text \
--query 'authorizationData[].authorizationToken' \
| base64 -D | cut -d: -f2`

#
# Create or replace registry secret
#
#eval `aws ecr get-login --region us-east-1 --no-include-email`
kubectl delete secret --ignore-not-found $SECRET_NAME
kubectl create secret docker-registry $SECRET_NAME \
 --docker-server=https://${AWS_ACCOUNT_NUMBER}.dkr.ecr.${AWS_REGION}.amazonaws.com \
 --docker-username=AWS \
 --docker-password="${TOKEN}" \
 --docker-email="${EMAIL}"
