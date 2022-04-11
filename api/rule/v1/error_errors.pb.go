// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	errors "github.com/tkeel-io/kit/errors"
	codes "google.golang.org/grpc/codes"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ego package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

var errUnknown *errors.TError
var errOkStatus *errors.TError
var errNotFound *errors.TError
var errInvalidArgument *errors.TError
var errInternalStore *errors.TError
var errInternalError *errors.TError
var errRuleNotFound *errors.TError
var errForbidden *errors.TError
var errUnauthorized *errors.TError
var errFailedKafkaConnection *errors.TError
var errOkKafkaConnection *errors.TError
var errDuplicateCreate *errors.TError
var errCantDeleteRunningRule *errors.TError
var errFailedMysqlConnection *errors.TError
var errFailedClickhouseConnection *errors.TError

func init() {
	errUnknown = errors.New(int(codes.Unknown), "rule.v1.ERR_UNKNOWN", "未知类型")
	errors.Register(errUnknown)
	errOkStatus = errors.New(int(codes.OK), "rule.v1.ERR_OK_STATUS", "成功")
	errors.Register(errOkStatus)
	errNotFound = errors.New(int(codes.NotFound), "rule.v1.ERR_NOT_FOUND", "未找到资源")
	errors.Register(errNotFound)
	errInvalidArgument = errors.New(int(codes.InvalidArgument), "rule.v1.ERR_INVALID_ARGUMENT", "请求参数无效")
	errors.Register(errInvalidArgument)
	errInternalStore = errors.New(int(codes.Internal), "rule.v1.ERR_INTERNAL_STORE", "请求后端存储错误")
	errors.Register(errInternalStore)
	errInternalError = errors.New(int(codes.Internal), "rule.v1.ERR_INTERNAL_ERROR", "内部错误")
	errors.Register(errInternalError)
	errRuleNotFound = errors.New(int(codes.NotFound), "rule.v1.ERR_RULE_NOT_FOUND", "未找到对应规则")
	errors.Register(errRuleNotFound)
	errForbidden = errors.New(int(codes.PermissionDenied), "rule.v1.ERR_FORBIDDEN", "请确保用户对该资源拥有足够的权限")
	errors.Register(errForbidden)
	errUnauthorized = errors.New(int(codes.PermissionDenied), "rule.v1.ERR_UNAUTHORIZED", "请确保用户权限")
	errors.Register(errUnauthorized)
	errFailedKafkaConnection = errors.New(int(codes.OK), "rule.v1.ERR_FAILED_KAFKA_CONNECTION", "建立连接失败")
	errors.Register(errFailedKafkaConnection)
	errOkKafkaConnection = errors.New(int(codes.OK), "rule.v1.ERR_OK_KAFKA_CONNECTION", "成功")
	errors.Register(errOkKafkaConnection)
	errDuplicateCreate = errors.New(int(codes.InvalidArgument), "rule.v1.ERR_DUPLICATE_CREATE", "重复创建")
	errors.Register(errDuplicateCreate)
	errCantDeleteRunningRule = errors.New(int(codes.InvalidArgument), "rule.v1.ERR_CANT_DELETE_RUNNING_RULE", "不能删除正在运行的规则")
	errors.Register(errCantDeleteRunningRule)
	errFailedMysqlConnection = errors.New(int(codes.OK), "rule.v1.ERR_FAILED_MYSQL_CONNECTION", "建立连接失败")
	errors.Register(errFailedMysqlConnection)
	errFailedClickhouseConnection = errors.New(int(codes.OK), "rule.v1.ERR_FAILED_CLICKHOUSE_CONNECTION", "建立连接失败")
	errors.Register(errFailedClickhouseConnection)
}

func ErrUnknown() errors.Error {
	return errUnknown
}

func ErrOkStatus() errors.Error {
	return errOkStatus
}

func ErrNotFound() errors.Error {
	return errNotFound
}

func ErrInvalidArgument() errors.Error {
	return errInvalidArgument
}

func ErrInternalStore() errors.Error {
	return errInternalStore
}

func ErrInternalError() errors.Error {
	return errInternalError
}

func ErrRuleNotFound() errors.Error {
	return errRuleNotFound
}

func ErrForbidden() errors.Error {
	return errForbidden
}

func ErrUnauthorized() errors.Error {
	return errUnauthorized
}

func ErrFailedKafkaConnection() errors.Error {
	return errFailedKafkaConnection
}

func ErrOkKafkaConnection() errors.Error {
	return errOkKafkaConnection
}

func ErrDuplicateCreate() errors.Error {
	return errDuplicateCreate
}

func ErrCantDeleteRunningRule() errors.Error {
	return errCantDeleteRunningRule
}

func ErrFailedMysqlConnection() errors.Error {
	return errFailedMysqlConnection
}

func ErrFailedClickhouseConnection() errors.Error {
	return errFailedClickhouseConnection
}
