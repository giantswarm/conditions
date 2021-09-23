package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func TestIsTrue(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Returns true for condition with Status=True",
			input:          &capi.Condition{Status: corev1.ConditionTrue},
			expectedOutput: true,
		},
		{
			name:           "case 1: Returns false for condition with Status=False",
			input:          &capi.Condition{Status: corev1.ConditionFalse},
			expectedOutput: false,
		},
		{
			name:           "case 2: Returns false for condition with Status=Unknown",
			input:          &capi.Condition{Status: corev1.ConditionUnknown},
			expectedOutput: false,
		},
		{
			name:           "case 3: Returns false for nil",
			input:          nil,
			expectedOutput: false,
		},
		{
			name:           "case 4: Returns false for condition with Status set to unsupported value",
			input:          &capi.Condition{Status: corev1.ConditionStatus("SomethingSomething")},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			output := IsTrue(tc.input)

			// assert
			if output != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, output)
				t.Fail()
			}
		})
	}
}

func TestIsFalse(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Returns false for condition with Status=True",
			input:          &capi.Condition{Status: corev1.ConditionTrue},
			expectedOutput: false,
		},
		{
			name:           "case 1: Returns true for condition with Status=False",
			input:          &capi.Condition{Status: corev1.ConditionFalse},
			expectedOutput: true,
		},
		{
			name:           "case 2: Returns false for condition with Status=Unknown",
			input:          &capi.Condition{Status: corev1.ConditionUnknown},
			expectedOutput: false,
		},
		{
			name:           "case 3: Returns false for nil",
			input:          nil,
			expectedOutput: false,
		},
		{
			name:           "case 4: Returns false for condition with Status set to unsupported value",
			input:          &capi.Condition{Status: corev1.ConditionStatus("NotGonnaWork")},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			output := IsFalse(tc.input)

			// assert
			if output != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, output)
				t.Fail()
			}
		})
	}
}

func TestIsUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Returns false for condition with Status=True",
			input:          &capi.Condition{Status: corev1.ConditionTrue},
			expectedOutput: false,
		},
		{
			name:           "case 1: Returns false for condition with Status=False",
			input:          &capi.Condition{Status: corev1.ConditionFalse},
			expectedOutput: false,
		},
		{
			name:           "case 2: Returns true for condition with Status=Unknown",
			input:          &capi.Condition{Status: corev1.ConditionUnknown},
			expectedOutput: true,
		},
		{
			name:           "case 3: Returns true for nil",
			input:          nil,
			expectedOutput: true,
		},
		{
			name:           "case 4: Returns false for condition with Status set to unsupported value",
			input:          &capi.Condition{Status: corev1.ConditionStatus("")},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			output := IsUnknown(tc.input)

			// assert
			if output != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, output)
				t.Fail()
			}
		})
	}
}

type withCheckTestInput struct {
	condition    *capi.Condition
	valueToCheck string
}

func TestWithReason(t *testing.T) {
	testCases := []struct {
		name           string
		input          withCheckTestInput
		expectedOutput bool
	}{
		{
			name: "case 0: Check for correct reason returns true",
			input: withCheckTestInput{
				condition:    &capi.Condition{Reason: "ForReasons"},
				valueToCheck: "ForReasons",
			},
			expectedOutput: true,
		},
		{
			name: "case 1: Check for incorrect reason returns false",
			input: withCheckTestInput{
				condition:    &capi.Condition{Reason: "SomethingElse"},
				valueToCheck: "Something",
			},
			expectedOutput: false,
		},
		{
			name: "case 2: Check for nil condition returns false",
			input: withCheckTestInput{
				condition:    nil,
				valueToCheck: "Something",
			},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			withReasonCheck := WithReason(tc.input.valueToCheck)
			output := withReasonCheck(tc.input.condition)

			// assert
			if output != tc.expectedOutput {
				if tc.input.condition != nil {
					t.Logf(
						"expected %t for %q (WithReason param) == %q (condition Reason field), got %t",
						tc.expectedOutput,
						tc.input.valueToCheck,
						tc.input.condition.Reason,
						output)
				} else {
					t.Logf(
						"expected %t for %q (WithReason param) when checking nil condition, got %t",
						tc.expectedOutput,
						tc.input.valueToCheck,
						output)
				}

				t.Fail()
			}
		})
	}
}

