package log

import (
	"github.com/Sirupsen/logrus"
	wrapper "github.com/evalphobia/go-log-wrapper/log"
	"golang.org/x/net/context"

	"github.com/eure/example-blog-golang/library/net/context/resource"
)

const (
	tagPrefix = "blog."
)

// init sets Sentry configurations
func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

// Packet is wrapper struct
type Packet wrapper.Packet

// Panic logs panic error
func (p Packet) Panic(v ...interface{}) {
	p = p.setOptionalData(v)
	p.setDefaultTag("panic")
	go (wrapper.Packet(p)).Panic()
}

// Error logs serious error
func (p Packet) Error(v ...interface{}) {
	p = p.setOptionalData(v)
	p.setDefaultTag("error")
	go (wrapper.Packet(p)).Error()
}

// Warn logs warning
func (p Packet) Warn(v ...interface{}) {
	p = p.setOptionalData(v)
	p.setDefaultTag("warn")
	go (wrapper.Packet(p)).Warn()
}

// Info logs information
func (p Packet) Info(v ...interface{}) {
	p = p.setOptionalData(v)
	p.setDefaultTag("info")
	go (wrapper.Packet(p)).Info()
}

// Debug logs development information
func (p Packet) Debug(v ...interface{}) {
	p = p.setOptionalData(v)
	p.setDefaultTag("debug")
	go (wrapper.Packet(p)).Debug()
}

// setOptionalData adds trace data
func (p Packet) setOptionalData(v interface{}) Packet {
	// set trace
	if p.TraceData == nil {
		if p.Trace == 0 {
			p.Trace = 20
		}
		if p.TraceSkip == 0 {
			p.TraceSkip = 4
		}
		p.TraceData = wrapper.GetTraces(p.Trace, p.TraceSkip)
	}

	// set *http.Request from context
	if ctx, ok := v[0].(context.Context); ok {
		p.Request = resource.GetRequest(ctx)
	}
	return p
}

// AddData adds logging data
func (p *Packet) AddData(v ...interface{}) *Packet {
	p.DataList = append(p.DataList, v...)
	return p
}

func (p *Packet) setDefaultTag(tag string) {
	if p.Tag != "" {
		return
	}
	p.Tag = tagPrefix + tag
}
