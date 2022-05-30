# appflow-activator

Self contained binary to (de)activate an AWS appflow.

## Usage

```
resource "null_resource" "activate_flow" {
  triggers = {
    flow_name = aws_appflow_flow.example.name
  }

  provisioner "local-exec" {
    command = "/tmp/appflow-activator -flow-name=${self.triggers.flow_name}"
  }

  provisioner "local-exec" {
    when = destroy
    command = "/tmp/appflow-activator -flow-name=${self.triggers.flow_name} -deactivate"
  } 

  depends_on = [
    aws_cloudwatch_event_bus.example
  ]
}
```

## But why?

This is a workaround for https://github.com/hashicorp/terraform-provider-aws/issues/25085.

Alternatively, you could also use the aws cli directly, but that requires you to install the aws cli when running this on terraform cloud first...

```
resource "null_resource" "activate_flow" {
  triggers = {
    flow_name = aws_appflow_flow.example.name
  }

  provisioner "local-exec" {
    command = "aws appflow start-flow --flow-name ${self.triggers.flow_name}"
  }

  provisioner "local-exec" {
    when = destroy
    command = "aws appflow stop-flow --flow-name ${self.triggers.flow_name}"
  } 

  depends_on = [
    aws_cloudwatch_event_bus.example
  ]
}
```
