package main

import "fmt"

type Report interface {
	Execute()
}

type Receiver struct {
}

type ConcreteReportA struct {
	receiver *Receiver
}

func (c *ConcreteReportA) Execute() {
	c.receiver.Action("ReportA")
}

type ConcreteReportB struct {
	receiver *Receiver
}

func (c *ConcreteReportB) Execute() {
	c.receiver.Action("ReportB")
}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

type Invoker struct {
	repository []Report
}

func (i *Invoker) Schedule(cmd Report) {
	i.repository = append(i.repository, cmd)
}

func (i *Invoker) Run() {
	for _, cmd := range i.repository {
		cmd.Execute()
	}
}

func main() {
	receiver := new(Receiver)
	ReportA := &ConcreteReportA{receiver}
	ReportB := &ConcreteReportB{receiver: receiver}
	invoker := new(Invoker)
	invoker.Schedule(ReportA)
	invoker.Run()
	invoker.Schedule(ReportB)
	invoker.Run()

}
