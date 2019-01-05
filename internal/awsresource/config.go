package awsresource

import "time"

type Config struct {
	Updated time.Time `json:"updated"`
}

func (c *Config) GetUpdatedTime() time.Time {
	return c.Updated
}

func (c *Config) SetUpdatedTime(newUpdatedTime time.Time) {
	if newUpdatedTime.After(c.GetUpdatedTime()) {
		c.Updated = newUpdatedTime
	}
}
