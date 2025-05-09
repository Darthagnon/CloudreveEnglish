// Code generated by ent, DO NOT EDIT.

package group

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cloudreve/Cloudreve/v4/ent/predicate"
	"github.com/cloudreve/Cloudreve/v4/pkg/boolset"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// MaxStorage applies equality check predicate on the "max_storage" field. It's identical to MaxStorageEQ.
func MaxStorage(v int64) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxStorage, v))
}

// SpeedLimit applies equality check predicate on the "speed_limit" field. It's identical to SpeedLimitEQ.
func SpeedLimit(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldSpeedLimit, v))
}

// Permissions applies equality check predicate on the "permissions" field. It's identical to PermissionsEQ.
func Permissions(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldPermissions, v))
}

// StoragePolicyID applies equality check predicate on the "storage_policy_id" field. It's identical to StoragePolicyIDEQ.
func StoragePolicyID(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldStoragePolicyID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldName, v))
}

// MaxStorageEQ applies the EQ predicate on the "max_storage" field.
func MaxStorageEQ(v int64) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxStorage, v))
}

// MaxStorageNEQ applies the NEQ predicate on the "max_storage" field.
func MaxStorageNEQ(v int64) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldMaxStorage, v))
}

// MaxStorageIn applies the In predicate on the "max_storage" field.
func MaxStorageIn(vs ...int64) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldMaxStorage, vs...))
}

// MaxStorageNotIn applies the NotIn predicate on the "max_storage" field.
func MaxStorageNotIn(vs ...int64) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldMaxStorage, vs...))
}

// MaxStorageGT applies the GT predicate on the "max_storage" field.
func MaxStorageGT(v int64) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldMaxStorage, v))
}

// MaxStorageGTE applies the GTE predicate on the "max_storage" field.
func MaxStorageGTE(v int64) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldMaxStorage, v))
}

// MaxStorageLT applies the LT predicate on the "max_storage" field.
func MaxStorageLT(v int64) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldMaxStorage, v))
}

// MaxStorageLTE applies the LTE predicate on the "max_storage" field.
func MaxStorageLTE(v int64) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldMaxStorage, v))
}

// MaxStorageIsNil applies the IsNil predicate on the "max_storage" field.
func MaxStorageIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldMaxStorage))
}

// MaxStorageNotNil applies the NotNil predicate on the "max_storage" field.
func MaxStorageNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldMaxStorage))
}

// SpeedLimitEQ applies the EQ predicate on the "speed_limit" field.
func SpeedLimitEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldSpeedLimit, v))
}

// SpeedLimitNEQ applies the NEQ predicate on the "speed_limit" field.
func SpeedLimitNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldSpeedLimit, v))
}

// SpeedLimitIn applies the In predicate on the "speed_limit" field.
func SpeedLimitIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldSpeedLimit, vs...))
}

// SpeedLimitNotIn applies the NotIn predicate on the "speed_limit" field.
func SpeedLimitNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldSpeedLimit, vs...))
}

// SpeedLimitGT applies the GT predicate on the "speed_limit" field.
func SpeedLimitGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldSpeedLimit, v))
}

// SpeedLimitGTE applies the GTE predicate on the "speed_limit" field.
func SpeedLimitGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldSpeedLimit, v))
}

// SpeedLimitLT applies the LT predicate on the "speed_limit" field.
func SpeedLimitLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldSpeedLimit, v))
}

// SpeedLimitLTE applies the LTE predicate on the "speed_limit" field.
func SpeedLimitLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldSpeedLimit, v))
}

// SpeedLimitIsNil applies the IsNil predicate on the "speed_limit" field.
func SpeedLimitIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldSpeedLimit))
}

// SpeedLimitNotNil applies the NotNil predicate on the "speed_limit" field.
func SpeedLimitNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldSpeedLimit))
}

// PermissionsEQ applies the EQ predicate on the "permissions" field.
func PermissionsEQ(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldPermissions, v))
}

// PermissionsNEQ applies the NEQ predicate on the "permissions" field.
func PermissionsNEQ(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldPermissions, v))
}

// PermissionsIn applies the In predicate on the "permissions" field.
func PermissionsIn(vs ...*boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldPermissions, vs...))
}

// PermissionsNotIn applies the NotIn predicate on the "permissions" field.
func PermissionsNotIn(vs ...*boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldPermissions, vs...))
}

// PermissionsGT applies the GT predicate on the "permissions" field.
func PermissionsGT(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldPermissions, v))
}

// PermissionsGTE applies the GTE predicate on the "permissions" field.
func PermissionsGTE(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldPermissions, v))
}

// PermissionsLT applies the LT predicate on the "permissions" field.
func PermissionsLT(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldPermissions, v))
}

// PermissionsLTE applies the LTE predicate on the "permissions" field.
func PermissionsLTE(v *boolset.BooleanSet) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldPermissions, v))
}

// SettingsIsNil applies the IsNil predicate on the "settings" field.
func SettingsIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldSettings))
}

// SettingsNotNil applies the NotNil predicate on the "settings" field.
func SettingsNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldSettings))
}

// StoragePolicyIDEQ applies the EQ predicate on the "storage_policy_id" field.
func StoragePolicyIDEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldStoragePolicyID, v))
}

// StoragePolicyIDNEQ applies the NEQ predicate on the "storage_policy_id" field.
func StoragePolicyIDNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldStoragePolicyID, v))
}

// StoragePolicyIDIn applies the In predicate on the "storage_policy_id" field.
func StoragePolicyIDIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldStoragePolicyID, vs...))
}

// StoragePolicyIDNotIn applies the NotIn predicate on the "storage_policy_id" field.
func StoragePolicyIDNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldStoragePolicyID, vs...))
}

// StoragePolicyIDIsNil applies the IsNil predicate on the "storage_policy_id" field.
func StoragePolicyIDIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldStoragePolicyID))
}

// StoragePolicyIDNotNil applies the NotNil predicate on the "storage_policy_id" field.
func StoragePolicyIDNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldStoragePolicyID))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStoragePolicies applies the HasEdge predicate on the "storage_policies" edge.
func HasStoragePolicies() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StoragePoliciesTable, StoragePoliciesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStoragePoliciesWith applies the HasEdge predicate on the "storage_policies" edge with a given conditions (other predicates).
func HasStoragePoliciesWith(preds ...predicate.StoragePolicy) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newStoragePoliciesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(sql.NotPredicates(p))
}
