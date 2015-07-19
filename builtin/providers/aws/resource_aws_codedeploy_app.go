package aws

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

func resourceAwsCodeDeployApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsCodeDeployAppCreate,
		Read:   resourceAwsCodeDeployAppRead,
		Update: resourceAwsCodeDeployUpdate,
		Delete: resourceAwsCodeDeployAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsCodeDeployAppCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).codedeployconn

	application := d.Get("name").(string)
	log.Printf("[DEBUG] Creating CodeDeploy application %s", application)

	resp, err := conn.CreateApplication(&codedeploy.CreateApplicationInput{
		ApplicationName: aws.String(application),
	})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] CodeDeploy application %s created", *resp.ApplicationID)

	d.SetId(*resp.ApplicationID)
	d.Set("name", application)
	return nil
}

func resourceAwsCodeDeployAppRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).codedeployconn

	application := d.Get("name").(string)
	log.Printf("[DEBUG] Reading CodeDeploy application %s", application)
	resp, err := conn.GetApplication(&codedeploy.GetApplicationInput{
		ApplicationName: aws.String(application),
	})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] CodeDeploy application %s created", *resp.Application.ApplicationName)

	d.SetId(*resp.Application.ApplicationID)
	d.Set("name", *resp.Application.ApplicationName)

	return nil
}

func resourceAwsCodeDeployUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).codedeployconn

	o, n := d.GetChange("name")

	_, err := conn.UpdateApplication(&codedeploy.UpdateApplicationInput{
		ApplicationName:    aws.String(o.(string)),
		NewApplicationName: aws.String(n.(string)),
	})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] CodeDeploy application %s updated", n)

	d.Set("name", n)

	return nil
}

func resourceAwsCodeDeployAppDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).codedeployconn

	_, err := conn.DeleteApplication(&codedeploy.DeleteApplicationInput{
		ApplicationName: aws.String(d.Get("name").(string)),
	})
	if err != nil {
		if cderr, ok := err.(awserr.Error); ok && cderr.Code() == "InvalidApplicationNameException" {
			d.SetId("")
			return nil
		} else {
			log.Printf("[ERROR] Error deleting CodeDeploy application: %s", err)
			return err
		}
	}

	return nil
}
