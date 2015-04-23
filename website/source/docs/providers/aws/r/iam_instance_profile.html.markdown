---
layout: "aws"
page_title: "AWS: aws_iam_instance_profile"
sidebar_current: "docs-aws-resource-iam-instance-profile"
description: |-
  Provides an IAM instance profile.
---

# aws\_iam\_instance\_profile

Provides an IAM instance profile.

## Example Usage

```
resource "aws_iam_role" "role" {
    name = "test_role"
    path = "/"
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

resource "aws_iam_instance_profile" "test_profile" {
    name = "test_profile"
    roles = ["${aws_iam_role.role.name}"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The profile's name.
* `path` - (Optional, default "/") Path in which to create the profile.
* `roles` - (Required) A list of role names to include in the profile.

## Attribute Reference

This resource has no attributes.
