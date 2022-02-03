/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalSubscription writes a value of the 'subscription' type to the given writer.
func MarshalSubscription(object *Subscription, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeSubscription(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeSubscription writes a value of the 'subscription' type to the given stream.
func writeSubscription(object *Subscription, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(SubscriptionLinkKind)
	} else {
		stream.WriteString(SubscriptionKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.capabilities != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("capabilities")
		writeCapabilityList(object.capabilities, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_account_id")
		stream.WriteString(object.cloudAccountID)
		count++
	}
	present_ = object.bitmap_&32 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_provider_id")
		stream.WriteString(object.cloudProviderID)
		count++
	}
	present_ = object.bitmap_&64 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_id")
		stream.WriteString(object.clusterID)
		count++
	}
	present_ = object.bitmap_&128 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_billing_model")
		stream.WriteString(string(object.clusterBillingModel))
		count++
	}
	present_ = object.bitmap_&256 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("console_url")
		stream.WriteString(object.consoleURL)
		count++
	}
	present_ = object.bitmap_&512 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("consumer_uuid")
		stream.WriteString(object.consumerUUID)
		count++
	}
	present_ = object.bitmap_&1024 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cpu_total")
		stream.WriteInt(object.cpuTotal)
		count++
	}
	present_ = object.bitmap_&2048 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("created_at")
		stream.WriteString((object.createdAt).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&4096 != 0 && object.creator != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("creator")
		writeAccount(object.creator, stream)
		count++
	}
	present_ = object.bitmap_&8192 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("display_name")
		stream.WriteString(object.displayName)
		count++
	}
	present_ = object.bitmap_&16384 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("external_cluster_id")
		stream.WriteString(object.externalClusterID)
		count++
	}
	present_ = object.bitmap_&32768 != 0 && object.labels != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("labels")
		writeLabelList(object.labels, stream)
		count++
	}
	present_ = object.bitmap_&65536 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("last_reconcile_date")
		stream.WriteString((object.lastReconcileDate).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&131072 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("last_released_at")
		stream.WriteString((object.lastReleasedAt).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&262144 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("last_telemetry_date")
		stream.WriteString((object.lastTelemetryDate).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&524288 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("managed")
		stream.WriteBool(object.managed)
		count++
	}
	present_ = object.bitmap_&1048576 != 0 && object.metrics != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("metrics")
		writeSubscriptionMetricsList(object.metrics, stream)
		count++
	}
	present_ = object.bitmap_&2097152 != 0 && object.notificationContacts != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("notification_contacts")
		writeAccountList(object.notificationContacts, stream)
		count++
	}
	present_ = object.bitmap_&4194304 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("organization_id")
		stream.WriteString(object.organizationID)
		count++
	}
	present_ = object.bitmap_&8388608 != 0 && object.plan != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("plan")
		writePlan(object.plan, stream)
		count++
	}
	present_ = object.bitmap_&16777216 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("product_bundle")
		stream.WriteString(object.productBundle)
		count++
	}
	present_ = object.bitmap_&33554432 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("provenance")
		stream.WriteString(object.provenance)
		count++
	}
	present_ = object.bitmap_&67108864 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("region_id")
		stream.WriteString(object.regionID)
		count++
	}
	present_ = object.bitmap_&134217728 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("released")
		stream.WriteBool(object.released)
		count++
	}
	present_ = object.bitmap_&268435456 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("service_level")
		stream.WriteString(object.serviceLevel)
		count++
	}
	present_ = object.bitmap_&536870912 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("socket_total")
		stream.WriteInt(object.socketTotal)
		count++
	}
	present_ = object.bitmap_&1073741824 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("status")
		stream.WriteString(object.status)
		count++
	}
	present_ = object.bitmap_&2147483648 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("support_level")
		stream.WriteString(object.supportLevel)
		count++
	}
	present_ = object.bitmap_&4294967296 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("system_units")
		stream.WriteString(object.systemUnits)
		count++
	}
	present_ = object.bitmap_&8589934592 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("trial_end_date")
		stream.WriteString((object.trialEndDate).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&17179869184 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("updated_at")
		stream.WriteString((object.updatedAt).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&34359738368 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("usage")
		stream.WriteString(object.usage)
	}
	stream.WriteObjectEnd()
}

