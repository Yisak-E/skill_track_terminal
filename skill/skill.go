package skill

import "time"

/**
A Skill struct with these fields:
- Name: the lookup key, e.g. git commit
- Command: e.g. git commit -m "message"
- Description: e.g. "used to commit changes"
- Notes or Flags: e.g. "can be used with -a, --amend ..."
- Category (optional): e.g. git or linux
- CreatedAt
*/

type Skill struct {
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
}
