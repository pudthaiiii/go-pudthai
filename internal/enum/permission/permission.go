package permission

type Action string
type Subject string

const (
	MANAGER Action = "Manager"
	VIEW    Action = "View"
	NONE    Action = "None"
	COOKIE  Action = "Cookie"
)

const (
	CONSOLE Subject = "Console"
	AUTH    Subject = "Auth"
	USERS   Subject = "Users"
	ROLES   Subject = "Roles"
)
