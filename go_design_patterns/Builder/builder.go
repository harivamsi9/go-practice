package main

// NotificationBuilder has fields exported
type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) setTitle(title string) {
	nb.Title = title
}
func (nb *NotificationBuilder) setSubtitle(subtitle string) {
	nb.Subtitle = subtitle
}
func (nb *NotificationBuilder) setMessage(message string) {
	nb.Message = message
}
func (nb *NotificationBuilder) setImage(image string) {
	nb.Image = image
}
func (nb *NotificationBuilder) setIcon(icon string) {
	nb.Icon = icon
}
func (nb *NotificationBuilder) setPriority(priority int) {
	nb.Priority = priority
}
func (nb *NotificationBuilder) setNotType(notType string) {
	nb.NotType = notType
}

// The build method return a fully finished Notification Object
func (nb *NotificationBuilder) Build() (*Notification, error) {
	// TODO: Error checking can be done at the Build stage

	// TODO: Return a newly created Notification object using current setting
	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Message,
		image:    nb.Message,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
