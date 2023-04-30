// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package loadbalancer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/x/idx"
)

const (
	// Label holds the string label denoting the loadbalancer type in the database.
	Label = "load_balancer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIPAddressID holds the string denoting the ip_address_id field in the database.
	FieldIPAddressID = "ip_address_id"
	// Table holds the table name of the loadbalancer in the database.
	Table = "load_balancers"
)

// Columns holds all SQL columns for loadbalancer fields.
var Columns = []string{
	FieldID,
	FieldLocationID,
	FieldTenantID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldIPAddressID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	LocationIDValidator func(string) error
	// TenantIDValidator is a validator for the "tenant_id" field. It is called by the builders before save.
	TenantIDValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() idx.PrefixedID
)

// OrderOption defines the ordering options for the LoadBalancer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIPAddressID orders the results by the ip_address_id field.
func ByIPAddressID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddressID, opts...).ToFunc()
}
