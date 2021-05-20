package detectionmodes

import "github.com/dtcookie/dynatrace/api/config/anomalydetection/services"

// DetectAutomatically has no documenation
const DetectAutomatically = services.DetectionMode("DETECT_AUTOMATICALLY")

// DetectUsingFixedThresholds has no documenation
const DetectUsingFixedThresholds = services.DetectionMode("DETECT_USING_FIXED_THRESHOLDS")

// DontDetect has no documenation
const DontDetect = services.DetectionMode("DONT_DETECT")
