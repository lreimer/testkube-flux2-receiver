package main

import (
	"github.com/gofiber/fiber/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Event struct {
	// The object that this event is about.
	// +required
	InvolvedObject corev1.ObjectReference `json:"involvedObject"`

	// Severity type of this event (info, error)
	// +required
	Severity string `json:"severity"`

	// The time at which this event was recorded.
	// +required
	Timestamp metav1.Time `json:"timestamp"`

	// A human-readable description of this event.
	// Maximum length 39,000 characters
	// +required
	Message string `json:"message"`

	// A machine understandable string that gives the reason
	// for the transition into the object's current status.
	// +required
	Reason string `json:"reason"`

	// Metadata of this event, e.g. apply change set.
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name of the controller that emitted this event, e.g. `source-controller`.
	// +required
	ReportingController string `json:"reportingController"`

	// ID of the controller instance, e.g. `source-controller-xyzf`.
	// +optional
	ReportingInstance string `json:"reportingInstance,omitempty"`
}

const (
	EventSeverityInfo  string = "info"
	EventSeverityError string = "error"
)

func ProcessEvent(c *fiber.Ctx) error {
	event := new(Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// TODO

	return c.Status(200).SendString("")
}
