package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
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

type withReasonTestInput struct {
	condition     *capi.Condition
	reasonToCheck string
}

func TestWithReason(t *testing.T) {
	testCases := []struct {
		name           string
		input          withReasonTestInput
		expectedOutput bool
	}{
		{
			name: "case 0: Check for correct reason returns true",
			input: withReasonTestInput{
				condition:     &capi.Condition{Reason: "ForReasons"},
				reasonToCheck: "ForReasons",
			},
			expectedOutput: true,
		},
		{
			name: "case 1: Check for incorrect reason returns false",
			input: withReasonTestInput{
				condition:     &capi.Condition{Reason: "SomethingElse"},
				reasonToCheck: "Something",
			},
			expectedOutput: false,
		},
		{
			name: "case 1: Check for nil condition returns false",
			input: withReasonTestInput{
				condition:     nil,
				reasonToCheck: "Something",
			},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			withReasonCheck := WithReason(tc.input.reasonToCheck)
			output := withReasonCheck(tc.input.condition)

			// assert
			if output != tc.expectedOutput {

				if tc.input.condition != nil {
					t.Logf(
						"expected %t for %q (WithReason param) == %q (condition Reason field), got %t",
						tc.expectedOutput,
						tc.input.reasonToCheck,
						tc.input.condition.Reason,
						output)
				} else {
					t.Logf(
						"expected %t for %q (WithReason param) when checking nil condition, got %t",
						tc.expectedOutput,
						tc.input.reasonToCheck,
						output)
				}

				t.Fail()
			}
		})
	}
}
