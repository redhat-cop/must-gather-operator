/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	redhatcopv1alpha1 "github.com/redhat-cop/must-gather-operator/api/v1alpha1"
	"github.com/redhat-cop/operator-utils/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var _ = Describe("MustGather controller", func() {
	Context("MustGather controller test", func() {

		const MustGatherName = "test-mustgather"

		ctx := context.Background()

		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:      MustGatherName,
				Namespace: MustGatherName,
			},
		}
		typeNamespaceName := types.NamespacedName{Name: MustGatherName, Namespace: MustGatherName}

		It("should successfully reconcile a custom resource for MustGather", func() {
			By("Creating the custom resource for the Kind MustGather")
			mustgather := &redhatcopv1alpha1.MustGather{}
			err := k8sClient.Get(ctx, typeNamespaceName, mustgather)
			if err != nil && errors.IsNotFound(err) {
				// Let's mock our custom resource at the same way that we would
				// apply on the cluster the manifest under config/samples
				mustgather := &redhatcopv1alpha1.MustGather{
					ObjectMeta: metav1.ObjectMeta{
						Name:      MustGatherName,
						Namespace: namespace.Name,
					},
					Spec: redhatcopv1alpha1.MustGatherSpec{},
				}

				err = k8sClient.Create(ctx, mustgather)
				Expect(err).To(Not(HaveOccurred()))
			}

			By("Checking if the custom resource was successfully created")
			Eventually(func() error {
				found := &redhatcopv1alpha1.MustGather{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Reconciling the custom resource created")
			mustgatherReconciler := &MustGatherReconciler{
				ReconcilerBase: util.NewReconcilerBase(k8sClient, k8sClient.Scheme(), cfg, nil, k8sClient),
				Log:            ctrl.Log.WithName("controllers").WithName("MustGather"),
			}
			mustgatherReconciler.init()

			result, err := mustgatherReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))
			Expect(result.Requeue).To(BeFalse())

		})

	})
})