type withSeverityCheckTestCase struct {
	name           string
	input          withSeverityInput
	expectedOutput bool
}
type withSeverityInput struct {
	condition *capi.Condition
	check     capi.ConditionSeverity
}

func newWithSeverityInput(condition *capi.Condition, check capi.ConditionSeverity) withSeverityInput {
	return withSeverityInput{
		condition: condition,
		check:     check,
	}
}

var (
	conditionWithSeverityInfo    = &capi.Condition{Severity: capi.ConditionSeverityInfo}
	conditionWithSeverityWarning = &capi.Condition{Severity: capi.ConditionSeverityWarning}
	conditionWithSeverityError   = &capi.Condition{Severity: capi.ConditionSeverityError}
	conditionWithSeverityNone    = &capi.Condition{Severity: capi.ConditionSeverityNone}
)

func TestWithSeverity(t *testing.T) {
	testCases := []withSeverityCheckTestCase{
		{
			name:           "case 0: Check for correct severity returns true",
			input:          newWithSeverityInput(conditionWithSeverityInfo, capi.ConditionSeverityInfo),
			expectedOutput: true,
		},
		{
			name:           "case 1: Check for correct severity returns true",
			input:          newWithSeverityInput(conditionWithSeverityNone, capi.ConditionSeverityNone),
			expectedOutput: true,
		},
		{
			name:           "case 2: Check for incorrect severity returns false",
			input:          newWithSeverityInput(conditionWithSeverityError, capi.ConditionSeverityWarning),
			expectedOutput: false,
		},
		{
			name:           "case 3: Check for incorrect severity returns false",
			input:          newWithSeverityInput(conditionWithSeverityNone, capi.ConditionSeverityInfo),
			expectedOutput: false,
		},
		{
			name:           "case 4: Check for incorrect severity returns false",
			input:          newWithSeverityInput(conditionWithSeverityWarning, capi.ConditionSeverityNone),
			expectedOutput: false,
		},
		{
			name:           "case 5: Check for severity Info for nil condition returns false",
			input:          newWithSeverityInput(nil, capi.ConditionSeverityInfo),
			expectedOutput: false,
		},
		{
			name:           "case 6: Check for severity Warning for nil condition returns false",
			input:          newWithSeverityInput(nil, capi.ConditionSeverityWarning),
			expectedOutput: false,
		},
		{
			name:           "case 7: Check for severity Error for nil condition returns false",
			input:          newWithSeverityInput(nil, capi.ConditionSeverityError),
			expectedOutput: false,
		},
		{
			name:           "case 8: Check for severity None for nil condition returns false",
			input:          newWithSeverityInput(nil, capi.ConditionSeverityNone),
			expectedOutput: false,
		},
		{
			name:           "case 9: Check for unsupported severity returns false",
			input:          newWithSeverityInput(conditionWithSeverityWarning, "Fatal"),
			expectedOutput: false,
		},
		{
			name:           "case 10: Check when unsupported severity is set returns false",
			input:          newWithSeverityInput(&capi.Condition{Severity: "Burning"}, capi.ConditionSeverityNone),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// arrange
			condition := tc.input.condition
			expectedSeverity := tc.input.check

			// act
			check := WithSeverity(expectedSeverity)
			output := check(condition)

			// assert
			if output != tc.expectedOutput {
				if condition != nil {
					t.Logf(
						"expected %t for WithSeverity(%q) when checking condition with Severity=%q), got %t",
						tc.expectedOutput,
						expectedSeverity,
						condition.Severity,
						output)
				} else {
					t.Logf(
						"expected %t for WithSeverity(%q) when checking nil condition, got %t",
						tc.expectedOutput,
						expectedSeverity,
						output)
				}

				t.Fail()
			}
		})
	}
}

func TestWithSeverityInfo(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Check for condition with severity Info returns true",
			input:          conditionWithSeverityInfo,
			expectedOutput: true,
		},
		{
			name:           "case 1: Check for condition with severity Warning returns false",
			input:          conditionWithSeverityWarning,
			expectedOutput: false,
		},
		{
			name:           "case 2: Check for condition with severity Error returns false",
			input:          conditionWithSeverityError,
			expectedOutput: false,
		},
		{
			name:           "case 3: Check for condition with severity None returns false",
			input:          conditionWithSeverityNone,
			expectedOutput: false,
		},
		{
			name:           "case 4: Check for nil condition returns false",
			input:          nil,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// arrange
			condition := tc.input

			// act
			check := WithSeverityInfo()
			output := check(condition)

			// assert
			if output != tc.expectedOutput {
				if condition != nil {
					t.Logf(
						"expected %t for WithSeverityInfo() when checking condition with Severity=%q, got %t",
						tc.expectedOutput,
						condition.Severity,
						output)
				} else {
					t.Logf(
						"expected %t for WithSeverityInfo() when checking nil condition, got %t",
						tc.expectedOutput,
						output)
				}

				t.Fail()
			}
		})
	}
}

