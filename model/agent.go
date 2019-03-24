package model

import "fmt"

type ReadAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func (r *ReadAgent) String() string {
	return fmt.Sprintf("<name:%s version:%s status:%s timestamp:%v> ,",
		r.Name, r.Version, r.Status, r.Timestamp)
}

type Command struct {
	Start       string `json:"start"`
	Stop        string `json:"stop"`
	HealthCheck string `json:"health_check"`
}

func (c *Command) String() string {
	return fmt.Sprintf("<start:%s stop:%s health check:%s",
		c.Start, c.Stop, c.HealthCheck)
}

type DesiredAgent struct {
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Status    string  `json:"status"`
	Timestamp int64   `json:"timestamp"`
	Tarball   string  `json:"tarball"`
	Md5       string  `json:"md_5"`
	Cmd       Command `json:"cmd"`
}

func (d *DesiredAgent) String() string {
	return fmt.Sprintf("<name:%s version:%s,status:%s,timestamp:%d,tarball:%s md5:%s cmd:%s>",
		d.Name, d.Version, d.Status, d.Timestamp, d.Tarball, d.Md5, d.Cmd.String())
}
