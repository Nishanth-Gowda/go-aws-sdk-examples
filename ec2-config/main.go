package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/nishanth-gowda/go-aws-sdk-examples/ec2-config/Util"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	fmt.Println("Please enter the desired EC2 Operation 🚀")
	fmt.Printf("\n 1. Create Instance \n 2. Describe Instance \n 3. Terminate Instance \n")

	var choice int16
	fmt.Scanln(&choice)
	switch choice {
		case 1: {
				fmt.Println("Creating an instance ✨")

				var name string
				fmt.Println("Enter the name of the tag to attach to the instance")
				fmt.Scanln(&name)

				var value string
				fmt.Println("Enter the value of the tag to attach to the instance")
				fmt.Scanln(&value)

				if name == "" || value == "" {
					fmt.Println("You must supply a name and value for the tag ❕")
					return
				}

				namePtr := &name
				valuePtr := &value

				res, err := Util.CreateInstance(svc, namePtr, valuePtr)
				if err != nil {
					fmt.Println("Got an error creating an instance with tag " + name)
					return
				}

							
				fmt.Println("Instance created successfully ✅")
				fmt.Println("Created tagged instance with ID " + *res.Instances[0].InstanceId)
			}

		case 2: {
			result, err := Util.GetInstanceDetails(svc, sess)
    		if err != nil {
			fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
			fmt.Println(err)
			return
    		}

			for _, r := range result.Reservations {
				fmt.Println("Reservation ID: " + *r.ReservationId)
				for _, i := range r.Instances {
					fmt.Println("InstanceId" + "   " + *i.InstanceId)
					fmt.Println("InstanceType" + "   " + *i.InstanceType)
					fmt.Println("State" + "   " + *i.State.Name)
					fmt.Println("Private IP" + "   " + *i.PrivateIpAddress)
					fmt.Println("Public IP" + "   " + *i.PublicIpAddress)
					fmt.Println("Tags" + "   " + fmt.Sprint(i.Tags))

				}
		
				fmt.Println("")
			}
		}

	case 3 : {
			fmt.Println("Terminating an instance ✨")
			var instanceId string
			fmt.Println("Enter the instance ID to terminate")
			fmt.Scanln(&instanceId)
			if instanceId == "" {
				fmt.Println("You must supply an instance ID to terminate ❕")
				return
			}

			/*Terminate instance only if it is in Running State*/
			result, err := Util.GetInstanceDetails(svc, sess)
    		if err != nil {
			fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
			fmt.Println(err)
			return
    		}

			for _, r := range result.Reservations {
				for _, i := range r.Instances {
					if *i.InstanceId == instanceId && *i.State.Name == "running" {
						fmt.Println("Instance is in Running State")
						err := Util.TerminateInstance(svc, instanceId)
						if err != nil {
							fmt.Println("Got an error terminating the instance with ID " + instanceId)
							return
						}
					} else {
						fmt.Println("Instance is not in Running State")
						return
					}
				}
			}
			fmt.Printf("Instance with id %v terminated successfully ✅\n", instanceId)
		}
	}

	
}