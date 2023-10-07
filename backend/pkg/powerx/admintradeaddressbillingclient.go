package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminTradeAddressBilling struct {
	*PowerX
}

func (c *AdminTradeAddressBilling) ListBillingAddressesPage(ctx context.Context, req *powerxtypes.ListBillingAddressesPageRequest) (*powerxtypes.ListBillingAddressesPageReply, error) {
	res := &powerxtypes.ListBillingAddressesPageReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/address/billing/page-list").
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) GetBillingAddress(ctx context.Context, req *powerxtypes.GetBillingAddressRequest) (*powerxtypes.GetBillingAddressReply, error) {
	res := &powerxtypes.GetBillingAddressReply{}
	err := c.H.Df().Method(http.MethodGet).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		BindQuery(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) CreateBillingAddress(ctx context.Context, req *powerxtypes.CreateBillingAddressRequest) (*powerxtypes.CreateBillingAddressReply, error) {
	res := &powerxtypes.CreateBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/trade/address/billing").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PutBillingAddress(ctx context.Context, req *powerxtypes.PutBillingAddressRequest) (*powerxtypes.PutBillingAddressReply, error) {
	res := &powerxtypes.PutBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPut).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) PatchBillingAddress(ctx context.Context, req *powerxtypes.PatchBillingAddressRequest) (*powerxtypes.PatchBillingAddressReply, error) {
	res := &powerxtypes.PatchBillingAddressReply{}
	err := c.H.Df().Method(http.MethodPatch).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminTradeAddressBilling) DeleteBillingAddress(ctx context.Context, req *powerxtypes.DeleteBillingAddressRequest) (*powerxtypes.DeleteBillingAddressReply, error) {
	res := &powerxtypes.DeleteBillingAddressReply{}
	err := c.H.Df().Method(http.MethodDelete).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/trade/address/billing/%v", req.BillingAddressId)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
