package splunk

const (
	ACLProperty = "eai:acl"
)

struct {
App            string `json:"app"`
CanChangePerms bool   `json:"can_change_perms"`
CanList        bool   `json:"can_list"`
CanShareApp    bool   `json:"can_share_app"`
CanShareGlobal bool   `json:"can_share_global"`
CanShareUser   bool   `json:"can_share_user"`
CanWrite       bool   `json:"can_write"`
Modifiable     bool   `json:"modifiable"`
Owner          string `json:"owner"`
Perms          struct {
Read  []string `json:"read"`
Write []string `json:"write"`
} `json:"perms"`
Removable bool   `json:"removable"`
Sharing   string `json:"sharing"`
}

type ACL struct {
	// The app context for the resource. Required for updating saved search ACL properties.
	// Allowed values are: The name of an app or system
	App string `schema:"app" json:"app"`

	// Indicates if the active user can change permissions for this object.
	CanChangePerms bool `schema:"can_change_perms" json:"can_change_perms"`

	// Indicates if the active user can change sharing to app level.
	CanShareApp bool `schema:"can_share_app" json:"can_share_app"`

	// Indicates if the active user can change sharing to system level
	CanShareGlobal bool `schema:"can_share_global" json:"can_share_global"`

	// Indicates if the active user can change sharing to user level.
	CanShareUser bool `schema:"can_share_user" json:"can_share_user"`

	// Indicates if the active user can edit this object. Defaults to true.
	CanWrite bool  `schema:"can_write" json:"can_write"`

	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties.
	// nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner string  `schema:"owner" json:"owner"`

	Perms          struct {
		// Properties that indicate resource read permissions.
		Read  []string `schema:"perms_read" json:"read"`

		// Properties that indicate write permissions of the resource.
		Write []string `schema:"perms_write" json:"write"`
	} `json:"perms"`

	// Indicates whether an admin or user with sufficient permissions can delete the entity.
	Removable bool `schema:"removable" json:"removable"`

	// Indicates how the resource is shared. Required for updating any knowledge object ACL properties.
	// app: Shared within a specific app
	// global: (Default) Shared globally to all apps.
	// user: Private to a user
	Sharing string `schema:"sharing" json:"sharing"`
}

func (c *Client) ACLPost(acl *ACL, path string) (r ACL, e error) {
	params, e := encode(acl)
	if e != nil {
		return
	}

	// fmt.Printf("ACL PARAMS: %v\n", params)
	f, e := c.Post(path, params)
	if e != nil {
		return
	}

	r = feedToACL(f)
	return
}

func feedToACL(f Feed) (a ACL) {
	if len(f.Entry) == 0 {
		return ACL{}
	}

	e := f.Entry[0]

	acl := e.PropertyLookup(ACLProperty).Dict()
	perms := acl.PropertyLookup("perms").Dict()
	a.PermsRead = perms.PropertyLookup("read").List().Item
	a.PermsWrite = perms.PropertyLookup("write").List().Item

	decoder.Decode(&a, acl.PropertyMap())

	return
}
