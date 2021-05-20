package sensitivity

// Sensitivity Sensitivity of the threshold.
// With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.
// With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert.
type Sensitivity string

const High = Sensitivity("HIGH")
const Low = Sensitivity("LOW")
const Medium = Sensitivity("MEDIUM")
