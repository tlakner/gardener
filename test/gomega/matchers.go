// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gomega

import (
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// DeepDerivativeEqual is similar to DeepEqual except that unset fields in actual are
// ignored (not compared). This allows us to focus on the fields that matter to
// the semantic comparison.
func DeepDerivativeEqual(expected interface{}) types.GomegaMatcher {
	// if CharactersAroundMismatchToInclude is to small, then format.MessageWithDiff will be unable to output our
	// mismatch message
	if format.CharactersAroundMismatchToInclude < 10 {
		format.CharactersAroundMismatchToInclude = 10
	}

	return &deepDerivativeMatcher{
		expected: expected,
	}
}

// BeNotFoundError checks if error is NotFound.
func BeNotFoundError() types.GomegaMatcher {
	return &kubernetesErrors{
		checkFunc: apierrors.IsNotFound,
		message:   "NotFound",
	}
}

// BeForbiddenError checks if error is Forbidden.
func BeForbiddenError() types.GomegaMatcher {
	return &kubernetesErrors{
		checkFunc: apierrors.IsForbidden,
		message:   "Forbidden",
	}
}

// BeBadRequestError checks if error is BadRequest.
func BeBadRequestError() types.GomegaMatcher {
	return &kubernetesErrors{
		checkFunc: apierrors.IsBadRequest,
		message:   "BadRequest",
	}
}