// UnmarshalSubscription reads a value of the 'subscription' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalSubscription(source interface{}) (object *Subscription, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readSubscription(iterator)
	err = iterator.Error
	return
}

// readSubscription reads a value of the 'subscription' type from the given iterator.
func readSubscription(iterator *jsoniter.Iterator) *Subscription {
	object := &Subscription{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == SubscriptionLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "capabilities":
			value := readCapabilityList(iterator)
			object.capabilities = value
			object.bitmap_ |= 8
		case "cloud_account_id":
			value := iterator.ReadString()
			object.cloudAccountID = value
			object.bitmap_ |= 16
		case "cloud_provider_id":
			value := iterator.ReadString()
			object.cloudProviderID = value
			object.bitmap_ |= 32
		case "cluster_id":
			value := iterator.ReadString()
			object.clusterID = value
			object.bitmap_ |= 64
		case "cluster_billing_model":
			text := iterator.ReadString()
			value := BillingModel(text)
			object.clusterBillingModel = value
			object.bitmap_ |= 128
		case "console_url":
			value := iterator.ReadString()
			object.consoleURL = value
			object.bitmap_ |= 256
		case "consumer_uuid":
			value := iterator.ReadString()
			object.consumerUUID = value
			object.bitmap_ |= 512
		case "cpu_total":
			value := iterator.ReadInt()
			object.cpuTotal = value
			object.bitmap_ |= 1024
		case "created_at":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.createdAt = value
			object.bitmap_ |= 2048
		case "creator":
			value := readAccount(iterator)
			object.creator = value
			object.bitmap_ |= 4096
		case "display_name":
			value := iterator.ReadString()
			object.displayName = value
			object.bitmap_ |= 8192
		case "external_cluster_id":
			value := iterator.ReadString()
			object.externalClusterID = value
			object.bitmap_ |= 16384
		case "labels":
			value := readLabelList(iterator)
			object.labels = value
			object.bitmap_ |= 32768
		case "last_reconcile_date":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.lastReconcileDate = value
			object.bitmap_ |= 65536
		case "last_released_at":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.lastReleasedAt = value
			object.bitmap_ |= 131072
		case "last_telemetry_date":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.lastTelemetryDate = value
			object.bitmap_ |= 262144
		case "managed":
			value := iterator.ReadBool()
			object.managed = value
			object.bitmap_ |= 524288
		case "metrics":
			value := readSubscriptionMetricsList(iterator)
			object.metrics = value
			object.bitmap_ |= 1048576
		case "notification_contacts":
			value := readAccountList(iterator)
			object.notificationContacts = value
			object.bitmap_ |= 2097152
		case "organization_id":
			value := iterator.ReadString()
			object.organizationID = value
			object.bitmap_ |= 4194304
		case "plan":
			value := readPlan(iterator)
			object.plan = value
			object.bitmap_ |= 8388608
		case "product_bundle":
			value := iterator.ReadString()
			object.productBundle = value
			object.bitmap_ |= 16777216
		case "provenance":
			value := iterator.ReadString()
			object.provenance = value
			object.bitmap_ |= 33554432
		case "region_id":
			value := iterator.ReadString()
			object.regionID = value
			object.bitmap_ |= 67108864
		case "released":
			value := iterator.ReadBool()
			object.released = value
			object.bitmap_ |= 134217728
		case "service_level":
			value := iterator.ReadString()
			object.serviceLevel = value
			object.bitmap_ |= 268435456
		case "socket_total":
			value := iterator.ReadInt()
			object.socketTotal = value
			object.bitmap_ |= 536870912
		case "status":
			value := iterator.ReadString()
			object.status = value
			object.bitmap_ |= 1073741824
		case "support_level":
			value := iterator.ReadString()
			object.supportLevel = value
			object.bitmap_ |= 2147483648
		case "system_units":
			value := iterator.ReadString()
			object.systemUnits = value
			object.bitmap_ |= 4294967296
		case "trial_end_date":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.trialEndDate = value
			object.bitmap_ |= 8589934592
		case "updated_at":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.updatedAt = value
			object.bitmap_ |= 17179869184
		case "usage":
			value := iterator.ReadString()
			object.usage = value
			object.bitmap_ |= 34359738368
		default:
			iterator.ReadAny()
		}
	}
	return object
}
