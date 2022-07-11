FROM public.ecr.aws/aws-ec2/amazon-ec2-metadata-mock:v1.11.1

ENTRYPOINT ["/ec2-metadata-mock"]