func TestWithSeverityWarning(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Check for condition with severity Info returns false",
			input:          conditionWithSeverityInfo,
			expectedOutput: false,
		},
		{
			name:           "case 1: Check for condition with severity Warning returns true",
			input:          conditionWithSeverityWarning,
			expectedOutput: true,
		},
		{
			name:           "case 2: Check for condition with severity Error returns false",
			input:          conditionWithSeverityError,
			expectedOutput: false,
		},
		{
			name:           "case 3: Check for condition with severity None returns false",
			input:          conditionWithSeverityNone,
			expectedOutput: false,
		},
		{
			name:           "case 4: Check for nil condition returns false",
			input:          nil,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// arrange
			condition := tc.input

			// act
			check := WithSeverityWarning()
			output := check(condition)

			// assert
			if output != tc.expectedOutput {
				if condition != nil {
					t.Logf(
						"expected %t for WithSeverityWarning() when checking condition with Severity=%q, got %t",
						tc.expectedOutput,
						condition.Severity,
						output)
				} else {
					t.Logf(
						"expected %t for WithSeverityWarning() when checking nil condition, got %t",
						tc.expectedOutput,
						output)
				}

				t.Fail()
			}
		})
	}
}

func TestWithSeverityError(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Check for condition with severity Info returns false",
			input:          conditionWithSeverityInfo,
			expectedOutput: false,
		},
		{
			name:           "case 1: Check for condition with severity Warning returns false",
			input:          conditionWithSeverityWarning,
			expectedOutput: false,
		},
		{
			name:           "case 2: Check for condition with severity Error returns true",
			input:          conditionWithSeverityError,
			expectedOutput: true,
		},
		{
			name:           "case 3: Check for condition with severity None returns false",
			input:          conditionWithSeverityNone,
			expectedOutput: false,
		},
		{
			name:           "case 4: Check for nil condition returns false",
			input:          nil,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// arrange
			condition := tc.input

			// act
			check := WithSeverityError()
			output := check(condition)

			// assert
			if output != tc.expectedOutput {
				if condition != nil {
					t.Logf(
						"expected %t for WithSeverityError() when checking condition with Severity=%q, got %t",
						tc.expectedOutput,
						condition.Severity,
						output)
				} else {
					t.Logf(
						"expected %t for WithSeverityError() when checking nil condition, got %t",
						tc.expectedOutput,
						output)
				}

				t.Fail()
			}
		})
	}
}

func TestWithoutSeverity(t *testing.T) {
	testCases := []struct {
		name           string
		input          *capi.Condition
		expectedOutput bool
	}{
		{
			name:           "case 0: Check for condition with severity Info returns false",
			input:          conditionWithSeverityInfo,
			expectedOutput: false,
		},
		{
			name:           "case 1: Check for condition with severity Warning returns false",
			input:          conditionWithSeverityWarning,
			expectedOutput: false,
		},
		{
			name:           "case 2: Check for condition with severity Error returns false",
			input:          conditionWithSeverityError,
			expectedOutput: false,
		},
		{
			name:           "case 3: Check for condition with severity None returns true",
			input:          conditionWithSeverityNone,
			expectedOutput: true,
		},
		{
			name:           "case 4: Check for nil condition returns false",
			input:          nil,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// arrange
			condition := tc.input

			// act
			check := WithoutSeverity()
			output := check(condition)

			// assert
			if output != tc.expectedOutput {
				if condition != nil {
					t.Logf(
						"expected %t for WithoutSeverity() when checking condition with Severity=%q, got %t",
						tc.expectedOutput,
						condition.Severity,
						output)
				} else {
					t.Logf(
						"expected %t for WithoutSeverity() when checking nil condition, got %t",
						tc.expectedOutput,
						output)
				}

				t.Fail()
			}
		})
	}
}
