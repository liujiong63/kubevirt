/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 *
 */

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"

	"kubevirt.io/api/instancetype/v1beta1"
)

/*
 * The following functions are manually defined to workaround conversion-gen
 * warnings about attributes in newer versions not being present in older
 * versions of the API.
 *
 * No custom code should be needed in such cases with each attribute
 * automatically being documented in generated comments within the used
 * autoConvert funcs.
 */

func Convert_v1beta1_VirtualMachinePreferenceSpec_To_v1alpha1_VirtualMachinePreferenceSpec(in *v1beta1.VirtualMachinePreferenceSpec, out *VirtualMachinePreferenceSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachinePreferenceSpec_To_v1alpha1_VirtualMachinePreferenceSpec(in, out, s)
}
