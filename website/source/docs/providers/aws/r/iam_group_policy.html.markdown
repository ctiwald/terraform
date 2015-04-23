---
layout: "aws"
page_title: "AWS: aws_group_policy"
sidebar_current: "docs-aws-resource-iam-group-policy"
description: |-
  Provides an IAM policy attached to a group.
---

# aws\_iam\_group\_policy

Provides an IAM policy attached to a group.

## Example Usage

```
resource "aws_iam_group" "my_developers" {
    name = "developers"
    path = "/users/"
}

resource "iam_group_policy" "my_developer_policy" {
    name = "my_developer_policy"
    group = "${aws_iam_group.my_developers.id}"
    policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `policy` - (Required) The policy document. This is a JSON formatted string.
  The heredoc syntax or `file` funciton is helpful here.
* `name` - (Required) Name of the policy.
* `user` - (Required) IAM group to which to attach this policy.

## Attributes Reference

This resource has no attributes.
