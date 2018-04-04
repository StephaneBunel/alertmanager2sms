package handler

import (
	// Force init() of all handlers
	_ "github.com/StephaneBunel/alertmanager2sms/pkg/recipient/handler/fromcsv"
)
