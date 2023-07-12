import { Construct } from "constructs";
import { App, TerraformStack } from "cdktf";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
import { AwsProvider } from "./.gen/providers/aws/provider";

class MyStack extends TerraformStack {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    new AwsProvider(this, `example-aws-provider`, {
      region: "us-east-1",
    })

    new S3Bucket(this, "example-s3-bucket",{
      bucket: "example-cdktf-infra-cna128zasde",
      versioning: {
        enabled: true
      },
      tags: {
        project: "example"
      }
    })
  }
}

const app = new App();
new MyStack(app, "cdktf-infrastructure");
app.synth();
