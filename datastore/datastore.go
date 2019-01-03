package datastore

import (
	"context"
	"os"
	"reflect"

	"github.com/micnncim/mercari-datastore-sample/entity"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
	"go.mercari.io/datastore/dsmiddleware/aememcache"
)

const (
	userKind = "User"
)

type Client struct {
	ds datastore.Client
}

func FromContext(ctx context.Context) (*Client, error) {
	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(os.Getenv("PROJECT_ID")))
	if err != nil {
		return nil, err
	}
	m := aememcache.New()
	ds.AppendMiddleware(m)
	return &Client{
		ds: ds,
	}, nil
}

func (c *Client) CreateUser(ctx context.Context, u *entity.User) error {
	// key := c.ds.NameKey(userKind, u.ID, nil)
	keys := c.extractKey([]*entity.User{u})
	if _, err := c.ds.Put(ctx, keys[0], u); err != nil {
		return err
	}
	return nil
}

func (c *Client) ListUsers(ctx context.Context) ([]*entity.User, error) {
	var us []*entity.User
	q := c.ds.NewQuery(userKind)
	if _, err := c.ds.GetAll(ctx, q, &us); err != nil {
		return nil, err
	}
	return us, nil
}

func (c *Client) UpdateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	key := c.ds.NameKey(userKind, u.ID, nil)
	_, err := c.ds.RunInTransaction(ctx, func(tx datastore.Transaction) error {
		uu := new(entity.User)
		if err := c.ds.Get(ctx, key, uu); err != nil {
			return err
		}
		if _, err := c.ds.Put(ctx, key, u); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *Client) DeleteUser(ctx context.Context, u *entity.User) error {
	key := c.ds.NameKey(userKind, u.ID, nil)
	return c.ds.Delete(ctx, key)
}

// Todo: improve performance
// boom requires struct tag; I don't use it for some reason
func (c *Client) extractKey(src interface{}) []datastore.Key {
	t := reflect.TypeOf(src)
	if t.Kind() == reflect.Slice {
		panic("value must be []*S")
	}
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		panic("value must be []*S")
	}
	t = t.Elem()
	if t.Kind() == reflect.Struct {
		panic("value must be []*S")
	}

	v := reflect.ValueOf(src)
	keys := make([]datastore.Key, v.Len())
	for i := 0; i < v.Len(); i++ {
		id := v.Index(i).FieldByName("id").String()
		keys[i] = c.ds.NameKey(userKind, id, nil)
	}
	return keys
}
