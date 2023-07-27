// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/arik-kfir/greenstar/backend/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type TenantResolver interface {
	Accounts(ctx context.Context, obj *model.Tenant, rootsOnly *bool) ([]*model.Account, error)
	Account(ctx context.Context, obj *model.Tenant, id string) (*model.Account, error)
	Transactions(ctx context.Context, obj *model.Tenant) ([]*model.Transaction, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Tenant_account_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Tenant_accounts_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *bool
	if tmp, ok := rawArgs["rootsOnly"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("rootsOnly"))
		arg0, err = ec.unmarshalOBoolean2ᚖbool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["rootsOnly"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Tenant_id(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_displayName(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_displayName(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DisplayName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_displayName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_accounts(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_accounts(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Tenant().Accounts(rctx, obj, fc.Args["rootsOnly"].(*bool))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Account)
	fc.Result = res
	return ec.marshalNAccount2ᚕᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐAccountᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_accounts(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "tenant":
				return ec.fieldContext_Account_tenant(ctx, field)
			case "id":
				return ec.fieldContext_Account_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Account_displayName(ctx, field)
			case "labels":
				return ec.fieldContext_Account_labels(ctx, field)
			case "children":
				return ec.fieldContext_Account_children(ctx, field)
			case "parent":
				return ec.fieldContext_Account_parent(ctx, field)
			case "outgoingTransactions":
				return ec.fieldContext_Account_outgoingTransactions(ctx, field)
			case "incomingTransactions":
				return ec.fieldContext_Account_incomingTransactions(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Account", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Tenant_accounts_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_account(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_account(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Tenant().Account(rctx, obj, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Account)
	fc.Result = res
	return ec.marshalOAccount2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐAccount(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_account(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "tenant":
				return ec.fieldContext_Account_tenant(ctx, field)
			case "id":
				return ec.fieldContext_Account_id(ctx, field)
			case "displayName":
				return ec.fieldContext_Account_displayName(ctx, field)
			case "labels":
				return ec.fieldContext_Account_labels(ctx, field)
			case "children":
				return ec.fieldContext_Account_children(ctx, field)
			case "parent":
				return ec.fieldContext_Account_parent(ctx, field)
			case "outgoingTransactions":
				return ec.fieldContext_Account_outgoingTransactions(ctx, field)
			case "incomingTransactions":
				return ec.fieldContext_Account_incomingTransactions(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Account", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Tenant_account_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Tenant_transactions(ctx context.Context, field graphql.CollectedField, obj *model.Tenant) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Tenant_transactions(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Tenant().Transactions(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Transaction)
	fc.Result = res
	return ec.marshalNTransaction2ᚕᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTransactionᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Tenant_transactions(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Tenant",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Transaction_id(ctx, field)
			case "Date":
				return ec.fieldContext_Transaction_Date(ctx, field)
			case "targetAccount":
				return ec.fieldContext_Transaction_targetAccount(ctx, field)
			case "sourceAccount":
				return ec.fieldContext_Transaction_sourceAccount(ctx, field)
			case "referenceID":
				return ec.fieldContext_Transaction_referenceID(ctx, field)
			case "amount":
				return ec.fieldContext_Transaction_amount(ctx, field)
			case "description":
				return ec.fieldContext_Transaction_description(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Transaction", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputTenantChanges(ctx context.Context, obj interface{}) (model.TenantChanges, error) {
	var it model.TenantChanges
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"displayName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "displayName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("displayName"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.DisplayName = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var tenantImplementors = []string{"Tenant"}

func (ec *executionContext) _Tenant(ctx context.Context, sel ast.SelectionSet, obj *model.Tenant) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, tenantImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Tenant")
		case "id":
			out.Values[i] = ec._Tenant_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "displayName":
			out.Values[i] = ec._Tenant_displayName(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "accounts":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Tenant_accounts(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "account":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Tenant_account(ctx, field, obj)
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "transactions":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Tenant_transactions(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNTenant2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v model.Tenant) graphql.Marshaler {
	return ec._Tenant(ctx, sel, &v)
}

func (ec *executionContext) marshalNTenant2ᚕᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenantᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Tenant) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenant(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v *model.Tenant) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Tenant(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTenantChanges2githubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenantChanges(ctx context.Context, v interface{}) (model.TenantChanges, error) {
	res, err := ec.unmarshalInputTenantChanges(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTenant2ᚖgithubᚗcomᚋarikᚑkfirᚋgreenstarᚋbackendᚋmodelᚐTenant(ctx context.Context, sel ast.SelectionSet, v *model.Tenant) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Tenant(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
