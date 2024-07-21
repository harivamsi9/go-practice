package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	bldr := newNotificationBuilder()

	// TODO: Use the builder to set some properties
	bldr.setTitle("New Notification")
	bldr.setIcon("icon.jgp")
	bldr.setImage("image.jpg")
	bldr.setPriority(10)
	bldr.setMessage("This is a basic Notification with some text in it")
	bldr.setNotType("alert")

	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("ERROR creating the notification ", err)
	} else {
		fmt.Printf("Notification: %#v\n", *notif)
	}

}
